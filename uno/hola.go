package main

import (
  "fmt"
  "holaworld/stringutil" 
  "holaworld/conversion"
)


func main() {
  p:= "this is the letter p" 
	fmt.Println("Hello world!" + p)
	fmt.Println(&p)
	fmt.Println(stringutil.Name, stringutil.LastName)
  conversion.Convert()
}
