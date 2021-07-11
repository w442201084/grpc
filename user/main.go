package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
	"github.com/w442201084/grpc/domain/repository"
	service2 "github.com/w442201084/grpc/domain/service"
	"github.com/w442201084/grpc/handler"
	pb "github.com/w442201084/grpc/proto/user"
	"log"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.first.grpc.user"),
		micro.Version("latest"),
	)
	srv.Init()
	// root:root@(localhost)/test?charset=utf8mb4&parseTime=True&loc=Local
	db ,err := gorm.Open("mysql" ,"root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if nil != err {
		log.Fatalln("gorm.Open error ..." , err)
	}
	defer db.Close()
	db.SingularTable(true)

	//rp := repository.NewUserRepository(db)
	//rp.InitTable()

	// Register handler
	pb.RegisterUserHandler(srv.Server(), new(handler.Users))

	userService := service2.NewUserService(repository.NewUserRepository(db))
	err = pb.RegisterUserHandler(srv.Server() , &handler.Users{
		UserService: userService,
	})
	if nil != err {
		log.Println("pb.RegisterUserHandler error ..." , err)
	}

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatalln(err)
	}
}
