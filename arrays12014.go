/* Alta3 Research - Multidimensional Arrays */

package main

import (
    "fmt"
)

func main() {

    // var variable_name [SIZE1][SIZE2][SIZEN] variable_type
    z := [3][4]int{
        {0, 1, 2, 3},
        {4, 5, 6, 7},
        {8, 9, 10, 11},
    }

    val := z[1][1]
    fmt.Println(val) // returns 5

    val = z[2][3]
    fmt.Println(val) // returns 11
}
