package main

import (
	"github.com/spf13/cobra"
)

var restApiCmd = &cobra.Command{
	Use:   "rest-api",
	Short: "run rest api",
	Run: func(cmd *cobra.Command, args []string) {
		//appConf := conf.LoadAppConf()

		//dependency, closeFn := services.NewDependency(appConf)
		//
		//server := presenter.New(&presenter.Presenter{
		//	Dependency: dependency,
		//	Port:       appConf.AppPort,
		//})
		//
		//ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		//defer stop()
		//
		//go func() {
		//	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		//		log.Err(err).Msg("failed listen serve")
		//		ctx.Done()
		//	}
		//}()
		//
		//<-ctx.Done()
		//log.Info().Msg("Received shutdown signal, shutting down server gracefully...")
		//
		//if err := server.Shutdown(context.Background()); err != nil {
		//	log.Err(err).Msg("failed shutdown server")
		//}
		//
		//closeFn()
		//log.Info().Msg("Shutdown complete. Exiting.")
		//return
	},
}
