# SQL-CRUD – CLI Tool for Golang Backend Projects

SQL-CRUD is a command-line tool designed to structure Golang backend projects with GORM and Echo for SQL-based applications. It also generates OpenAPI documentation using Golang’s Swag.

## Installation

To install SQL-CRUD, run:

```bash
go install github.com/bushubdegefu/sql-crud@latest
```

## Quick Start

To get started quickly, you can use the following commands:

```bash
sql-crud-sample/
├── manager/
│   ├── [echo.go](http://_vscodecontentref_/0)
│   ├── consumer.go
│   ├── [manager.go](http://_vscodecontentref_/1)
│   └── migrate.go
├── helper/
├── messages/
├── observe/
├── configs/
├── database/
├── bluetasks/
├── testsetting/
├── tests/
└── controllers/


mkdir sql-crud-rest
cd sql-crud-rest
sql-crud init --name=github.com/username/sql-crud-rest
```

You can place this in your `~/.bashrc` or `~/.zshrc`:

```bash
export GOPROXY=https://proxy.golang.org,direct
```