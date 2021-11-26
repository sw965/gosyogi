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

func (board *Board) GetPiece(position *Position) Piece {
  row := position.Row
  column := position.Column
  return board[row][column]
}

func (board1 *Board) Equal(board2 *Board) bool {
  for _, position := range BOARD_ALL_POSITIONS {
    row := position.Row
    column := position.Column
    if board1[row][column] != board2[row][column] {
      return false
    }
  }
  return true
}

func (board *Board) Transpose() Board {
  result := Board{}
  for _, position := range BOARD_ALL_POSITIONS {
    result[position.Row][position.Column] = board[position.Column][position.Row]
  }
  return result
}

func (board *Board) IsNiHu(turn Turn) []bool {
  result := make([]bool, BOARD_COLUMN_SIZE)

  inHu := func(pieces [BOARD_COLUMN_SIZE]Piece) bool {
    for _, piece := range pieces {
      if piece.Name == HU && piece.Turn == turn {
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

func (board *Board) NewTurnPositions(turn Turn) Positions {
  result := make(Positions, 0, INIT_BOARD_PIECE_NUM / 2)
  for _, position := range BOARD_ALL_POSITIONS {
    piece := board[position.Row][position.Column]
    if piece.Turn == turn {
      result = append(result, position)
    }
  }
  return result
}

func (board *Board) NewKingPosition(turn Turn) Position {
  king := TURN_TO_KING[turn]
  for _, position := range BOARD_ALL_POSITIONS {
    if board[position.Row][position.Column].Name == king {
      return position
    }
  }
  return Position{}
}

func (board *Board) NewKingHeadPosition(turn Turn) Position {
  kingPosition := board.NewKingPosition(turn)
  if turn == FIRST {
    return kingPosition.Add(&RELATIVE_UP_POSITION)
  } else {
    relativeUpPosition := RELATIVE_UP_POSITION.ReversePointOfView()
    return kingPosition.Add(&relativeUpPosition)
  }
}

//自らが王を取られに行く手や王手放置の反則は考慮していない
func (board *Board) NewLegalMoves(turn Turn) Moves {
  result := Moves{}
  enemyRegionPositions := BOARD_ENEMY_REGION_POSITIONS[turn]

  for _, tp := range board.NewTurnPositions(turn) {
    piece := board[tp.Row][tp.Column]
    byDirectionRelativeMovePositions := *PIECE_NAME_TO_BY_DIRECTION_RELATIVE_MOVE_POSITIONS[piece.Name]
    if turn == SECOND {
      byDirectionRelativeMovePositions = byDirectionRelativeMovePositions.ReversePointOfView()
    }
    for _, bdrmps := range byDirectionRelativeMovePositions.ToSlice() {
      for _, bdrmp := range bdrmps {
        positionAfterMove := tp.Add(&bdrmp)

        //ボードの範囲外に出たら
        if positionAfterMove.IsOutOBoardRange() {
          break
        }
        movePosPiece := board[positionAfterMove.Row][positionAfterMove.Column]

        //自分の駒にぶつかったら
        if movePosPiece.Turn == turn {
          break
        }

        canPromotion := CAN_PROMOTION[piece.Name] && (enemyRegionPositions.In(tp) || enemyRegionPositions.In(positionAfterMove))
        foulPositions, existsFoulPos := FOUL_POSITIONS[turn][piece.Name]
        noPromotionMove := Move{PieceName:piece.Name, BeforePosition:tp, AfterPosition:positionAfterMove, IsPromotion:false}
        promotionMove := Move{PieceName:piece.Name, BeforePosition:tp, AfterPosition:positionAfterMove, IsPromotion:true}

        //相手の駒にぶつかったら
        if movePosPiece.Turn == REVERSE_TURN[turn] {
          //歩・桂・香が成らないと禁じ手になる場合
          if existsFoulPos && foulPositions.In(positionAfterMove) {
            result = append(result, promotionMove)
            break
          }

          result = append(result, noPromotionMove)
          if canPromotion {
            result = append(result, promotionMove)
          }
          break
        }

        //駒にぶつからなかったら
        //歩・桂・香が成らないと禁じ手になる場合
        if existsFoulPos && foulPositions.In(positionAfterMove) {
          result = append(result, promotionMove)
        } else {
          result = append(result, noPromotionMove)
          if canPromotion {
            result = append(result, promotionMove)
          }
        }
      }
    }
  }
  return result
}

func (board *Board) IsCheck(currentTurn Turn) bool {
  rTurn := REVERSE_TURN[currentTurn]
  for _, rtp := range board.NewTurnPositions(rTurn) {
    piece := board[rtp.Row][rtp.Column]
    byDirectionRelativeMovePositions := *PIECE_NAME_TO_BY_DIRECTION_RELATIVE_MOVE_POSITIONS[piece.Name]
    if rTurn == SECOND {
      byDirectionRelativeMovePositions = byDirectionRelativeMovePositions.ReversePointOfView()
    }
    for _, bdrmps := range byDirectionRelativeMovePositions.ToSlice() {
      for _, bdrmp := range bdrmps {
        positionAfterMove := rtp.Add(&bdrmp)

        //ボードの範囲外に出たら
        if positionAfterMove.IsOutOBoardRange() {
          break
        }
        movePosPiece := board[positionAfterMove.Row][positionAfterMove.Column]

        //相手の駒が相手自身の駒にぶつかったら
        if movePosPiece.Turn == rTurn {
          break
        }

        //相手の駒が自分の駒にぶつかったら
        if movePosPiece.Turn == currentTurn {
          if movePosPiece.Name == TURN_TO_KING[currentTurn] {
            return true
          }
          break
        }
      }
    }
  }
  return false
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

func (boards1 Boards) Equal(boards2 Boards) bool {
  for i, board1 := range boards1 {
    board2 := boards2[i]
    if !board1.Equal(&board2) {
      return false
    }
  }
  return true
}
