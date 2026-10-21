# Запуск проекта

### Gateway

перейдити в папку `gateway`:

```bash
cd gateway
go run .
```
API
GET /api/budgets - Получение списка бюджетов

POST /api/budgets - Создание бюджета

GET /api/transactions - Получение списка транзакций

POST /api/transactions - Создание транзакции

```curl http://localhost:8080/api/budgets```

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