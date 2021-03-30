//Package srv host clinic search function as http api service
package srv

import (
	"net/http"
	"strings"
	"time"

	"clinics-apis/modules"

	"github.com/gin-gonic/gin"
)

//Start start api server
func Start(addr string) error {
	go func() {
		//delay to load clinics, let http server start first
		time.Sleep(1 * time.Second)
		modules.LoadClinics()
	}()
	return setupRouter().Run(addr)
}

//setupRouter setup global route for api server
func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/search", func(c *gin.Context) {
		name, _ := c.GetQuery("name")
		state, _ := c.GetQuery("state")
		opening, _ := c.GetQuery("opening")
		list := modules.Search(strings.ToLower(name), strings.ToLower(state), opening)

		c.JSON(http.StatusOK, list)
	})

	r.Static("/data", "./data")

	return r
}
