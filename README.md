# go-restful-server-template

## Dependencies

```
<!-- Database -->
gorm.io/gorm
gorm.io/driver/postgres

<!-- API -->
github.com/gin-gonic/gin
github.com/gin-contrib/cors

<!-- API (Auth) -->
github.com/dgrijalva/jwt-go

<!-- Others -->
github.com/joho/godotenv
go.uber.org/zap
github.com/rs/xid
```

## Prepare

- `.env`

  ```
  ENV="development"
  HOST=0.0.0.0
  PORT=8080
  ACCESS_ALLOW_ORIGIN="*"
  DATABASE_DSN="host=<host> port=<port> user=<user> password=<password> dbname=<dbname> sslmode=disable"
  LOG_FOLDER_PATH=<local_path>

  # Not necessary
  DISABLE_LOG=true
  ```
