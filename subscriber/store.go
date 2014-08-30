package subscriber

import (
	"encoding/csv"
	"fmt"
	"os"
)

// Store is a place to store subscribers
type Store struct{}

// Save the subscriber to the list
func (s Store) Save(subscriber Subscriber) error {
	list := subscriber.Referer
	if list == "" {
		list = "default"
	}

	listPath := fmt.Sprintf("lists/%s.list", list)
	fileOptions := os.O_RDWR | os.O_CREATE | os.O_APPEND
	file, err := os.OpenFile(listPath, fileOptions, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write(subscriber.ToCSV())
	writer.Flush()
	return nil
}
