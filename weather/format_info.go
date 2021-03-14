package weather

import (
	"fmt"
	"time"
)

func formatOutputWeatherInfo(fullWeatherInfo *AnswerWeatherServer) string {
	return fmt.Sprintf("Time: %s. Region: %s\n%s\n%s\t%s\n",
		formatOutputTime(fullWeatherInfo.Dt), formatOutputWeaterRegion(*fullWeatherInfo.Sys, fullWeatherInfo.Name),
		formatOutputTemp(*fullWeatherInfo.Main, *fullWeatherInfo.Weather[0]),
		formatOutputWind(*fullWeatherInfo.Wind),
		formatOutputVisibility(fullWeatherInfo.Visibility))
}

func formatOutputTime(tm int64) string {
	curTime := time.Unix(tm, 0)
	return fmt.Sprintf("%d.%d.%d %d:%d:%d",
		curTime.Day(), curTime.Month(), curTime.Year(),
		curTime.Hour(), curTime.Minute(), curTime.Second())
}

func formatOutputWeaterRegion(country SysInfo, name string) string {
	return fmt.Sprintf("%s:%s", country.Country, name)
}

func formatOutputTemp(temperatures MainInfo, condition WeatherInfo) string {
	return fmt.Sprintf("Temp:\n\tCurrent: %.0f\n\tMin-Max: %.0f-%.0f\n\tFeels like: %.0f\n\tCondition of weather: %s",
		temperatures.Temp, temperatures.TempMin, temperatures.TempMax, temperatures.FeelsLike, condition.Description)
}

func formatOutputWind(wind WindInfo) string {
	return fmt.Sprintf("Wind speed: %.1f km/h", wind.Speed)
}

func formatOutputVisibility(visibility int) string {
	return fmt.Sprintf("Visibility: %.1f km.", float32(visibility/1000))
}
