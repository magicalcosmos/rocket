package main

import (
	"log"
	"net/http"
	"rocket-refactor-with-DDD/application"
	"rocket-refactor-with-DDD/domain"
	"rocket-refactor-with-DDD/infrastructure-mongo"
	"rocket-refactor-with-DDD/interfaces/restapi"
)

func main() {
	// 初始化 MongoDB 仓储
	mongoRepo, err := infrastructure.NewMongoRocketRepository("mongodb://localhost:27017", "rocketsdb", "rockets")
	if err != nil {
		log.Fatal(err)
	}

	// 初始化领域服务
	rocketService := domain.NewRocketService(mongoRepo)

	// 初始化应用服务
	appService := application.NewRocketAppService(rocketService)

	// 初始化 API 处理器
	rocketHandler := restapi.NewRocketHandler(appService)

	// 设置路由
	http.HandleFunc("/rockets/update", rocketHandler.UpdateRocketCapacity)

	// 启动服务器
	log.Fatal(http.ListenAndServe(":8080", nil))
}
