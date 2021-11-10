package gosyogi

import (
  "fmt"
)

type PieceName string

const (
  HU_NAME = PieceName("歩")
  HI_NAME = PieceName("飛")
  KAKU_NAME = PieceName("角")
  KIN_NAME = PieceName("金")
  GIN_NAME = PieceName("銀")
  KEI_NAME = PieceName("桂")
  KYOU_NAME = PieceName("香")

  TO_NAME = PieceName("と")
  RYUU_NAME = PieceName("竜")
  UMA_NAME = PieceName("馬")
  NARI_GIN_NAME = PieceName("全")
  NARI_KEI_NAME = PieceName("圭")
  NARI_KYOU_NAME = PieceName("杏")

  OU_NAME = PieceName("王")
  GYOKU_NAME = PieceName("玉")
)

var PIECE_NAME_TO_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = map[PieceName]*ByDirectionPositions{
  HU_NAME:&HU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  HI_NAME:&HI_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  KAKU_NAME:&KAKU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  KIN_NAME:&KIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  GIN_NAME:&GIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  KEI_NAME:&KEI_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  KYOU_NAME:&KYOU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,

  TO_NAME:&TO_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  RYUU_NAME:&RYUU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  UMA_NAME:&UMA_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  NARI_GIN_NAME:&NARI_GIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  NARI_KEI_NAME:&NARI_KEI_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  NARI_KYOU_NAME:&NARI_KYOU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,

  OU_NAME:&OU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  GYOKU_NAME:&GYOKU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
}

type PieceNames []PieceName

func (pieceNames PieceNames) Remove(pieceName PieceName) (PieceNames, error) {
  result := make(PieceNames, 0, len(pieceNames) - 1)
  ok := false
  for _, iPieceName := range pieceNames {
    if iPieceName == pieceName {
      ok = true
      continue
    }
    result = append(result, iPieceName)
  }

  if ok {
    return result, nil
  } else {
    errMsg := fmt.Sprintf("%vは駒名の中には存在しなかった", pieceName)
    return result, fmt.Errorf(errMsg)
  }
}

type Piece struct {
  Name PieceName
  Normal *Piece
  Promoted *Piece
}

var HU = Piece{Name:HU_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var HI = Piece{Name:HI_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var KAKU = Piece{Name:KAKU_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var KIN = Piece{Name:KIN_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var GIN = Piece{Name:GIN_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var KEI = Piece{Name:KEI_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var KYOU = Piece{Name:KYOU_NAME, Normal:&Piece{}, Promoted:&Piece{}}

var TO = Piece{Name:TO_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var RYUU = Piece{Name:RYUU_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var UMA = Piece{Name:UMA_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var NARI_GIN = Piece{Name:NARI_GIN_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var NARI_KEI = Piece{Name:NARI_KEI_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var NARI_KYOU = Piece{Name:NARI_KYOU_NAME, Normal:&Piece{}, Promoted:&Piece{}}

var OU = Piece{Name:OU_NAME, Normal:&Piece{}, Promoted:&Piece{}}
var GYOKU = Piece{Name:GYOKU_NAME, Normal:&Piece{}, Promoted:&Piece{}}

var PIECE_NAME_TO_PIECE = map[PieceName]*Piece{
  HU_NAME:&HU, HI_NAME:&HI, KAKU_NAME:&KAKU, KIN_NAME:&KIN, GIN_NAME:&GIN, KEI_NAME:&KEI, KYOU_NAME:&KYOU,
  TO_NAME:&TO, RYUU_NAME:&RYUU, UMA_NAME:&UMA, NARI_GIN_NAME:&NARI_GIN, NARI_KEI_NAME:&NARI_KEI, NARI_KYOU_NAME:&NARI_KYOU,
  OU_NAME:&OU, GYOKU_NAME:&GYOKU,
}

type PieceWithTurn struct {
  Piece Piece
  Turn Turn
}

var FIRST_HU = PieceWithTurn{Piece:HU, Turn:FIRST}
var FIRST_HI = PieceWithTurn{Piece:HI, Turn:FIRST}
var FIRST_KAKU = PieceWithTurn{Piece:KAKU, Turn:FIRST}
var FIRST_KIN = PieceWithTurn{Piece:KIN, Turn:FIRST}
var FIRST_GIN = PieceWithTurn{Piece:GIN, Turn:FIRST}
var FIRST_KEI = PieceWithTurn{Piece:KEI, Turn:FIRST}
var FIRST_KYOU = PieceWithTurn{Piece:KYOU, Turn:FIRST}

var FIRST_TO = PieceWithTurn{Piece:TO, Turn:FIRST}
var FIRST_RYUU = PieceWithTurn{Piece:RYUU, Turn:FIRST}
var FIRST_UMA = PieceWithTurn{Piece:UMA, Turn:FIRST}
var FIRST_NARI_GIN = PieceWithTurn{Piece:NARI_GIN, Turn:FIRST}
var FIRST_NARI_KEI = PieceWithTurn{Piece:NARI_KEI, Turn:FIRST}
var FIRST_NARI_KYOU = PieceWithTurn{Piece:NARI_KYOU, Turn:FIRST}
var FIRST_OU = PieceWithTurn{Piece:OU, Turn:FIRST}

var SECOND_HU = FIRST_HI.ReverseTurn()
var SECOND_HI = FIRST_HI.ReverseTurn()
var SECOND_KAKU = FIRST_KAKU.ReverseTurn()
var SECOND_KIN = FIRST_KIN.ReverseTurn()
var SECOND_GIN = FIRST_GIN.ReverseTurn()
var SECOND_KEI = FIRST_KEI.ReverseTurn()
var SECOND_KYOU = FIRST_KYOU.ReverseTurn()

var SECOND_TO = FIRST_TO.ReverseTurn()
var SECOND_RYUU = FIRST_RYUU.ReverseTurn()
var SECOND_UMA = FIRST_UMA.ReverseTurn()
var SECOND_NARI_GIN = FIRST_NARI_GIN.ReverseTurn()
var SECOND_NARI_KEI = FIRST_NARI_KEI.ReverseTurn()
var SECOND_NARI_KYOU = FIRST_NARI_KYOU.ReverseTurn()
var SECOND_GYOKU = PieceWithTurn{Piece:GYOKU, Turn:SECOND}

func (pieceWithTurn *PieceWithTurn) ReverseTurn() PieceWithTurn {
  piece := pieceWithTurn.Piece
  turn := pieceWithTurn.Turn
  turn = REVERSE_TURN[turn]
  return PieceWithTurn{Piece:piece, Turn:turn}
}

func (pieceWithTurn *PieceWithTurn) ToSimple() string {
  var turnMark string
  if pieceWithTurn.Turn == FIRST {
    turnMark = "f"
  } else if pieceWithTurn.Turn == SECOND {
    turnMark = "s"
  } else {
    return " 　"
  }
  return turnMark + string(pieceWithTurn.Piece.Name)
}

type PieceWithTurns []PieceWithTurn
type TurnToPieceNames map[Turn]PieceNames

func (turnToPieceNames TurnToPieceNames) Copy() TurnToPieceNames {
  result := TurnToPieceNames{}
  for turn, pieceNames := range turnToPieceNames {
    result[turn] = pieceNames
  }
  return result
}

func init() {
  HU.Promoted = &TO
  HI.Promoted = &RYUU
  KAKU.Promoted = &UMA
  GIN.Promoted = &NARI_GIN
  KEI.Promoted = &NARI_KEI
  KYOU.Promoted = &NARI_KYOU

  TO.Normal = &HU
  RYUU.Normal = &HI
  UMA.Normal = &KAKU
  NARI_GIN.Normal = &GIN
  NARI_KEI.Normal = &KEI
  NARI_KYOU.Normal = &KYOU
}
