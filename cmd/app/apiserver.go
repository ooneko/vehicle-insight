package app

import (
	"github.com/spf13/cobra"

	"ooneko.github.com/vehicle-insight/cmd/app/options"
	"ooneko.github.com/vehicle-insight/pkg/util/signals"
)

func NewAPIServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
	cmd := cobra.Command{
		Use:  "car-pacemaker",
		Long: "A Recorder for car maintianance information",
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(s, signals.SetupSignalHandler())
		},
	}

	fs := cmd.Flags()
	namedFlagSets := s.Flags()
	for _, f := range namedFlagSets.FlagSets {
		fs.AddFlagSet(f)
	}

	return &cmd
}

func run(s *options.ServerRunOptions, stopCh <-chan struct{}) error {
	// _ = apiserver.NewAPIServer()
	s.NewAPIServer()
	return nil
}
