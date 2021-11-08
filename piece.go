package gosyogi

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
  NARI_GIN_NAME = PieceName("成銀")
  NARI_KEI_NAME = PieceName("成桂")
  NARI_KYOU_NAME = PieceName("成香")

  OU_NAME = PieceName("王")
  GYOKU_NAME = PieceName("玉")
)

type Piece struct {
  Name PieceName
  MovePositions Positions
  BeforeNari *Piece
  Nari *Piece
}

var HU = Piece{Name:HU_NAME, MovePositions:HU_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var HI = Piece{Name:HI_NAME, MovePositions:HI_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var KAKU = Piece{Name:KAKU_NAME, MovePositions:KAKU_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var KIN = Piece{Name:KIN_NAME, MovePositions:KIN_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var GIN = Piece{Name:GIN_NAME, MovePositions:GIN_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var KEI = Piece{Name:KEI_NAME, MovePositions:KEI_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var KYOU = Piece{Name:KYOU_NAME, MovePositions:KYOU_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}

var TO = Piece{Name:TO, MovePositions:TO_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var RYUU = Piece{Name:RYUU, MovePositions:RYUU_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var UMA = Piece{Name:UMA, MovePositions: UMA_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var NARI_GIN = Piece{Name:NARI_GIN, MovePositions:NARI_GIN_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var NARI_KEI = Piece{Name:NARI_KEI, MovePositions:NARI_KEI_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var NARI_KYOU = Piece{Name:NARI_KYOU, MovePositions:NARI_KYOU_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}

var OU = Piece{Name:OU_NAME, MovePositions:OU_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}
var GYOKU = Piece{Name:GYOKU_NAME, MovePositions:GYOKU_MOVE_POSITIONS, BeforeNari:&Piece{}, Nari:&Piece{}}

type PieceWithHand struct {
  Piece Piece
  Hand Hand
}

var FIRST_HU = PieceWithHand{Piece:HU, Hand:FIRST}
var FIRST_HI = PieceWithHand{Piece:HI, Hand:FIRST}
var FIRST_KAKU = PieceWithHand{Piece:KAKU, Hand:FIRST}
var FIRST_KIN = PieceWithHand{Piece:KIN, Hand:FIRST}
var FIRST_GIN = PieceWithHand{Piece:GIN, Hand:FIRST}
var FIRST_KEI = PieceWithHand{Piece:KEI, Hand:FIRST}
var FIRST_KYOU = PieceWithHand{Piece:KYOU, Hand:FIRST}

var FIRST_TO = PieceWithHand{Piece:TO, Hand:FIRST}
var FIRST_RYUU = PieceWithHand{Piece:RYUU, Hand:FIRST}
var FIRST_UMA = PieceWithHand{Piece:UMA, Hand:FIRST}
var FIRST_NARI_GIN = PieceWithHand{Piece:NARI_GIN, Hand:FIRST}
var FIRST_NARI_KEI = PieceWithHand{Piece:NARI_KEI, Hand:FIRST}
var FIRST_NARI_KYOU = PieceWithHand{Piece:NARI_KYOU, Hand:FIRST}
var FIRST_OU = PieceWithHand{Piece:OU, Hand:FIRST}

var SECOND_HU = FIRST_HI.ReverseHand()
var SECOND_HI = FIRST_HI.ReverseHand()
var SECOND_KAKU = FIRST_KAKU.ReverseHand()
var SECOND_KIN = FIRST_KIN.ReverseHand()
var SECOND_GIN = FIRST_GIN.ReverseHand()
var SECOND_KEI = FIRST_KEI.ReverseHand()
var SECOND_KYOU = FIRST_KYOU.ReverseHand()

var SECOND_TO = FIRST_TO.ReverseHand()
var SECOND_RYUU = FIRST_RYUU.ReverseHand()
var SECOND_UMA = FIRST_UMA.ReverseHand()
var SECOND_NARI_GIN = FIRST_NARI_GIN.ReverseHand()
var SECOND_NARI_KEI = FIRST_NARI_KEI.ReverseHand()
var SECOND_NARI_KYOU = FIRST_NARI_KYOU.ReverseHand()
var SECOND_GYOKU = PieceWithHand{Piece:GYOKU, Hand:SECOND}

func (pieceWithHand *PieceWithHand) ReverseHand() PieceWithHand {
  piece := pieceWithHand.Piece
  piece.MovePositions = piece.MovePositions.ReversePointOfView()
  hand := pieceWithHand.Hand
  hand = REVERSE_HAND[hand]
  return PieceWithHand{Piece:piece, Hand:hand}
}

type PieceWithHands []PieceWithHand

func init() {
  HU.Nari = &TO
  HI.Nari = &RYUU
  KAKU.Nari = &UMA
  GIN.Nari = &NARI_GIN
  KEI.Nari = &NARI_KEI
  KYOU.Nari = &NARI_KYOU

  TO.BeforeNari = &HU
  RYUU.BeforeNari = &HI
  UMA.BeforeNari = &KAKU
  NARI_GIN.BeforeNari = &GIN
  NARI_KEI.BeforeNari = &KEI
  NARI_KYOU.BeforeNari = &KYOU
}

type PieceWithHandWithPosition struct {
  PieceWithHand PieceWithHand
  Position Position
}

type PieceWithHandWithPositions []PieceWithHandWithPosition
