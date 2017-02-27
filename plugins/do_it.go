package main

import (
	"encoding/json"
	"time"
	"errors"
	"fmt"
	"strconv"
)

type Resp struct {
	Aint int
	Abool bool
	Astring string
	Ts time.Time
}

func Process(args ...string) ([]byte, error) {

	if len(args) < 3 {
		return nil, errors.New("Not enough args")
	}

	a_int, err := strconv.Atoi(args[0])
	a_bool, err := strconv.ParseBool(args[1])
	a_string := args[2]

	if err != nil {
		return nil, err
	}

	resp := Resp{a_int, a_bool, a_string, time.Now()}

	fmt.Println(resp)

	json_resp, err := json.Marshal(&resp)

	if err != nil {
		return nil, err
	}

	return json_resp, nil
}
