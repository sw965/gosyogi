package gosyogi

import (
  "testing"
  "fmt"
)

func TestNewLegalMoves(t *testing.T) {
  aspect := NewAspect()
  move := Move{BeforePosition:Position{Row:6, Column:2}, AfterPosition:Position{Row:5, Column:2}}
  aspect, err := aspect.Put(&move)
  if err != nil {
    panic(err)
  }
  aspect.Board.PrintSimple()

  move = Move{BeforePosition:Position{Row:2, Column:6}, AfterPosition:Position{Row:3, Column:6}}
  aspect, err = aspect.Put(&move)
  if err != nil {
    panic(err)
  }
  aspect.Board.PrintSimple()

  move = Move{BeforePosition:Position{Row:7, Column:1}, AfterPosition:Position{Row:1, Column:7}, IsPromotion:true}
  aspect, err = aspect.Put(&move)
  if err != nil {
    panic(err)
  }
  aspect.Board.PrintSimple()
  fmt.Println(aspect.EachTurnCapturedPieceNames[FIRST])
}
