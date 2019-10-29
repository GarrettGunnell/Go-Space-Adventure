package main

import (
    "fmt"
    "math/rand"
    "time"
    "os"
    "encoding/json"
    "io/ioutil"
    "strings"
    "strconv"
)

type Planets struct {
    Planets []Planet `json:"planets"`
    Name string `json:"name"`
}

type Planet struct {
    Name string `json:"name"`
    Description string `json:"description"`
}

func parsePlanetData(filename string) Planets {
    planetsjson, err := os.Open(filename)
    if err != nil {
        fmt.Println(err)
    }
    jsondata, _ := ioutil.ReadAll(planetsjson)
    var planets Planets
    json.Unmarshal(jsondata, &planets)

    return planets
}

func printIntroduction(planetdata Planets) Planets {
    fmt.Println("Welcome to the " + planetdata.Name + "!")
    var numPlanets string = strconv.Itoa(len(planetdata.Planets))
    fmt.Println("There are " + numPlanets + " planets to explore.")
    fmt.Println("What is your name?")
    fmt.Println("Nice to meet you, " + takeInput() + ". Let's go for an adventure.")

    return planetdata
}

func manualOrRandomNavigation(planetdata Planets) (string, Planets) {
    fmt.Println("Shall I randomly choose a planet for you to visit? (Y / N)")
    var choice string
    for true {
        var temp string = takeInput()
        if temp == "Y" || temp == "N" {
            choice = temp
            break
        }
    }

    return choice, planetdata
}

func navigate(choice string, planetdata Planets) {
    if choice == "Y" {
        randomNavigation(planetdata)
    } else {
        manualNavigation(planetdata)
    }
}

func randomNavigation(planetdata Planets) {
    var randomindex = random(9)

    fmt.Println("Okay, traveling to " + planetdata.Planets[randomindex].Name + "...")
    fmt.Println("Arrived at " + planetdata.Planets[randomindex].Name + ". " + planetdata.Planets[randomindex].Description)
}

func manualNavigation(planetdata Planets) {
    fmt.Println("What planet would you like to travel to?")
    for true {
        var planet, exists, planetindex = findPlanet(planetdata, takeInput())
        if exists {
            fmt.Println("Traveling to " + planet + "...")
            fmt.Println("Arrived at " + planet + ". " + planetdata.Planets[planetindex].Description)
            break
        } else {
            fmt.Println("That planet doesn't exist.")
        }
    }
}

func takeInput() (input string) {
    fmt.Scanln(&input)
    input = strings.ToLower(input)
    input = strings.Title(input)
    return
}

func random(n int) int {
    seed := rand.NewSource(time.Now().UnixNano())
    temprand := rand.New(seed)
    return temprand.Intn(n)
}


func findPlanet(planetdata Planets, planet string) (name string, exists bool, planetindex int) {
    for i := 0; i < len(planetdata.Planets); i++ {
        if planetdata.Planets[i].Name == planet {
            name = planet
            exists = true
            planetindex = i
        }
    }
    return
}

func main() {
    navigate(manualOrRandomNavigation(printIntroduction(parsePlanetData(os.Args[1]))))
}
