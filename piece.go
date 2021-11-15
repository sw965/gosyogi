package gosyogi

import (
  "fmt"
)

type PieceName string

const (
  HU = PieceName("歩")
  HI = PieceName("飛")
  KAKU = PieceName("角")
  KIN = PieceName("金")
  GIN = PieceName("銀")
  KEI = PieceName("桂")
  KYOU = PieceName("香")

  TO = PieceName("と")
  RYUU = PieceName("竜")
  UMA = PieceName("馬")
  NARI_GIN = PieceName("全")
  NARI_KEI = PieceName("圭")
  NARI_KYOU = PieceName("杏")

  OU = PieceName("王")
  GYOKU = PieceName("玉")
)

var CAN_PROMOTION = map[PieceName]bool{
  HU:true, HI:true, KAKU:true,
  KIN:false, GIN:true, KEI:true, KYOU:true,

  TO:false, RYUU:false, UMA:false,
  NARI_GIN:false, NARI_KEI:false, NARI_KYOU:false,
  OU:false, GYOKU:false,
}

var (
  PIECE_NAME_TO_NORMAL_PIECE_NAMES = map[PieceName]PieceName{
    HU:HU, HI:HI, KAKU:KAKU,
    KIN:KIN, GIN:GIN, KEI:KEI, KYOU:KYOU,

    TO:HU, RYUU:HI, UMA:KAKU,
    NARI_GIN:GIN, NARI_KEI:KEI, NARI_KYOU:KYOU,
    OU:OU, GYOKU:GYOKU,
  }

  PIECE_NAME_TO_PROMOTION_PIECE_NAMES = map[PieceName]PieceName{
    HU:TO, HI:RYUU, KAKU:UMA,
    KIN:KIN, GIN:NARI_GIN, KEI:NARI_KEI, KYOU:NARI_KYOU,

    TO:TO, RYUU:RYUU, UMA:UMA,
    NARI_GIN:NARI_GIN, NARI_KEI:NARI_KEI, NARI_KYOU:NARI_KYOU,
    OU:OU, GYOKU:GYOKU,
  }
)

var PIECE_NAME_TO_BY_DIRECTION_RELATIVE_MOVE_POSITIONS = map[PieceName]*ByDirectionPositions{
  HU:&HU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  HI:&HI_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  KAKU:&KAKU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  KIN:&KIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  GIN:&GIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  KEI:&KEI_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  KYOU:&KYOU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,

  TO:&TO_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  RYUU:&RYUU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  UMA:&UMA_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  NARI_GIN:&NARI_GIN_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  NARI_KEI:&NARI_KEI_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  NARI_KYOU:&NARI_KYOU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,

  OU:&OU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
  GYOKU:&GYOKU_BY_DIRECTION_RELATIVE_MOVE_POSITIONS,
}

type PieceNames []PieceName

var PIECE_NAMES_OF_MOVE_FORWARD_ONLY = PieceNames{HU, KEI, KYOU}

func (pieceNames PieceNames) Remove(pieceName PieceName) (PieceNames, error) {
  result := make(PieceNames, 0, len(pieceNames) - 1)
  ok := false
  for _, iPieceName := range pieceNames {
    if iPieceName == pieceName && !ok {
      ok = true
      continue
    }
    result = append(result, iPieceName)
  }

  if ok {
    return result, nil
  } else {
    errMsg := fmt.Sprintf("%vはPieceNamesの中には存在しなかった", pieceName)
    return result, fmt.Errorf(errMsg)
  }
}

type Piece struct {
  Name PieceName
  Turn Turn
}

var (
  FIRST_HU = Piece{Name:HU, Turn:FIRST}
  FIRST_HI = Piece{Name:HI, Turn:FIRST}
  FIRST_KAKU = Piece{Name:KAKU, Turn:FIRST}
  FIRST_KIN = Piece{Name:KIN, Turn:FIRST}
  FIRST_GIN = Piece{Name:GIN, Turn:FIRST}
  FIRST_KEI = Piece{Name:KEI, Turn:FIRST}
  FIRST_KYOU = Piece{Name:KYOU, Turn:FIRST}

  FIRST_TO = Piece{Name:TO, Turn:FIRST}
  FIRST_RYUU = Piece{Name:RYUU, Turn:FIRST}
  FIRST_UMA = Piece{Name:UMA, Turn:FIRST}
  FIRST_NARI_GIN = Piece{Name:NARI_GIN, Turn:FIRST}
  FIRST_NARI_KEI = Piece{Name:NARI_KEI, Turn:FIRST}
  FIRST_NARI_KYOU = Piece{Name:NARI_KYOU, Turn:FIRST}
  FIRST_OU = Piece{Name:OU, Turn:FIRST}

  SECOND_HU = FIRST_HU.ReverseTurn()
  SECOND_HI = FIRST_HI.ReverseTurn()
  SECOND_KAKU = FIRST_KAKU.ReverseTurn()
  SECOND_KIN = FIRST_KIN.ReverseTurn()
  SECOND_GIN = FIRST_GIN.ReverseTurn()
  SECOND_KEI = FIRST_KEI.ReverseTurn()
  SECOND_KYOU = FIRST_KYOU.ReverseTurn()

  SECOND_TO = FIRST_TO.ReverseTurn()
  SECOND_RYUU = FIRST_RYUU.ReverseTurn()
  SECOND_UMA = FIRST_UMA.ReverseTurn()
  SECOND_NARI_GIN = FIRST_NARI_GIN.ReverseTurn()
  SECOND_NARI_KEI = FIRST_NARI_KEI.ReverseTurn()
  SECOND_NARI_KYOU = FIRST_NARI_KYOU.ReverseTurn()
  SECOND_GYOKU = Piece{Name:GYOKU, Turn:SECOND}
)

func (piece *Piece) ReverseTurn() Piece {
  return Piece{Name:piece.Name, Turn:REVERSE_TURN[piece.Turn]}
}

func (piece *Piece) ToSimple() string {
  var turnMark string
  if piece.Turn == FIRST {
    turnMark = "f"
  } else if piece.Turn == SECOND {
    turnMark = "s"
  } else {
    return " 　"
  }
  return turnMark + string(piece.Name)
}

type EachTurnCapturedPieceNames map[Turn]PieceNames

func (eachTurnCapturedPieceNames EachTurnCapturedPieceNames) Copy() EachTurnCapturedPieceNames {
  result := EachTurnCapturedPieceNames{}
  for turn, pieceNames := range eachTurnCapturedPieceNames {
    result[turn] = pieceNames
  }
  return result
}
