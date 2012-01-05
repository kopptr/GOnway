package main

import (
	"fmt"
	"./conway"
	"time"
	"flag"
	"os"
   "exec"
)

func main() {

   var (
      seed int64
      steps int
   )

   flag.Int64Var(&seed, "seed", time.Nanoseconds(),"Integer seed for the simulation")
   flag.IntVar(&steps, "steps", 1, "Number of steps to perform in the simulation")

   flag.Parse()

   fmt.Printf("Simulation seeded with %d.\n", seed)

	var b *conway.ConwaySim = conway.MakeConwaySim(12, 15)
	b.Randomize(seed)
   fmt.Printf("Simulation seeded with %d.\n", seed)
	fmt.Printf("Step %d:\n", 0)
	fmt.Print(b)

   for i := 1; i <= steps; i++ {
      b.Step()
      clearScreen()
      fmt.Printf("Simulation seeded with %d.\n", seed)
      fmt.Printf("Step %d:\n", i)
      fmt.Print(b)
      time.Sleep(1000000000)
   }
   fmt.Printf("Simulation complete.\n")
}

func clearScreen() {
   c := exec.Command("clear")
   c.Stdout = os.Stdout
   c.Run()
}
