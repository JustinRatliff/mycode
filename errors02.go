/* Alta3 Research | RZFeeser
   Errors - Creating basic errors with fmt.Errorf */

package main

import (
    "fmt"
    "time"
)

func main() {

    // build a error message with the current time
    err := fmt.Errorf("error occurred at: %v", time.Now())
    fmt.Println("An error happened:", err)
}

