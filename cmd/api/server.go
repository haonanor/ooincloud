package api

import (
	"fmt"
	"log"
	"net/http"
	"study/core"
	"study/initialize"

	"github.com/spf13/cobra"
)

var (
	configYml string
	StartCmd  = &cobra.Command{
		Use:          "server",
		Short:        "Start API server",
		Example:      "go server -c config/settings.yml",
		SilenceUsage: true,
		PreRun: func(cmd *cobra.Command, args []string) {
			setup()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func setup() {
	core.ROBOT_CONFIG = initialize.Viper("config/app.yaml")
	core.ROBOT_LOGGER = initialize.Zap()
	core.ROBOT_DB = initialize.Gorm()
	// fmt.Println(gin.DebugMode)
}

var AppRouters = make([]func(), 0)

func init() {
	StartCmd.PersistentFlags().StringVarP(&configYml, "config", "c", "config/app.yaml", "Start server with provided configuration file")
	// AppRouters = append(AppRouters, router.InitRouter)
	// boot.InitStatic()

}

func run() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", core.ROBOT_CONFIG.Get("app.host"), core.ROBOT_CONFIG.Get("app.port")),
		Handler: initialize.InitRouter(),
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("listen: ", err)
	}

	return nil
}
