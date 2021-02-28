package main

import (
	"fmt"
	"go-user/domain/repository"
	service2 "go-user/domain/service"
	"go-user/handler"
	user "go-user/proto/user"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/micro/go-micro/v2"
)

func main() {
	// 服务参数设置
	srv := micro.NewService(
		micro.Name("go-user.user"),
		micro.Version("latest"),
	)
	//初始化服务
	srv.Init()

	//创建数据库连接
	dsn := "root:FSD9di_dsG@tcp(127.0.0.1:3306)/go_micro?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接数据库失败: ", err)
	}

	rp := repository.NewUserRepository(db)
	//只执行一次
	//rp.InitTable()

	userDataService := service2.NewUserDataService(rp)

	// Register handler
	err = user.RegisterUserHandler(srv.Server(), &handler.User{UserDataService: userDataService})
	if err != nil {
		fmt.Println("注册handler失败: ", err)
	}

	// Run service
	if err = srv.Run(); err != nil {
		fmt.Println("运行失败: ", err)
	}
}
