# RPSSL-game-server
Rock Paper Scissors Lizard Spock!! Microservice API setup in GO, run your own RPSSL game **TODAY**

## Get Started
requirements: GO and Git

1. Clone the repostitory
    git clone https://github.com/jmiceter/RPSLS-game-server

2. go run main.go
    (This will run on port 8080 of your system)

3. Using a popular web browser you may use http://localhost:8080 (this port may be changed if the port is already in use)

## System Microservices
* Get list of options
    - **GET** request ULR + "/choices"
    - Returns
    [
        {
            “id": "integer",
            "name": "string *defualt on clone(rock, paper, scissors, lizard, spock)*"
        } {..} ...
    ]
* Get random choice
    - **GET** request URL + "/choice"
    - Returns:
        {
          "id": "integer (item id)",
          "name" : "string *defualt on clone(rock, paper, scissors, lizard, spock)*"
        }
* Play your hand against the CPU (random choice)
    - **POST** request URL + "/play" and post Data "{ player: *choice_id*}"
    - Returns:
    {
        "results": *win results (win lose tie)*
        "player": *choice_id*
    }

## Designed for Modifications
* Add or remove more hand shapes. An array of hand shapes are provided in the software and more may be added. Just make sure an odd
    number of gestures are provided. The software will auto integrate them in the system on compile
* Port modification. Change the port to your disired host port. just make sure you change your localhost in your web browser
    (http://localhost:xxxx)




Special thank to the following:

Setting up terminal to run ec2 forever on ssh close (credit to https://unix.stackexchange.com/questions/4004/how-can-i-run-a-command-which-will-survive-terminal-close)

Applying Cors to the server for sending data (credit to https://flaviocopes.com/golang-enable-cors/)

