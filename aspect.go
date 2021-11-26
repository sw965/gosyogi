package gosyogi

import(
  "fmt"
)

type Aspect struct {
  Board Board
  CapturedPieces CapturedPieces
  Turn Turn
}

func NewInitAspect() Aspect {
  return Aspect{Board:INIT_BOARD, CapturedPieces:CapturedPieces{}, Turn:FIRST}
}

func (aspect1 *Aspect) Equal(aspect2 *Aspect) bool {
  board1 := &aspect1.Board
  board2 := &aspect2.Board

  if !board1.Equal(board2) {
    return false
  }

  if !aspect1.CapturedPieces.Equal(aspect2.CapturedPieces) {
    return false
  }

  return aspect1.Turn == aspect2.Turn
}

//打ち歩詰めを考慮していない
func (aspect Aspect) newLegalMoves(history Aspects) Moves {
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

  selfTurnCapturedPieces := aspect.CapturedPieces[aspect.Turn]
  boardLegalMoves := aspect.Board.NewLegalMoves(aspect.Turn)
  handLegalMoves := make(Moves, 0, len(selfTurnCapturedPieces) * BOARD_ROW_SIZE * BOARD_COLUMN_SIZE)

  for pieceName, count := range selfTurnCapturedPieces {
    if count == 0 {
      continue
    }
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
    nextAspect, _ := aspect.Put(&move, history)
    //王手放置や自らが王手されにいくよう手を除外する
    if !nextAspect.Board.IsCheck(aspect.Turn) {
      result = append(result, move)
    }
  }
  return result
}

func (aspect Aspect) NewLegalMoves(history Aspects) Moves {
  legalMoves := aspect.newLegalMoves(history)
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
      nextAspect, nextHistory := aspect.Put(&move, history)
      //打ち歩詰めならば
      if len(nextAspect.newLegalMoves(nextHistory)) == 0 {
        continue
      }
    }
    result = append(result, move)
  }
  return result
}

func (aspect Aspect) Put(move *Move, history Aspects) (Aspect, Aspects) {
  positionBeforeMove := move.BeforePosition
  positionAfterMove := move.AfterPosition
  turn := aspect.Turn

  //参照透過を保つためにコピーする
  capturedPieces := aspect.CapturedPieces.Copy()
  aspect.CapturedPieces = capturedPieces

  if positionBeforeMove.IsCapturedPiece() {
    aspect.CapturedPieces[turn][move.PieceName] -= 1
    aspect.Board[positionAfterMove.Row][positionAfterMove.Column] = Piece{Name:move.PieceName, Turn:turn}
  } else {
    movePiece := aspect.Board[positionBeforeMove.Row][positionBeforeMove.Column]
    pieceNameOfMovePosition := aspect.Board[positionAfterMove.Row][positionAfterMove.Column].Name
    aspect.Board[positionBeforeMove.Row][positionBeforeMove.Column] = Piece{}

    if move.IsPromotion {
      promotionPieceName := PIECE_NAME_TO_PROMOTION_PIECE_NAME[movePiece.Name]
      movePiece.Name = promotionPieceName
    }
    aspect.Board[positionAfterMove.Row][positionAfterMove.Column] = movePiece

    if pieceNameOfMovePosition != "" {
      normalPieceName := PIECE_NAME_TO_NORMAL_PIECE_NAME[pieceNameOfMovePosition]
      aspect.CapturedPieces[turn][normalPieceName] += 1
    }
  }

  aspect.Turn = REVERSE_TURN[turn]
  history = append(history, aspect)
  return aspect, history
}

func (aspect *Aspect) IsRepetitionOfMoves(history Aspects) bool {
  return history.Count(aspect) == 4
}

func (aspect *Aspect) IsGameEnd(history Aspects) bool {
  return len(aspect.NewLegalMoves(history)) == 0 || aspect.IsRepetitionOfMoves(history)
}

func (aspect *Aspect) Winner(history Aspects) (Winner, error) {
  if aspect.IsGameEnd(history) {
    if aspect.IsRepetitionOfMoves(history) {
      //今の局面とまったく同じ局面が一番最初に現れたインデックスを取得する
      //すなわち千日手認定された局面の最初のインデックス
      equalFirstIndex := history.EqualFirstIndex(aspect)
      cutHistory := history[equalFirstIndex:]

      //先手側が連続王手の千日手をされた場合
      if cutHistory.IsALLCheck(FIRST) {
        return WINNER_P1, nil
      } else if cutHistory.IsALLCheck(SECOND) {
        return WINNER_P2, nil
      }
    }
    return DRAW, nil
  }
  return Winner{}, fmt.Errorf("ゲームが終了していない状態でWinnerは求められない")
}

type Aspects []Aspect

func NewHistory() Aspects {
  result := make(Aspects, 0, 256)
  result = append(result, NewInitAspect())
  return result
}

func (aspects Aspects) Count(aspect *Aspect) int {
  result := 0
  for _, iAspect := range aspects {
    if iAspect.Equal(aspect) {
      result += 1
    }
  }
  return result
}

func (aspects Aspects) EqualFirstIndex(aspect *Aspect) int {
  for i, iAspect := range aspects {
    if iAspect.Equal(aspect) {
      return i
    }
  }
  return -1
}

func (aspects Aspects) TurnFilter(turn Turn) Aspects {
  result := make(Aspects, len(aspects) / 2)
  for _, aspect := range aspects {
    if aspect.Turn == turn {
      result = append(result, aspect)
    }
  }
  return result
}

func (aspects Aspects) IsALLCheck(turn Turn) bool {
  cutAspects := aspects.TurnFilter(turn)
  for _, aspect := range cutAspects {
    if !aspect.Board.IsCheck(turn) {
      return false
    }
  }
  return true
}
