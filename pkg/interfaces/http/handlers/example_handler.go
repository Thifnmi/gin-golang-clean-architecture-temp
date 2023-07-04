package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"github.com/thifnmi/gin-golang-clean-architecture-temp/pkg/domain"
	"github.com/thifnmi/gin-golang-clean-architecture-temp/pkg/usecases"
)

type ExampleHandler interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
}

type exampleHandler struct {
	exampleUsecase usecases.ExampleUsecase
}

func NewExampleHandler(exampleUsecase usecases.ExampleUsecase) ExampleHandler {
	return &exampleHandler{
		exampleUsecase,
	}
}

func (lh *exampleHandler) GetAll(c *gin.Context) {
	var query domain.ExampleQuery
	err := c.BindQuery(&query)
	if err != nil {
		fmt.Printf("err %v'n", err)
	}
	examples, err := lh.exampleUsecase.GetAll(c, &query)
	if err != nil {
		fmt.Printf("err %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, examples)
}

func (lh *exampleHandler) GetByID(c *gin.Context) {
	id, err := uuid.FromString(c.Param("uuid"))
	if err != nil {
		fmt.Printf("err %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	example, err := lh.exampleUsecase.GetByID(c, id)
	if err != nil {
		fmt.Printf("err %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, example)
}
