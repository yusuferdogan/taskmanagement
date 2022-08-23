package main

import (
	"context"
	"example/taskmanagement/controllers"
	"example/taskmanagement/services"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	//"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.TaskService
	usercontroller controllers.TaskController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err := mongo.Connect(ctx, mongoconn)

	if err != nil {
		log.Fatal(err)
	}

	/*err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}*/

	fmt.Println("mongo connection has been established")

	usercollection = mongoclient.Database("userdb").Collection("users")
	userservice = services.NewTaskService(usercollection, ctx)
	usercontroller = controllers.New(userservice)
	server = gin.Default()

}

func main() {

	defer mongoclient.Disconnect(ctx)

	basepath := server.Group("/v1")
	usercontroller.RegisterTaskRoute(basepath)

	log.Fatal(server.Run(":9090"))
}
