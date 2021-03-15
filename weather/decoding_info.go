package weather

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func proccesingJSONWithCurrentWeatherAndCloseConnection(resp *http.Response) (*AnswerWeatherServer, error) {
	defer resp.Body.Close()

	if err := assetsHTTPStatus(resp); err != nil {
		return nil, err
	}

	return decodingJSONWithCurrentWeather(resp), nil
}

func decodingJSONWithCurrentWeather(resp *http.Response) *AnswerWeatherServer {
	var result AnswerWeatherServer
	err := json.NewDecoder(resp.Body).Decode(&result)
	assetsDecodingJSON(err)

	return &result
}

func decodingErrorMessage(resp *http.Response) ErrorMessage {
	var result ErrorMessage
	err := json.NewDecoder(resp.Body).Decode(&result)
	assetsDecodingJSON(err)

	return result
}

func assetsHTTPStatus(resp *http.Response) error {
	if resp.StatusCode != http.StatusOK {
		return errors.New(decodingErrorMessage(resp).Message)
	} else {
		return nil
	}
}

func assetsDecodingJSON(err error) {
	if err != nil {
		log.Fatalf("decodingJSON: Incorrect JSON struct: %v", err)
	}
}
