/* Alta3 Research | RZFeeser
   Custom Sorting - Sorting structs with custom functions   */

package main

import (
    "fmt"
    "sort"
)

// record a "Person" Name, Age
type Person struct {
    Name   string
    Age    int
}

// return a formatted string
func (p Person) String() string {
    return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

// ByAge implements sort.Interface for []Person based on
// the Age field.
type ByAge []Person

func (a ByAge) Len() int {
    return len(a)
}
func (a ByAge) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}
func (a ByAge) Less(i, j int) bool {
    return a[i].Age < a[j].Age
}

func main() {

    // Name, Age
    people := []Person{
        {"Bob", 31},
        {"John", 42},
        {"Michael", 17},
        {"Jenny", 26},
    }

    fmt.Println(people)

    // There are two (2) ways to sort a slice

    // 1)
    // Impliment the Interfact with Len, Swap, Less 
    sort.Sort(ByAge(people))
    fmt.Println(people)

    // 2)  
    // Use "sort.Slice" with a custom Less
    // function, which can be provided as a closure. In this
    // case no methods are needed. (And if they exist, they
    // are ignored.)
    // re-sort in reverse order: compare
    sort.Slice(people, func(i, j int) bool {
        return people[i].Age > people[j].Age
    })
    fmt.Println(people)

}

