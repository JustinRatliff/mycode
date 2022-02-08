/* Alta3 Research | RZFeeser
   Map - associative data types   */

package main

import "fmt"

func main() {

    // create an empty map using "make"
    // make(map[key-type]val-type)
    m := make(map[string]int)             // m is a map

    m["k1"] = 7                 // name[key] = val
    m["k2"] = 13                // name[key] = val

    // this will show all of the key/value pairs
    fmt.Println("map:", m)

    // return the value associated with "key", k1
    v1 := m["k1"]
    fmt.Println("v1: ", v1)

    // determine number of key/value associations
    fmt.Println("len:", len(m))

    // get rid of the "key", k2
    delete(m, "k2")
    
    // no more "key", k2
    fmt.Println("map:", m)

    // save only the optional second return value
    // this second value can be used to determine if a key exists
    // of if it is missing from the map
    _, prs := m["k2"]
    fmt.Println("prs:", prs)

    // declare and initialize a map all one line
    n := map[string]int{"foo": 1, "bar": 2}
    fmt.Println("map:", n)
}

