package gosyogi

import (
  "math/rand"
)

type Player func(*Aspect, Aspects) Move

func NewRandomPlayer(random *rand.Rand) Player {
  result := func(aspect *Aspect, history Aspects) Move {
    legalMoves := aspect.NewLegalMoves(history)
    return legalMoves.RandomChoice(random)
  }
  return result
}

type Players map[Turn]Player

func (players Players) OneGame(aspect Aspect, history Aspects) (Aspect, error) {
  for {
    player := players[aspect.Turn]
    move := player(&aspect, history)
    aspect, history = aspect.Put(&move, history)
    if aspect.IsGameEnd(history) {
      break
    }
  }
  return aspect, nil
}
