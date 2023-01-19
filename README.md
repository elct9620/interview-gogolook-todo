(Interview) Gogolook Todo
===

## RESTful API

### `GET /tasks`

List all todo items

### `POST /tasks`

Create a new task

```json
{"name": "New Item"}
```

### `PUT /tasks/:id`

Update an exists task

```json
{"name": "Rename", "status": 1}
```

> To ensure the behavior is simple, it will create a new one if not exists.

### `DELETE /tasks/:id`

Delete a task

## Limitations

* Not thread-safe
* No persist store
