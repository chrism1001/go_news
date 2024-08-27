package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func main() {
	// router := mux.NewRouter()

	apiKey := goDotEnvVariable("API_KEY")
	reqURL := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=us&apiKey=%s", apiKey)
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		log.Fatalf("Failed to create request object for /GET endpoint: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to send HTTP request: %v", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed to read response body: %v", err)
	}

	defer resp.Body.Close()

	data := make(map[string]interface{})
	json.Unmarshal(body, &data)
}
