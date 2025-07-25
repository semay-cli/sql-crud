# SQL-CRUD – CLI Tool for Golang Backend Projects

**SQL-CRUD** is a powerful command-line tool that streamlines the setup of Golang backend projects using **GORM** and **Echo** for SQL-based applications. It automates the generation of boilerplate code, caching logic, and documentation, significantly accelerating development.

## ✨ Features

- Generates **GORM models**, **services**, and **controllers** based on metadata provided in a `config.json` file.
- Automatically creates **RESTful CRUD APIs** for each model.
- Integrates with the **Echo** framework for routing and middleware.
- Adds **in-memory caching** using [Ristretto](https://github.com/dgraph-io/ristretto) for enhanced performance.
- Generates **OpenAPI (Swagger) documentation** using [Swag](https://github.com/swaggo/swag).
- Simplifies project scaffolding and enforces consistent architecture.
- Currently supports **MySQL**, **SQLite**, and **PostgreSQL** databases.

> ⚠️ Note: Models with **JSON datatypes** are not fully supported in service layer generation. The tool can generate the base model, but you'll need to manually handle business logic in the generated services.

## 📦 Use Case

Ideal for developers looking to:

- Quickly bootstrap a Golang backend project with a standardized structure.
- Reduce manual setup time for CRUD operations and documentation.
- Leverage caching for improved read performance.
- Enforce separation of concerns with clean service and controller layers.


## Installation

To install SQL-CRUD, run:

```bash
go install github.com/semay-cli/sql-crud@latest
```

## 📁 Project Structure (Generated)

```bash
sql-crud-sample/
├── manager/
│   ├── app.go
│   ├── middleware.go
│   ├── [manager.go](http://_vscodecontentref_/1)
│   └── migrate.go
├── blue_admin/
│   ├── scheduler
│   ├── controllers
│   ├── models
│   ├── services
│   ├── services
│   ├── testsetting/
│   ├── middleware.go
│   └── setup.go
├── observe/
├── configs/
│   ├── .dev.env
│   └── .env
├── logs/
├── database/
├── project.json
└── main.go

## 📁 Folder Structure Explanation

### `manager/`
Contains application-level setup and utility functions.
- **`app.go`** – Initializes core application components like database, cache, and router.
- **`middleware.go`** – Global middleware definitions (e.g., CORS, logging).
- **`manager.go`** – Handles app-wide logic such as dependency injection and route group registration.
- **`migrate.go`** – Runs GORM-based database migrations.

### `blue_admin/`
This directory represents a self-contained application module — in this case, "blue_admin". It contains all the business logic specific to that app, and follows a clean modular structure. Each application (like `blue_admin`, `crm`, etc.) can maintain its own version of this layout.

If you add another app such as `crm/`, it will mirror this structure and share the rest of the project components (like `manager/`, `configs/`, `database/`, etc.).

**Shared Structure Across Multiple Apps:**
Each app (`blue_admin/`, `crm/`, etc.) will typically include:

- **`scheduler/`** – Scheduled tasks (e.g., cron jobs) specific to the app.
- **`controllers/`** – HTTP handler functions generated from the app\'s models.
- **`models/`** – GORM model definitions based on `project.json` or scoped configs.
- **`services/`** – Business logic layer, caching via Ristretto, and DB access.
- **`testsetting/`** – App-specific test configurations or fixtures.
- **`middleware.go`** – Middlewares that apply to this specific app/module.
- **`setup.go`** – Initialization logic to wire up this app into the main application (register routes, dependencies, etc.).

> ✅ This modular approach makes it easy to manage and scale multiple apps in a monorepo-style project, while keeping shared concerns centralized (e.g., in `manager/`, `configs/`, `database/`, etc.).

### `observe/`
Handles observability features for the entire application, including monitoring, metrics, and tracing integrations.

- **`/metrics` endpoint** – Automatically exposed for **Prometheus** scraping. This provides runtime metrics such as HTTP request counts, response durations, memory usage, and more.
- **Tracing with Jaeger** – If tracing is initialized, the app is configured to:
  - Export spans to **Jaeger**.
  - **Only trace failed requests** by default (non-2xx responses).
  - Sample **10%** of the traffic using the default configuration.

> 💡 You can adjust the trace sampling rate using the `TRACER_SAMPLE` environment variable. If `TRACER_SAMPLE` is not defined, a 10% default sampling rate is used.

> ✅ This module ensures observability is centralized and configurable, without polluting business logic or service layers.


### `configs/`
Houses environment-specific configuration files used to initialize the application with the correct settings (e.g., database URLs, ports, caching parameters, etc.).

- **`.env`** – Base configuration that applies across all environments (common defaults).
- **`.dev.env`** – Environment-specific overrides for development.

> 🔄 The `.dev.env` file naming is **dynamic** based on the environment. For example:
> - `.staging.env`
> - `.production.env`
> - `.test.env`

These files follow the pattern `.<env>.env`, and the active environment is determined by the `--env` flag when starting your application.

### ✅ Example:
```bash
go run main.go  run --env=staging


### `logs/`
Directory where application logs are written (if logging to file is enabled).

### `database/`
Contains the logic for establishing and configuring database connections using **GORM**, with support for **MySQL**, **PostgreSQL**, and **SQLite**. It also integrates GORM with **OpenTelemetry** for tracing and uses structured logging for better observability.

#### 🔍 Summary of Responsibilities:
- **Dynamic Driver Selection**: Based on the environment variable `<APP_NAME>_DB_TYPE`, it dynamically selects and configures the appropriate GORM driver (`mysql`, `postgres`, or `sqlite`).
- **Connection Strings**: Reads DSN (connection string) from env variables like `<APP_NAME>_MYSQL_URI`, `<APP_NAME>_POSTGRES_URI`, or `<APP_NAME>_SQLLITE_URI` based on the selected DB type.
- **Custom GORM Logger**:
  - Logs slow queries.
  - Outputs logs to a file named `<app_name>_gorm.log`.
  - Uses `logger.Info` level with timestamped entries.
- **Connection Pool Configuration**:
  - Sets maximum open connections and connection lifetime for each DB type to optimize performance.
- **OpenTelemetry Tracing**:
  - Attaches the `otel` tracing plugin to the GORM session to support distributed tracing.
  - Helps trace slow or failing database calls when observability is enabled.

> ✅ The logic is designed to support multi-tenant or multi-app setups by using the app name as a prefix for environment variable lookup, making it scalable and modular.

### `project.json`
This file serves as the **main input metadata** for the SQL-CRUD CLI tool. It defines project-wide configuration and tracks individual application modules for code generation.

- The file is **automatically generated** by the CLI.
- It keeps track of the project name, registered apps, and auth configuration.
- If `auth_app_name` is not provided, it **defaults to `blue-admin`**.
- If `auth_app_type` is not specified, it **defaults to `sso`** (Single Sign-On).
- The `models` field is populated with your application\'s data models and is used as the source of truth for generating GORM models, services, and controllers.

#### 📄 Sample `project.json`

```json
{
  "project_name": "github.com/bushubdegefu/sql-play",
  "app_names": [
    "blue-admin"
  ],
  "current_app_name": "",
  "back_tick": "`",
  "auth_app_name": "blue-admin",
  "auth_app_type": "sso" 
}

### `config.json`
This file defines **model metadata** used by the SQL-CRUD CLI tool to auto-generate:

- GORM models
- RESTful controllers
- Service layers
- OpenAPI documentation

Each application module (e.g., `blue-admin`, `crm`) has its own `config.json` file that is **automatically generated** when the app is initialized.

#### 🔧 Key Notes:
- `app_name`: The name of the app the config belongs to.
- `project_name`: The Go module path (used for import paths).
- `models`: An array of model definitions, each with fields, types, relationships, and CRUD flags.
- This metadata drives **code scaffolding** (models, services, routes, docs).

---

#### 📄 Sample `config.json`
```json
{
  "project_name": "github.com/bushubdegefu/sql-play",
  "app_name": "blue-admin",
  "models": [
    {
      "name": "Group",
      "search_fields": ["name", "description", "active"],
      "rln_model": ["User$mtm$user_groups", "Scope$mtm$group_scopes"],
      "fields": [
        {
          "name": "ID",
          "type": "uint",
          "annotation": "gorm:\"primaryKey;autoIncrement:true\" json:\"id,omitempty\"",
          "curd_flag": "true$false$false$true$false$false"
        },
        {
          "name": "Name",
          "type": "string",
          "annotation": "gorm:\"not null; unique;\" json:\"name,omitempty\"",
          "curd_flag": "true$true$true$true$false$false"
        }
        // ...more fields...
      ]
    }
    // ...more models like App, User, Scope, Resource, JWTSalt...
  ]
}


## 🚀 Quick Start

Follow the steps below to get your project up and running:

### 1. Create a Project Directory

```bash
mkdir <folder-name>
```

### 2. Install `sql-crud` CLI

```bash
sql-crud-install -u <github-username> -a blue-admin -p <folder-name>
```

> **Note:** To use the admin UI that ships with `sql-crud`, you must set the app name to `blue-admin`, as it includes the built-in role management interface.

### 3. Configure the Database URI

Edit the `.dev.env` file in the project root and set the SQL database URI:

```env
DATABASE_URI=your-database-uri-here
```

### 4. Run Migrations

```bash
go run main.go migrate -t create -e dev
```

> **Note for Windows users using SQLite:**  
Before running the command above, set the environment variable:

```bash
set CGO_ENABLED=1
```

### 5. Create a Superuser

```bash
go run main.go superuser
```

This will create a default superuser with the following credentials:

- **Email:** `superuser@mail.com`
- **Password:** `default@123`

### 6. Start the Development Server

```bash
go run main.go run --env=dev
```


## 7. 🖥️ Accessing the UI

- The **Admin UI** is available at:  
  `http://localhost:<port>/admin/`

- The **Swagger Documentation UI** can be found at:  
  `http://localhost:<port>/<app_name>/docs/`

> **Note:**  
> If your app name contains hyphens (`-`), they will be converted to underscores (`_`) in the Swagger docs URL.


## 📝 Notes and Pitfalls

1. **App Cache Instance**  
   Each app has its own cache instance. The default cache uses Resiretto for local management.  
   You should carefully consider the total cache size you want for your global app environment, then divide that size equally among the number of apps. This ensures fair and efficient cache allocation.

2. **Adding New Apps**  
   Every time you add a new app, you must run the configuration and global echo app generation commands again. This step is necessary to register the new app correctly.

3. **File Overwrite Warning**  
   If a file with the same name already exists, running the `sql-crud` generation commands will completely replace the existing file.  
   Be cautious when running these commands on heavily customized files to avoid losing your changes.

4. **Unit Tests Coverage**  
   The current generated unit tests cover CRUD functionality only.  
   Unit test generation for relationships between entities is planned and will be coming soon.

## ℹ️ Additional Information

For more details and to explore all available commands, you can always run:

```bash
sql-crud --help
```
---

## 📄 License

This project is licensed under the [MIT License](https://opensource.org/licenses/MIT).
