package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func (n Todo) Display() {
	fmt.Println(n.Text)
}

func (n Todo) Save() error {
	fileName := "todo.json"

	//convert data to json
	json, err := json.Marshal(n)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)

}
func New(content string) (Todo, error) {

	if content == "" {
		return Todo{}, errors.New("invalid input")

	}

	return Todo{
		Text: content,
	}, nil
}
