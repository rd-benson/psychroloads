package services

import (
	"bufio"
	"bytes"
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
	stationLocation                       string
	state                                 string
	country                               string
	source                                string
	stationID                             string
	latitude                              float32
	longitude                             float32
	timeZone                              string
	elevation                             string
	year                                  []int16
	month                                 []int8
	day                                   []int8
	hour                                  []int8
	minute                                []int8
	uncertainty                           []string
	dryBulbTemperature                    []float32
	dewPointTemperature                   []float32
	relativeHumidity                      []float32
	atmosphericStationPressure            []int32 // valid range 31,000 to 120,000; missing value 999999
	extraterrestrialHorizontalRadiation   []int16
	extraterrestrialDirectNormalRadiation []int16
	horizontalInfraredRadiationIntensity  []int16
	globalHorizontalRadiation             []int16
	directNormalRadiation                 []int16
	diffuseHorizontalRadiation            []int16
	globalHorizontalIlluminance           []int32
	directNormalIlluminance               []int32
	diffuseHorizontalIlluminance          []int32
	zenithLuminance                       []int16
	windDirection                         []float32 // valid 0 to 360; missing value 999
	windSpeed                             []float32 // valid 0 to 40; missing value 999
	totalSkyCover                         []int8    // valid 0 to 10; missing value 99
	opaqueSkyCover                        []int8    // valid 0 to 10; missing value 99
	visibility                            []float32 // missing value 9999
	ceilingHeight                         []int32   // missing value 9999
	presentWeatherObservation             []int8    // valid 0 to 9; missing value 99
	presentWeatherCodes                   []string
	precipitableWater                     []float32 // missing value 999
	aerosolOpticalDepth                   []float32 // missing value 999
	snowDepth                             []float32 // missing value 999
	daysSinceLastSnowfall                 []float32 // missing value 99
	albedo                                []float32
	liquidPrecipitationDepth              []float32
	liquidPrecipitationQuantity           []float32
}

func NewEPW() *EPW {
	return &EPW{}
}

func (epw *EPW) Parse(fileName string) (EPW, error) {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// parse header (first line of EPW)
	scanner.Scan()
	header := strings.Split(scanner.Text(), ",")
	epw.stationID = header[1]
	epw.stationLocation = header[2]
	epw.state = header[3]
	epw.country = header[4]
	epw.source = header[5]
	epw.latitude = ParseFloat32(header[6])
	epw.longitude = ParseFloat32(header[7])
	epw.timeZone = header[8]
	epw.elevation = header[9]

	// parse data (line 9 onwards)
	// skip next 7 lines of header
	for i := 0; i < 7; i++ {
		scanner.Scan()
	}
	var data []string
	for scanner.Scan() {
		data = strings.Split(scanner.Text(), ",")
		epw.year = append(epw.year, ParseInt16(data[0]))
		epw.month = append(epw.month, ParseInt8(data[1]))
		epw.day = append(epw.day, ParseInt8(data[2]))
		epw.hour = append(epw.hour, ParseInt8(data[3]))
		epw.minute = append(epw.minute, ParseInt8(data[4]))
		epw.uncertainty = append(epw.uncertainty, data[5])
		epw.dryBulbTemperature = append(epw.dryBulbTemperature, ParseFloat32(data[7]))
		epw.dewPointTemperature = append(epw.dewPointTemperature, ParseFloat32(data[8]))
		epw.relativeHumidity = append(epw.relativeHumidity, ParseFloat32(data[9]))
		epw.atmosphericStationPressure = append(epw.atmosphericStationPressure, ParseInt32(data[10]))
		epw.extraterrestrialHorizontalRadiation = append(epw.extraterrestrialHorizontalRadiation, ParseInt16(data[10]))
		epw.extraterrestrialDirectNormalRadiation = append(epw.extraterrestrialHorizontalRadiation, ParseInt16(data[11]))
		epw.horizontalInfraredRadiationIntensity = append(epw.horizontalInfraredRadiationIntensity, ParseInt16(data[12]))
		epw.globalHorizontalRadiation = append(epw.globalHorizontalRadiation, ParseInt16(data[13]))
		epw.directNormalRadiation = append(epw.directNormalRadiation, ParseInt16(data[14]))
		epw.diffuseHorizontalRadiation = append(epw.diffuseHorizontalRadiation, ParseInt16(data[15]))
		epw.globalHorizontalIlluminance = append(epw.globalHorizontalIlluminance, ParseInt32(data[16]))
		epw.directNormalIlluminance = append(epw.directNormalIlluminance, ParseInt32(data[17]))
		epw.diffuseHorizontalIlluminance = append(epw.diffuseHorizontalIlluminance, ParseInt32(data[18]))
		epw.zenithLuminance = append(epw.zenithLuminance, ParseInt16(data[19]))
		epw.windDirection = append(epw.windDirection, ParseFloat32(data[20]))
		epw.windSpeed = append(epw.windSpeed, ParseFloat32(data[21]))
		epw.totalSkyCover = append(epw.totalSkyCover, ParseInt8(data[22]))
		epw.opaqueSkyCover = append(epw.opaqueSkyCover, ParseInt8(data[23]))
		epw.visibility = append(epw.visibility, ParseFloat32(data[24]))
		epw.ceilingHeight = append(epw.ceilingHeight, ParseInt32(data[25]))
		epw.presentWeatherObservation = append(epw.presentWeatherObservation, ParseInt8(data[26]))
		epw.presentWeatherCodes = append(epw.presentWeatherCodes, data[27])
		epw.precipitableWater = append(epw.precipitableWater, ParseFloat32(data[28]))
		epw.aerosolOpticalDepth = append(epw.aerosolOpticalDepth, ParseFloat32(data[29]))
		epw.snowDepth = append(epw.snowDepth, ParseFloat32(data[30]))
		epw.daysSinceLastSnowfall = append(epw.daysSinceLastSnowfall, ParseFloat32(data[31]))
		epw.albedo = append(epw.albedo, ParseFloat32(data[32]))
		epw.liquidPrecipitationDepth = append(epw.liquidPrecipitationDepth, ParseFloat32(data[33]))
		epw.liquidPrecipitationQuantity = append(epw.liquidPrecipitationQuantity, ParseFloat32(data[34]))
	}
	return *epw, err
}

func createKeyValuePairs(m map[string]string) string {
	b := new(bytes.Buffer)
	for key, value := range m {
		fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
	}
	return b.String()
}

func (epw *EPW) Header() string {
	h := make(map[string]string)
	h["station location"] = epw.stationLocation
	h["state"] = epw.state
	h["country"] = epw.country
	h["source"] = epw.source
	h["station ID"] = epw.stationID
	h["latitude"] = fmt.Sprintf("%.2f", epw.latitude)
	h["longitude"] = fmt.Sprintf("%.2f", epw.longitude)
	h["time zone"] = epw.timeZone
	h["elevation"] = epw.elevation

	return createKeyValuePairs(h)
}
