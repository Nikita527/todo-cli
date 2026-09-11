# todo-cli

Учебный CLI для задач на Go. Хранение — JSON-файл. Парсинг аргументов: стандартный пакет `flag`.

## Схема аргументов

Флаги **перед** подкомандой:

```text
todo [флаги] <команда> [аргументы]
```

Пример:

```bash
go run ./cmd/todo -file data.json list
go run ./cmd/todo -file data.json add "купить молоко"
```

Вариант `todo list -file data.json` **не** поддерживается.

### Флаги

| Флаг | По умолчанию | Описание |
|------|--------------|----------|
| `-file` | `todos.json` | путь к JSON-хранилищу |

## Команды

| Команда | Пример | Поведение |
|---------|--------|-----------|
| `add` | `go run ./cmd/todo add "купить молоко"` | добавить задачу и сохранить |
| `list` | `go run ./cmd/todo list` | показать задачи (`id`, текст, `done`) |
| `done` | `go run ./cmd/todo done 1` | отметить выполненной и сохранить |
| `delete` | `go run ./cmd/todo delete 1` | удалить и сохранить |

## Запуск

Из корня репозитория:

```bash
go run ./cmd/todo list
go test ./...
```

## Структура

```text
cmd/todo/          # CLI (package main, flag)
internal/todo/     # store, Load/Save JSON
```

Формат файла (вариант B): `{"next": N, "tasks": [...]}`.
