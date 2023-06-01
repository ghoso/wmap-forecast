# OpenWeatherMap Web APIクライアント

OpenWeatherMap(https://openweathermap.org)から天気情報を取得するクライアント。  
  
## 事前準備

環境変数WMAP_API_KEYにサイトで取得したAPIキーを設定する。  
  
export WMAP_API_KEY=（取得したAPIキー)

## 実行

wmap_forecast --location ロケーション名

--locationはオプションで指定しなかった場合はTokyoが指定される。
