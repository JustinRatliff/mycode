/* Alta3 Research | RZFeeser
   Methods - value and pointer receivers   */


package main
 
import (
    "fmt"
)
type Shark struct{
    name string
}

// pass by value - makes a copy of what is passed (changes WILL NOT persist)
func (a Shark)Swim(){
    fmt.Println(a.name, "shark is swimming...")
}

// pass by pointer - do this when you need to make changes (changes WILL persist)
func (a *Shark)SwimFaster(){
    fmt.Println(a.name, "shark is swimming...")
}
 
func main() {
    fish := Shark{"Tiger"}
    
    // both of these will work
    fish.Swim()                // Tiger shark is running...
    fish.SwimFaster()          // Tiger shark is running...
}

