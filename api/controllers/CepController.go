package controllers

import (
	"api/api/clients"
	"api/api/entities"
	"api/infra/config/cache"
	"api/infra/config/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	cepCache = cache.GetCache()
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

	logger.Info("Looking for cep in local: " + cep)
	
	var cachedData, exists = cepCache[cep]

	if exists && !cachedData.Expired {
		ctx.JSON(http.StatusOK, cachedData.Data)
		return
	}

	logger.Info("Not found in local, looking for in viacep: " + cep)
	
	cepResponse, err := clients.Get(cep)
	
	if err != nil {
		logger.Error("Error getting cep from viacep: " + err.Error())
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

	logger.Info("Found in viacep: " + cep)
	if cachedData.Expired {
		cache.UpdateCache(cep, newCep)
	} else {
		cache.SetCache(cep, newCep)
	}

	ctx.JSON(http.StatusOK, newCep)
}