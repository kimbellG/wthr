package weather

func GetCurrentWeather(city string) string {
	return formatOutputWeatherInfo(decodingJSONAndCloseConnection(getHTTPSRequest(getURLForCurrentWeatherData(city, defaultUnits))))
}
