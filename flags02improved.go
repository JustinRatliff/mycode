/* Alta3 Research | RZFeeser
   Flags - Using the init() function   */
   
package main

import (
    "flag"
    "fmt"
)

var (
    env  *string
    port *int
)

// the 'init()' function is often used to initialize state variables
func init() {
    env = flag.String("env", "development", "current environment")
    port = flag.Int("port", 3000, "port number")
}

func main() {

    flag.Parse()

    fmt.Println("env:", *env)
    fmt.Println("port:", *port)
}

