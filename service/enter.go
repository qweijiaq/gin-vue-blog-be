package service

import (
	"server/service/image"
	"server/service/user"
)

type ServiceGroup struct {
	ImageService image.ImageService
	UserService  user.UserService
}

var ServiceApp = new(ServiceGroup)
