// 	w.Write([]byte(message))
// }

// func main() {
// 	http.HandleFunc("/", sayHello)
// 	if err := http.ListenAndServe(":8080", nil); err != nil {
// 		panic(err)
// 	}

// }

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

// Get random number 1 - 100
func RNG() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return fmt.Print(r1.Intn(99)+1, ",")
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
	RNG()
}

func pong(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ping"))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	message := r.URL.Path
	message = strings.TrimPrefix(message, "/")
	message = "Hello " + message
	message += " how are you feeling?"
}

// Random Number Generator
// result:
// { "random_number": integer [1-100] }

// Choices: Get choices for the ui
// result:
// [
//   {
//     “id": integer [1-5],
//     "name": string [12] (rock, paper, scissors, lizard, spock)
//   }
// ]

// choice: Get a randomly generated choice
// result:
// {
//     "id": integer [1-5],
//     "name" : string [12] (rock, paper, scissors, lizard, spock)
//}

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
// }

func main() {
	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/pong", pong)
	//http.HandleFunc("/", sayHello)
	fmt.Println("server 8080 running")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
