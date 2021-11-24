package gosyogi

import (
  "fmt"
  "math/rand"
  "github.com/sw965/crow"
)

type Node struct {
  Aspect *Aspect
  History Aspects
  LegalMoves Moves
  UCB1s UCB1s
  NextNodes Nodes
  IsAllExpansion bool
  IsFirstNode bool
  SelectCount int
}

func NewNodePointer(aspect *Aspect, history Aspects) *Node {
  legalMoves := aspect.NewLegalMoves(history)
  return &Node{Aspect:aspect, History:history, LegalMoves:legalMoves, IsFirstNode:aspect.Turn == FIRST}
}

func (node *Node) NewNoExpansionMoves() Moves {
  result := make(Moves, 0, len(node.LegalMoves) - len(node.UCB1s))
  for _, move := range node.LegalMoves {
    if _, ok := node.UCB1s[move]; !ok {
      result = append(result, move)
    }
  }
  return result
}

func (node *Node) Select(aspect Aspect, history Aspects, X float64, random *rand.Rand) (*Node, Aspect, Aspects, Selects, error) {
  selects := Selects{}

  for {
    if node.IsAllExpansion {
      break
    }

    maxUCBMoves, err := node.UCB1s.NewMaxMoves(X)
    if err != nil {
      return &Node{}, Aspect{}, Aspects{}, Selects{}, err
    }

    selectMove := maxUCBMoves.RandomChoice(random)
    selects = append(selects, Select{Node:node, Move:selectMove})

    aspect, history = aspect.Put(&selectMove, history)

    if aspect.IsGameEnd(history) {
      break
    }
    node = node.NextNodes[selectMove]
  }
  return node, aspect, history, selects, nil
}

func (node *Node) ExpansionWithEvalY(aspect Aspect, history Aspects, eval *Eval, selects Selects, random *rand.Rand) (float64, Selects, error) {
  if node.IsAllExpansion {
    return 0.0, Selects{}, fmt.Errorf("展開済みのNodeである")
  }

  //未展開の合法手をランダムで展開する
  noExpansionMove := node.NewNoExpansionMoves().RandomChoice(random)
  aspect, _ = aspect.Put(&noExpansionMove, history)
  node.UCB1s[noExpansionMove] = &crow.UpperConfidenceBound1{}
  selects = append(selects, Select{Node:node, Move:noExpansionMove})

  if len(node.UCB1s) == len(node.LegalMoves) {
    node.IsAllExpansion = true
  }

  evalY, err := eval.Func(&aspect)
  return evalY, selects, err
}

type UCB1s map[Move]*crow.UpperConfidenceBound1

func (ucb1s UCB1s) NewKeyMoves() Moves {
  result := make(Moves, 0, len(ucb1s))
  for move, _ := range ucb1s {
    result = append(result, move)
  }
  return result
}

func (ucb1s UCB1s) TotalTrial() int {
  result := 0
  for _, ucb1 := range ucb1s {
    result += ucb1.Trial
  }
  return result
}

func (ucb1s UCB1s) Max(X float64) (float64, error) {
  totalTrial := ucb1s.TotalTrial()
  moves := ucb1s.NewKeyMoves()
  result, err := ucb1s[moves[0]].Get(totalTrial, X)
  if err != nil {
    return 0.0, err
  }

  for _, move := range moves[1:] {
    ucb1v, err := ucb1s[move].Get(totalTrial, X)
    if err != nil {
      return 0.0, err
    }

    if ucb1v > result {
      result = ucb1v
    }
  }
  return result, nil
}

func (ucb1s UCB1s) NewMaxMoves(X float64) (Moves, error) {
  max, err := ucb1s.Max(X)
  if err != nil {
    return Moves{}, err
  }
  totalTrial := ucb1s.TotalTrial()
  result := make(Moves, 0)

  for _, move := range ucb1s.NewKeyMoves() {
    ucb1v, err := ucb1s[move].Get(totalTrial, X)
    if err != nil {
      return Moves{}, err
    }

    if ucb1v == max {
      result = append(result, move)
    }
  }
  return result, nil
}

type Nodes map[Move]*Node

type Select struct {
  Node *Node
  Move Move
}

type Selects []Select

func (selects Selects) Backward(evalY float64, eval *Eval) {
  for _, select_ := range selects {
    node := select_.Node
    move := select_.Move

    if node.IsFirstNode {
      node.UCB1s[move].AccumReward += evalY
    } else {
      node.UCB1s[move].AccumReward += eval.ReverseFunc(evalY)
    }
    node.UCB1s[move].Trial += 1
  }
}

func RunMCTS(rootAspect Aspect, rootHistory Aspects, simuNum int, X float64, eval *Eval, random *rand.Rand) (*Node, error) {
  rootNode := NewNodePointer(&rootAspect, rootHistory)
  node := rootNode
  aspect := rootAspect
  history := rootHistory
  var evalY float64
  var err error

  for i := 0; i < simuNum; i++ {
    node, aspect, history, selects, err := node.Select(aspect, history, X, random)
    if err != nil {
      return &Node{}, err
    }

    if aspect.IsGameEnd(history) {
      evalY, err = eval.Func(&aspect)
    } else {
      evalY, selects, err = node.ExpansionWithEvalY(aspect, history, eval, selects, random)
    }

    if err != nil {
      return &Node{}, err
    }

    selects.Backward(evalY, eval)
    node = rootNode
    aspect = rootAspect
    history = rootHistory
  }
  return rootNode, nil
}
