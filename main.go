//Nicholas Larsen
//February 18, 2020
//Prompts the user to guess the computer's random number
package main

import (
    "fmt"
    "math/rand"
)

func main() {
  var max int
  var count int

  fmt.Println("Enter a max range for this guessing game")
  
fmt.Scanln(&max)
   //has the user set a max range for the random computer guess
    var computerGuess = rand.Intn(max)
    var userGuess int

    fmt.Println("Enter a guess for the computer number")

    fmt.Scanln(&userGuess)
    //the user's guess for the number

    for userGuess!=computerGuess {
      fmt.Println("That number is incorrect, try again!")

      count++
      //counts the tries

      fmt.Scanln(&userGuess)
      //has user input another number

      if userGuess==computerGuess{
        fmt.Println("You've guessed correctly!")
        //tells user they have guessed the number correctly
      }
    }
    fmt.Println("It took you",count,"times to guess the correct number")
    //tells the user how many tries it took

}