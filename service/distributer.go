package service

import (
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/roh4nyh/qube_challenge_2016/models"
	"github.com/roh4nyh/qube_challenge_2016/utils"
)

var DistributorCollection map[string]models.Distributor

func InitDistributorCollection() {
	DistributorCollection = make(map[string]models.Distributor)
}

func GetDistributors() map[string]models.Distributor {
	return DistributorCollection
}

func GetDistributor(distId string) models.Distributor {
	return DistributorCollection[distId]
}

func AddDistributor(req models.NewDistributorCmd) (models.Distributor, error) {

	log.Printf("req: %+v", req)
	isValidInclude := utils.ValidateIncludeExclude(req.Includes)
	if !isValidInclude {
		return models.Distributor{}, fmt.Errorf("invalid include")
	}

	isValidExclude := utils.ValidateIncludeExclude(req.Excludes)
	if !isValidExclude {
		return models.Distributor{}, fmt.Errorf("invalid exclude")
	}

	if req.Includes == nil {
		req.Includes = []string{}
	}

	if req.Excludes == nil {
		req.Excludes = []string{}
	}

	distributor := models.Distributor{
		DistributorID:     uuid.New().String()[0:8],
		Name:              req.Name,
		Includes:          req.Includes,
		Excludes:          req.Excludes,
		ParentDistributor: req.ParentDistributorID,
		Level: func() int {
			if req.ParentDistributorID == "" {
				return 1
			}

			parentDistributor, ok := DistributorCollection[req.ParentDistributorID]
			if !ok {
				return 1
			}

			return parentDistributor.Level + 1
		}(),
	}

	// check parent distributor permission
	if distributor.ParentDistributor != "" {
		parentDistributor, ok := DistributorCollection[distributor.ParentDistributor]
		if !ok {
			return models.Distributor{}, fmt.Errorf("parent distributor not found")
		}

		for _, include := range distributor.Includes {
			if !CheckDistributorPermissionforLocation(parentDistributor.DistributorID, include) {
				return models.Distributor{}, fmt.Errorf("unauthorized access for include: %s", include)
			}
		}

		for _, exclude := range distributor.Excludes {
			if !CheckDistributorPermissionforLocation(parentDistributor.DistributorID, exclude) {
				return models.Distributor{}, fmt.Errorf("unauthorized access for exclude: %s", exclude)
			}
		}
	}

	DistributorCollection[distributor.DistributorID] = distributor
	return distributor, nil
}

func CheckDistributorPermission(distId string, locations []string) (models.CheckDistributorPermissionResponse, error) {
	distributor, ok := DistributorCollection[distId]
	if !ok {
		return models.CheckDistributorPermissionResponse{}, fmt.Errorf("distributor not found")
	}

	permissionMap := make(map[string]bool)

	for _, location := range locations {
		permissionMap[location] = CheckDistributorPermissionforLocation(distributor.DistributorID, location)
	}

	return models.CheckDistributorPermissionResponse{
		DistributorID:   distId,
		DistributorName: distributor.Name,
		PermissionMap:   permissionMap,
	}, nil
}

func CheckDistributorPermissionforLocation(distId string, location string) bool {
	distributor, ok := DistributorCollection[distId]
	if !ok {
		return false
	}

	for _, include := range distributor.Includes {
		if include == location {
			return true
		}
	}

	for _, exclude := range distributor.Excludes {
		if strings.Contains(location, exclude) {
			return false
		}
	}

	for _, include := range distributor.Includes {
		if strings.HasSuffix(location, include) {
			return true
		}
	}

	return false
}

func UpdateDistributor(distributorID string, req models.UpdateDistributorCmd) (models.Distributor, error) {
	distributor, ok := DistributorCollection[distributorID]
	if !ok {
		return models.Distributor{}, fmt.Errorf("distributor not found")
	}

	isValidInclude := utils.ValidateIncludeExclude(req.Includes)
	if !isValidInclude {
		return models.Distributor{}, fmt.Errorf("invalid include")
	}

	isValidExclude := utils.ValidateIncludeExclude(req.Excludes)
	if !isValidExclude {
		return models.Distributor{}, fmt.Errorf("invalid exclude")
	}

	if req.Includes == nil {
		req.Includes = []string{}
	}

	if req.Excludes == nil {
		req.Excludes = []string{}
	}

	distributor.Name = req.Name
	distributor.Includes = append(distributor.Includes, req.Includes...)
	distributor.Excludes = append(distributor.Excludes, req.Excludes...)
	// distributor.ParentDistributor = req.ParentDistributorID

	utils.RemoveDuplicateLocations(distributor.Includes)
	utils.RemoveDuplicateLocations(distributor.Excludes)

	DistributorCollection[distributor.DistributorID] = distributor
	return distributor, nil
}
