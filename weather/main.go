package weather

func GetCurrentWeather(req WeatherRequest) string {
	return formatOutputWeatherInfo(decodingJSONAndCloseConnection(getHTTPSRequest(getURLForCurrentWeatherData(req))))
}
