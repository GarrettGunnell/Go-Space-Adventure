package main

import (
  "fmt"
  "math/rand"
  "time"
  "os"
  "encoding/json"
  "io/ioutil"
  "strings"
)

func main() {

  planetsjson, err := os.Open(os.Args[1])
  if err != nil {
    fmt.Println(err)
  }

  jsondata, _ := ioutil.ReadAll(planetsjson)
  var planets Planets

  json.Unmarshal(jsondata, &planets)
  for i := 0; i < len(planets.Planets); i++ {
    fmt.Println("Planet Name: " + planets.Planets[i].Name)
  }


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
      seed := rand.NewSource(time.Now().UnixNano())
      newrand := rand.New(seed)
      var randomindex = newrand.Intn(10)

      fmt.Println("Okay, traveling to " + planets.Planets[randomindex].Name + "...")
      fmt.Println("Arrived at " + planets.Planets[randomindex].Name + ". " + planets.Planets[randomindex].Description)
      break;
    } else if choose == "N" {
      fmt.Println("What planet would you like to travel to?")
      for true {
        var planet string
        fmt.Scanln(&planet)
        planet = strings.ToLower(planet)
        planet = strings.Title(planet)
        var exists bool
        var planetindex int
        for i := 0; i < len(planets.Planets); i++ {
          if planets.Planets[i].Name == planet {
            exists = true;
            planetindex = i;
          }
        }
        if exists {
          fmt.Println("Traveling to " + planet + "...")
          fmt.Println("Arrived at " + planet + ". " + planets.Planets[planetindex].Description)
          break;
        } else {
          fmt.Println("That planet doesn't exist.")
        }
      }
      break;
    } else {
      fmt.Println("Excuse me?")
    }
  }
}

type Planets struct {
  Planets []Planet `json:"planets"`
}

type Planet struct {
  Name string `json:"name"`
  Description string `json:"description"`
}
