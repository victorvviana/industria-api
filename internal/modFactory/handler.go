package modfactory

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegistrarRotas(r *gin.RouterGroup, service ProducaoService) {
	r.POST("/produtos", func(c *gin.Context) {
		var p Produto
		if err := c.ShouldBindJSON(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}

		id, err := service.CadastrarProduto(p)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	})

	r.POST("/ordens-producao", func(c *gin.Context) {
		var input CriarOPInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}

		id, err := service.CriarOrdemProducao(input)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"erro": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"id": id})
	})
}
