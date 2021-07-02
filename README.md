# Простое банковское обслуживание

Сервис, который мы собираемся создать, представляет собой простой банк. Он предоставит API-интерфейсы для веб-интерфейса, чтобы делать следующие вещи:

1. Создавайте и управляйте банковскими счетами, которые включают имя владельца, баланс и валюту.
2. Запишите все изменения баланса для каждой учетной записи. Таким образом, каждый раз, когда деньги добавляются или вычитаются из учетной записи, создается запись для входа в учетную запись.
3. Выполните перевод денег между 2-мя счетами. Это должно происходить в рамках транзакции, чтобы баланс обоих аккаунтов обновлялся успешно или ни один из них не обновлялся.

## Настроить локальную разработку

### Установить инструменты

- [Docker desktop](https://www.docker.com/products/docker-desktop)
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/)
- [Homebrew](https://brew.sh/)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
   scoop install migrate
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```go get
    go get github.com/kyleconroy/sqlc/cmd/sqlc
    ```

### Настроить инфраструктуру

- Запустить контейнер postgres:

    ```bash
    make postgres
    ```

- Создать базу данных simple_bank:

    ```bash
    make createdb
    ```

-Запустите миграцию базы данных для всех версий:

    ```bash
    make migrateup
    ```

- Запустите миграцию базы данных до 1 версии:

    ```bash
    make migrateup1
    ```

- Запустите миграцию базы данных на все версии:

    ```bash
    make migratedown
    ```

- Запустите миграцию базы данных на 1 версию:

    ```bash
    make migratedown1
    ```

### Как сгенерировать код

- Создание SQL CRUD с помощью sqlc:

    ```bash
    make sqlc
    ```

- Создать макет БД с помощью gomock:

    ```bash
    make mock
    ```

-Создать новую миграцию БД:

    ```bash
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```

### How to run

- Run server:

    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```
