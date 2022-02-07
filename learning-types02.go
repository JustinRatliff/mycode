/* Alta3 Research | Author: RZFeeser
  Types - transformations   */

package main

import (
    "fmt"
    "math"
)

func main() {
    
    // x and y are integers
    var x, y int = 3, 4
    
                            // transform x and y
    var f float64 = math.Sqrt(float64(x*x + y*y))
    
    // tranform into unsigned integer
    var z uint = uint(f)
    
    // display f
    fmt.Println(f)
    
    // display x, y and z
    fmt.Println(x, y, z)
}

