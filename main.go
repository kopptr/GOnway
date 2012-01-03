package main

import (
   "fmt"
   "os"
   "./board"
)

func main() {
   fmt.Printf("Hello ¤, world!\n")

   var (
      b board.Board
   )

   b = board.MakeBoard(5,5)

   b[1][4] = board.Alive

   b.Print(os.Stdout)
}


