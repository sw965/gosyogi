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

func (aspect Aspect) LegalMoves() Moves {
  return Moves{}
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
