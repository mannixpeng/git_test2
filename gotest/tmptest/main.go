package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Player struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	processData("data.txt")
}

func processData(file string) error {
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()
	return unmarshalAndPrint(f)
}

func unmarshalAndPrint(f io.Reader) error {
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}
	var players []Player

	err = json.Unmarshal(data, &players)
	if err != nil {
		return err
	}

	for _, p := range players {
		fmt.Println("Player name: ", p.Name)
	}
	return nil
}
