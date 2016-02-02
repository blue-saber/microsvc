package microsvc

import "github.com/gin-gonic/gin"

type IResource interface {
	PathPrefix() string
	GetParam() string
	GetAll(*gin.Context)
	DoGet(*gin.Context)
	DoPost(*gin.Context)
	DoPut(*gin.Context)
	DoDelete(*gin.Context)
	DeleteAll(*gin.Context)
}
