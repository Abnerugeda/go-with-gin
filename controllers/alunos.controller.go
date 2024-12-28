package controllers

import (
	"net/http"

	"github.com/abnerugeda/go-with-gin/database"
	"github.com/abnerugeda/go-with-gin/models"
	"github.com/gin-gonic/gin"
)

func FindAlunos(c *gin.Context) {
	var alunos []models.Aluno

	database.DB.Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}

func FindOneAluno(c *gin.Context) {
	id := c.Params.ByName("id")

	var aluno models.Aluno
	database.DB.First(&aluno, id)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}
	c.JSON(http.StatusOK, aluno)

}

func Saudacoes(c *gin.Context) {
	nome := c.Params.ByName("nome")

	c.JSON(http.StatusOK, gin.H{
		"API DIZ: ": "E ai " + nome + ", tudo bem?",
	})
}

func CreateAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

func DeleteAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{
		"data": "Aluno deletado com sucesso",
	})
}

func EditAluno(c *gin.Context) {
	id := c.Params.ByName("id")
	var aluno models.Aluno
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	database.DB.Model(&aluno).UpdateColumns(aluno)

	c.JSON(http.StatusOK, aluno)
}

func SearchAlunoByCPF(c *gin.Context) {
	var alunos []models.Aluno

	cpf := c.Param("cpf")
	database.DB.Where("cpf LIKE ?", "%"+cpf+"%").Find(&alunos)
	c.JSON(http.StatusOK, alunos)
}
