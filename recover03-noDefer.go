/* Alta3 Research | RZFeeser
   Refer - panic() and Defer (only)  */


package main

import "fmt"

func main() {
    f()
    fmt.Println("Returned normally from f.")      // we will never see this message
}

func f() {
    // removed the defer function

    fmt.Println("Calling g.")
    g(0)
    
    // we will never reach this point
    fmt.Println("Returned normally from g.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panicking!")
        panic(fmt.Sprintf("%v", i))       // this time, will be displayed
    }
    defer fmt.Println("Defer in g", i)    // defer statements will still be processed
    fmt.Println("Printing in g", i)
    g(i + 1)
}

