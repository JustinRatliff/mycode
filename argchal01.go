/* Alta3 Research | RZFeeser
   CHALLENGE 01 - Display the number of arguments passed to the program */
   
package main
 
import (
    "fmt"
    "os"
)
 
func main() {

    // use the len() function to return length
    argLength := len(os.Args[1:])
    
    // use fmt.Printf() to format string
    fmt.Printf("Arg length is %d", argLength) 
}

