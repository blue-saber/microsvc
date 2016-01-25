package microsvc

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

type MicroService struct {
	resource *IResource `@Autowired:"*"`
}

type microServiceStatus struct {
	Sys      uint64 `json:"sys"`
	Alloc    uint64 `json:"alloc"`
	Idle     uint64 `json:"idle"`
	Released uint64 `json:"released"`
}

func (svc *MicroService) Run(addr string) {
	res := *svc.resource

	if res.InReleaseMode() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	prefix := res.PathPrefix()
	param := res.GetParam()

	aurl := "/" + prefix
	iurl := aurl + "/:" + param

	r.GET(aurl, res.GetAll)
	r.GET(iurl, res.DoGet)
	r.POST(aurl, res.DoPost)
	r.PUT(iurl, res.DoPut)
	r.DELETE(iurl, res.DoDelete)
	r.DELETE(aurl, res.DeleteAll)

	r.GET("/status", func(c *gin.Context) {
		var m runtime.MemStats
		var status microServiceStatus

		runtime.ReadMemStats(&m)

		status.Sys = m.HeapSys
		status.Alloc = m.HeapAlloc
		status.Idle = m.HeapIdle
		status.Released = m.HeapReleased

		c.JSON(http.StatusOK, status)
	})

	// we'll pass in configuration later
	r.Run(addr)
}

func (svc *MicroService) Setresource(r interface{}) {
	if origional, ok := r.(IResource); ok {
		svc.resource = &origional
	}
}
