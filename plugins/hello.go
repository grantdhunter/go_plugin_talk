package main

import (
	"encoding/json"
	"time"
	"errors"
	"fmt"
)

type Greeting struct {
	Msg string
	Ts time.Time
}

func Process(args ...string) ([]byte, error) {

	if len(args) < 1 {
		return nil, errors.New("Not enough args")
	}
	
	greet := Greeting{"Hello " + args[0] + " from plugin", time.Now()}

	fmt.Println(greet)

	json_resp, err := json.Marshal(&greet)

	if err != nil {
		return nil, err
	}

	return json_resp, nil
}
