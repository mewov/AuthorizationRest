# Authorization Service 🔐

**Auth Service** — микросервис авторизации и управления пользователями, написанный на **Golang (Gin)**.
Он обеспечивает **регистрацию**, **вход**, **освежение и валидацию JWT токенов**, а также **логаут**.
Использует **PostgreSQL** для хранения пользователей и **rate limiter** для защиты от спама.

**Auth Service** is a microservice for user authentication and session management, built with **Golang (Gin)**.
It supports **registration**, **login**, **token refresh and validation**, and **logout**,
backed by **PostgreSQL** and protected by a built-in **rate limiter**.

---

## 🧩 Core Features / Основные возможности

| Feature / Функция           | Description (EN)                                          | Описание (RU)                                              |
| --------------------------- | --------------------------------------------------------- | ---------------------------------------------------------- |
| **User Registration**       | Create user with login, password, email, role, and client | Регистрация пользователя с логином, паролем, email и ролью |
| **Login**                   | Authenticate user and issue JWT tokens                    | Авторизация пользователя и выдача JWT                      |
| **Access / Refresh Tokens** | Short-lived access + long-lived refresh pair              | Короткоживущий access и долговечный refresh токен          |
| **Token Refresh**           | Get a new access token via refresh token                  | Обновление токена доступа по refresh токену                |
| **User Info**               | Get decoded user data from access token                   | Получение информации о пользователе из access токена       |
| **Logout**                  | Revoke session and remove refresh token                   | Выход и удаление активной сессии                           |
| **Rate Limiting**           | Middleware limits request frequency                       | Ограничение частоты запросов через middleware              |
| **Structured Logging**      | Custom logger (slog-based) for each request               | Логирование всех запросов через собственный логгер         |

---

## ⚙️ Technologies / Технологии

| Layer              | Stack                                      |
| ------------------ | ------------------------------------------ |
| **Language**       | Go 1.25+                                   |
| **Framework**      | Gin                                        |
| **Database**       | PostgreSQL                                 |
| **Auth**           | JWT (access / refresh)                     |
| **Middleware**     | RateLimiter, Logger                        |
| **Config**         | Custom service config via `LocalConfig`    |
| **Sessions**       | Stored in PostgreSQL via `SessionsService` |
| **Docker Compose** | Used for container orchestration           |

---

## 🚀 API Endpoints / Эндпойнты

| Method | Path                | Description (EN)                                         | Описание (RU)                                  |
| ------ | ------------------- | -------------------------------------------------------- | ---------------------------------------------- |
| `GET`  | `/v1/status`        | Check service health                                     | Проверка статуса сервиса                       |
| `POST` | `/v1/auth/register` | Register new user (login, password, email, client, role) | Регистрация нового пользователя                |
| `POST` | `/v1/auth/login`    | Authenticate user and receive tokens                     | Авторизация и получение пары токенов           |
| `POST` | `/v1/auth/info`     | Decode access token and return user info                 | Получение данных пользователя по access токену |
| `POST` | `/v1/auth/refresh`  | Refresh tokens (new access + refresh)                    | Обновление пары токенов                        |
| `POST` | `/v1/auth/logout`   | Remove session and revoke refresh token                  | Выход и завершение сессии                      |

---

## 🔒 Authentication Flow / Процесс авторизации

1. **Register** — пользователь отправляет `login`, `password`, `email`, `client`, `role`.
   → получает `access_token` и `refresh_token`.

2. **Login** — вводит `login` и `password`.
   → снова получает пару токенов.

3. **Access Token** — короткоживущий (обычно 15 мин), используется для API-запросов.
   **Refresh Token** — живёт 7 дней, хранится в БД в таблице `sessions`.

4. **Refresh** — при истечении access токена отправляется запрос с refresh токеном.
   → получает новую пару токенов и старая сессия удаляется.

5. **Logout** — удаляет refresh токен из хранилища, полностью обнуляя сессию.

---

## ⚡ Quick Start / Быстрый старт

### Local

```bash
go run cmd/server.go
```

### Docker

```bash
docker-compose up --build
```

### Default URL

```
http://localhost:8080
```

---

## 📄 License / Лицензия

**MIT License** — свободное использование и модификация проекта.
