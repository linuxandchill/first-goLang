package main
import "fmt"

func wrappy() func() int {
  t:=0
  return func() int{
    t++
    return t
  }
}

func main(){
  increment:= wrappy()
  fmt.Println(increment())
  fmt.Println(increment())
}
