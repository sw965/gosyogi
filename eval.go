package gosyogi

type Eval struct {
  Func func(*Aspect) (float64, error)
  ReverseFunc func(float64) float64
}

func NewRandomPlayoutEval() Eval {
  return Eval{}
}
