package actions

import (
	"go-mongo/controllers"
	"go-mongo/infrastructures"
	"go-mongo/repositories"
	"go-mongo/routes"
	"go-mongo/services"

	config "github.com/spf13/viper"
)

// StartServer start http serve
func StartServer() {
	// mongo connections
	conn := infrastructures.MongoConfig{
		Host:         config.GetString("database.mongo.host"),
		Port:         config.GetInt("database.mongo.port"),
		DatabaseName: config.GetString("database.mongo.database_name"),
		DBTimeout:    config.GetInt("database.mongo.timeout"),
		User:         config.GetString("database.mongo.username"),
		Password:     config.GetString("database.mongo.password"),
	}
	mongoConnect := infrastructures.MongoOpen(conn)

	// dependencies injection
	commentRepository := new(repositories.CommentRepository)
	commentRepository.MongoConnect = mongoConnect

	commentService := new(services.CommentService)
	commentService.CommentRepository = commentRepository

	memberRepository := new(repositories.MemberRepository)
	memberRepository.MongoConnect = mongoConnect

	memberService := new(services.MemberService)
	memberService.MemberRepository = memberRepository

	commentController := new(controllers.CommentController)
	commentController.CommentService = commentService

	memberController := new(controllers.MemberController)
	memberController.MemberService = memberService

	// route & run serve
	route := new(routes.Route)
	route.CommentController = commentController
	route.MemberController = memberController
	infrastructures.ServerListen(route.GetRoute())
}
