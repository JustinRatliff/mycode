/* Alta3 Research | Author: RZFeeser
   Type transformations - string to integer  */

package main

import (
    "fmt"
    "strconv"    // built in package for string conversions
)

func main() {

    // create a string via inference
    s := "33"
    
    // string to int
    i, err := strconv.Atoi(s)
    if err != nil {
        // handle error
        panic(err)
    }
    fmt.Println("String:", s)
    fmt.Println("Integer:", i)
    
    // transform an integer into a string
    st := strconv.Itoa(42)   /// st == "42"      note the quotes! st is now a string
    fmt.Printf("st is now of type, %T", st)      // place the "type" of st in place of %T
    fmt.Println("String", st)
}

