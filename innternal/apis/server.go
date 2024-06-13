package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx *cli.Context, app *AppProvider) error {
	if !app.Config.Debug() {
		gin.SetMode(gin.ReleaseMode)
	}
	eg, groupCtx := errgroup.WithContext(ctx.Context)
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	log.Printf("HTTP Listen Port:%d", app.Config)
	return run(c)
}
