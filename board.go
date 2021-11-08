package gosyogi

const (
  BOARD_ROW_SIZE = 9
  BOARD_COLUMN_SIZE = 9
)

type Board [BOARD_ROW_SIZE][BOARD_COLUMN_SIZE]PieceWithHand

var INIT_BOARD = Board{
  [BOARD_COLUMN_SIZE]PieceWithHand{
    SECOND_KYOU, SECOND_KEI, SECOND_GIN, SECOND_KIN, SECOND_GYOKU, SECOND_KIN, SECOND_GIN, SECOND_KEI, SECOND_KYOU,
  },

  [BOARD_COLUMN_SIZE]PieceWithHand{
    PieceWithHand{}, SECOND_HI, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, SECOND_KAKU, PieceWithHand{},
  },

  [BOARD_COLUMN_SIZE]PieceWithHand{
    SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU,
  },

  [BOARD_COLUMN_SIZE]PieceWithHand{
    PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{},
  },

  [BOARD_COLUMN_SIZE]PieceWithHand{
    PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{},
  },

  [BOARD_COLUMN_SIZE]PieceWithHand{
    PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{},
  },

  [BOARD_COLUMN_SIZE]PieceWithHand{
    FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU,
  },

  [BOARD_COLUMN_SIZE]PieceWithHand{
    PieceWithHand{}, FIRST_HI, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, PieceWithHand{}, FIRST_KAKU, PieceWithHand{},
  },

  [BOARD_COLUMN_SIZE]PieceWithHand{
    FIRST_KYOU, FIRST_KEI, FIRST_GIN, FIRST_KIN, FIRST_OU, FIRST_KIN, FIRST_GIN, FIRST_KEI, FIRST_KYOU,
  },
}

func (board *Board) NewPieceWithHandWithPositions(position *Position, directions Positions) PieceWithHandWithPositions {
  result := make(PieceWithHandWithPositions, 0, BOARD_ROW_SIZE)
  for _, iPositon := range directions {
    positionAfterMove := position.Add(&iPositon)
    if iPositon.IsOutOBoardRange() {
      break
    }
    pieceWithHand := board[positionAfterMove.X][positionAfterMove.Y]
    result = append(result, PieceWithHandWithPosition{PieceWithHand:pieceWithHand, Position:positionAfterMove})
  }
  return result
}

func (board Board) NewKingPosition(hand Hand) Position {
  var pieceName
}

func (board Board) IsCheck(hand Hand) bool {

}

func (board Board) NewLegalMoves(hand Hand) Moves {
  newLegalMoves := func(piece *Piece) Moves {
    movePositions := piece.MovePositions

  }
}
