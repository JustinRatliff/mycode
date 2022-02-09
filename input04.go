/* Alta3 Research | RZFeeser
   Improved -- Reading a line of text with bufio.NewReader()   */

package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    fmt.Print("Enter text: ")
    reader := bufio.NewReader(os.Stdin)
    
    // ReadString will block until the delimiter is entered
    input, err := reader.ReadString('\n')   // this returns (string, error) and we capture both
    
    // simple error handling
    if err != nil {
        fmt.Println("An error occurred while reading input. Please try again", err)
        return
    }

    // remove the delimiter from the string
    input = strings.TrimSuffix(input, "\n")   // A "Suffix" means added to "the end of"
    fmt.Println(input)   // "fmt.Println" has a built in "\n" it will append to any string
}

