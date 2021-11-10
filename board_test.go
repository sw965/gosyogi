package gosyogi

import (
  "testing"
)

func TestNewLegalMoves(t *testing.T) {
  aspect := NewAspect()
  aspect, err := aspect.Put(&Move{BeforePosition:Position{Row:6, Column:2}, AfterPosition:Position{Row:5, Column:2}})
  if err != nil {
    panic(err)
  }
  aspect.Board.PrintSimple()

  aspect, err = aspect.Put(&Move{BeforePosition:Positions{}})
}
