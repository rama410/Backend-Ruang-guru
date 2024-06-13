package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

// AIModelConnector struct to handle the AI model connection
type AIModelConnector struct {
	Client *http.Client
}

// Inputs struct for AI model input
type Inputs struct {
	Table map[string][]string `json:"table"`
	Query string              `json:"query"`
}

// Response struct for AI model output
type Response struct {
	Answer      string   `json:"answer"`
	Coordinates [][]int  `json:"coordinates"`
	Cells       []string `json:"cells"`
	Aggregator  string   `json:"aggregator"`
}

// CsvToSlice function to convert CSV data to map
func CsvToSlice(data string) (map[string][]string, error) {
	reader := csv.NewReader(strings.NewReader(data))
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	result := make(map[string][]string)
	headers := records[0]
	for i, header := range headers {
		result[header] = []string{}
		for _, record := range records[1:] {
			result[header] = append(result[header], record[i])
		}
	}
	return result, nil
}

// ConnectAIModel function to interact with the AI model
func (c *AIModelConnector) ConnectAIModel(payload interface{}, token string) (Response, error) {
	url := "https://api-inference.huggingface.co/models/google/tapas-base-finetuned-wtq"
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return Response{}, err
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonPayload)))
	if err != nil {
		return Response{}, err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.Client.Do(req)
	if err != nil {
		return Response{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Response{}, err
	}

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return Response{}, err
	}

	return response, nil
}

func main() {
	// Load CSV data
	data, err := ioutil.ReadFile("data-series.csv")
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	// Convert CSV data to map
	table, err := CsvToSlice(string(data))
	if err != nil {
		fmt.Println("Error converting CSV to slice:", err)
		return
	}

	// Prepare AI model input
	query := "What is the energy consumption of the Refrigerator in the Kitchen?"
	input := Inputs{
		Table: table,
		Query: query,
	}

	// Connect to AI model
	client := &http.Client{}
	connector := AIModelConnector{Client: client}
	token := os.Getenv("HUGGINGFACE_API_TOKEN") // Set your Hugging Face API token in environment variables
	response, err := connector.ConnectAIModel(input, token)
	if err != nil {
		fmt.Println("Error connecting to AI model:", err)
		return
	}

	// Print the response
	fmt.Printf("Answer: %s\n", response.Answer)
	fmt.Printf("Coordinates: %v\n", response.Coordinates)
	fmt.Printf("Cells: %v\n", response.Cells)
	fmt.Printf("Aggregator: %s\n", response.Aggregator)
}
