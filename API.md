# Score API

Base URL for local development:

```text
http://localhost:8080
```

Set `PORT` in `.env` if you want to run the server on another port.

## Environment Setup

Create a `.env` file in the project root:

```env
PORT=8080
DB_CONNECTION=user:password@tcp(127.0.0.1:3306)/score?charset=utf8mb4&parseTime=True&loc=Local
```

The app uses MySQL through GORM. On startup, it connects using `DB_CONNECTION`. On the first request, the app auto-creates the `formats` and `teams` tables when they are missing.

## Response Shapes

Successful list response:

```json
{
  "total": 1,
  "data": []
}
```

Error response:

```json
{
  "code": "SC01",
  "message": "error details"
}
```

## Created APIs

### GET /formats

Returns all game formats stored in the database.

Example response:

```json
{
  "total": 2,
  "data": [
    {
      "entityId": 1,
      "name": "T20"
    },
    {
      "entityId": 2,
      "name": "ODI"
    }
  ]
}
```

Current table/model fields:

| Field | Type | Notes |
| --- | --- | --- |
| `entityId` | number | Primary key |
| `name` | string | Unique format name |

### GET /teams

Returns all teams stored in the database, including their related format.

Example response:

```json
{
  "total": 1,
  "data": [
    {
      "entity_id": 1,
      "name": "India",
      "logo": "https://example.com/india.png",
      "format": {
        "name": "ODI"
      }
    }
  ]
}
```

Current table/model fields:

| Field | Type | Notes |
| --- | --- | --- |
| `entity_id` | number | Primary key |
| `name` | string | Unique team name |
| `logo` | string | Logo URL or path |
| `format` | object | Loaded from the related `formats` row |

### GET /tournaments

Returns a temporary hard-coded tournament list. This API does not use the database yet.

Example response:

```json
[
  {
    "name": "IPL",
    "format": "T20",
    "league": "Domestic"
  },
  {
    "name": "World Cup",
    "format": "ODI",
    "league": "International"
  }
]
```

## What To Do Next

1. Create the MySQL database, for example `score`.
2. Add the `.env` file with `PORT` and `DB_CONNECTION`.
3. Run the app with:

```bash
go run ./src
```

4. Call `GET /formats` once so the middleware can create missing tables.
5. Insert sample rows into `formats` and `teams`, or build POST APIs next.
6. Test the current APIs:

```bash
curl http://localhost:8080/formats
curl http://localhost:8080/teams
curl http://localhost:8080/tournaments
```

## Remaining APIs To Build

Recommended next APIs:

| Method | Path | Purpose |
| --- | --- | --- |
| `POST` | `/formats` | Create a format |
| `GET` | `/formats/:id` | Get one format |
| `PUT` | `/formats/:id` | Update a format |
| `DELETE` | `/formats/:id` | Delete a format |
| `POST` | `/teams` | Create a team |
| `GET` | `/teams/:id` | Get one team |
| `PUT` | `/teams/:id` | Update a team |
| `DELETE` | `/teams/:id` | Delete a team |
| `POST` | `/tournaments` | Create a tournament |
| `GET` | `/tournaments/:id` | Get one tournament |
| `PUT` | `/tournaments/:id` | Update a tournament |
| `DELETE` | `/tournaments/:id` | Delete a tournament |
| `POST` | `/matches` | Create a match |
| `PUT` | `/matches/:id/score` | Update score |
| `GET` | `/matches/:id/score` | Get current score |

The highest-value next step is adding POST APIs for `formats` and `teams`, because the existing GET APIs need database records before they become useful.
