/* Alta3 Research | RZFeeser
   CHALLENGE 01 - Mapping programming languages to extension names */

package main

import "fmt"

func main() {
    var fileExtensions = map[string]string{
        "Python": ".py",
        "C++":    ".cpp",
        "Java":   ".java",
        "Golang": ".go",
        "Ansible": ".yml",
    }

    fmt.Println(fileExtensions)

    // remove the key C++ and its associated value
    delete(fileExtensions, "C++")

    // because the map is initalized, we can add new keys to it
    fileExtensions["Julia"] = ".jl"

    // display the state of the map
    fmt.Println(fileExtensions)
}

