package gosyogi

import (
  "fmt"
)

const (
  BOARD_ROW_SIZE = 9
  BOARD_COLUMN_SIZE = 9
)

type Board [BOARD_ROW_SIZE][BOARD_COLUMN_SIZE]Piece

var INIT_BOARD = Board{
  [BOARD_COLUMN_SIZE]Piece{
    SECOND_KYOU, SECOND_KEI, SECOND_GIN, SECOND_KIN, SECOND_GYOKU, SECOND_KIN, SECOND_GIN, SECOND_KEI, SECOND_KYOU,
  },

  [BOARD_COLUMN_SIZE]Piece{
    Piece{}, SECOND_HI, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, SECOND_KAKU, Piece{},
  },

  [BOARD_COLUMN_SIZE]Piece{
    SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU, SECOND_HU,
  },

  [BOARD_COLUMN_SIZE]Piece{
    Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{},
  },

  [BOARD_COLUMN_SIZE]Piece{
    Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{},
  },

  [BOARD_COLUMN_SIZE]Piece{
    Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, Piece{},
  },

  [BOARD_COLUMN_SIZE]Piece{
    FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU, FIRST_HU,
  },

  [BOARD_COLUMN_SIZE]Piece{
    Piece{}, FIRST_KAKU, Piece{}, Piece{}, Piece{}, Piece{}, Piece{}, FIRST_HI, Piece{},
  },

  [BOARD_COLUMN_SIZE]Piece{
    FIRST_KYOU, FIRST_KEI, FIRST_GIN, FIRST_KIN, FIRST_OU, FIRST_KIN, FIRST_GIN, FIRST_KEI, FIRST_KYOU,
  },
}

var INIT_BOARD_PIECE_NUM = func() int {
  result := 0
  empty := Piece{}
  for _, position := range BOARD_ALL_POSITIONS {
    piece := INIT_BOARD[position.Row][position.Column]
    if piece != empty {
      result += 1
    }
  }
  return result
}()

func (board *Board) Transpose() Board {
  result := Board{}
  for _, position := range BOARD_ALL_POSITIONS {
    p := position.Flip()
    result[p.Row][p.Column] = board[position.Row][position.Column]
  }
  return result
}

func (board *Board) InHu() []bool {
  result := make([]bool, BOARD_COLUMN_SIZE)

  inHu := func(pieces [BOARD_COLUMN_SIZE]Piece) bool {
    for _, piece := range pieces {
      if piece.Name == HU {
        return true
      }
    }
    return false
  }

  for i, pieces := range board.Transpose() {
    result[i] = inHu(pieces)
  }
  return result
}

func (board *Board) NewSelfTurnPositions(turn Turn) Positions {
  result := make(Positions, 0, INIT_BOARD_PIECE_NUM / 2)
  for _, position := range BOARD_ALL_POSITIONS {
    piece := board[position.Row][position.Column]
    if piece.Turn == turn {
      result = append(result, position)
    }
  }
  return result
}

func (board Board) NewLegalMoves(turn Turn) Moves {
  result := Moves{}
  selfTurnPositions := board.NewSelfTurnPositions(turn)
  enemyRegionPositions := BOARD_ENEMY_REGION_POSITIONS[turn]

  for _, sp := range selfTurnPositions {
    piece := board[sp.Row][sp.Column]
    byDirectionRelativeMovePositions := *PIECE_NAME_TO_BY_DIRECTION_RELATIVE_MOVE_POSITIONS[piece.Name]
    if turn == SECOND {
      byDirectionRelativeMovePositions = byDirectionRelativeMovePositions.ReverseTurn()
    }
    for _, bdrmps := range byDirectionRelativeMovePositions.ToSlice() {
      for _, bdrmp := range bdrmps {
        positionAfterMove := sp.Add(&bdrmp)
        if positionAfterMove.IsOutOBoardRange() {
          break
        }

        //自分の駒にぶつかったら
        if board[positionAfterMove.Row][positionAfterMove.Column].Turn == turn {
          break
        }

        canPromotion := CAN_PROMOTION[piece.Name] && (enemyRegionPositions.In(sp) || enemyRegionPositions.In(positionAfterMove))

        //相手の駒にぶつかったら
        if board[positionAfterMove.Row][positionAfterMove.Column].Turn == REVERSE_TURN[turn] {
          result = append(result, Move{PieceName:piece.Name, BeforePosition:sp, AfterPosition:positionAfterMove, IsPromotion:false})
          if canPromotion {
            result = append(result, Move{PieceName:piece.Name, BeforePosition:sp, AfterPosition:positionAfterMove, IsPromotion:true})
          }
          break
        }

        //駒にぶつからなかったら
        result = append(result, Move{PieceName:piece.Name, BeforePosition:sp, AfterPosition:positionAfterMove, IsPromotion:false})
        if canPromotion {
          result = append(result, Move{PieceName:piece.Name, BeforePosition:sp, AfterPosition:positionAfterMove, IsPromotion:true})
        }
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
