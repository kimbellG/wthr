package weather

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func getHTTPSRequest(url string) *http.Response {
	resp, err := createClient().Get("https://" + url)
	assetsGETrequest(err)
	return resp
}

func getURLForCurrentWeatherDataByCityName(req WeatherRequestByCityName) string {
	return getURLRequestWithParametrs(APIURL, requestParametrsToString(map[string]string{"q": req.CityName, "appid": req.APIKey, "units": defaultUnits}))
}

func getURLForCurrentWeatherDataByGeoCoord(req WeatherRequestByGeoCoord) string {
	return getURLRequestWithParametrs(APIURL, requestParametrsToString(map[string]string{"lat": fmt.Sprint(req.Coordinate.Latitude), "lon": fmt.Sprint(req.Coordinate.Longitude),
		"appid": req.APIKey, "units": defaultUnits}))
}

func getURLRequestWithParametrs(APIURL string, params string) string {
	return APIURL + "?" + params
}

func assetsGETrequest(err error) {
	if err != nil {
		log.Fatalf("Connection failed. %v", err)
	}
}

func createClient() *http.Client {
	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}

	return &http.Client{Transport: transport}
}

func requestParametrsToString(parametrsMap map[string]string) string {
	var ReqParamtersString string
	for key, value := range parametrsMap {
		ReqParamtersString += key + "=" + value + "&"
	}

	//return url.QueryEscape(ReqParamtersString[:len(ReqParamtersString)-1])
	return ReqParamtersString[:len(ReqParamtersString)-1]
}
