package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abnerugeda/go-with-gin/controllers"
	"github.com/abnerugeda/go-with-gin/database"
	"github.com/abnerugeda/go-with-gin/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupRotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock() {
	aluno := models.Aluno{Nome: "Nome aluno teste", CPF: "12312312312", RG: "123123123"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)
}

func DeletaAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestVerificaStatusCodeSaudacoes(t *testing.T) {
	r := SetupRotasTeste()
	r.GET("/:nome", controllers.Saudacoes)

	req, _ := http.NewRequest("GET", "/abner", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")
	mockResposta := `{"API DIZ: ":"E ai abner, tudo bem?"}`
	respostaBody, _ := io.ReadAll(resposta.Body)
	assert.Equal(t, mockResposta, string(respostaBody))
}

func TestListandoAlunosHandler(t *testing.T) {
	database.ConnectDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos", controllers.FindAlunos)
	req, _ := http.NewRequest("GET", "/alunos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoCPF(t *testing.T) {
	database.ConnectDB()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupRotasTeste()
	r.GET("/alunos/cpf/:cpf", controllers.SearchAlunoByCPF)
	req, _ := http.NewRequest("GET", "/alunos/cpf/12312312312", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
