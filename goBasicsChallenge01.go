/* Alta3 Research | RZFeeser
   Multiple functions   */

package main

import "fmt"

//               (param, param) return
func CityCalc(x int, y int,) int {    // the return type is int
    return x + y 
}

func main() {
	fmt.Println("Enter City: 16")    // we'll learn about inputting data later
	fmt.Println("Enter Zipcode: 45050")
    fmt.Println("Total: ", CityCalc(16, 45050))   // print the results of the function
}

