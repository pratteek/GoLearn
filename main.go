package main

import (
	"fmt"
  "strings"
  "os"
)
func helloSay(){
  fmt.Println("Hello Test!")
}
func hello(names []string)string{
  return fmt.Sprint("Hello ",strings.Join(names,", ")," !")
}

func main() {
  var names = []string {"Prateek","Nidhi","paksboi"}
  fmt.Println(hello(names))
  fmt.Println(hello(os.Args))
  helloSay()
}
