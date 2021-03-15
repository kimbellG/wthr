package weather

import (
	"encoding/json"
	"net/http"
)

const GEOAPIURL = "api.ipgeolocation.io/ipgeo"
const GEOAPIKEY = "4f7fc8c2a920430da78d9729043bc514"

func GetGeolocationCoordinates() *AnswerIpGeolocationServer {
	return decodingJSONByGeolocationAndCloseConnection(getHTTPSRequest(getURLForGeolocationRequest()))
}

func getURLForGeolocationRequest() string {
	return getURLRequestWithParametrs(GEOAPIURL, requestParametrsToString(map[string]string{"apiKey": GEOAPIKEY, "fields": "latitude,longitude"}))
}

func decodingJSONByGeolocationAndCloseConnection(resp *http.Response) *AnswerIpGeolocationServer {
	defer resp.Body.Close()

	var result AnswerIpGeolocationServer
	err := json.NewDecoder(resp.Body).Decode(&result)
	assetsDecodingJSON(err)
	return &result
}
