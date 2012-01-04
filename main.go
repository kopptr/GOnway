package main

import (
	"fmt"
	"os"
	"./board"
)

func main() {
	fmt.Printf("Hello ¤, world!\n")

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
