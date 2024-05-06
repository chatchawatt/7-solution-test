package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"unicode"
)

type MeatSummary struct {
	Beef map[string]int `json:"beef"`
}

func main() {
	http.HandleFunc("/beef/summary", beefSummaryHandler)
	http.ListenAndServe(":8080", nil)
}

func beefSummaryHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		http.Error(w, "Failed to fetch meat text", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	meatBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	meatString := string(meatBytes)

	words := strings.Fields(meatString)

	meatCounts := make(map[string]int)

	for _, word := range words {

		word = strings.ToLower(word)
		word = strings.TrimRightFunc(word, func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		meatCounts[word]++
	}

	summary := MeatSummary{
		Beef: meatCounts,
	}

	jsonData, err := json.Marshal(summary)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}
