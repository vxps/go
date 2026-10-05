# Запуск проекта

### Gateway

перейдити в папку `gateway`:

```bash
cd gateway
go run .
```

Gateway запустится на порту `8080`.

Проверка:

```bash
curl http://localhost:8080/ping
```

Ожидаемый ответ:

```text
pong
```

### Ledger

Откройте второй терминал и перейдите в папку `ledger`:

```bash
cd ledger
go run .
```

Ledger загрузит бюджеты из `budgets.json`, добавит тестовые транзакции из `main.go` и выведет результат в консоль.

Исходные данные для проверки находятся в:

- `ledger/budgets.json`
- `ledger/main.go`