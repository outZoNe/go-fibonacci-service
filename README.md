0) Склонируйте репозиторий
1) Установить и запустить Redis
2) Сделать `cp .env.example .env` и настроить `.env` файл (если redis запущен на порту: `6379`, то по идее ничего
   настраивать не нужно. Достаточно просто скопировать `.env.example` в `.env`)
3) Запустить gRPC модуль: `go run main.go` либо собрать и запустить бинарник: `go build main.go && ./<bin_file>`
4) Запустить gRPC-Gateway модуль для работы по http: `go run proxy.go` либо собрать и запустить
   бинарник: `go build proxy.go && ./<bin_file>`
5) Проверить работу RESTful можно
   тут: [http://127.0.0.1:8080/fibonacci?startNum=5&endNum=10](http://127.0.0.1:8080/fibonacci?startNum=5&endNum=10)
6) Для проверки gRPC можно использовать клиент: [BloomRPC](https://github.com/uw-labs/bloomrpc/releases)
