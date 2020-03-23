package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// RNG Random Number Generator
// result:
// { "random_number": integer [1-100] }
func RNG() int {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return (r1.Intn(99) + 1)
}

// gameShapes have a max string lenght of 12
var gameShapes = [...]string{"rocker", "paper", "scissors", "lizard", "spock"}
var shapeChoices = [len(gameShapes)]handShape{handShape{}}

type handShape struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func InitShapes() {
	for i := 0; i < len(gameShapes); i++ {
		println(gameShapes[i])
		shapeChoices[i] = handShape{
			Id:   i + 1, //Hand shapes start at 1
			Name: gameShapes[i],
		}
	}
}

// Choices: Get :respond with choices for the ui
// result:
// [
//   {
//     “id": integer [1-5],
//     "Name": string [12] (rock, paper, scissors, lizard, spock)
//   }
// ]
func GetAllChoices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("system call")
	if r.Method != "GET" {
		fmt.Fprintf(w, "Only GET method is supported.")
	}
	js, err := json.Marshal(shapeChoices)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// choice: Get a randomly generated choice
// result:
// {
//     "Id": integer [1-5],
//     "Name" : string [12] (rock, paper, scissors, lizard, spock)
//}
func GetRandChoice(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Only GET method is supported.")
	}
	rn := RNG()
	fmt.Print(rn, ",")
	rn = rn % len(gameShapes)
	randomChoice := shapeChoices[rn]
	js, err := json.Marshal(randomChoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// play: Post a round against a computer opponent
// intput:
// {
//     “player”: choice_id
// }
// result:
// {
//     "results": string [12] (win, lose, tie),
//     “player”: choice_id,
//     “computer”:  choice_id
func CompareChoices(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		fmt.Fprintf(w, "Only GET method is supported.")
	}
	rn := RNG()
	fmt.Print(rn, ",")
	rn = rn % len(gameShapes)
	randomChoice := shapeChoices[rn]
	js, err := json.Marshal(randomChoice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	InitShapes()
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/choices", GetAllChoices)
	http.HandleFunc("/choice", GetRandChoice)
	http.HandleFunc("/play", CompareChoices)
	//http.HandleFunc("/", sayHello)
	fmt.Println("server 8080 running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	message += " how are you feeling?"
}
