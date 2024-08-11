# films-api

REST API на Go.

### Запуск приложения

```
make build && make run
```

Миграция базы данных (нужна при первом запуске приложения):

```
make migrate
```
Запустить тесты:

```
make test
```
Документация REST API будет доступна локально после запуска приложения:

http://localhost:8000/swagger/index.html
