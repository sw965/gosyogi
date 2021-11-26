package gosyogi

import (
  "fmt"
  "testing"
  "strconv"
)

func TestAspect(t *testing.T) {
  v, _ := strconv.Atoi("２")
  fmt.Println(v)
  hugous := Hugous{}
  hugou := Hugou("▲２四銀(32)")
  fmt.Println(hugou)
  move, err := hugou.ToMove(hugous)
  if err != nil {
    panic(err)
  }
  fmt.Println(move)

  hugous = append(hugous, hugou)

  hugou = "△同金(11)"
  move, err = hugou.ToMove(hugous)
  if err != nil {
    panic(err)
  }
  fmt.Println(move)
}
