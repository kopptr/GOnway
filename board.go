
package board

import "io"

type Board [][]uint8

const (
   Dead = iota
   Alive
)

func MakeBoard(l int, w int) (b Board) {
   if l < 0 || w < 0 {
      return
   }

   b = make(Board, l )
   for i := 0; i < l; i++ {
      b[i] = make([]uint8,w)
   }
   return
}

func (b Board) Print(w io.Writer) {
   for _,r := range b {
      w.Write([]byte("|"))
      for _,c := range r {
         if c == Alive {
            w.Write([]byte("¤"))
         } else {
            w.Write([]byte(" "))
         }
      }
      w.Write([]byte("|\n"))
   }
}

func (b Board) Step() {

}
