package gosyogi

import (
  "math/rand"
)

type Player func(*Aspect) Move

func NewRandomPlayer(random *rand.Rand) Player {
  result := func(aspect *Aspect) Move {
    legalMoves := aspect.NewLegalMoves()
    return legalMoves.RandomChoice(random)
  }
  return result
}

type Players map[Turn]Player

func (players Players) OneGame(aspect Aspect) (Aspect, error) {
  var err error
  for {
    player := players[aspect.Turn]
    move := player(&aspect)
    aspect, err = aspect.Put(&move)
    if err != nil {
      return Aspect{}, err
    }

    _, err := aspect.IsFirstWin()
    if err == nil {
      break
    }
  }
  return aspect, nil
}
