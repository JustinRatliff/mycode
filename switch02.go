/* RZFeeser | Alta3 Research
   switch - Exploring 'switch true' with a default */
   
package main

import (
    "fmt"
    "time"
)

func main() {
    
    // determine what time it is
    watch := time.Now() 
    
    // there is no condition here
    // read as, "switch true"
    switch {
    case watch.Hour() < 6:
        fmt.Println("Go back to sleep.")
    case watch.Hour() < 12:
        fmt.Println("Good morning!")
    case watch.Hour() < 18:
        fmt.Println("Good afternoon.")
    default:
        fmt.Println("Counting sheep. Z-z-z-z-z-z-z")
    }
}

