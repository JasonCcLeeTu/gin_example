package main

import (
	handle "gin/internal/delivery"
	"log"

	envconfig "gin/internal/pkg/env"
	repopostgres "gin/internal/repository/db/postgres"
	usecaseusr "gin/internal/usecase/user"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @BasePath /v1
// @host localhost:8070
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("load env error:%s \n", err.Error())
		return
	}

	envConfig := envconfig.GetEnv()
	//log.Printf("%+v \n", envConfig.PgConfig)
	repoPg, err := repopostgres.NewRepoPostgres(envConfig.PgConfig.User, envConfig.PgConfig.Pwd, envConfig.PgConfig.DB,
		envConfig.PgConfig.Host, envConfig.PgConfig.Port, envConfig.PgConfig.Sslmode)

	if err != nil {
		log.Printf("初始化Postgres repo error:%s \n", err.Error())
		return
	}
	repoPostgres := usecaseusr.NewUsecaseUser(repoPg)

	engine := gin.Default()

	handle.NewRouter(engine, repoPostgres)

	engine.Run(":8070")
}
