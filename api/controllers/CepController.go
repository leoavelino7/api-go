package controllers

import (
	"api/api/clients"
	"api/api/entities"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	cepCache = make(map[string]entities.Cep)
	mutex sync.RWMutex
)

type cepController struct {
	ceps []entities.Cep
}

func NewCepController() *cepController {
	return &cepController{}
}

func (value *cepController) FindAll(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, value.ceps)
}	

func (value *cepController) FindByCep(ctx *gin.Context) {
	cep := ctx.Param("cep")

	fmt.Println("Looking for cep in local: ", cep)

	mutex.RLock()
	
	if cachedData, exists := cepCache[cep]; exists {
		mutex.RUnlock()
		ctx.JSON(http.StatusOK, cachedData)
		return
	}
	mutex.RUnlock()

	fmt.Println("Not found in local, looking for in viacep: ", cep)
	
	cepResponse, err := clients.Get(cep)
	
	if err != nil {
		fmt.Println(err.Error())
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "CEP not found",
		})
		return
	}

	newCep := entities.NewCep()
	
	newCep.Cep = cepResponse.Cep
	newCep.Street = cepResponse.Street
	newCep.Complement = cepResponse.Complement
	newCep.Neighborhood = cepResponse.Neighborhood
	newCep.City = cepResponse.City
	newCep.State = cepResponse.State

	value.ceps = append(value.ceps, *newCep)

	fmt.Println("Found in viacep: ", cep)
	mutex.Lock()
	cepCache[cep] = *newCep
	mutex.Unlock()

	ctx.JSON(http.StatusOK, newCep)
}