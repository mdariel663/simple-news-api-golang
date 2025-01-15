package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const defaultApiKey = "YOUR DEFAULT API KEY"
const defaultCountryCode = "US"
const defaultPort = "8080"

type NewsAPIResponse struct {
	Articles []struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	} `json:"articles"`
	TotalResults int `json:"totalResults"`
}

func fetchNews(apiKey string, countryCode string) (*NewsAPIResponse, error) {
	// Construir la URL para la API
	apiUrl := fmt.Sprintf("https://newsapi.org/v2/top-headlines?country=%s&apiKey=%s", countryCode, apiKey)
	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var newsResponse NewsAPIResponse
	if err := json.Unmarshal(body, &newsResponse); err != nil {
		return nil, err
	}

	return &newsResponse, nil
}

func handler(w http.ResponseWriter, r *http.Request, apiKey string, countryCode string) {
	// Habilitar CORS
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	if r.Method == http.MethodOptions {
		return
	}

	if apiKey == "" {
		apiKey = defaultApiKey
	}
	if countryCode == "" {
		countryCode = defaultCountryCode
	}

	news, err := fetchNews(apiKey, countryCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func main() {
	port := flag.String("port", defaultPort, "Puerto en el que se ejecutará el servidor")
	apiKey := flag.String("api-key", defaultApiKey, "Clave de API para NewsAPI")
	countryCode := flag.String("country", defaultCountryCode, "Código de país para las noticias")
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, *apiKey, *countryCode)
	})

	address := fmt.Sprintf(":%s", *port)
	fmt.Printf("Server started on %s\n", address)

	// Iniciar el servidor web
	log.Fatal(http.ListenAndServe(address, nil))
}
