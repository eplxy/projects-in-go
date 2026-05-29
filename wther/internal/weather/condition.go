package weather

import (
	"fmt"
	"strings"
)

type Condition struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
	Code int    `json:"code"`
}

var conditionPrepositionByCode = map[int]string{
	1000: "and",  // Sunny
	1003: "and",  // Partly cloudy
	1006: "and",  // Cloudy
	1009: "and",  // Overcast
	1030: "with", // Mist
	1063: "with", // Patchy rain possible
	1066: "with", // Patchy snow possible
	1069: "with", // Patchy sleet possible
	1072: "with", // Patchy freezing drizzle possible
	1087: "with", // Thundery outbreaks possible
	1114: "with", // Blowing snow
	1117: "with", // Blizzard
	1135: "with", // Fog
	1147: "with", // Freezing fog
	1150: "with", // Patchy light drizzle
	1153: "with", // Light drizzle
	1168: "with", // Freezing drizzle
	1171: "with", // Heavy freezing drizzle
	1180: "with", // Patchy light rain
	1183: "with", // Light rain
	1186: "with", // Moderate rain at times
	1189: "with", // Moderate rain
	1192: "with", // Heavy rain at times
	1195: "with", // Heavy rain
	1198: "with", // Light freezing rain
	1201: "with", // Moderate or heavy freezing rain
	1204: "with", // Light sleet
	1207: "with", // Moderate or heavy sleet
	1210: "with", // Patchy light snow
	1213: "with", // Light snow
	1216: "with", // Patchy moderate snow
	1219: "with", // Moderate snow
	1222: "with", // Patchy heavy snow
	1225: "with", // Heavy snow
	1237: "with", // Ice pellets
	1240: "with", // Light rain shower
	1243: "with", // Moderate or heavy rain shower
	1246: "with", // Torrential rain shower
	1249: "with", // Light sleet showers
	1252: "with", // Moderate or heavy sleet showers
	1255: "with", // Light snow showers
	1258: "with", // Moderate or heavy snow showers
	1261: "with", // Light showers of ice pellets
	1264: "with", // Moderate or heavy showers of ice pellets
	1273: "with", // Patchy light rain with thunder
	1276: "with", // Moderate or heavy rain with thunder
	1279: "with", // Patchy light snow with thunder
	1282: "with", // Moderate or heavy snow with thunder
}

func (c Condition) Phrase() string {
	preposition, ok := conditionPrepositionByCode[c.Code]

	if !ok {
		//fallback is and
		return fmt.Sprintf("and %s", strings.ToLower(c.Text))
	}

	return fmt.Sprintf("%s %s", preposition, strings.ToLower(c.Text))

}
