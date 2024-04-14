# [Avito-Banner](https://github.com/avito-tech/backend-trainee-assignment-2024)

Приложение запускается командой:
```
make up
```

Приложение микросервисное, общение между сервисами происходит по GRPC API. Общение между браузером и сервисом происходит через REST API. Присутствуют 2 сервиса: авторизации и баннеров. В сервисах реализована Чистая архитектура. Api отвечает за обработку запросов, usecase за бизнес логику, repo за работу с БД.
Также присутствует контейнер c Nginx.
Для системы авторизации и сохранения сессий была выбрана бд кэширования Redis. Для получения данных пользователя и баннеров была выбрана бд PostgreSQL.

## Сервис авторизации
### Авторизация.
#### POST /signin
Результатом успешной авторизации является выдача cookie. Пример: <br/>
![img_1.png](img_readme/img_1.png)

### Регистрация
#### POST /signup
Результатом успешной регистрации является создание нового пользователя в БД. Пример: <br/>
![img.png](img_readme/img.png)

### Выход
#### DELETE /logout
Для выхода из аккаунта необходима кука session_id, которая была получена при авторизации. <br/>
![img_3.png](img_readme/img_3.png)

### Проверка авторизации
#### GET /authcheck
Аутентификация пользователя. Проверка проходит по куке session_id. <br/>
![img_2.png](img_readme/img_2.png)

## Сервис баннеров
### Получение баннеров.
#### GET /api/v1/banner 
Получение баннеров<br/>
![img_4.png](img_readme/img_4.png)

### Создание нового баннера.
#### POST /api/v1/banner
Создание баннера<br/>
![img_5.png](img_readme/img_5.png)

### Изменение баннера.
#### PATCH /api/v1/banner/{id}
Изменение баннера<br/>
![img_6.png](img_readme/img_6.png)

### Удаление баннера.
#### DELETE /api/v1/banner/{id}
Удаление баннера<br/>
![img_7.png](img_readme/img_7.png)