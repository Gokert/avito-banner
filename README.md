# Banner

Приложение запускается командой:
```
docker-compose --env-file .env up --build
```

Приложение микросервисное, общение между сервисами происходит по GRPC API. Общение между браузером и сервсиом происходит через REST API. Присутствуют 2 сервиса: авторизации и баннеров. В сервисах реализована Чистая архитектура. Api отвечает за обработку запросов, usecase за бизнес логику, repo за работу с БД.
Для системы авторизации и сохранения сессий была выбрана бд кэширования Redis. Для получения данных пользователя и баннеров была выбрана бд PostgreSQL.

## Сервис авторизации (127.0.0.1:8081)
### Авторизация.
#### POST /signin
Результатом успешной авторизации является отдача cookie. Пример запроса:
```
{
    "login":"admin",
    "passsword": "admin"
}
```
### Регистрация
#### POST /signup
Результатом успешной регистрации является создание нового пользователя в БД. Пример запроса:
```
{
    "login":"admin",
    "passsword": "admin"
}
```

### Выход
#### DELETE /logout
Для выхода из аккаунта необходима кука session_id, которая была получена при авторизации.

### Проверка авторизации
#### GET /authcheck
Аутентификация пользователя. Проверка проходит по куке session_id.
