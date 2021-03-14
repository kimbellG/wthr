package weather

type WeatherInfo struct {
	Description string
}

type MainInfo struct {
	Temp      float64
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
}

type WindInfo struct {
	Speed float64
}

type SysInfo struct {
	Country string
}

type AnswerWeatherServer struct {
	Weather    []*WeatherInfo
	Main       *MainInfo
	Visibility int
	Wind       *WindInfo
	Dt         int64
	Sys        *SysInfo
	Name       string
}

type WeatherRequest struct {
	CityName string
	APIKey   string
}

const APIURL = "api.openweathermap.org/data/2.5/weather"
const defaultUnits = "metric"
