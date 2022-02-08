/* RZFeeser | Alta3 Research
   Create a struct type named Coord to contain
   X and Y coordinates */

package main

import "fmt"

type Coord struct {
    X int
    Y int
}

func main() {
    fmt.Println(Coord{1, 2})
    
    z := Coord{42, 100}
    z.Y = 180
    fmt.Println(z)
    fmt.Println("The X coordinate is:", z.X)
    fmt.Println("The Y Coordinate is:", z.Y)
    
}

