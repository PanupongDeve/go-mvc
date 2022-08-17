package healthz

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthz(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
