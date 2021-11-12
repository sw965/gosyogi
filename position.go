package gosyogi

type Position struct {
  Row int
  Column int
}

var RELATIVE_UP_POSITION = Position{Row:-1, Column:0}
var RELATIVE_RIGHT_UP_POSITION = Position{Row:-1, Column:1}
var RELATIVE_RIGHT_POSITION = Position{Row:0, Column:1}
var RELATIVE_RIGHT_DOWN_POSITION = Position{Row:1, Column:1}
var RELATIVE_DOWN_POSITION = Position{Row:1, Column:0}
var RELATIVE_LEFT_DOWN_POSITION = Position{Row:1, Column:-1}
var RELATIVE_LEFT_POSITION = Position{Row:0, Column:-1}
var RELATIVE_LEFT_UP_POSITION = Position{Row:-1, Column:-1}

func (position *Position) Flip() Position {
  return Position{Row:position.Column, Column:position.Row}
}

func (position *Position) IsOutOBoardRange() bool {
  return position.Row < 0 || position.Row > BOARD_ROW_SIZE - 1 || position.Column < 0 || position.Column > BOARD_COLUMN_SIZE - 1
}

func (position *Position) IsCapturedPiece() bool {
  return position.Row == -1 && position.Column == -1
}

func (position1 *Position) Add(position2 *Position) Position {
  row1 := position1.Row
  column1 := position1.Column
  row2 := position2.Row
  column2 := position2.Column
  return Position{Row:row1 + row2, Column:column1 + column2}
}

func(position *Position) ReverseTurn() Position {
  x := position.Row
  y := position.Column
  return Position{Row:x * -1, Column:y * -1}
}

type Positions []Position

var BOARD_ALL_POSITIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE * BOARD_COLUMN_SIZE)
  for i := 0; i < BOARD_ROW_SIZE; i++ {
    for j := 0; j < BOARD_COLUMN_SIZE; j++ {
      result = append(result, Position{Row:i, Column:j})
    }
  }
  return result
}()

var BOARD_FIRST_REGION_POSITIONS = func() Positions {
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

var BOARD_SECOND_REGION_POSITIONS = func() Positions {
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

var BOARD_ENEMY_REGION_POSITIONS = map[Turn]Positions{
  FIRST:BOARD_SECOND_REGION_POSITIONS, SECOND:BOARD_FIRST_REGION_POSITIONS,
}

var ALL_RELATIVE_UP_POSITIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    row := RELATIVE_UP_POSITION.Row * (i + 1)
    column := RELATIVE_UP_POSITION.Column * (i + 1)
    result = append(result, Position{Row:row, Column:column})
  }
  return result
}()

var ALL_RELATIVE_RIGHT_UP_POSITIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    row := RELATIVE_RIGHT_UP_POSITION.Row * (i + 1)
    column := RELATIVE_RIGHT_UP_POSITION.Column * (i + 1)
    result = append(result, Position{Row:row, Column:column})
  }
  return result
}()

var ALL_RELATIVE_RIGHT_POSITIONS = func() Positions {
  result := make(Positions, 0, BOARD_COLUMN_SIZE)
  for i := 0; i < BOARD_COLUMN_SIZE - 1; i++ {
    row := RELATIVE_RIGHT_POSITION.Row * (i + 1)
    column := RELATIVE_RIGHT_POSITION.Column * (i + 1)
    result = append(result, Position{Row:row, Column:column})
  }
  return result
}()

var ALL_RELATIVE_RIGHT_DOWN_POSITIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    row := RELATIVE_RIGHT_DOWN_POSITION.Row * (i + 1)
    column := RELATIVE_RIGHT_DOWN_POSITION.Column * (i + 1)
    result = append(result, Position{Row:row, Column:column})
  }
  return result
}()

var ALL_RELATIVE_DOWN_POSITIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    row := RELATIVE_DOWN_POSITION.Row * (i + 1)
    column := RELATIVE_DOWN_POSITION.Column * (i + 1)
    result = append(result, Position{Row:row, Column:column})
  }
  return result
}()

var ALL_RELATIVE_LEFT_DOWN_POSITIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    row := RELATIVE_LEFT_DOWN_POSITION.Row * (i + 1)
    column := RELATIVE_LEFT_DOWN_POSITION.Column * (i + 1)
    result = append(result, Position{Row:row, Column:column})
  }
  return result
}()

var ALL_RELATIVE_LEFT_POSITIONS = func() Positions {
  result := make(Positions, 0, BOARD_COLUMN_SIZE)
  for i := 0; i < BOARD_COLUMN_SIZE - 1; i++ {
    row := RELATIVE_LEFT_POSITION.Row * (i + 1)
    column := RELATIVE_LEFT_POSITION.Column * (i + 1)
    result = append(result, Position{Row:row, Column:column})
  }
  return result
}()

var ALL_RELATIVE_LEFT_UP_POSITIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    row := RELATIVE_LEFT_UP_POSITION.Row * (i + 1)
    column := RELATIVE_LEFT_UP_POSITION.Column * (i + 1)
    result = append(result, Position{Row:row, Column:column})
  }
  return result
}()

func (positions Positions) ReverseTurn() Positions {
  result := make(Positions, len(positions))
  for i, position := range positions {
    result[i] = position.ReverseTurn()
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
