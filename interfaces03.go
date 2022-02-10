/* Alta3 Research | RZFeeser
   Interface - Abstract class */
   
package main
   
import (
    "fmt"
)

// define the interface
type Animal interface{
    Sound() string
    MakeSound()
}

// partially define the 
type abstractAnimal struct {Animal}

func (a abstractAnimal) MakeSound() {
    fmt.Printf("%v", a.Sound())
}


type Chicken struct {abstractAnimal}
func (c Chicken) Sound() string {
    return "bwaaak bwak bwak bwak\n"
}

func NewChicken() *Chicken {
    chicken := Chicken{abstractAnimal{}}
    chicken.abstractAnimal.Animal = chicken
    return &chicken
}

type Lion struct {abstractAnimal}
func (d Lion) Sound() string {
    return "RAAAAWWWWWRRRR\n"
}

func NewLion() *Lion{
    lion := Lion{abstractAnimal{}}
    lion.abstractAnimal.Animal = lion
    return &lion
}

func main(){
    c := NewChicken()
    c.MakeSound()
    d := NewLion()
    d.MakeSound()
}  

