package services

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseInt8(s string) int8 {
	f, err := strconv.ParseInt(s, 10, 8)
	check(err)
	return int8(f)
}
func ParseInt16(s string) int16 {
	f, err := strconv.ParseInt(s, 10, 16)
	check(err)
	return int16(f)
}
func ParseInt32(s string) int32 {
	f, err := strconv.ParseInt(s, 10, 32)
	check(err)
	return int32(f)
}

func ParseFloat32(s string) float32 {
	f, err := strconv.ParseFloat(s, 32)
	check(err)
	return float32(f)
}

/*
type weatherCode struct {
	thunderstorm_tornado_squall              byte
	rain_rainShowers_freezingRain            byte
	rain_rainSqualls_drizzle_freezingDrizzle byte
	snow_snowPellets_iceCrystals             byte
	snowShowers_snowSqualls_snowGrains       byte
	sleet_sleetShowers_hail                  byte
	fog_blowingDust_blowingSand              byte
	smoke_haze_smokeHaze_blowingSnow_dust    byte
	icePellets                               byte
}

func parseWeatherCodes(weatherCodeString string) weatherCode {
	return weatherCode{
		thunderstorm_tornado_squall:              weatherCodeString[0],
		rain_rainShowers_freezingRain:            weatherCodeString[1],
		rain_rainSqualls_drizzle_freezingDrizzle: weatherCodeString[2],
		snow_snowPellets_iceCrystals:             weatherCodeString[3],
		snowShowers_snowSqualls_snowGrains:       weatherCodeString[4],
		sleet_sleetShowers_hail:                  weatherCodeString[5],
		fog_blowingDust_blowingSand:              weatherCodeString[6],
		smoke_haze_smokeHaze_blowingSnow_dust:    weatherCodeString[7],
		icePellets:                               weatherCodeString[8],
	}
}
*/

type EPW struct {
	StationLocation                       string    `json:"station_location,omitempty"`
	State                                 string    `json:"state,omitempty"`
	Country                               string    `json:"country,omitempty"`
	Source                                string    `json:"source,omitempty"`
	StationID                             string    `json:"station_id,omitempty"`
	Latitude                              float32   `json:"latitude,omitempty"`
	Longitude                             float32   `json:"longitude,omitempty"`
	TimeZone                              string    `json:"time_zone,omitempty"`
	Elevation                             string    `json:"elevation,omitempty"`
	Year                                  []int16   `json:"year,omitempty"`
	Month                                 []int8    `json:"month,omitempty"`
	Day                                   []int8    `json:"day,omitempty"`
	Hour                                  []int8    `json:"hour,omitempty"`
	Minute                                []int8    `json:"minute,omitempty"`
	Uncertainty                           []string  `json:"uncertainty,omitempty"`
	DryBulbTemperature                    []float32 `json:"dry_bulb_temperature,omitempty"`
	DewPointTemperature                   []float32 `json:"dew_point_temperature,omitempty"`
	RelativeHumidity                      []float32 `json:"relative_humidity,omitempty"`
	AtmosphericStationPressure            []int32   `json:"atmospheric_station_pressure,omitempty"` // valid range 31,000 to 120,000; missing value 999999
	ExtraterrestrialHorizontalRadiation   []int16   `json:"extraterrestrial_horizontal_radiation,omitempty"`
	ExtraterrestrialDirectNormalRadiation []int16   `json:"extraterrestrial_direct_normal_radiation,omitempty"`
	HorizontalInfraredRadiationIntensity  []int16   `json:"horizontal_infrared_radiation_intensity,omitempty"`
	GlobalHorizontalRadiation             []int16   `json:"global_horizontal_radiation,omitempty"`
	DirectNormalRadiation                 []int16   `json:"direct_normal_radiation,omitempty"`
	DiffuseHorizontalRadiation            []int16   `json:"diffuse_horizontal_radiation,omitempty"`
	GlobalHorizontalIlluminance           []int32   `json:"global_horizontal_illuminance,omitempty"`
	DirectNormalIlluminance               []int32   `json:"direct_normal_illuminance,omitempty"`
	DiffuseHorizontalIlluminance          []int32   `json:"diffuse_horizontal_illuminance,omitempty"`
	ZenithLuminance                       []int16   `json:"zenith_luminance,omitempty"`
	WindDirection                         []float32 `json:"wind_direction,omitempty"`              // valid 0 to 360; missing value 999
	WindSpeed                             []float32 `json:"wind_speed,omitempty"`                  // valid 0 to 40; missing value 999
	TotalSkyCover                         []int8    `json:"total_sky_cover,omitempty"`             // valid 0 to 10; missing value 99
	OpaqueSkyCover                        []int8    `json:"opaque_sky_cover,omitempty"`            // valid 0 to 10; missing value 99
	Visibility                            []float32 `json:"visibility,omitempty"`                  // missing value 9999
	CeilingHeight                         []int32   `json:"ceiling_height,omitempty"`              // missing value 9999
	PresentWeatherObservation             []int8    `json:"present_weather_observation,omitempty"` // valid 0 to 9; missing value 99
	PresentWeatherCodes                   []string  `json:"present_weather_codes,omitempty"`
	PrecipitableWater                     []float32 `json:"precipitable_water,omitempty"`       // missing value 999
	AerosolOpticalDepth                   []float32 `json:"aerosol_optical_depth,omitempty"`    // missing value 999
	SnowDepth                             []float32 `json:"snow_depth,omitempty"`               // missing value 999
	DaysSinceLastSnowfall                 []float32 `json:"days_since_last_snowfall,omitempty"` // missing value 99
	Albedo                                []float32 `json:"albedo,omitempty"`
	LiquidPrecipitationDepth              []float32 `json:"liquid_precipitation_depth,omitempty"`
	LiquidPrecipitationQuantity           []float32 `json:"liquid_precipitation_quantity,omitempty"`
}

type EPWService struct {
	Ctx context.Context
	EPW *EPW
}

func NewEPWService() *EPWService {
	return &EPWService{}
}

func (fs *EPWService) Parse(fileName string) (*EPW, error) {
	epw := &EPW{}
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// parse header (first line of EPW)
	scanner.Scan()
	header := strings.Split(scanner.Text(), ",")
	epw.StationID = header[1]
	epw.StationLocation = header[2]
	epw.State = header[3]
	epw.Country = header[4]
	epw.Source = header[5]
	epw.Latitude = ParseFloat32(header[6])
	epw.Longitude = ParseFloat32(header[7])
	epw.TimeZone = header[8]
	epw.Elevation = header[9]

	// parse data (line 9 onwards)
	// skip next 7 lines of header
	for i := 0; i < 7; i++ {
		scanner.Scan()
	}
	var data []string
	for scanner.Scan() {
		data = strings.Split(scanner.Text(), ",")
		epw.Year = append(epw.Year, ParseInt16(data[0]))
		epw.Month = append(epw.Month, ParseInt8(data[1]))
		epw.Day = append(epw.Day, ParseInt8(data[2]))
		epw.Hour = append(epw.Hour, ParseInt8(data[3]))
		epw.Minute = append(epw.Minute, ParseInt8(data[4]))
		epw.Uncertainty = append(epw.Uncertainty, data[5])
		epw.DryBulbTemperature = append(epw.DryBulbTemperature, ParseFloat32(data[7]))
		epw.DewPointTemperature = append(epw.DewPointTemperature, ParseFloat32(data[8]))
		epw.RelativeHumidity = append(epw.RelativeHumidity, ParseFloat32(data[9]))
		epw.AtmosphericStationPressure = append(epw.AtmosphericStationPressure, ParseInt32(data[10]))
		epw.ExtraterrestrialHorizontalRadiation = append(epw.ExtraterrestrialHorizontalRadiation, ParseInt16(data[10]))
		epw.ExtraterrestrialDirectNormalRadiation = append(epw.ExtraterrestrialHorizontalRadiation, ParseInt16(data[11]))
		epw.HorizontalInfraredRadiationIntensity = append(epw.HorizontalInfraredRadiationIntensity, ParseInt16(data[12]))
		epw.GlobalHorizontalRadiation = append(epw.GlobalHorizontalRadiation, ParseInt16(data[13]))
		epw.DirectNormalRadiation = append(epw.DirectNormalRadiation, ParseInt16(data[14]))
		epw.DiffuseHorizontalRadiation = append(epw.DiffuseHorizontalRadiation, ParseInt16(data[15]))
		epw.GlobalHorizontalIlluminance = append(epw.GlobalHorizontalIlluminance, ParseInt32(data[16]))
		epw.DirectNormalIlluminance = append(epw.DirectNormalIlluminance, ParseInt32(data[17]))
		epw.DiffuseHorizontalIlluminance = append(epw.DiffuseHorizontalIlluminance, ParseInt32(data[18]))
		epw.ZenithLuminance = append(epw.ZenithLuminance, ParseInt16(data[19]))
		epw.WindDirection = append(epw.WindDirection, ParseFloat32(data[20]))
		epw.WindSpeed = append(epw.WindSpeed, ParseFloat32(data[21]))
		epw.TotalSkyCover = append(epw.TotalSkyCover, ParseInt8(data[22]))
		epw.OpaqueSkyCover = append(epw.OpaqueSkyCover, ParseInt8(data[23]))
		epw.Visibility = append(epw.Visibility, ParseFloat32(data[24]))
		epw.CeilingHeight = append(epw.CeilingHeight, ParseInt32(data[25]))
		epw.PresentWeatherObservation = append(epw.PresentWeatherObservation, ParseInt8(data[26]))
		epw.PresentWeatherCodes = append(epw.PresentWeatherCodes, data[27])
		epw.PrecipitableWater = append(epw.PrecipitableWater, ParseFloat32(data[28]))
		epw.AerosolOpticalDepth = append(epw.AerosolOpticalDepth, ParseFloat32(data[29]))
		epw.SnowDepth = append(epw.SnowDepth, ParseFloat32(data[30]))
		epw.DaysSinceLastSnowfall = append(epw.DaysSinceLastSnowfall, ParseFloat32(data[31]))
		epw.Albedo = append(epw.Albedo, ParseFloat32(data[32]))
		epw.LiquidPrecipitationDepth = append(epw.LiquidPrecipitationDepth, ParseFloat32(data[33]))
		epw.LiquidPrecipitationQuantity = append(epw.LiquidPrecipitationQuantity, ParseFloat32(data[34]))
	}
	return epw, err
}

func createKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func (fs *EPWService) Header(epw EPW) string {
	h := make(map[string]string)
	h["station location"] = epw.StationLocation
	h["state"] = epw.State
	h["country"] = epw.Country
	h["source"] = epw.Source
	h["station ID"] = epw.StationID
	h["latitude"] = fmt.Sprintf("%.2f", epw.Latitude)
	h["longitude"] = fmt.Sprintf("%.2f", epw.Longitude)
	h["time zone"] = epw.TimeZone
	h["elevation"] = epw.Elevation

	return createKeyValuePairs(h)
}
