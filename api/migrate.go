package api

import (
	mQuasar "github.com/elgatito/elementum/migrate/quasar"

	"github.com/gin-gonic/gin"
)

// MigratePlugin gin proxy for /migrate/:plugin ...
func MigratePlugin(ctx *gin.Context) {
	plugin := ctx.Params.ByName("plugin")
	if plugin == "quasar" {
		mQuasar.Migrate()
	}
}
