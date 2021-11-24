package gosyogi

import (
  "testing"
  "fmt"
  "time"
  "math/rand"
  "github.com/seehuhn/mt19937"
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

  mtRandom := rand.New(mt19937.New())
  mtRandom.Seed(time.Now().UnixNano())

  for i:= 0; i < 64; i++ {
    aspect := NewAspect()
    players := Players{FIRST:NewRandomPlayer(mtRandom), SECOND:NewRandomPlayer(mtRandom)}
    gameEndAspect, err := players.OneGame(aspect)
    if err != nil {
      panic(err)
    }
    gameEndAspect.Board.PrintSimple()
    fmt.Println(gameEndAspect.Turn)
    fmt.Println(len(gameEndAspect.History))
    fmt.Println(gameEndAspect.CapturedPieces[FIRST])
    fmt.Println(gameEndAspect.CapturedPieces[SECOND])
  }
  //board := gameEndAspect.Board.Transpose()
  //board.PrintSimple()
}
