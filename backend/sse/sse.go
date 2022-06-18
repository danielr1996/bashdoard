package sse

import (
	"danielr1996/bashdoard/types"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gopkg.in/antage/eventsource.v1"
	"log"
	"net/http"
)

type SSE struct {
	es      eventsource.EventSource
	Entries map[string]types.DashboardEntry
}

func New() *SSE {
	c := new(SSE)
	c.Entries = make(map[string]types.DashboardEntry)
	c.es = eventsource.New(
		eventsource.DefaultSettings(),
		func(req *http.Request) [][]byte {
			return [][]byte{
				[]byte("Access-Control-Allow-Origin: *"),
			}
		},
	)
	return c
}

func (s *SSE) Add(id string, entry types.DashboardEntry) {
	s.Entries[id] = entry
	fmt.Print("ADD: ")
	fmt.Println(entry)
	b, err := json.Marshal(entry)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	s.es.SendEventMessage(string(b), "add", uuid.New().String())

}

func (s *SSE) Delete(id string) {
	fmt.Print("DEL: ")
	fmt.Println(s.Entries[id])
	b, err := json.Marshal(s.Entries[id])
	delete(s.Entries, id)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	s.es.SendEventMessage(string(b), "delete", uuid.New().String())
}

func (s *SSE) Update(id string, entry types.DashboardEntry) {
	old := s.Entries[id]
	newEntry := entry
	fmt.Print("OLD: ")
	fmt.Println(old)
	fmt.Print("NEW: ")
	fmt.Println(newEntry)
	b, err := json.Marshal(newEntry)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	s.es.SendEventMessage(string(b), "update", uuid.New().String())
}

func (s *SSE) Contains(id string) bool {
	if _, ok := s.Entries[id]; ok {
		return true
	}
	return false
}
func (s *SSE) Get(id string) types.DashboardEntry {
	return s.Entries[id]
}

func (s *SSE) Serve(stopCh <-chan struct{}) {
	http.Handle("/api/dashboardentries/sse", s.es)
	http.HandleFunc("/api/dashboardentries", func(w http.ResponseWriter, req *http.Request) {

		var entries []types.DashboardEntry
		for _, entry := range s.Entries {
			entries = append(entries, entry)
		}
		b, err := json.Marshal(entries)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		w.Header().Set("Access-Control-Allow-Origin", "*")
		fmt.Fprintf(w, string(b))
	})
	log.Fatal(http.ListenAndServe(":3040", nil))
	<-stopCh
	s.es.Close()
}
