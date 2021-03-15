package weather

import (
	"encoding/json"
	"net/http"
)

const GEOAPIURL = "api.ipgeolocation.io/ipgeo"

func GetGeolocationCoordinates() *AnswerIpGeolocationServer {
	return decodingJSONByGeolocationAndCloseConnection(getHTTPSRequest(getURLForGeolocationRequest()))
}

func getURLForGeolocationRequest() string {
	return getURLRequestWithParametrs(GEOAPIURL, requestParametrsToString(map[string]string{"apiKey": Conf.GEOAPIKEY, "fields": "latitude,longitude"}))
}

func decodingJSONByGeolocationAndCloseConnection(resp *http.Response) *AnswerIpGeolocationServer {
	defer resp.Body.Close()

	var result AnswerIpGeolocationServer
	err := json.NewDecoder(resp.Body).Decode(&result)
	assetsDecodingJSON(err)
	return &result
}
