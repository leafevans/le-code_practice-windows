# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build / Run / Develop

```bash
# Run the server (reads conf/app.ini for MySQL + Redis config)
go run main.go

# Build a binary
go build -o ./tmp/main.exe .

# Live-reload during development (requires air)
air

# Run a specific test file
go test ./models/...

# Verify all packages compile
go build ./...
```

The project uses [air](https://github.com/air-verse/air) for live reload (config in `.air.toml`). When `air` is running, it watches `.go`, `.html`, `.tpl`, and `.tmpl` files and rebuilds on change.

## Architecture

This is a **Go web application** built with **Gin** + **GORM** (MySQL) + **Redis sessions**, following an MVC-ish structure.

### Request flow

```
main.go
  ├── loads conf/app.ini (via models/db.go init())
  ├── sets up Gin engine with template funcs, session middleware (Redis)
  ├── registers route groups: WebRoutersInit → APIRoutersInit → AdminRoutersInit
  └── r.Run()
```

### Directory layout

| Directory | Purpose |
|---|---|
| `routers/` | Route registration, split into 3 groups: **web** (`/`), **api** (`/api`), **admin** (`/admin`). Each group applies `middlewares.InitMiddleware`. |
| `controller/` | Request handlers. **admin** controllers all embed `BaseController`. **web** controllers handle public pages and sessions/cookies. |
| `models/` | GORM model structs + DB init (`db.go`) + shared helpers (`tools.go`). The `init()` in `db.go` opens the MySQL connection and stores it in the package-level `models.DB` variable. |
| `middlewares/` | Custom Gin middleware. `InitMiddleware` runs on every route group and sets a `"username"` context value. |
| `templates/` | Go `html/template` files, loaded via `r.LoadHTMLGlob("templates/**/*")`. Templates are defined with `{{define "name"}}`. |
| `static/` | Static assets (CSS, images, uploaded files). File uploads from `DoAdd` handler land here under `static/upload/`. |
| `conf/` | INI configuration (MySQL credentials, Redis host). Parsed by `gopkg.in/ini.v1`. |

### Key dependencies

- **Web framework**: `github.com/gin-gonic/gin`
- **ORM**: `gorm.io/gorm` + `gorm.io/driver/mysql`
- **Sessions**: `github.com/gin-contrib/sessions` with Redis backend
- **Logging**: `go.uber.org/zap` (global logger via `zap.ReplaceGlobals`)
- **Config**: `gopkg.in/ini.v1`

### Model relationships (GORM)

- `Article` belongs to `ArticleCate` via `CateID` foreign key; `ArticleCate` has many `Articles`
- `Student` ↔ `Lesson` is a many-to-many relationship through the `lesson_students` join table
- All models define explicit table names via `TableName()` method
- `Article.TableName()` is a package-level function (not a method) — this may be unintentional

### Template functions

Registered in `main.go`:
- `UnixToTime(int64) string` — converts Unix timestamp to `"2006-01-02 15:04:05"` format

### Session & auth

Sessions use Redis via `redis.NewStore` with a hardcoded secret. The `web.WebController.Index` sets `"username"` in the session. The `middlewares.InitMiddleware` sets a `"username"` Gin context value (separate from sessions).

### main.go history

The bulk of `main.go` (~800 lines) is commented-out code showing the progression from basic Gin routing → templates → middleware → sessions → config parsing. The active code is the final ~40 lines at the bottom.
