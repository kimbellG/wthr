package weather

import (
	"log"
	"net/http"
	"time"
)

func getHTTPSRequest(url string) *http.Response {
	resp, err := createClient().Get("https://" + url)
	assetsGETrequest(err)
	assetsHTTPStatus(resp)
	return resp
}

func getURLForCurrentWeatherData(req WeatherRequest) string {
	return getURLRequest(requestParametrsToString(map[string]string{"q": req.CityName, "appid": req.APIKey, "units": defaultUnits}))
}

func getURLRequest(params string) string {
	return APIURL + "?" + params
}

func assetsHTTPStatus(resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Faild http GET-request: %v", resp.Status)
	}
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
