/* Alta3 Research | RZFeeser
   Intro to panic
   
package main

import (
    "fmt"
)

func main() {

    // panic produces a quick exit
    panic("Jim, we have a problem.")

    fmt.Println("You will not even see this line. The panic creates a fast fail.")
}

