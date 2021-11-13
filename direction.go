package gosyogi

type ByDirectionPositions struct {
  Up Positions
  RightUp Positions
  Right Positions
  RightDown Positions
  Down Positions
  LeftDown Positions
  Left Positions
  LeftUp Positions
}

func (byDirectionPositions ByDirectionPositions) ReverseTurn() ByDirectionPositions {
  byDirectionPositions.Up = byDirectionPositions.Up.ReverseTurn()
  byDirectionPositions.RightUp = byDirectionPositions.RightUp.ReverseTurn()
  byDirectionPositions.Right = byDirectionPositions.Right.ReverseTurn()
  byDirectionPositions.RightDown = byDirectionPositions.RightDown.ReverseTurn()
  byDirectionPositions.Down = byDirectionPositions.Down.ReverseTurn()
  byDirectionPositions.LeftDown = byDirectionPositions.LeftDown.ReverseTurn()
  byDirectionPositions.Left = byDirectionPositions.Left.ReverseTurn()
  byDirectionPositions.LeftUp = byDirectionPositions.LeftUp.ReverseTurn()
  return byDirectionPositions
}

func (byDirectionPositions *ByDirectionPositions) ToSlice() []Positions {
  return []Positions{
    byDirectionPositions.Up,
    byDirectionPositions.RightUp, byDirectionPositions.Right, byDirectionPositions.RightDown,
    byDirectionPositions.Down,
    byDirectionPositions.LeftDown, byDirectionPositions.Left, byDirectionPositions.LeftUp,
  }
}

var (
  HU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = ByDirectionPositions{
    Up:Positions{RELATIVE_UP_POSITION},
    RightUp:Positions{}, Right:Positions{}, RightDown:Positions{},
    Down:Positions{},
    LeftDown:Positions{}, Left:Positions{}, LeftUp:Positions{},
  }

  HI_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = ByDirectionPositions{
    Up:ALL_RELATIVE_UP_POSITIONS,
    RightUp:Positions{}, Right:ALL_RELATIVE_RIGHT_POSITIONS, RightDown:Positions{},
    Down:ALL_RELATIVE_DOWN_POSITIONS,
    LeftDown:Positions{}, Left:ALL_RELATIVE_LEFT_POSITIONS, LeftUp:Positions{},
  }

  KAKU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = ByDirectionPositions{
    Up:Positions{},
    RightUp:ALL_RELATIVE_RIGHT_UP_POSITIONS, Right:Positions{}, RightDown:ALL_RELATIVE_RIGHT_DOWN_POSITIONS,
    Down:Positions{},
    LeftDown:ALL_RELATIVE_LEFT_DOWN_POSITIONS, Left:Positions{}, LeftUp:ALL_RELATIVE_LEFT_UP_POSITIONS,
  }

  KIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = ByDirectionPositions{
    Up:Positions{RELATIVE_UP_POSITION},
    RightUp:Positions{RELATIVE_RIGHT_UP_POSITION}, Right:Positions{RELATIVE_RIGHT_POSITION}, RightDown:Positions{},
    Down:Positions{RELATIVE_DOWN_POSITION},
    LeftDown:Positions{}, Left:Positions{RELATIVE_LEFT_POSITION}, LeftUp:Positions{RELATIVE_LEFT_UP_POSITION},
  }

  GIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = ByDirectionPositions{
    Up:Positions{RELATIVE_UP_POSITION},
    RightUp:Positions{RELATIVE_RIGHT_UP_POSITION}, Right:Positions{}, RightDown:Positions{RELATIVE_RIGHT_DOWN_POSITION},
    Down:Positions{},
    LeftDown:Positions{RELATIVE_LEFT_DOWN_POSITION}, Left:Positions{}, LeftUp:Positions{RELATIVE_LEFT_UP_POSITION},
  }

  KEI_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = ByDirectionPositions{
    Up:Positions{},
    RightUp:Positions{Position{Row:-2, Column:1}}, Right:Positions{}, RightDown:Positions{},
    Down:Positions{},
    LeftDown:Positions{}, Left:Positions{}, LeftUp:Positions{Position{Row:-2, Column:-1}},
  }

  KYOU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = ByDirectionPositions{
    Up:ALL_RELATIVE_UP_POSITIONS,
    RightUp:Positions{}, Right:Positions{}, RightDown:Positions{},
    Down:Positions{},
    LeftDown:Positions{}, Left:Positions{}, LeftUp:Positions{},
  }

  TO_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = KIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS
  RYUU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = ByDirectionPositions{
    Up:ALL_RELATIVE_UP_POSITIONS,
    RightUp:Positions{RELATIVE_RIGHT_UP_POSITION}, Right:ALL_RELATIVE_RIGHT_POSITIONS, RightDown:Positions{RELATIVE_RIGHT_DOWN_POSITION},
    Down:ALL_RELATIVE_DOWN_POSITIONS,
    LeftDown:Positions{RELATIVE_LEFT_DOWN_POSITION}, Left:ALL_RELATIVE_LEFT_POSITIONS, LeftUp:Positions{RELATIVE_LEFT_UP_POSITION},
  }

  UMA_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = ByDirectionPositions{
    Up:Positions{RELATIVE_UP_POSITION},
    RightUp:ALL_RELATIVE_RIGHT_UP_POSITIONS, Right:Positions{RELATIVE_RIGHT_POSITION}, RightDown:ALL_RELATIVE_RIGHT_DOWN_POSITIONS,
    Down:Positions{RELATIVE_DOWN_POSITION},
    LeftDown:ALL_RELATIVE_LEFT_DOWN_POSITIONS, Left:Positions{RELATIVE_LEFT_POSITION}, LeftUp:ALL_RELATIVE_LEFT_UP_POSITIONS,
  }

  NARI_GIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = KIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS
  NARI_KEI_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = KIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS
  NARI_KYOU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = KIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS

  OU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = ByDirectionPositions{
    Up:Positions{RELATIVE_UP_POSITION},
    RightUp:Positions{RELATIVE_RIGHT_UP_POSITION}, Right:Positions{RELATIVE_RIGHT_POSITION}, RightDown:Positions{RELATIVE_RIGHT_DOWN_POSITION},
    Down:Positions{RELATIVE_DOWN_POSITION},
    LeftDown:Positions{RELATIVE_LEFT_DOWN_POSITION}, Left:Positions{RELATIVE_LEFT_POSITION}, LeftUp:Positions{RELATIVE_LEFT_UP_POSITION},
  }

  GYOKU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = OU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS
)
