package weather

import (
	"fmt"
	"time"
)

func formatOutputWeatherInfo(fullWeatherInfo *AnswerWeatherServer) string {
	return fmt.Sprintf("at: %s\nin: %s:\n%s\n\t%s\t%s\n",
		formatOutputTime(fullWeatherInfo.Dt), formatOutputWeaterRegion(*fullWeatherInfo.Sys, fullWeatherInfo.Name),
		formatOutputTemp(*fullWeatherInfo.Main, *fullWeatherInfo.Weather[0]),
		formatOutputWind(*fullWeatherInfo.Wind),
		formatOutputVisibility(fullWeatherInfo.Visibility))
}

func formatOutputTime(tm int64) string {
	t := time.Unix(tm, 0)
	return fmt.Sprintf("%02d.%02d.%d, %02d:%02d",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute())
}

func formatOutputWeaterRegion(country SysInfo, name string) string {
	return fmt.Sprintf("%s:%s", country.Country, name)
}

func formatOutputTemp(temperatures MainInfo, condition WeatherInfo) string {
	return fmt.Sprintf("\tTemp:\n\t\tCurrent: %.0f \u2103\n\t\tMin-Max: %.0f \u2103 -- %.0f \u2103\n\t\tFeels like: %.0f \u2103\n\t\tCondition of weather: %s",
		temperatures.Temp, temperatures.TempMin, temperatures.TempMax, temperatures.FeelsLike, condition.Description)
}

func formatOutputWind(wind WindInfo) string {
	return fmt.Sprintf("Wind speed: %.1f km/h", wind.Speed)
}

func formatOutputVisibility(visibility int) string {
	return fmt.Sprintf("Visibility: %.1f km.", float32(visibility/1000))
}

func formatOutputError(err error) string {
	return fmt.Sprint(err)
}
