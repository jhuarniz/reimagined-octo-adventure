We use: https://github.com/cosmtrek/air for hot reloading

# Installs
go get github.com/gofiber/fiber/v2
go get github.com/joho/godotenv
go get github.com/go-playground/validator/v10
go get github.com/thoas/go-funk
go get github.com/satori/go.uuid
go get github.com/jinzhu/copier
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres

# Run with go
go run .

# Run with air
# once (create a .air.toml config file)
air init
# for init your app
air
# Command to clean dependecies 
go mod tidy