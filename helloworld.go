package main

import (
  "fmt"
  "os"
)

func main(){
  args := os.Args

  fmt.Println("The arguments are : ", args)
  fmt.Println("Hello ", args[1])
}
