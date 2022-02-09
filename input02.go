/* Alta3 Research | RZFeeser
   Reading multiple inputs in a single line with Scanf   */

package main

import "fmt"

func main() {

    var name string                 // define a string variable (for tracking user name)
    var age int                     // define an int var (for tracking user age)

    fmt.Print("Enter your name & age: ")            // "fmt.Print" is not followed by a /n
    fmt.Scanf("%s %d", &name, &age)                 // "fmt.Scanf" function reads a string and an integer to the two provided variables
    fmt.Printf("%s is %d years old\n", name, age)   // "fmt.Printf" allows us to "format" text
                                                    //    push our captured data to standard out using "formatters"
}

