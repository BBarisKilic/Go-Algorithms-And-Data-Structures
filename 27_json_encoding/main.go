package main

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

type Manager struct {
	FullName       string
	Position       string
	Age            int32
	YearsInCompany int32
}

func main() {

	// manager := Manager{"Baris Kilic", "Sofware Dev", 27, 5}
	manager := &Manager{}

	manager.FullName = "Baris"
	manager.Position = "Sofware Dev"
	manager.Age = 27
	manager.YearsInCompany = 5

	fmt.Println(EncodeManager(manager))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func EncodeManager(manager *Manager) (result io.Reader, err error) {

	jsonString, err := json.Marshal(manager)

	result = strings.NewReader(string(jsonString))

	return result, err
}
