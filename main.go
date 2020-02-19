
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
    var computerGuess = rand.Intn(max)
    var userGuess int

    fmt.Println("Enter a guess for the computer number")

    fmt.Scanln(&userGuess)

    for userGuess!=computerGuess {
      fmt.Println("That number is incorrect, try again!")

      count++

      fmt.Scanln(&userGuess)

      if userGuess==computerGuess{
        fmt.Println("You've guessed correctly!")
      }
    }
    fmt.Println("It took you",count,"times to guess the correct number")

}