package main

import (
  "fmt"
//  "holaworld/stringutil" 
 // "holaworld/conversion"
)

var x = 0

func main() {
//  p:= "this is the letter p" 
//	fmt.Println("Hello world!" + p)
//	fmt.Println(&p)
//	fmt.Println(stringutil.Name, stringutil.LastName)
//	fmt.Println(stringutil.Testy())
//  conversion.Convert()
fmt.Println(increment())
fmt.Println(x)

fmt.Println(printXplusOne())
}

func increment() int {
  x++
  return x
}

func printXplusOne() int {
  return x + 1
}
