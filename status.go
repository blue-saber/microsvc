package microsvc

import (
	// "fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"runtime"
)

type MicroServiceStatus struct {
	G_engine *gin.Engine `@Autowired:"*"`
}

type microServiceStatus struct {
	Sys      uint64 `json:"sys"`
	Alloc    uint64 `json:"alloc"`
	Idle     uint64 `json:"idle"`
	Released uint64 `json:"released"`
}

func (svc *MicroServiceStatus) PostSummerConstruct() {
	engine := svc.G_engine

	engine.GET("/status", func(c *gin.Context) {
		var m runtime.MemStats
		var status microServiceStatus

		runtime.ReadMemStats(&m)

		status.Sys = m.HeapSys
		status.Alloc = m.HeapAlloc
		status.Idle = m.HeapIdle
		status.Released = m.HeapReleased

		c.JSON(http.StatusOK, status)
	})

}
