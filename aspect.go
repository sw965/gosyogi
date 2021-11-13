package gosyogi

type Aspect struct {
  Board Board
  EachTurnCapturedPieceNames EachTurnCapturedPieceNames
  Turn Turn
  History Boards
}

func NewAspect() Aspect {
  return Aspect{Board:INIT_BOARD, EachTurnCapturedPieceNames:EachTurnCapturedPieceNames{}, Turn:FIRST, History:Boards{}}
}

func (aspect Aspect) NewLegalMoves() Moves {
  isNihu := aspect.Board.IsNiHu()
  filter := func(position Position) bool {
    pieceName := aspect.Board[position.Row][position.Column].Name

    if pieceName != "" {
      return false
    }

    if isNihu[position.Column] && pieceName == HU {
      return false
    }

    foulPositions, existsFoulPos := FOUL_POSITIONS[aspect.Turn][pieceName]

    if existsFoulPos && foulPositions.In(position) {
      return false
    }

    return true
  }

  capturedPieceNames := aspect.EachTurnCapturedPieceNames[aspect.Turn]
  result := make(Moves, 0, len(capturedPieceNames) * BOARD_ROW_SIZE * BOARD_COLUMN_SIZE)
  for _, pieceName := range capturedPieceNames {
    for _, position := range BOARD_ALL_POSITIONS {
      move := Move{PieceName:pieceName, BeforePosition:CAPTURED_PIECE_POSITION, AfterPosition:position}
      if filter(position) {
        result = append(result, move)
      }
    }
  }
  return result.Add(aspect.Board.NewLegalMoves(aspect.Turn))
}

func (aspect Aspect) Put(move *Move) (Aspect, error) {
  positionBeforeMove := move.BeforePosition
  positionAfterMove := move.AfterPosition
  turn := aspect.Turn
  aspect.EachTurnCapturedPieceNames = aspect.EachTurnCapturedPieceNames.Copy()

  if positionBeforeMove.IsCapturedPiece() {
    newEachTurnCapturedPieceNames, err := aspect.EachTurnCapturedPieceNames[turn].Remove(move.PieceName)
    if err != nil {
      return Aspect{}, err
    }
    aspect.EachTurnCapturedPieceNames[turn] = newEachTurnCapturedPieceNames
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
