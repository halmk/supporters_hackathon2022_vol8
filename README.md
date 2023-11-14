# supporters_hackathon2022_vol8の継続開発用レポジトリ

fork元： https://github.com/halmk/supporters_hackathon2022_vol8


# 構成
## Frontend
- Nuxt.js
- Vuetify
- Firebase Hosting
## API
- Golang
- Gin
- AWS Lambda
- AWS API Gateway
## Databese
- Firebase Firestore Database
## Authentication
- Firebase Authentication
## Infrastructure
- GitHub
- Docker
- AWS
- Firebase

# Command
## Deploy
- Frontend
    -
    - `npm install -g firebase-tools`  
        (初回のみ)
    - `npm run generate`
    - `firebase login --no-localhost`
    - `firebase deploy`
- API
    -
    - `go get github.com/aws/aws-lambda-go/lambda`  
        (初回のみ)
    - `GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go`
    - `bootstrap`ファイルをzipにする
    - AWS Lambdaの該当関数にzipアップロード


## Development environment
-  `docker-compose up -d`
---
- Frontend
    -
    - `docker-compose exec app /bin/bash`
    - `cd frontend/`
    - `yarn`  
    (初回のみ)
    - `yarn dev`
- Backend
    -
    - `docker-compose exec api /bin/bash`
    - `cd api/`
    - `go run main.go`
