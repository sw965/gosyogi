package gosyogi

import (
  "testing"
  "fmt"
)

func Helper(aspect Aspect, move *Move) Aspect {
  legalMoves := aspect.NewLegalMoves()
  if !aspect.NewLegalMoves().In(*move) {
    for _, move := range legalMoves {
      fmt.Println(move)
    }

    panic(fmt.Errorf("n"))
  }

  var err error
  aspect, err = aspect.Put(move)
  if err != nil {
    panic(err)
  }
  aspect.Board.PrintSimple()
  fmt.Println(aspect.Turn)
  fmt.Println("先手", aspect.EachTurnCapturedPieceNames[FIRST])
  fmt.Println("後手", aspect.EachTurnCapturedPieceNames[SECOND])
  fmt.Println("isCheckBoard1", aspect.Board.IsCheck(aspect.Turn))
  fmt.Println("isCheckBoard2", aspect.Board.IsCheck(REVERSE_TURN[aspect.Turn]))
  fmt.Println("")
  return aspect
}

func TestNewLegalMoves(t *testing.T) {
  aspect := NewAspect()
  move := Move{BeforePosition:Position{Row:6, Column:2}, AfterPosition:Position{Row:5, Column:2}, PieceName:HU}
  aspect = Helper(aspect, &move)

  move = Move{BeforePosition:Position{Row:2, Column:6}, AfterPosition:Position{Row:3, Column:6}, PieceName:HU}
  aspect = Helper(aspect, &move)

  //move = Move{BeforePosition:Position{Row:, Column:}, AfterPosition:Position{Row:, Column:}, PieceName:}
  move = Move{BeforePosition:Position{Row:7, Column:1}, AfterPosition:Position{Row:1, Column:7}, PieceName:KAKU, IsPromotion:true}
  aspect = Helper(aspect, &move)

  move = Move{BeforePosition:Position{Row:0, Column:6}, AfterPosition:Position{Row:1, Column:7}, PieceName:GIN}
  aspect = Helper(aspect, &move)

  move = Move{BeforePosition:CAPTURED_PIECE_POSITION, AfterPosition:Position{Row:4, Column:8}, PieceName:KAKU}
  aspect = Helper(aspect, &move)

  move = Move{BeforePosition:Position{Row:0, Column:0}, AfterPosition:Position{Row:1, Column:0}, PieceName:KYOU}
  aspect = Helper(aspect, &move)
}
