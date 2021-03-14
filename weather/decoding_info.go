package weather

import (
	"encoding/json"
	"log"
	"net/http"
)

func decodingJSONAndCloseConnection(resp *http.Response) *AnswerWeatherServer {
	defer resp.Body.Close()

	var result AnswerWeatherServer
	err := json.NewDecoder(resp.Body).Decode(&result)
	assetsDecodingJSON(err)
	return &result
}

func assetsDecodingJSON(err error) {
	if err != nil {
		log.Fatalf("decodingJSON: Incorrect JSON struct: %v", err)
	}
}
