package main
import "fmt"

//narrow variable scope is better
func main(){
//  var t int
  t:=0
  increment:= func() int{
    t++
    return t
  }
  fmt.Println(increment())
  fmt.Println(increment())
}
