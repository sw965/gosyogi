package gosyogi

type Move struct {
  PieceName PieceName
  BeforePosition Position
  AfterPosition Position
  IsPromotion bool
}

type Moves []Move

func (moves1 Moves) Add(moves2 Moves) Moves {
  result := make(Moves, 0, len(moves1) + len(moves2))
  for _, move := range moves1 {
    result = append(result, move)
  }

  for _, move := range moves2 {
    result = append(result, move)
  }
  return result
}
