/* RZFeeser | Alta3 Research
   CHALLENGE 01 - match on any minor version of Go  */

package main
  
import (
    "fmt"
    "runtime"
    "strings"
)

func main() {

    // init gove
    var gove string = runtime.Version()  // this returns the version of Go
    
    switch {
    case strings.Contains(gove, "go1.17"):      // if matches "go1.17", do the code below then STOP
        fmt.Println("You are using the latest version of GoLang")
    case strings.Contains(gove, "go1.16"), strings.Contains(gove, "go1.15"):       // if matches "go1.16" OR "go1.15"
        fmt.Println("You are using an older, but acceptable version of GoLang")
    default:                       // in all other cases run the code below
        fmt.Println("Upgrade GoLang before you continue")
    }
}

