/* RZFeeser | Alta3 Research
   Arrays - basics              */
   
package main

import "fmt"

func main() {

    // Create an array that holds 5 integers
    var scores [5]int
    fmt.Println("emp:", scores)

    // set a value at a particular index
    scores[4] = 100
    fmt.Println("set:", scores)
    fmt.Println("get:", scores[4])

    // the builtin len() will return the length
    fmt.Println("len:", len(scores))


    // declare and initialize in a single line
    highScores := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", highScores)

}

