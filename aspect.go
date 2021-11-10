package gosyogi

type Aspect struct {
  Board Board
  CapturedPieceNames TurnToPieceNames
  Turn Turn
  History Boards
}

func NewAspect() Aspect {
  return Aspect{Board:INIT_BOARD, CapturedPieceNames:TurnToPieceNames{}, Turn:FIRST, History:Boards{}}
}

func (aspect Aspect) Put(move *Move) (Aspect, error) {
  positionBeforeMove := move.BeforePosition
  positionAfterMove := move.AfterPosition
  turn := aspect.Turn
  aspect.CapturedPieceNames = aspect.CapturedPieceNames.Copy()

  if positionBeforeMove.IsCapturedPiece() {
    newCapturedPieceNames, err := aspect.CapturedPieceNames[turn].Remove(move.PieceName)
    if err != nil {
      return Aspect{}, err
    }
    aspect.CapturedPieceNames[turn] = newCapturedPieceNames
    piece := PIECE_NAME_TO_PIECE[move.PieceName]
    aspect.Board[positionAfterMove.Row][positionAfterMove.Column] = PieceWithTurn{Piece:*piece, Turn:turn}
  } else {
    movePiece := aspect.Board[positionBeforeMove.Row][positionBeforeMove.Column]
    pieceNameOfMovePosition := aspect.Board[positionAfterMove.Row][positionAfterMove.Column].Piece.Name
    aspect.Board[positionBeforeMove.Row][positionBeforeMove.Column] = PieceWithTurn{}
    aspect.Board[positionAfterMove.Row][positionBeforeMove.Column] = movePiece
    if pieceNameOfMovePosition != "" {
      aspect.CapturedPieceNames[turn] = append(aspect.CapturedPieceNames[turn], pieceNameOfMovePosition)
    }
  }

  aspect.Turn = REVERSE_TURN[turn]
  aspect.History = append(aspect.History, aspect.Board)
  return aspect, nil
}
