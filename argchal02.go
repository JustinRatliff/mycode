/* Alta3 Research | RZFeeser
   CHALLENGE 02 - Iterate across arguments passed in via the CLI */
   
package main
 
import (
    "fmt"
    "os"
)
 
func main() {

    // the first argument i.e. program name is excluded via [1:]
    argLength := len(os.Args[1:])                   // determine the length
    fmt.Printf("Arg length is %d\n", argLength)     // display number of args
 
    // this is called a "range" loop
    // i is an integer starting at 0
    // a is the string (argument in this case) corresponding to the integer
    for i, a := range os.Args[1:] {                 // loop across args
        fmt.Printf("Arg %d is %s\n", i+1, a)        // use a format string
    }                                               // %d == int and %s = string
}

