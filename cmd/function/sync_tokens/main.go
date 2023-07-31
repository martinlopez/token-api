package main

import (
	"context"
	"fmt"
	"log"
	taskPkg "tokens-api/cmd/function/sync_tokens/pkg"
	"tokens-api/internal/tokens"
	"tokens-api/pkg"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var err error
	err = godotenv.Load(".env")
	if err != nil {
		log.Default().Print(err)
	}

	ctx := context.Background()

	//db connection
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

	tokenUseCaseRW := tokens.NewUseCase(db)

	task := taskPkg.NewTask(tokenUseCaseRW, tokenUseCaseRW)
	err = task.Execute(ctx)
	if err != nil {
		log.Fatal(err)
	}
}
