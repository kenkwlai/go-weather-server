package models

type Coord struct {
  Lon         float32   `json:"lon"`
  Lat         float32   `json:"lat"`
}

type Weather struct {
  Id          int       `json:"id"`
  Main        string    `json:"main"`
  Description string    `json:"description"`
  Icon        string    `json:"icon"`
}

type Main struct {
  Temp        float32   `json:"temp"`
  Pressure    int       `json:"pressure"`
  Humidity    int       `json:"humidity"`
  TempMin     float32   `json:"temp_min"`
  TempMax     float32   `json:"temp_max"`
  SeaLevel    int       `json:"sea_level"`
  GrndLevel   int       `json:"grnd_level"`
}

type Wind struct {
  Speed       float32   `json:"speed"`
  Deg         int       `json:"deg"`
}

type Cloud struct {
  All         int       `json:"all"`
}

type Rain struct {
  OneHour     int       `json:"1h"`
  ThreeHour   int       `json:"3h"`
}

type Snow struct {
  OneHour     int       `json:"1h"`
  ThreeHour   int       `json:"3h"`
}

type Sys struct {
  Type        int       `json:"type"`
  Id          int       `json:"id"`
  Message     float32   `json:"message"`
  Country     string    `json:"country"`
  Sunrise     uint64    `json:"sunrise"`
  Sunset      uint64    `json:"sunset"`
}

type CurrentWeatherData struct {
  Coord       Coord     `json:"coord"`
  Weather     []Weather `json:"weather"`
  Base        string    `json:"base"`
  Main        Main      `json:"main"`
  Visibility  int       `json:"visibility"`
  Wind        Wind      `json:"wind"`
  Clouds      Cloud     `json:"clouds"`
  Rain        Rain      `json:"rain"`
  Snow        Snow      `json:"snow"`
  Dt          uint64    `json:"dt"`
  Sys         Sys       `json:"sys"`
  Timezone    int       `json:"timezone"`
  Id          int       `json:"id"`
  Name        string    `json:"name"`
  Cod         int       `json:"cod"`
}
