package gosyogi

type Position struct {
  Row int
  Column int
}

var UP_DIRECTION = Position{Row:0, Column:-1}
var LEFT_UP_DIRECTION = Position{Row:-1, Column:-1}
var LEFT_DIRECTION = Position{Row:-1, Column:0}
var LEFT_DOWN_DIRECTION = Position{Row:-1, Column:1}
var DOWN_DIRECTION = Position{Row:0, Column:1}
var RIGHT_DOWN_DIRECTION = Position{Row:1, Column:1}
var RIGHT_DIRECTION = Position{Row:1, Column:0}
var RIGHT_UP_DIRECTION = Position{Row:1, Column:-1}

func (position *Position) IsOutOBoardRange() bool {
  return position.Row < 0 || position.Row > BOARD_ROW_SIZE - 1 || position.Column < 0 || position.Column > BOARD_COLUMN_SIZE - 1
}

func (position1 *Position) Add(position2 *Position) Position {
  x1 := position1.Row
  x2 := position2.Row
  y1 := position2.Column
  y2 := position2.Column
  return Position{Row:x1 + x2, Column:y1 + y2}
}

func(position *Position) ReversePointOfView() Position {
  x := position.Row
  y := position.Column
  return Position{Row:x * -1, Column:y * -1}
}


type Positions []Position

var ALL_UP_DIRECTIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    result = append(result, Position{Row:0, Column:-1 * (i + 1)})
  }
  return result
}()

var All_LEFT_UP_DIRECTIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    result = append(result, Position{Row:-1 * (i + 1), Column:-1 * (i + 1)})
  }
  return result
}()

var ALL_LEFT_DIRECTIONS = func() Positions {
  result := make(Positions, 0, BOARD_COLUMN_SIZE)
  for i := 0; i < BOARD_COLUMN_SIZE - 1; i++ {
    result = append(result, Position{Row:-1 * (i + 1), Column:0})
  }
  return result
}()

var ALL_LEFT_DOWN_DIRECTIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    result = append(result, Position{Row:-1 * (i + 1), Column:i + 1})
  }
  return result
}()

var ALL_DOWN_DIRECTIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    result = append(result, Position{Row:0, Column:i + 1})
  }
  return result
}()

var ALL_RIGHT_DOWN_DIRECTIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    result = append(result, Position{Row:i + 1, Column:i + 1})
  }
  return result
}()

var ALL_RIGHT_DIRECTIONS = func() Positions {
  result := make(Positions, 0, BOARD_COLUMN_SIZE)
  for i := 0; i < BOARD_COLUMN_SIZE - 1; i++ {
    result = append(result, Position{Row:i + 1, Column:0})
  }
  return result
}()

var ALL_RIGHT_UP_DIRECTIONS = func() Positions {
  result := make(Positions, 0, BOARD_ROW_SIZE)
  for i := 0; i < BOARD_ROW_SIZE - 1; i++ {
    result = append(result, Position{Row:i +1, Column:-1 * (i + 1)})
  }
  return result
}()

var HU_MOVE_POSITIONS = Positions{
  UP_DIRECTION,
}

var HI_MOVE_POSITIONS = func() Positions {
  result := ALL_UP_DIRECTIONS
  result = result.AddAll(ALL_LEFT_DIRECTIONS)
  result = result.AddAll(ALL_DOWN_DIRECTIONS)
  result = result.AddAll(ALL_RIGHT_DIRECTIONS)
  return result
}()

var KAKU_MOVE_POSITIONS = func() Positions {
  result := All_LEFT_UP_DIRECTIONS
  result = result.AddAll(ALL_LEFT_DOWN_DIRECTIONS)
  result = result.AddAll(ALL_RIGHT_DOWN_DIRECTIONS)
  result = result.AddAll(ALL_RIGHT_UP_DIRECTIONS)
  return result
}()

var KIN_MOVE_POSITIONS = Positions{
  UP_DIRECTION,
  LEFT_UP_DIRECTION,
  LEFT_DIRECTION,
  DOWN_DIRECTION,
  RIGHT_DIRECTION,
  RIGHT_UP_DIRECTION,
}

var GIN_MOVE_POSITIONS = Positions{
  UP_DIRECTION,
  LEFT_UP_DIRECTION,
  LEFT_DOWN_DIRECTION,
  RIGHT_DOWN_DIRECTION,
  RIGHT_UP_DIRECTION,
}

var KEI_MOVE_POSITIONS = Positions{
  Position{Row:-1, Column:-2},
  Position{Row:1, Column:-2},
}

var KColumnOU_MOVE_POSITIONS = ALL_UP_DIRECTIONS
var TO_MOVE_POSITIONS = KIN_MOVE_POSITIONS

var RColumnUU_MOVE_POSITIONS = func() Positions {
  result := Positions{
    LEFT_UP_DIRECTION,
    LEFT_DOWN_DIRECTION,
    RIGHT_DOWN_DIRECTION,
    RIGHT_UP_DIRECTION,
  }

  return result.AddAll(HI_MOVE_POSITIONS)
}()

var UMA_MOVE_POSITIONS = func() Positions {
  result := Positions{
    UP_DIRECTION,
    LEFT_DIRECTION,
    DOWN_DIRECTION,
    RIGHT_DIRECTION,
  }
  return result.AddAll(KAKU_MOVE_POSITIONS)
}()

var NARI_GIN_MOVE_POSITIONS = KIN_MOVE_POSITIONS
var NARI_KEI_MOVE_POSITIONS = KIN_MOVE_POSITIONS
var NARI_KColumnOU_MOVE_POSITIONS = KIN_MOVE_POSITIONS

var OU_MOVE_POSITIONS = Positions{
  UP_DIRECTION,
  LEFT_UP_DIRECTION,
  LEFT_DIRECTION,
  LEFT_DOWN_DIRECTION,
  DOWN_DIRECTION,
  RIGHT_DOWN_DIRECTION,
  RIGHT_DIRECTION,
  RIGHT_UP_DIRECTION,
}

var GColumnOKU_MOVE_POSITIONS = OU_MOVE_POSITIONS

func (positions1 Positions) Concatenation(positions2 Positions) Positions {
  result := make(Positions, 0, len(positions1) + len(positions2))
  for _, position := range positions1 {
    result = append(result, position)
  }

  for _, position := range positions2 {
    result = append(result, position)
  }
  return result
}

func (positions Positions) ReversePointOfView() Positions {
  result := make(Positions, len(positions))
  for i, position := range positions {
    result[i] = position.ReversePointOfView()
  }
  return result
}

type HugouPosition struct {
  Row int
  Column int
}
