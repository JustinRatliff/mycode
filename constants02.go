/* Alta3 Research | Author: RZFeeser
   Constants - Introduction to numerical constants   */

package main

import "fmt"

// this is a group of constants being declared at once
const (
    // Create a huge number by shifting a 1 bit left 100 places.
    // In other words, the binary number that is 1 followed by 100 zeroes.
    Big     = 1 << 100 // go will determine this is an int
    Country = "USA"
)

func oven(i string) {
    return
}

func main() {

    // display our global constants
    // fmt.Println(Big)  // fails when golang tries to express the integer
    // fmt.Println(Big * .0001) // works, we shrink the size of the expression back down
    fmt.Println(Country)

    // Note: constants may also be lower case
    // they do not have to be uppercase
    const Small = 3
    fmt.Println(Small)

    // if uncommented, this will fail
    // Small = 10  // reassignment is NOT allowed

    // if uncommented, this will fail
    // the value of constants needs to be known at compile time
    // therefore values returned by function calls are illegal
    // const cupcake = oven("vanilla")
}

