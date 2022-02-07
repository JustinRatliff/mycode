/* RZFeeser | Alta3 Research
   Arrays - pre-defining values   */

package main

import "fmt"

func main() {

    // var x [5]int = [5]int{10, 20, 30, 40, 50}    // Initialized with vales
    x := [5]int{10, 20, 30, 40, 50}                 // Initialized with values - shorthand

    // var y [5]int = [5]int{10, 20, 30}    // Partial assignment
    y := [5]int{10, 20, 30}                 // Partial assignment - shorthand

    fmt.Println(x) // show x
    fmt.Println(y) // show y
}

