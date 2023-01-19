(Interview) Gogolook Todo
===
[![Go](https://github.com/elct9620/interview-gogolook-todo/actions/workflows/go.yml/badge.svg)](https://github.com/elct9620/interview-gogolook-todo/actions/workflows/go.yml)

## RESTful API

### `GET /tasks`

List all todo items

#### Response Body

```json
{"result": [{"id": 1, "name": "Todo Item", "status": 0}]}
```

### `POST /tasks`

Create a new task

#### Request Body

```json
{"name": "New Item"}
```

#### Response Body

```json
{"result": {"id": 1, "name": "New Item", "status": 0}}
```

### `PUT /tasks/:id`

Update an exists task

#### Request Body

```json
{"name": "Rename", "status": 1}
```
#### Response Body

```json
{"result": {"id": 1, "name": "Rename", "status": 1}}
```

> To ensure the behavior is simple, it will create a new one if not exists.

### `DELETE /tasks/:id`

Delete a task

## Limitations

* Not thread-safe
* No persist store
