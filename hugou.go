package gosyogi

import (
  "fmt"
  "strconv"
)

type Hugou string

func (hugou Hugou) Split() SplittedHugou {
  rune := []rune(hugou)
  result := make([]string, 0, len(rune))
  for _, c := range rune {
    result = append(result, string(c))
  }
  return result
}

func (hugou Hugou) ToMove(hugous Hugous) (Move, error) {
  splittedHugou := hugou.Split()
  hugouPositionAfterMove, err := splittedHugou.NewHugouPositionAfterMove(hugous)
  if err != nil {
    return Move{}, err
  }
  fmt.Println("hugouPositionAfterMove", hugouPositionAfterMove)
  pieceName := splittedHugou.NewPieceName()

  positionAfterMove := hugouPositionAfterMove.ToPosition()
  fmt.Println("positionAfterMove", positionAfterMove)
  positionBeforeMove, err := splittedHugou.NewuPositionBeforeMove()
  fmt.Println("positionBeforeMove", positionBeforeMove)
  isPromotion := splittedHugou.In("成")
  result := Move{PieceName:pieceName, BeforePosition:positionBeforeMove,
    AfterPosition:positionAfterMove, IsPromotion:isPromotion,
  }
  return result, err
}

type Hugous []Hugou

func (hugous Hugous) LastHugou() Hugou {
  index := len(hugous) - 1
  return hugous[index]
}

type SplittedHugou []string

func (splittedHugou SplittedHugou) In(s string) bool {
  for _, ele := range splittedHugou {
    if ele == s {
      return true
    }
  }
  return false
}

func (splittedHugou SplittedHugou) NewTurn() Turn {
  return Turn(splittedHugou[0])
}

func (splittedHugou SplittedHugou) FirstIndex(s string) int {
  for i, ele := range splittedHugou {
    if ele == s {
      return i
    }
  }
  return -1
}

func (splittedHugou SplittedHugou) NewHugouPositionAfterMove(hugous Hugous) (HugouPosition, error) {
  if splittedHugou.In("同") {
    lastHugou := hugous.LastHugou()
    splittedLastHugou := lastHugou.Split()
    return splittedLastHugou.NewHugouPositionAfterMove(hugous[:len(hugous)-1])
  } else {
    suji := splittedHugou[1]
    dan := splittedHugou[2]
    return HugouPosition{Suji:suji, Dan:dan}, nil
  }
}

func (splittedHugou SplittedHugou) NewPieceName() PieceName {
  if splittedHugou.In("同") {
    return PieceName(splittedHugou[2])
  } else {
    return PieceName(splittedHugou[3])
  }
}

func (splittedHugou SplittedHugou) NewuPositionBeforeMove() (Position, error) {
  index := splittedHugou.FirstIndex("(")

  if index == -1 {
    if splittedHugou.In("打") {
      return CAPTURED_PIECE_POSITION, nil
    } else {
      return Position{}, fmt.Errorf("符号の形式が不適")
    }
  }

  suji, err := strconv.Atoi(splittedHugou[index + 1])
  if err != nil {
    return Position{}, err
  }

  dan, err := strconv.Atoi(splittedHugou[index + 2])
  row := dan - 1
  column := 9 - suji
  return Position{Row:row, Column:column}, err
}

type HugouPosition struct {
  Suji string
  Dan string
}

var CAPTURED_PIECE_HUGOU_POSITION = HugouPosition{Suji:"128", Dan:"128"}

func (hugouPosition *HugouPosition) ToPosition() Position {
  column := SUJI_TO_COLUMN[hugouPosition.Suji]
  row := DAN_TO_ROW[hugouPosition.Dan]
  return Position{Row:row, Column:column}
}

var SUJI_TO_COLUMN = map[string]int{
  "１":8, "２":7, "３":6, "４":5, "５":4, "６":3, "７":2, "８":1, "９":0,
}

var DAN_TO_ROW = map[string]int{
  "一":0, "二":1, "三":2, "四":3, "五":4, "六":5, "七":6, "八":7, "九":8,
}
