## URL do Cloud Run

`https://desafio-google-cloud-run-293510040052.us-central1.run.app`

## Endpoint

`GET /weather?cep=01001000`

## Obtendo a chave da WeatherAPI

A aplicação consulta a [WeatherAPI](https://www.weatherapi.com/) e exige uma chave (gratuita) passada via variável de ambiente `WEATHER_API_KEY`

## Rodando localmente via Docker

1. Construa a imagem:

    ```
    docker build -t desafio-cloud-run .
    ```

2. Execute o container substituindo `sua_chave` pelo valor copiado do painel da WeatherAPI:

    ```
    docker run --rm -p 8080:8080 -e WEATHER_API_KEY=sua_chave desafio-cloud-run
    ```

3. Teste:

    ```
    curl "http://localhost:8080/weather?cep=01001000"
    ```

## Rodando os testes

```
go test ./...
```
