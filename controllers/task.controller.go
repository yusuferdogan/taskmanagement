package controllers

import (
	"example/taskmanagement/models"
	"example/taskmanagement/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TaskController struct {
	TaskService services.TaskService
}

func New(taskservice services.TaskService) TaskController {
	return TaskController{
		TaskService: taskservice,
	}
}

func (tc *TaskController) CreateTask(ctx *gin.Context) {
	var task models.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err := tc.TaskService.CreateTask(&task)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) GetTask(ctx *gin.Context) {
	id := ctx.Param("id")
	task, err := tc.TaskService.GetTask(&id)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.IndentedJSON(http.StatusOK, task)

}

func (tc *TaskController) GetAll(ctx *gin.Context) {
	tasks, err := tc.TaskService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, tasks)
}

func (tc *TaskController) UpdateTask(ctx *gin.Context) {
	var task models.Task

	if err := ctx.ShouldBindJSON(&task); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := tc.TaskService.UpdateTask(&task)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, task)
}

func (tc *TaskController) DeleteTask(ctx *gin.Context){
	id := ctx.Param("id")
	err := tc.TaskService.DeleteTask(&id)
	if err != nil{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message" : "successed"})
}

func (tc *TaskController) RegisterTaskRoute(rg *gin.RouterGroup){
	taskroute := rg.Group("/task")
	taskroute.POST("/create", tc.CreateTask)
	taskroute.GET("/task/:id", tc.GetTask)
	taskroute.GET("/tasks", tc.GetAll)
	taskroute.PATCH("/update", tc.UpdateTask)
	taskroute.DELETE("/delete/:id", tc.DeleteTask)
}