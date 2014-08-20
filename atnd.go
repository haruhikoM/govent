package go_events

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	partition = "==========================="
)

type Event struct {
	EventID  int `json:"event_id"`
	Title    string
	Catch    string
	EventURL string `json:"event_url"`
	URL      string
	Address  string
	// StartedAt
}

type Response struct {
	Data struct {
		ResultsReturned int `json:"results_returned"`
		Children        []struct {
			Data Event
		}
	}
}

func (e Event) String() string {
	return fmt.Sprintf("%s\n%s%s\n%s\n\n%s\n%s", partition, e.Title, e.Catch, e.Address, e.EventURL, partition)
}

func Get(event string) ([]Event, error) {
	url := fmt.Sprintf("http://api.atnd.org/events/?keyword=%s&format=json", event)
	// url := fmt.Sprintf("http://reddit.com/r/%s.json", reddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(Response)
	err = json.NewDecoder(resp.Body).Decode(r)
	if err != nil {
		return nil, err
	}
	events := make([]Event, len(r.Data.Children))
	for i, child := range r.Data.Children {
		events[i] = child.Data
	}
	return events, nil
}

func main() {
	// http://api.atnd.org/events/?keyword_or=google,cloud&format=atom
	resp, err := http.Get("http://api.atnd.org/events/?keyword=golang&format=json")

	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		log.Fatal(resp.Status)
	}

	_, err = io.Copy(os.Stdout, resp.Body)
	if err != nil {
		log.Fatal(err)
	}
}
