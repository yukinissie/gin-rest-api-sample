# gin-rest-api-sample
Ginの勉強用個人レポジトリ

## build & run

```
docker-compose up -d --build
```

run api

Create Scenario.

```
curl --location --request POST 'localhost:3000/api/scenario' \
--header 'Content-Type: application/json' \
--data-raw '{
    "Id": "utyu",
    "Name": "宇宙"
}'
```

Read Scenario.

```
curl --location --request GET 'localhost:3000/api/scenario'
```
