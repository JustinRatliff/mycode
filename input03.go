/* Alta3 Research | RZFeeser
   Reading a line of text with bufio.NewReader()   */

package main

import (
     "os"
     "bufio"
     "fmt"
)

func main() {

     reader := bufio.NewReader(os.Stdin)   // create a "Reader" object

     fmt.Print("Enter your name: ")        // print to the screen

     name, _ := reader.ReadString('\n')    // Read up to the "\n" character. This returns "string, error"
                                           //    so we capture the string in "name", and throw out "error" via "_"
     fmt.Printf("Hello %s\n", name)
}

