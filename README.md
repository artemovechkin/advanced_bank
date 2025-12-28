# advanced_bank

Репозиторий для студентов Бруноям

Команда сборки образа с приложением:
```
docker build -t bank-app .
```

Команда запуска контейнера из собранного на предыдущем шаге образа:
```
docker run -d -p 8080:8080 --name bank-container bank-app
```

```
docker run -d -p 8080:8080 -v $(pwd)/accounts.json:/app/accounts.json --name bank-container bank-app
```
