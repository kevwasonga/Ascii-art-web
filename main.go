package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"as/asciiArt"
)

func main() {
	// Serve the static HTML page
	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.HandleFunc("/ascii-art", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
			return
		}

		var request struct {
			Banner string `json:"banner"`
			Input  string `json:"input"`
		}

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		fileName := asciiArt.BannerFile(request.Banner)

		// Load the banner map from the file
		bannerMap, err := asciiArt.LoadBannerMap(fileName)
		if err != nil {
			http.Error(w, "Failed to load banner file", http.StatusInternalServerError)
			return
		}

		response := generateASCIIArt(request.Input, bannerMap)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(response))
	})

	fmt.Println("Server is listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func generateASCIIArt(input string, bannerMap map[int][]string) string {
	lines := make([]string, 8)

	for _, char := range input {
		banner, exists := bannerMap[int(char)]
		if !exists {
			banner = bannerMap[32] // Fallback to space if character not found
		}
		for i := 0; i < 8; i++ {
			lines[i] += banner[i]
		}
	}

	return strings.Join(lines, "\n")
}
