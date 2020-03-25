package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

// gameShapes have a max string lenght of 12 and must have an even number of items
// The order of winners, choice_1(player) will win to every other item starting with itself and
// lose to every other item starting with the next item (item1 beats item3 but loses to item2)
var gameShapes = [...]string{"rock", "paper", "scissors", "spock", "lizard"}
var shapeChoices = [len(gameShapes)]handShape{handShape{}}

type handShape struct {
    Id   int    `json:"id"`
    Name string `json:"name"`
}

type player struct {
    PlayerChoice int `json:"player"`
}

type playResults struct {
    Results        string `json:"results"`
    PlayerChoice   int    `json:"player"`
    ComputerChoice int    `json:"computer"`
}

type rng struct {
    RandomNum int `json:"random_number"`
}

// RNG Random Number Generator
// imports: API results of { "random_number": integer [1-100] }
// result: index of random chosen shape
func RNG() int {
    var newRNG rng

    //get random number from boohma api.
    res, err := http.Get("https://codechallenge.boohma.com/random")
    if err != nil {
        panic(err)
    }
    responseData, err := ioutil.ReadAll(res.Body)
    if err != nil {
        panic(err)
    }
    // unpack results
    json.Unmarshal([]byte(responseData), &newRNG)

    //modify value into shape choice
    result := newRNG.RandomNum % len(gameShapes)
    return result
}

// InitShapes initializes all shape choices
// result: none
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
//     “id": integer [1-(shapes count)],
//     "Name": string
//   }, ...
// ]
func GetAllChoices(res http.ResponseWriter, req *http.Request) {
    setupResponse(&res, req)

    if req.Method != "GET" {
        fmt.Fprintf(res, "Only GET method is supported.")
    }
    allChoicesjs, _ := json.Marshal(shapeChoices)

    res.Header().Set("Content-Type", "application/json")
    res.Write(allChoicesjs)
}

// choice: Get a randomly generated choice
// result:
// {
//     "Id": integer [1-(shapes count)],
//     "Name" : string
//}
func GetRandChoice(res http.ResponseWriter, req *http.Request) {
    setupResponse(&res, req)

    if req.Method != "GET" {
        fmt.Fprintf(res, "Only GET method is supported.")
    }
    rn := RNG()
    randomChoice := shapeChoices[rn]
    RNG_js, err := json.Marshal(randomChoice)
    if err != nil {
        http.Error(res, err.Error(), http.StatusInternalServerError)
        return
    }
    res.Header().Set("Content-Type", "application/json")
    res.Write(RNG_js)
}

// play: Post a round against a computer opponent
// input: { “player”: choice_id }
// result:
// {
//     "results": string [12] (win, lose, tie),
//     “player”: choice_id,
//     “computer”:  choice_id
// }
func CompareChoices(res http.ResponseWriter, req *http.Request) {
    setupResponse(&res, req)
    // get req data from POST
    if (*req).Method != "POST" {
        fmt.Fprintf(res, "Only POST method is supported.")
        return
    }
    decoder := json.NewDecoder(req.Body)
    var newPlay player
    err := decoder.Decode(&newPlay)
    if err != nil {
        panic(err)
    }
    choiceid := newPlay.PlayerChoice

    // get random chice ID
    randomID := shapeChoices[RNG()].Id

    var playResults playResults
    playResults.PlayerChoice = choiceid
    playResults.ComputerChoice = randomID

    //prep send data
    res.Header().Set("Content-Type", "application/json")
    if(choiceid == randomID){
            playResults.Results = "tie"
    // newPlay wins (allows for game expansion in the future without needing update)
    } else if ((choiceid > randomID) && ((choiceid+randomID)%2 == 1)) ||
              ((choiceid < randomID) && ((choiceid+randomID)%2 == 0)) {
            playResults.Results = "win"
    } else { //player loses
            playResults.Results = "lose"
    }

    Responsejs, _ := json.Marshal(playResults)
    res.Header().Set("Content-Type", "application/json")
    res.Write(Responsejs)
}

func main() {
    InitShapes()

    http.HandleFunc("/choices", GetAllChoices)
    http.HandleFunc("/choice", GetRandChoice)
    http.HandleFunc("/play", CompareChoices)
    http.HandleFunc("/", sayHello)

    fmt.Println("server 8080 running")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        panic(err)
    }
}

func sayHello(res http.ResponseWriter, req *http.Request) {
    message := "Hello Billups"
    message += " hope you are well"
    res.Write([]byte(message))
}

// cors setup
func setupResponse(w *http.ResponseWriter, req *http.Request) {
    (*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
