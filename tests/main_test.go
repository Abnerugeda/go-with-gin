package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/abnerugeda/go-with-gin/controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupRotasTeste() *gin.Engine {
	rotas := gin.Default()
	return rotas
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
