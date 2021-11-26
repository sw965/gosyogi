package gosyogi

type Position struct {
  Row int
  Column int
}

var (
  RELATIVE_UP_POSITION = Position{Row:-1, Column:0}
  RELATIVE_RIGHT_UP_POSITION = Position{Row:-1, Column:1}
  RELATIVE_RIGHT_POSITION = Position{Row:0, Column:1}
  RELATIVE_RIGHT_DOWN_POSITION = Position{Row:1, Column:1}
  RELATIVE_DOWN_POSITION = Position{Row:1, Column:0}
  RELATIVE_LEFT_DOWN_POSITION = Position{Row:1, Column:-1}
  RELATIVE_LEFT_POSITION = Position{Row:0, Column:-1}
  RELATIVE_LEFT_UP_POSITION = Position{Row:-1, Column:-1}
  CAPTURED_PIECE_POSITION = Position{Row:-128, Column:-128}
)

func (position *Position) Flip() Position {
  return Position{Row:position.Column, Column:position.Row}
}

func (position *Position) IsOutOBoardRange() bool {
  return position.Row < 0 || position.Row > BOARD_ROW_SIZE - 1 || position.Column < 0 || position.Column > BOARD_COLUMN_SIZE - 1
}

func (position Position) IsCapturedPiece() bool {
  return position == CAPTURED_PIECE_POSITION
}

func (position1 *Position) Add(position2 *Position) Position {
  row1 := position1.Row
  column1 := position1.Column
  row2 := position2.Row
  column2 := position2.Column
  return Position{Row:row1 + row2, Column:column1 + column2}
}

//先手番から見た駒の動きを後手番から見た駒の動きにする為の関数
func(position *Position) ReversePointOfView() Position {
  x := position.Row
  y := position.Column
  return Position{Row:x * -1, Column:y * -1}
}

type Positions []Position

func NewBeforeMovePositions(pieceName PieceName, positionAfterMove *Position, turn Turn) Positions {
  result := make(Positions, 0)

  byDirectionRelativeMovePositions := PIECE_NAME_TO_BY_DIRECTION_RELATIVE_MOVE_POSITIONS[pieceName]

  if turn == SECOND {
    tmp := byDirectionRelativeMovePositions.ReversePointOfView()
    byDirectionRelativeMovePositions = &tmp
  }

  for _, bdrmps := range byDirectionRelativeMovePositions.ToSlice() {
    for _, bdrmp := range bdrmps {
      rPos := bdrmp.ReversePointOfView()
      positionBeforeMove := positionAfterMove.Add(&rPos)
      if !positionBeforeMove.IsOutOBoardRange() {
        result = append(result, positionBeforeMove)
      }
    }
  }
  result = append(result, CAPTURED_PIECE_POSITION)
  return result
}

var BOARD_ALL_POSITIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE * BOARD_COLUMN_SIZE)
  for i := 0; i < BOARD_ROW_SIZE; i++ {
    for j := 0; j < BOARD_COLUMN_SIZE; j++ {
      result = append(result, Position{Row:i, Column:j})
    }
  }
  return result
}()

var (
  BOARD_FIRST_REGION_POSITIONS = func() Positions {
    regionRowNum := 3
    result := make(Positions, regionRowNum * BOARD_COLUMN_SIZE)
    for i := 0; i < regionRowNum; i++ {
      for j := 0; j < BOARD_COLUMN_SIZE; j++ {
        row := i + (BOARD_ROW_SIZE - regionRowNum)
        position := Position{Row:row, Column:j}
        result = append(result, position)
      }
    }
    return result
    }()


  BOARD_SECOND_REGION_POSITIONS = func() Positions {
    regionRowNum := 3
    result := make(Positions, regionRowNum * BOARD_COLUMN_SIZE)
    for i := 0; i < regionRowNum; i++ {
      for j := 0; j < BOARD_COLUMN_SIZE; j++ {
        position := Position{Row:i, Column:j}
        result = append(result, position)
      }
    }
    return result
  }()

  BOARD_ENEMY_REGION_POSITIONS = map[Turn]Positions{
    FIRST:BOARD_SECOND_REGION_POSITIONS, SECOND:BOARD_FIRST_REGION_POSITIONS,
  }
)

var (
  FIRST_HU_FOUL_POSITIONS = func() Positions {
    result := make(Positions, BOARD_COLUMN_SIZE)
    for i := 0; i < BOARD_COLUMN_SIZE; i++ {
      result[i] = Position{Row:0, Column:i}
    }
    return result
    }()

  SECOND_HU_FOUL_POSITIONS = func() Positions {
    result := make(Positions, BOARD_COLUMN_SIZE)
    for i := 0; i < BOARD_COLUMN_SIZE; i++ {
      result[i] = Position{Row:BOARD_ROW_SIZE - 1, Column:i}
    }
    return result
  }()

  FIRST_KYOU_FOUL_POSITIONS = func() Positions {
    result := make(Positions, BOARD_COLUMN_SIZE)
    for i := 0; i < BOARD_COLUMN_SIZE; i++ {
      result[i] = Position{Row:0, Column:i}
    }
    return result
  }()

  SECOND_KYOU_FOUL_POSITIONS = func() Positions {
    result := make(Positions, BOARD_COLUMN_SIZE)
    for i := 0; i < BOARD_COLUMN_SIZE; i++ {
      result[i] = Position{Row:BOARD_ROW_SIZE - 1, Column:i}
    }
    return result
  }()

  FIRST_KEI_FOUL_POSITIONS = func() Positions {
    result := make(Positions, 0, 2 * BOARD_COLUMN_SIZE)
    for i := 0; i < 2; i++ {
      for j := 0; j < BOARD_COLUMN_SIZE; j++ {
        position := Position{Row:i, Column:j}
        result = append(result, position)
      }
    }
    return result
  }()

  SECOND_KEI_FOUL_POSITIONS = func() Positions {
    result := make(Positions, 0, 2 * BOARD_COLUMN_SIZE)
    for i := 0; i < 2; i++ {
      for j := 0; j < BOARD_COLUMN_SIZE; j++ {
        position := Position{Row:BOARD_ROW_SIZE - (i + 1), Column:j}
        result = append(result, position)
      }
    }
    return result
  }()

  FOUL_POSITIONS = map[Turn]map[PieceName]Positions{
    FIRST:map[PieceName]Positions{
      HU:FIRST_HU_FOUL_POSITIONS, KYOU:FIRST_KYOU_FOUL_POSITIONS, KEI:FIRST_KEI_FOUL_POSITIONS,
    },

    SECOND:map[PieceName]Positions{
      HU:SECOND_HU_FOUL_POSITIONS, KYOU:SECOND_KYOU_FOUL_POSITIONS, KEI:SECOND_KEI_FOUL_POSITIONS,
    },
  }
)

var (
  ALL_RELATIVE_UP_POSITIONS = func() Positions {
    result := make(Positions, 0, BOARD_ROW_SIZE)
    for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
      row := RELATIVE_UP_POSITION.Row * (i + 1)
      column := RELATIVE_UP_POSITION.Column * (i + 1)
      result = append(result, Position{Row:row, Column:column})
    }
    return result
  }()

  ALL_RELATIVE_RIGHT_UP_POSITIONS = func() Positions {
    result := make(Positions, 0, BOARD_ROW_SIZE)
    for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
      row := RELATIVE_RIGHT_UP_POSITION.Row * (i + 1)
      column := RELATIVE_RIGHT_UP_POSITION.Column * (i + 1)
      result = append(result, Position{Row:row, Column:column})
    }
    return result
  }()

  ALL_RELATIVE_RIGHT_POSITIONS = func() Positions {
    result := make(Positions, 0, BOARD_COLUMN_SIZE)
    for i := 0; i < BOARD_COLUMN_SIZE - 1; i++ {
      row := RELATIVE_RIGHT_POSITION.Row * (i + 1)
      column := RELATIVE_RIGHT_POSITION.Column * (i + 1)
      result = append(result, Position{Row:row, Column:column})
    }
    return result
  }()

  ALL_RELATIVE_RIGHT_DOWN_POSITIONS = func() Positions {
    result := make(Positions, 0, BOARD_ROW_SIZE)
    for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
      row := RELATIVE_RIGHT_DOWN_POSITION.Row * (i + 1)
      column := RELATIVE_RIGHT_DOWN_POSITION.Column * (i + 1)
      result = append(result, Position{Row:row, Column:column})
    }
    return result
  }()

  ALL_RELATIVE_DOWN_POSITIONS = func() Positions {
    result := make(Positions, 0, BOARD_ROW_SIZE)
    for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
      row := RELATIVE_DOWN_POSITION.Row * (i + 1)
      column := RELATIVE_DOWN_POSITION.Column * (i + 1)
      result = append(result, Position{Row:row, Column:column})
    }
    return result
  }()

  ALL_RELATIVE_LEFT_DOWN_POSITIONS = func() Positions {
    result := make(Positions, 0, BOARD_ROW_SIZE)
    for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
      row := RELATIVE_LEFT_DOWN_POSITION.Row * (i + 1)
      column := RELATIVE_LEFT_DOWN_POSITION.Column * (i + 1)
      result = append(result, Position{Row:row, Column:column})
    }
    return result
  }()

  ALL_RELATIVE_LEFT_POSITIONS = func() Positions {
    result := make(Positions, 0, BOARD_COLUMN_SIZE)
    for i := 0; i < BOARD_COLUMN_SIZE - 1; i++ {
      row := RELATIVE_LEFT_POSITION.Row * (i + 1)
      column := RELATIVE_LEFT_POSITION.Column * (i + 1)
      result = append(result, Position{Row:row, Column:column})
    }
    return result
  }()

  ALL_RELATIVE_LEFT_UP_POSITIONS = func() Positions {
    result := make(Positions, 0, BOARD_ROW_SIZE)
    for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
      row := RELATIVE_LEFT_UP_POSITION.Row * (i + 1)
      column := RELATIVE_LEFT_UP_POSITION.Column * (i + 1)
      result = append(result, Position{Row:row, Column:column})
    }
    return result
  }()
)

func (positions Positions) ReversePointOfView() Positions {
  result := make(Positions, len(positions))
  for i, position := range positions {
    result[i] = position.ReversePointOfView()
  }
  return result
}

func (positions Positions) In(position Position) bool {
  for _, iPosition := range positions {
    if iPosition == position {
      return true
    }
  }
  return false
}
