package gosyogi

type Turn string

const(
  FIRST = Turn("▲")
  SECOND = Turn("△")
)

var REVERSE_TURN = map[Turn]Turn{
  FIRST:SECOND, SECOND:FIRST,
}

var TURN_TO_KING_NAME = map[Turn]PieceName{
  FIRST:OU_NAME, SECOND:GYOKU_NAME,
}
