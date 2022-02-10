// Justin's Game | by Justin Ratliff 
// GO language class 2.10.2022

package main

import (
	"fmt"  	      // import the fmt package which allows text, formatting, reading input & printing output functions
	"math/rand"   // import the rand package to generate random numbers
	"time"        // import the package to allow time functions
)

func main() {
	fmt.Println("This is Justin's game in GO")
	fmt.Println("The Game: Your mission is to guess the branch of military service")
	fmt.Println("You have two(2) tries")
	fmt.Println("If you guess wrong, you will be given a clue")
	fmt.Println("Each branch of service will be represented by a number")
	fmt.Println("0 = Army")
	fmt.Println("1 = Air Force")
	fmt.Println("2 = Navy")
	fmt.Println("3 = Marines")
	fmt.Println("4 = Coast Guard")
	fmt.Println("5 = Space Force")

// generating random numbers
   source := rand.NewSource(time.Now().UnixNano())        
	       
   //  The default number generator is predictable, so it will produce the same sequence of numbers each time. To 
   //  produce varying range of numbers, give it a seed that changes (in this case: time would ensure it changes ). 
   //  Note that this is not safe to use for random numbers you want to be secret; use crpyto/rand for those.

    // now we can generate some chaos with our "source" seed
            randomizer := rand.New(source)
	    secretNumber := randomizer.Intn(5)  // generate numbers between 0 and 5 only

    // we declare a variable of type int called "guess"
    var guess int

    // create a for loop that begins at 1 and loops to 3
    for try := 1; try <= 2; try++  {
// print out the number of times the player has made a guess
        fmt.Println("Round:", try)

        // the program will prompt the player to make a guess and enter a number
        fmt.Println("Please enter your number")
        fmt.Scan(&guess)   // this function makes it possible for the program to receive the input

        if guess < secretNumber {
            // if the guessed number is less than or greater than the correct number; give the player a hint

            fmt.Printf("Sorry, wrong guess ; number is too small\n ")
		if secretNumber == 0 {
		fmt.Println("Hint: Be All You Can Be!\n")}
                if secretNumber == 1 {
		fmt.Println("Hint: Aim High!\n ")}
		if secretNumber == 2 {
		fmt.Println("Hint: Always Courageous!\n ")}
                if secretNumber == 3 {
		fmt.Println("Hint: Sempre Fidelis!\n ")}
		if secretNumber == 4 {
		fmt.Println("Hint: Always Ready!\n ")}
                if secretNumber == 5 {
		fmt.Println("Hint: Always Above!\n ")}
        } else if guess > secretNumber {
            fmt.Printf("Sorry, wrong guess ; number is too large\n ")
		if secretNumber == 0 {
		fmt.Println("Hint: Be All You Can Be!\n ")}
                if secretNumber == 1 {
		fmt.Println("Hint: Aim High!\n ")}
		if secretNumber == 2 {
		fmt.Println("Hint: Always Courageous!\n ")}
                if secretNumber == 3 {
		fmt.Println("Hint: Sempre Fidelis!\n ")}
		if secretNumber == 4 {
		fmt.Println("Hint: Always Ready!\n ")}
                if secretNumber == 5 {
		fmt.Println("Hint: Always Above!\n ")}
        } else {
            fmt.Printf("You win!\n")
            break
            // Print out "you win" message when the player guesses the correct number
        }

        if try == 2 {
            // if the number of tries is equal to 2, print game over and also the correct number

            fmt.Printf("Game over!!\n ")
            fmt.Printf("The correct number is %d\n", secretNumber)
            break

        }
    // every time the game ends, print this message
    fmt.Println("Thanks for playing, try again!")

    } 
    
}
