# URL Shortener

[![Go Report Card](https://goreportcard.com/badge/github.com/Str1kez/url-shortener)](https://goreportcard.com/report/github.com/Str1kez/url-shortener)

## Запуск
```commandline
make env
make up
```

## Load Testing
1. Заполнить userlist.txt
2. Запустить приложение
```commandline
cd tests
docker compose up -d --remove-orphans
docker logs -f yandex-tank
```

[Держит 50000 RPS](https://overload.yandex.net/577254)

