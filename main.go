package main
// forecast
// OpenWeather Web APIから天気情報を取得する
// https://openwethermap.org

import (
  "os"
  "flag"
  "fmt"
  "io"
  "net/http"
  "encoding/json"
)

// Web API URL
const Url = "http://api.openweathermap.org/data/2.5/weather"
// ロケーション
const defaultCity = "Tokyo"
// ケルビン摂氏変換
const cKelvin = 273.15

// お天気データ型
type WeatherData struct {
  Id string          `json:"id"`
  Main string        `json:"main"`
  Description string `json:"description"`
  Icon string        `json:"icon"`
}

type TemperatureData struct {
  Temp float64     `json:"temp"`
  FeelLike float64 `json:"feel_like"`
  TempMin float64  `json:"temp_min"`
  TempMax float64  `json:"temp_max"`
  Pressure int     `json:"pressure"`
  Humidity int     `json:"humidity"`
}

type ForecastData struct {
  Weather [5]WeatherData `json:"weather"`
  Temperature TemperatureData `json:"main"`
}

// サーバー接続
func connectService(apikey string, locationName string) (*http.Response, error){
  var city string

  if len(locationName) == 0 {
    city = defaultCity
  } else {
    city = locationName
  }
  // fmt.Printf("Location = %v\n", city)
  apiUrl := fmt.Sprintf("%v?q=%v&appid=%v", Url, city, apikey)
  // fmt.Printf("URL=%v\n", apiUrl)
  req, err := http.NewRequest(http.MethodGet, apiUrl, nil)
  resp, err := http.DefaultClient.Do(req)
  return resp, err
}

// レスポンスJSONデータをForecastData型に変換
func decodeForcast(resp *http.Response) (*ForecastData, error){
  var fdata *ForecastData

  fdata = new(ForecastData)
  body, _ := io.ReadAll(resp.Body)
  // fmt.Printf("body = %s\n", body)
  json.Unmarshal(body, fdata)
  // fmt.Printf("fdata = %v\n", fdata.Weather[0])
  return fdata,nil
}

func main() {
  // APIキー取得
  apiKey := os.Getenv("WMAP_API_KEY")
  if len(apiKey) == 0{
    fmt.Println("Error: WMAP_API_KEY does not defined")
    os.Exit(1)
  }

  // ロケーション指定
  location := flag.String("location", "", "Location")
  flag.Parse()

  // サーバー接続
  http_resp, err :=connectService(apiKey, *location)
  if (err != nil) {
    fmt.Printf("error: %v\n", err)
  }
  defer http_resp.Body.Close()
  // レスポンス取り出し
  data, _ := decodeForcast(http_resp)
  fmt.Printf("weather = %v\n", data.Weather[0].Main)
  fmt.Printf("Temperature Max = %.2f\n",data.Temperature.TempMax - cKelvin)
  fmt.Printf("Temperature Min = %.2f\n",data.Temperature.TempMin - cKelvin)
  fmt.Printf("Temperature Humidity = %v\n",data.Temperature.Humidity)
}
