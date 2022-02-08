/* RZFeeser | Alta3 Research
   switch - using fallthrough */

package main

import (
    "fmt"
)

// our use of fallthrough creates the following pattern
// A -> B -> C -> D -> E

func main() {

    // Batcave
    nextstop := "Batcave"

    fmt.Println("Stops ahead of us:")

    switch nextstop {

    case "Superman":
        fmt.Println("Superman's Fortress of Solitude")
        fallthrough

    case "Batcave":
        fmt.Println("Batcave")
        fallthrough

    case "Strange":
        fmt.Println("Dr. Strange's Sanctum Sanctorum")
        fallthrough

    case "Justice":
        fmt.Println("Justice League's Hall of Justice")
        fallthrough

    case "Xavier":
        fmt.Println("Xavier's School for Gifted Youngsters")
    }
}

