# Тестовое задание на позицию Junior Backend Developer 

## Микросервис, предоставляющий API для управления списком заметок.

**Решение:**

**Конфигурирование:**

Настройка реквизитов доступа к СУБД осуществляется через переменные окружения. Чтобы их задать, необходимо создать и заполнить `.env` файл. Пример его заполнения в файле `.env.example`. 

**Развёртывание:**

Запуск сервера и СУБД:
```shell
make run
```
**Примеры:**

***Добавление заметки***

```shell
curl -X POST http://localhost:8080/notes -H "Content-Type: application/json" -d "{\"content\":\"This a very first note\"}"
```

```json
{
  "id": 1
}
```


```shell
curl -X POST http://localhost:8080/notes -H "Content-Type: application/json" -d "{\"content\":\"Hi, Sports. It's rainy outside\"}"
```

```json
{
  "id": 2
}
```


```shell
curl -X POST http://localhost:8080/notes -H "Content-Type: application/json" -d "{\"content\":\"Gnabry's better than Messi\"}"
```

```json
{
  "id": 3
}
```


```shell
curl -X POST http://localhost:8080/notes -H "Content-Type: application/json" -d "{\"content\":\"Earth's population = 7 million people\"}"
```

```json
{
  "id": 4
}
```

***Удаление заметки***

```shell
curl -X DELETE http://localhost:8080/notes/2 
```

```json
{
  "status": "deleted"
}
```

```shell
curl -X DELETE http://localhost:8080/notes/4
```

```json
{
  "status": "not deleted, note with id = 4 doesn't exist"
}
```

***Получение всех заметок***

```shell
curl -X GET http://localhost:8080/notes
```

```json
[
    {
        "id": 1,
        "content": "This a very first note"
    },
    {
        "id": 3,
        "content": "Gnabry's better than Messi"
    },
    {
        "id": 4,
        "content": "Earth's population = 7 million people"
    }
]
```
