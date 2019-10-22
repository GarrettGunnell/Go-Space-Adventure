package main

import (
  "fmt"
  rand "math/rand"
  "time"
)

func main() {
  seed := rand.NewSource(time.Now().UnixNano())
  newrand := rand.New(seed)
  var x = newrand.Intn(10)
  fmt.Print(x)

  /*
  fmt.Println("Welcome to the Solar System!")
  fmt.Println("There are 9 planets to explore.")
  fmt.Println("What is your name?")
  var name string
  fmt.Scanln(&name)
  fmt.Println("Nice to meet you, " + name + ". Let's go for an adventure.")
  fmt.Println("Shall I randomly choose a planet for you to visit? (Y / N)")

  for true {
    var choose string
    fmt.Scanln(&choose)
    if choose == "Y" {
      fmt.Println("Okay, traveling to RANDOM")
      break;
    } else if choose == "N" {
      fmt.Println("What planet would you like to travel to?")
      break;
    } else {
      fmt.Println("Excuse me?")
    }
  }
  */
}
