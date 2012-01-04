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
      steps int
      verbose bool
   )

   flag.Int64Var(&seed, "seed", time.Nanoseconds(),"Integer seed for the simulation")
   flag.IntVar(&steps, "steps", 1, "Number of steps to perform in the simulation")
   flag.BoolVar(&verbose, "v", false, "Verbose output")

   flag.Parse()

   rand.Seed(seed)
   fmt.Printf("Simulation seeded with %d.\n", seed)

	var (
		b *board.Board = board.MakeBoard(12, 15)
	)
	b.Randomize()
	fmt.Printf("Initial State:\n")
	b.Print(os.Stdout)

   for i := 1; i <= steps; i++ {
      b.Step()
      if verbose {
         fmt.Printf("Step %d:\n", i)
         b.Print(os.Stdout)
      }
   }
   if !verbose {
      fmt.Printf("Simulation after %d steps:\n", steps)
      b.Print(os.Stdout)
   }
}
