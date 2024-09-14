package service

import (
	"HunCoding/src/main/configuration/rest_err"
	"HunCoding/src/main/model"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *rest_err.RestErr {
	userDomain.EncryptPassword()
	
	return nil
}