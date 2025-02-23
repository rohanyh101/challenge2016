package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/roh4nyh/qube_challenge_2016/models"
	"github.com/roh4nyh/qube_challenge_2016/service"
)

var DistributorCollectionMap map[string]models.Distributor

func GetDistributors(c *gin.Context) {
	collection := service.GetDistributors()
	c.JSON(http.StatusOK, collection)
}

func GetDistributor(c *gin.Context) {
	distributorID := c.Param("distributor_id")
	distributor := service.GetDistributor(distributorID)
	c.JSON(http.StatusOK, distributor)
}

func AddDistributor(c *gin.Context) {
	var newDistributor models.NewDistributorCmd

	if err := c.BindJSON(&newDistributor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.AddDistributor(newDistributor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func CheckDistributorPermission(c *gin.Context) {
	distributorID := c.Param("distributor_id")
	var checkDistributorPermissionCmd models.CheckDistributorPermissionCmd

	if err := c.BindJSON(&checkDistributorPermissionCmd); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.CheckDistributorPermission(distributorID, checkDistributorPermissionCmd.Locations)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func UpdateDistributor(c *gin.Context) {
	distributorID := c.Param("distributor_id")
	var updateDistributor models.UpdateDistributorCmd

	if err := c.BindJSON(&updateDistributor); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := service.UpdateDistributor(distributorID, updateDistributor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
