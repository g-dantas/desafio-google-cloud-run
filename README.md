# Desafio Google Cloud Run

Sistema em Go que recebe um CEP, identifica a cidade (ViaCEP) e retorna a temperatura atual em Celsius, Fahrenheit e Kelvin (WeatherAPI).

## URL do Cloud Run

`https://<preencher-apos-deploy>.run.app`

## Endpoint

`GET /weather?cep=01001000`

Respostas:

- `200` — `{ "temp_C": 28.5, "temp_F": 83.3, "temp_K": 301.5 }`
- `422` — `invalid zipcode` (formato inválido)
- `404` — `can not find zipcode` (CEP inexistente)

## Obtendo a chave da WeatherAPI

A aplicação consulta a [WeatherAPI](https://www.weatherapi.com/) e exige uma chave (gratuita) passada via variável de ambiente `WEATHER_API_KEY`

## Rodando localmente via Docker

1. Garanta que você já tem a `WEATHER_API_KEY` (passo acima).
2. Construa a imagem:

    ```
    docker build -t desafio-cloud-run .
    ```

3. Execute o container substituindo `sua_chave` pelo valor copiado do painel da WeatherAPI:

    ```
    docker run --rm -p 8080:8080 -e WEATHER_API_KEY=sua_chave desafio-cloud-run
    ```

4. Teste:

    ```
    curl "http://localhost:8080/weather?cep=01001000"
    ```

## Rodando os testes

```
go test ./...
```
