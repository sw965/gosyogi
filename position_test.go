package gosyogi

import (
  "testing"
  "fmt"
)

func TestAllPositionsPrint(t *testing.T) {
  fmt.Println("ALL_UP_POSITIONS = ", ALL_RELATIVE_UP_POSITIONS)
  fmt.Println("All_RIGHT_UP_POSITIONS = ", ALL_RELATIVE_RIGHT_UP_POSITIONS)
  fmt.Println("ALL_RIGHT_POSITIONS = ", ALL_RELATIVE_RIGHT_POSITIONS)
  fmt.Println("ALL_RIGHT_DOWN_POSITIONS = ", ALL_RELATIVE_RIGHT_DOWN_POSITIONS)
  fmt.Println("ALL_DOWN_POSITIONS = ", ALL_RELATIVE_DOWN_POSITIONS)
  fmt.Println("ALL_LEFT_DOWN_POSITIONS = ", ALL_RELATIVE_LEFT_DOWN_POSITIONS)
  fmt.Println("ALL_LEFT_POSITIONS = ", ALL_RELATIVE_LEFT_POSITIONS)
  fmt.Println("ALL_LEFT_UP_POSITIONS = ", ALL_RELATIVE_LEFT_UP_POSITIONS)

  foulPositions := FOUL_POSITIONS[FIRST]
  for pieceName, position := range foulPositions {
    fmt.Println(pieceName, position)
  }

  foulPositions = FOUL_POSITIONS[SECOND]
  for pieceName, position := range foulPositions {
    fmt.Println(pieceName, position)
  }
}
