## Go Practice: Simple Web Server to fetch weather data

This project is written in [Go](https://github.com/golang/go) with [Gin Web Framework](https://github.com/gin-gonic/gin) that runs on [Docker](https://www.docker.com/) containers.

It provides REST API to fetch current weather data in Hong Kong from [OpenWeatherMap](https://openweathermap.org/current). If [OpenWeatherMap](https://openweathermap.org/current) is down, the app gets weather data from the [MongoDB](https://www.mongodb.com/) instance. 

## Installation

1. Clone this repository ```git clone https://github.com/kenkwlai/go-weather-server.git```

2. Have [Docker](https://www.docker.com/) installed on your local machine 

3. Setup `.env` file for your local environment

4. Build and run the app with Docker ```docker-compose up --build```

## Usage

The server exposes two APIs:
- `POST  /authorize` to get a [JWT](https://jwt.io/) for authorization, currently there is no validation on the credentials, just to illustrate the flow to obtain the token.

Example request with body:
```json
{
  "username": "some trivial name",
  "password": "some trivial password"
}
```

Example response:
```json
{
  "expire": "2019-07-30T00:44:01+0800",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjQ0MTg2NDEsImlhdCI6MTU2NDQxNTA0MSwiaXNzIjoia2VuLmsudy5sYWlAZ21haWwuY29tIiwidXNlcm5hbWUiOiJzb21lIn0.32MlR6M69sR3VmtzLgcYYw3aenRLCyRHlan3mfUHhBI"
}
```

- `GET  /weather` to get current weather data in Hong Kong, use the token obtained from `/authorize` to authenticate

Example request:
```curl -i -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjQ0MTg2NDEsImlhdCI6MTU2NDQxNTA0MSwiaXNzIjoia2VuLmsudy5sYWlAZ21haWwuY29tIiwidXNlcm5hbWUiOiJzb21lIn0.32MlR6M69sR3VmtzLgcYYw3aenRLCyRHlan3mfUHhBI" http://localhost:8000/weather```

Example response:
```json
{
    "coord": {
        "lon": 114.16,
        "lat": 22.29
    },
    "weather": [
        {
            "id": 801,
            "main": "Clouds",
            "description": "few clouds",
            "icon": "02n"
        }
    ],
    "base": "stations",
    "main": {
        "temp": 301.39,
        "pressure": 1007,
        "humidity": 79,
        "temp_min": 299.82,
        "temp_max": 302.15,
        "sea_level": 0,
        "grnd_level": 0
    },
    "visibility": 10000,
    "wind": {
        "speed": 4.6,
        "deg": 110
    },
    "clouds": {
        "all": 20
    },
    "rain": {
        "1h": 0,
        "3h": 0
    },
    "snow": {
        "1h": 0,
        "3h": 0
    },
    "dt": 1564337755,
    "sys": {
        "type": 1,
        "id": 9154,
        "message": 0.0067,
        "country": "HK",
        "sunrise": 1564350819,
        "sunset": 1564398349
    },
    "timezone": 28800,
    "id": 1819729,
    "name": "Hong Kong",
    "cod": 200
}
```

## TODOs
- Add business logic to validate user credentials
- Add more unit tests
- Refactor the code to have more structure
