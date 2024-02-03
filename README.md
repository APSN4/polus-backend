# polus-backend

### Stack
**Database**: _PostgreSQL_  
**Programming language**: _Go_  
**DI**: _google/wire_

### How to start

`go run main.go`

### Dependency
`go get -u github.com/antonfisher/nested-logrus-formatter`  
`go get -u github.com/gin-gonic/gin`  
`go get -u golang.org/x/crypto`  
`go get -u gorm.io/gorm`  
`go get -u gorm.io/driver/postgres`  
`go get -u github.com/sirupsen/logrus`  
`go get -u github.com/joho/godotenv`  

### ENV structure

```
PORT=8080
# Application
APPLICATION_NAME=polus-restful-api

# Database
DB_DSN="host=localhost user=root password=root dbname=polus-api port=5432"

# Logging
LOG_LEVEL=DEBUG
```

### Database structure

base_models  
```
CREATE TABLE IF NOT EXISTS base_models (
id SERIAL PRIMARY KEY NOT NULL,
created_at TIMESTAMPTZ DEFAULT current_timestamp,
updated_at TIMESTAMPTZ DEFAULT current_timestamp,
deleted_at TIMESTAMPTZ
);
```

roles  
```
CREATE TABLE IF NOT EXISTS roles (
id SERIAL PRIMARY KEY NOT NULL,
role VARCHAR(255),
base_model_id INT,
CONSTRAINT fk_base_model_role FOREIGN KEY (base_model_id) REFERENCES base_models(id),
created_at TIMESTAMPTZ DEFAULT current_timestamp,
updated_at TIMESTAMPTZ DEFAULT current_timestamp,
deleted_at TIMESTAMPTZ
);
```

users  
```
CREATE TABLE IF NOT EXISTS users (
id SERIAL PRIMARY KEY NOT NULL,
name VARCHAR(255),
email VARCHAR(255),
password VARCHAR(255) CHECK (password <> '') NOT NULL,
status BIGINT,
role_id INT NOT NULL,
base_model_id INT,
CONSTRAINT fk_base_model_user FOREIGN KEY (base_model_id) REFERENCES base_models(id),
CONSTRAINT fk_role_user FOREIGN KEY (role_id) REFERENCES roles(id),
created_at TIMESTAMPTZ DEFAULT current_timestamp,
updated_at TIMESTAMPTZ DEFAULT current_timestamp,
deleted_at TIMESTAMPTZ
);
```