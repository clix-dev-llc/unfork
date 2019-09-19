package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/replicatedhq/unfork/pkg/version"
)

func VersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the current version and exit",
		Long:  `Print the current version and exit`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// print basic version info
			fmt.Printf("Replicated Unfork %s\n", version.Version())

			// check if this is the latest release, and display possible upgrade instructions
			isLatest, latestVer, err := version.IsLatestRelease()
			if err != nil {
				fmt.Printf("\nUnable to check for newer releases: %s\n", err.Error())
			} else if !isLatest {
				fmt.Printf("\nVersion %s is available for unfork. To install updates, run\n  $ curl https://unfork.io/install | bash\n", latestVer)
			}

			return nil
		},
	}
	return cmd
}
