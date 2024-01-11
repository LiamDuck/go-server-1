package messages

import (
	"encoding/json"
	"os"
)

type FileStore struct {
	file string
}

func NewFileStore(f string) *FileStore {
	return &FileStore{
		f,
	}

}

func (fs FileStore) Add(message Message) error {
	content, err := os.ReadFile(fs.file)
	if err != nil {
		return err
	}

	var messages []Message
	if len(content) > 0 {

		err = json.Unmarshal(content, &messages)
		if err != nil {
			return err
		}
	}
	messages = append(messages, message)

	outputData, err := json.MarshalIndent(messages, "", "	")
	if err != nil {
		return err
	}
	err = os.WriteFile(fs.file, outputData, 0644)
	if err != nil {
		return err
	}

	return nil

}

func (fs FileStore) List() ([]Message, error) {
	content, err := os.ReadFile(fs.file)
	if err != nil {
		return nil, err
	}

	var messages []Message
	if len(content) > 0 {

		err = json.Unmarshal(content, &messages)
		if err != nil {
			return nil, err
		}
	}
	return messages, nil
}

func (fs FileStore) Remove(id string) error {
	content, err := os.ReadFile(fs.file)
	if err != nil {
		return err
	}
	var messages []Message
	err = json.Unmarshal(content, &messages)
	if err != nil {
		return err
	}
	var finMessages []Message
	for _, data := range messages {
		if data.ID != id {
			finMessages = append(finMessages, data)
		}
	}
	return nil
}
