/* Alta3 Research | RZFeeser
   Looping - for loop behaving like a while loop (improved) */

package main

import "fmt"
import "math/rand"

func main() {
    
    drive := 0 // drive is defined at the function level
    fmt.Print("Get a mulligan on any drive under 60 yards.\n")

    // with no init or post statements, we can completely drop the requirements
    // for use of semi colons
    for drive <= 60 {           // "as long as" driver is less than or equal to 60
        drive = rand.Intn(251)
        fmt.Print("SWING!\n")
    }
    fmt.Println("Your longest drive was", drive, "yards")
}

