package players

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func getPlayerNames() []string {
	jsonFile, err := os.Open("players/players.json")
	if err != nil {
		fmt.Println(err)
		log.Fatal("could not load player names.")
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var players map[string][]string
	json.Unmarshal([]byte(byteValue), &players)

	return players["playerNames"]
}

func getRandomPlayer(players []string) string {
	var amountOfPlayers = len(players)
	return players[rand.Intn(amountOfPlayers)]
}

func spliceTwoNames(name1 string, name2 string) string {
	return name1[0:int(math.Round(float64(len(name1)/2)))] + name2[int(math.Round(float64(len(name2)/2))):]
}

func Rname() string {
	var players []string = getPlayerNames()

	var name1 = getRandomPlayer(players)
	var name2 = getRandomPlayer(players)

	return spliceTwoNames(name1, name2)
}
