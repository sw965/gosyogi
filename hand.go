package gosyogi

type Hand string

const(
  FIRST = Hand("▲")
  SECOND = Hand("△")
)

var REVERSE_HAND = map[Hand]Hand{
  FIRST:SECOND, SECOND:FIRST,
}

var HAND_TO_KING_NAME = map[Hand]PieceName{
  FIRST:OU_NAME, SECOND:GYOKU_NAME,
}
