package gosyogi

import (
  "math/rand"
)

type Eval struct {
  Func func(*Aspect, Aspects) (float64, error)
  ReverseFunc func(float64) float64
}

func NewRandomPlayoutEval(random *rand.Rand) Eval {
  resultFunc := func(aspect *Aspect, history Aspects) (float64, error) {
    aspectV := *aspect
    randomPlayers := Players{FIRST:NewRandomPlayer(random), SECOND:NewRandomPlayer(random),}
    gameEndAspect, err := randomPlayers.OneGame(aspectV, history)
    if err != nil {
      return 0.0, err
    }
    winner, err := gameEndAspect.Winner(history)
    reward := WINNER_TO_FLOAT64[winner]
    return reward, err
  }

  reverseFunc := func(reward float64) float64 {
    return 1.0 - reward
  }

  return Eval{Func:resultFunc, ReverseFunc:reverseFunc}
}
