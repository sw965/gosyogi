package gosyogi

type Turn string

const(
  FIRST = Turn("▲")
  SECOND = Turn("△")
)

var REVERSE_TURN = map[Turn]Turn{
  FIRST:SECOND, SECOND:FIRST,
}

var TURN_TO_KING = map[Turn]PieceName{
  FIRST:OU, SECOND:GYOKU,
}
