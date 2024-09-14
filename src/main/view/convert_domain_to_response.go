package view

import (
	"HunCoding/src/main/model"
	"HunCoding/src/main/controller/model/response"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
	) response.UserResponse { 
		return response.UserResponse {
			ID : "",
			Email: userDomain.GetEmail(),
			Name:  userDomain.GetName(),
			Age:   userDomain.GetAge(),
		}
}