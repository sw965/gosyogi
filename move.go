package gosyogi

type Move struct {
  PieceName PieceName
  BeforePosition Position
  AfterPosition Position
}

type Moves []Move
