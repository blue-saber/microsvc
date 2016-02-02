package microsvc

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)

type MicroService struct {
	resource *IResource  `@Autowired:"*"`
	G_engine *gin.Engine `@Autowired:"*"`
}

func (svc *MicroService) PostSummerConstruct() {
	res := *svc.resource

	engine := svc.G_engine
	prefix := res.PathPrefix()
	param := res.GetParam()

	aurl := "/" + prefix
	iurl := aurl + "/:" + param

	engine.GET(aurl, res.GetAll)
	engine.GET(iurl, res.DoGet)
	engine.POST(aurl, res.DoPost)
	engine.PUT(iurl, res.DoPut)
	engine.DELETE(iurl, res.DoDelete)
	engine.DELETE(aurl, res.DeleteAll)
}

func (svc *MicroService) Setresource(r interface{}) {
	if origional, ok := r.(IResource); ok {
		svc.resource = &origional
	}
}
