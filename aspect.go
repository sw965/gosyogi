package gosyogi

import(
  "fmt"
)

type Aspect struct {
  Board Board
  EachTurnCapturedPieceNames EachTurnCapturedPieceNames
  Turn Turn
  History Boards
}

func NewAspect() Aspect {
  return Aspect{Board:INIT_BOARD, EachTurnCapturedPieceNames:EachTurnCapturedPieceNames{},
    Turn:FIRST, History:make(Boards, 0, 255)}
}

//打ち歩詰めを考慮していない関数
func (aspect Aspect) newLegalMoves() Moves {
  isNihu := aspect.Board.IsNiHu(aspect.Turn)
  handMoveFilter := func(position Position, move *Move) bool {
    pieceName := aspect.Board[position.Row][position.Column].Name

    if pieceName != "" {
      return false
    }

    if isNihu[position.Column] && move.PieceName == HU {
      return false
    }

    foulPositions, existsFoulPos := FOUL_POSITIONS[aspect.Turn][move.PieceName]
    //歩・桂・香が動けない場所に打つ場合
    if existsFoulPos && foulPositions.In(position) {
      return false
    }

    return true
  }

  capturedPieceNames := aspect.EachTurnCapturedPieceNames[aspect.Turn]
  boardLegalMoves := aspect.Board.NewLegalMoves(aspect.Turn)
  handLegalMoves := make(Moves, 0, len(capturedPieceNames) * BOARD_ROW_SIZE * BOARD_COLUMN_SIZE)

  for _, pieceName := range capturedPieceNames {
    for _, position := range BOARD_ALL_POSITIONS {
      move := Move{PieceName:pieceName, BeforePosition:CAPTURED_PIECE_POSITION, AfterPosition:position}
      if handMoveFilter(position, &move) {
        handLegalMoves = append(handLegalMoves, move)
      }
    }
  }

  handAndBoardLegalMoves := boardLegalMoves.Add(handLegalMoves)
  result := make(Moves, 0, len(handAndBoardLegalMoves))

  for _, move := range handAndBoardLegalMoves {
    nextAspect, err := aspect.Put(&move)
    if err != nil {
      panic(err)
    }

    if !nextAspect.Board.IsCheck(aspect.Turn) {
      result = append(result, move)
    }
  }
  return result
}

func (aspect Aspect) NewLegalMoves() Moves {
  legalMoves := aspect.newLegalMoves()
  enemyKingHeadPostion := aspect.Board.NewKingHeadPosition(REVERSE_TURN[aspect.Turn])

  //相手の玉頭に既に駒が存在する場合は、打ち歩詰めを考慮する必要がないので、処理を終了する
  if enemyKingHeadPostion.IsOutOBoardRange() || aspect.Board[enemyKingHeadPostion.Row][enemyKingHeadPostion.Column].Name != "" {
    return legalMoves
  }
  result := make(Moves, 0, len(legalMoves))

  for _, move := range legalMoves {
    positionBeforeMove := move.BeforePosition
    isHandMoveHu := positionBeforeMove == CAPTURED_PIECE_POSITION && move.PieceName == HU

    if isHandMoveHu && positionBeforeMove == enemyKingHeadPostion {
      nextAspect, err := aspect.Put(&move)

      if err != nil {
        panic(err)
      }

      //打ち歩詰めなら
      if len(nextAspect.newLegalMoves()) == 0 {
        continue
      }
    }
    result = append(result, move)
  }
  return result
}

func (aspect Aspect) Put(move *Move) (Aspect, error) {
  positionBeforeMove := move.BeforePosition
  positionAfterMove := move.AfterPosition
  turn := aspect.Turn
  aspect.EachTurnCapturedPieceNames = aspect.EachTurnCapturedPieceNames.Copy()

  if positionBeforeMove.IsCapturedPiece() {
    capturedPieceNames, err := aspect.EachTurnCapturedPieceNames[turn].Remove(move.PieceName)
    if err != nil {
      return Aspect{}, err
    }
    aspect.EachTurnCapturedPieceNames[turn] = capturedPieceNames
    aspect.Board[positionAfterMove.Row][positionAfterMove.Column] = Piece{Name:move.PieceName, Turn:turn}
  } else {
    movePiece := aspect.Board[positionBeforeMove.Row][positionBeforeMove.Column]
    pieceNameOfMovePosition := aspect.Board[positionAfterMove.Row][positionAfterMove.Column].Name
    aspect.Board[positionBeforeMove.Row][positionBeforeMove.Column] = Piece{}

    if move.IsPromotion {
      promotionPieceName := PIECE_NAME_TO_PROMOTION_PIECE_NAMES[movePiece.Name]
      movePiece.Name = promotionPieceName
    }
    aspect.Board[positionAfterMove.Row][positionAfterMove.Column] = movePiece

    if pieceNameOfMovePosition != "" {
      normalPieceName := PIECE_NAME_TO_NORMAL_PIECE_NAMES[pieceNameOfMovePosition]
      aspect.EachTurnCapturedPieceNames[turn] = append(aspect.EachTurnCapturedPieceNames[turn], normalPieceName)
    }
  }

  aspect.Turn = REVERSE_TURN[turn]
  aspect.History = append(aspect.History, aspect.Board)
  return aspect, nil
}

func (aspect *Aspect) IsGameEnd() bool {
  return len(aspect.NewLegalMoves()) == 0
}

func (aspect *Aspect) IsFirstWin() (bool, error) {
  if aspect.IsGameEnd() {
    return aspect.Turn == FIRST, nil
  } else {
    return false, fmt.Errorf("ゲームが終了していない状態でWinnerは求められない")
  }
}
