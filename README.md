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

## Тестирование

Для запуска unit-тестов Ledger перейдите в папку `ledger`:

```bash
go test -v ./...
```

```
go test -count 1 -p 1 -coverprofile cover.out .
go tool cover -func cover.out
```

Ledger (бизнес-логика) — 69,0%


Ledger (весь модуль, включая CLI) — 33,6%

Для тестов Gateway `gateway`:

```bash
go test -count 1 -p 1 -coverpkg ./... -coverprofile cover.out ./test
go tool cover -func cover.out
go tool cover -html cover.out -o cover.html
```

**Покрытие Gateway: 86,2%.**

HTML-отчёт доступен в файле `gateway/cover.html` после запуска соответствующей команды.