package weather

import "fmt"

func GetCurrentWeather(req WeatherRequestByCityName) string {
	currentWeather, err := proccesingJSONWithCurrentWeatherAndCloseConnection(getHTTPSRequest(getURLForCurrentWeatherDataByCityName(req)))
	if err != nil {
		return fmt.Sprint(err)
	}

	return formatOutputWeatherInfo(currentWeather)
}

func GetCurrentWeatherForGeolocation(req WeatherRequestByGeoCoord) string {
	currentWeather, err := proccesingJSONWithCurrentWeatherAndCloseConnection(getHTTPSRequest(getURLForCurrentWeatherDataByGeoCoord(req)))
	if err != nil {
		return formatOutputError(err)
	}

	return formatOutputWeatherInfo(currentWeather)
}
