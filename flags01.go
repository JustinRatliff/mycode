/* Alta3 Research | RZFeeser
   Flags - integer   */

package main

import (
    "flag"      // required for setting flags
    "fmt"
)

func main() {

    // accept an integer flag from the user
    // first parameter is the name of the flag ("n")
    // second parameter is the default value (5)
    // the third is the description of the flag
    num := flag.Int("n", 5, "How angry is the Hulk? (# of iterations)")
    flag.Parse() // process the flags

    // we used a function that returns a pointer to a variable
    // we need to dereference the pointer and get the value
    n := *num    // dereference the pointer and get the value

    // create a loop starting at 0 up to (not including) the value passed
    i := 0
    for i < n {
        fmt.Println("HULK SMASH!")
        i++
    }
}

