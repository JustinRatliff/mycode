/* Alta3 Research | RZFeeser
   Errors - Creating basic errors with errors.New*/

package main

import (
    "errors"
    "fmt"
)

func main() {

    // this is the error we are creating
    err := errors.New("we don't have the power")
    fmt.Println("Scotty says:", err)
}

