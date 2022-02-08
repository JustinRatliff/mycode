/* RZFeeser | Alta3 Research
   simple if statements */

package main

import "fmt"
import "math/rand"
import "time"

func main() {

    // ensure we produce random values
    rand.Seed(time.Now().UnixNano())


    // simple if statement    
    if 10%5 == 0 {                   // if 10 mod 5 is 0
        fmt.Println("10 is divisible by 5")            
    }
    
    
    // if can state with a short statement to execute
    if v := rand.Intn(101); v > 60 {
        fmt.Println("You tested in the top 60%! Your score was", v)
    }
    fmt.Println("We want to thank everyone for taking the test.")


    // Println function is used to display output in the next line
    fmt.Println("Enter your desired name: ")
  
    // var then variable name then variable type
    var first string
  
    // Taking input from user
    fmt.Scanln(&first)

    // Determine the length of the var first
    if length := len(first); length > 11 {
        fmt.Println("Screen names have a limit of 18 characters \n only the first 11 will be displayed")
    }

}

