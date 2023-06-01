#!/bin/sh

echo "Geocode:"
curl "http://api.openweathermap.org/geo/1.0/direct?q=Tokyo&limit=5&appid={API_KEY}"
echo "\nWeather:"
curl "http://api.openweathermap.org/data/2.5/weather?q=Tokyo&appid={API_KEY}"
