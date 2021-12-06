package weatherHandler

type Weather struct {
	Current []Current `json:"current"`
}

type Current struct {
	Temp_c     float32      `json:"temp_c"`
	Temp_f     float32      `json:"temp_f"`
	Condition  []Condition  `json:"condition"`
	AirQuality []AirQuality `json:"air_quality`
}

type Condition struct {
	Text string `json:"text"`
}

type AirQuality struct {
	Co           float64 `json:"co"`
	Pm10         float32 `json:"pm10"`
	GbDefraIndex int     `json:"gb-defra-index"`
}
