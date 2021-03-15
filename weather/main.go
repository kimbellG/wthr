package weather

func GetCurrentWeather(req WeatherRequestByCityName) string {
	return formatOutputWeatherInfo(decodingJSONBCurrentWeatherAndCloseConnection(getHTTPSRequest(getURLForCurrentWeatherDataByCityName(req))))
}

func GetCurrentWeatherForGeolocation(req WeatherRequestByGeoCoord) string {
	return formatOutputWeatherInfo(decodingJSONBCurrentWeatherAndCloseConnection(getHTTPSRequest(getURLForCurrentWeatherDataByGeoCoord(req))))
}
