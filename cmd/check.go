package cmd

import (
	"context"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/spf13/cobra"

	"github.com/bavix/gripmock/v3/internal/deps"
	"github.com/bavix/gripmock/v3/internal/domain/waiter"
)

var (
	silenceErrorsFlag     bool
	pingTimeout           time.Duration
	errServerIsNotRunning = errors.New("server is not running")
)

const serviceName = "gripmock"

var checkCmd = &cobra.Command{
	Use:          "check",
	Args:         cobra.NoArgs,
	SilenceUsage: true,
	Short:        "The command checks whether the gripmock server is alive or dead by accessing it via the API",
	RunE: func(cmd *cobra.Command, _ []string) error {
		cmd.SilenceErrors = silenceErrorsFlag

		builder := deps.NewBuilder(deps.WithDefaultConfig())
		ctx, cancel := builder.SignalNotify(cmd.Context())
		defer cancel()

		defer builder.Shutdown(context.WithoutCancel(ctx))

		ctx = builder.Logger(ctx)

		pingService, err := builder.PingService()
		if err != nil {
			return errors.WithStack(err)
		}

		code, err := pingService.PingWithTimeout(ctx, pingTimeout, serviceName)
		if err != nil {
			return errors.WithStack(err)
		}

		if code != waiter.Serving {
			return errors.Wrapf(errServerIsNotRunning, "code: %d", code)
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	const defaultPingTimeout = time.Second * 5

	checkCmd.Flags().DurationVarP(&pingTimeout, "timeout", "t", defaultPingTimeout, "timeout")
	checkCmd.Flags().BoolVar(&silenceErrorsFlag, "silent", false, "silence errors")
}
