package main

import (
	"fmt"
	"os"
	"./board"
	"rand"
	"time"
	"flag"
)

func main() {

   var (
      seed int64
   )
   flag.Int64Var(&seed, "seed", time.Nanoseconds(),"Integer seed for the simulation")
   flag.Parse()
   rand.Seed(seed)

	var (
		b *board.Board = board.MakeBoard(12, 15)
	)

	fmt.Printf("Initial State:\n")
	b.Print(os.Stdout)

	b.Randomize()
	fmt.Printf("After Randomization:\n")
	b.Print(os.Stdout)

	b.Step()
	fmt.Printf("Step 1:\n")
	b.Print(os.Stdout)

	b.Step()
	fmt.Printf("Step 2:\n")
	b.Print(os.Stdout)
}
