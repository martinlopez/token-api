package main

import (
	"fmt"
	"log"
	"net/http"
	"tokens-api/internal/tokens"
	"tokens-api/pkg"
	"tokens-api/pkg/handler"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Default().Print(err)
	}

	//db connection
	dbSecrets, err := pkg.GetDBSecrets()
	if err != nil {
		log.Fatal(err)
	}
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/core?charset=utf8mb4&parseTime=True&loc=Local",
		dbSecrets.Username,
		dbSecrets.Password,
		dbSecrets.Host,
		dbSecrets.Port,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to database: ", err.Error())
		log.Fatal(err)
	}

	tokenUseCase := tokens.NewUseCase(db)
	tokenUseCaseHandler := handler.NewCrudHandler("tokens", tokenUseCase, tokenUseCase)
	tokenUseCaseHandler.Setup(r)

	r.Run()
}
