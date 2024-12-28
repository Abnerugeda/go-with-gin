package routes

import (
	"github.com/abnerugeda/go-with-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/")
	r.GET("/alunos", controllers.FindAlunos)
	r.GET("/alunos/:id", controllers.FindOneAluno)
	r.GET("/alunos/cpf/:cpf", controllers.SearchAlunoByCPF)
	r.DELETE("/alunos/:id", controllers.DeleteAluno)
	r.PATCH("/alunos/:id", controllers.EditAluno)
	r.POST("/alunos", controllers.CreateAluno)
	r.GET("/:nome", controllers.Saudacoes)

	r.Run()
}
