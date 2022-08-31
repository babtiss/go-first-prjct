# Приложение трекер задач.

### О приложении:
- Состояние: В процессе разработки ...
- Планируется сделать простой трекер задач.
- Возможность создавать дедлайны
- Возможность пользователей в группы и создавать общие задачи

### Используемые утилиты:
- Работа с фреймворком <a href="https://github.com/gin-gonic/gin">gin-gonic/gin</a>.
- Конфигурация приложения с помощью библиотеки <a href="https://github.com/spf13/viper">spf13/viper</a>.
- Работа с БД, библиотека <a href="https://github.com/jmoiron/sqlx">sqlx</a>.

### Для запуска приложения:
```
docker-compose up --build go-application
```
Контейнер запустится но не успеет подключится к бд. Прописать ещё команду
```
docker-compose up go-application
```