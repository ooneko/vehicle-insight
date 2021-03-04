package app

import (
	"fmt"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"ooneko.github.com/vehicle-insight/cmd/app/options"
	"ooneko.github.com/vehicle-insight/pkg/apiserver"
	"ooneko.github.com/vehicle-insight/pkg/log"
	"ooneko.github.com/vehicle-insight/pkg/util/signals"
)

func NewAPIServerCommand() *cobra.Command {
	apiserverFlags := options.NewAPIServerFlags()
	cmd := &cobra.Command{
		Use:  "vehicle-insight",
		Long: "A Recorder for car maintianance information",
		Run: func(cmd *cobra.Command, args []string) {
			config := zap.NewProductionConfig()
			config.Level = zap.NewAtomicLevelAt(log.ConvertToZapLevel(apiserverFlags.LogLevel))
			logger, err := config.Build()
			if err != nil {
				panic(err)
			}
			defer logger.Sync()

			apiServer := &options.APIServer{
				APIServerFlags: *apiserverFlags,
				Logger:         logger,
			}

			stopCh := signals.SetupSignalHandler()
			if err := run(apiServer, stopCh); err != nil {
				logger.Sugar().Errorf("run apiserver: %v", err.Error())
			}
		},
	}

	fs := cmd.Flags()
	apiserverFlags.AddFlags(fs)

	// ugly, but necessary, because Cobra's default UsageFunc and HelpFunc pollute the flagset with global flags
	const usageFmt = "Usage:\n  %s\n\nFlags:\n%s"
	cmd.SetUsageFunc(func(cmd *cobra.Command) error {
		fmt.Fprintf(cmd.OutOrStderr(), usageFmt, cmd.UseLine())
		return nil
	})
	cmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(cmd.OutOrStdout(), "%s\n\n"+usageFmt, cmd.Long, cmd.UseLine())
	})

	return cmd
}

func run(s *options.APIServer, stopCh <-chan struct{}) error {
	if err := runAPIServer(s); err != nil {
		return err
	}
	select {
	case <-stopCh:
		break
	}
	return nil
}

func runAPIServer(s *options.APIServer) error {
	apiserver := apiserver.NewMainAPIServer(s)
	if err := apiserver.Run(); err != nil {
		return err
	}
	s.Logger.Sugar().Infof("apiserver started")
	return nil
}
