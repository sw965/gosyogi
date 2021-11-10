package gosyogi

import (
  "fmt"
)

const (
  BOARD_ROW_SIZE = 9
  BOARD_COLUMN_SIZE = 9
)

type Board [BOARD_ROW_SIZE][BOARD_COLUMN_SIZE]PieceWithTurn

var INIT_BOARD = Board{
  [BOARD_COLUMN_SIZE]PieceWithTurn{
    SECOND_KYOU, SECOND_KEI, SECOND_GIN, SECOND_KIN, SECOND_GYOKU, SECOND_KIN, SECOND_GIN, SECOND_KEI, SECOND_KYOU,
  },

  [BOARD_COLUMN_SIZE]PieceWithTurn{
    PieceWithTurn{}, SECOND_HI, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, SECOND_KAKU, PieceWithTurn{},
  },

  [BOARD_COLUMN_SIZE]PieceWithTurn{
    SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU,
  },

  [BOARD_COLUMN_SIZE]PieceWithTurn{
    PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{},
  },

  [BOARD_COLUMN_SIZE]PieceWithTurn{
    PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{},
  },

  [BOARD_COLUMN_SIZE]PieceWithTurn{
    PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{},
  },

  [BOARD_COLUMN_SIZE]PieceWithTurn{
    FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU,
  },

  [BOARD_COLUMN_SIZE]PieceWithTurn{
    PieceWithTurn{}, FIRST_KAKU, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, PieceWithTurn{}, FIRST_HI, PieceWithTurn{},
  },

  [BOARD_COLUMN_SIZE]PieceWithTurn{
    FIRST_KYOU, FIRST_KEI, FIRST_GIN, FIRST_KIN, FIRST_OU, FIRST_KIN, FIRST_GIN, FIRST_KEI, FIRST_KYOU,
  },
}

var INIT_BOARD_PIECE_NUM = func() int {
  result := 0
  empty := PieceWithTurn{}
  for _, position := range BOARD_ALL_POSITIONS {
    pieceWithTurn := INIT_BOARD[position.Row][position.Column]
    if pieceWithTurn != empty {
      result += 1
    }
  }
  return result
}()

func (board *Board) NewSelfTurnPositions(turn Turn) Positions {
  result := make(Positions, 0, INIT_BOARD_PIECE_NUM / 2)
  for _, position := range BOARD_ALL_POSITIONS {
    pieceWithTurn := board[position.Row][position.Column]
    if pieceWithTurn.Turn == turn {
      result = append(result, position)
    }
  }
  return result
}

func (board Board) NewLegalMoves(turn Turn) Moves {
  result := Moves{}
  selfTurnPositions := board.NewSelfTurnPositions(turn)

  for _, sp := range selfTurnPositions {
    pieceWithTurn := board[sp.Row][sp.Column]
    pieceName := pieceWithTurn.Piece.Name
    byDirectionRelativeMovePositions := *PIECE_NAME_TO_BY_DIRECTION_RELATIVE_MOVE_POSITIONS[pieceWithTurn.Piece.Name]
    if turn == SECOND {
      byDirectionRelativeMovePositions = byDirectionRelativeMovePositions.ReverseTurn()
    }
    for _, bdrmps := range byDirectionRelativeMovePositions.ToSlice() {
      for _, bdrmp := range bdrmps {
        positionAfterMove := sp.Add(&bdrmp)
        if positionAfterMove.IsOutOBoardRange() {
          break
        }

        if board[positionAfterMove.Row][positionAfterMove.Column].Turn == turn {
          break
        }

        if board[positionAfterMove.Row][positionAfterMove.Column].Turn == REVERSE_TURN[turn] {
          result = append(result, Move{PieceName:pieceName, BeforePosition:sp, AfterPosition:positionAfterMove})
          break
        }

        result = append(result, Move{PieceName:pieceName, BeforePosition:sp, AfterPosition:positionAfterMove})
      }
    }
  }
  return result
}

func (board *Board) ToSimple() [][]string {
  result := make([][]string, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_COLUMN_SIZE; i++ {
    result[i] = make([]string, BOARD_COLUMN_SIZE)
  }

  for i := 0; i < BOARD_ROW_SIZE; i++ {
    for j := 0; j < BOARD_COLUMN_SIZE; j++ {
      result[i][j] = board[i][j].ToSimple()
    }
  }
  return result
}

func (board *Board) PrintSimple() {
  for _, ele := range board.ToSimple() {
    fmt.Println(ele)
  }
}


type Boards []Board
