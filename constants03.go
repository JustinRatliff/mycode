/* Alta3 Research | RZFeeser
   Const - untyped vs typed constants    */
package main

import (
    "fmt"
)

func main() {

    var myFloat32 float32          = 4.5          // 4.5 is a "floating point untyped constant" and can be used to define float32
    var myComplex64 complex64      = 4.5          // 4.5 is a "floating point untyped constant" and can be used to define complex64
    fmt.Println(myFloat32)
    fmt.Println(myComplex64)


    type AltaString string                        // create an alias for the type string

    // This will NOT work because of Go's "Highly Typed" framework
    /*
    var myString string            = "Hello"
    var zString AltaString         = myString     // this will not work, we cannot MIX types
    */

    // This WILL work, "myUntypedString" is a untyped variable
    const myUntypedString          = "Alta3 Research"
    var uts AltaString             = myUntypedString

    fmt.Println(uts)


    // Typed constants will NOT work
    // the const 'typedInt' can ONLY be used with type int
    /*
    const typedInt int             = 1
    var myFloat64 float64          = typedInt      // compiler error
    */
}

