// Copyright 2020 by Red Hat, Inc. All rights reserved.
// Use of this source is goverend by the Apache License
// that can be found in the LICENSE file.

package compose

import (
	"github.com/spf13/cobra"

	"weldr-client/cmd/composer-cli/root"
)

var (
	composeCmd = &cobra.Command{
		Use:   "compose ...",
		Short: "Manage composes",
		Long:  "Manage composes on the server",
	}
)

func init() {
	root.AddRootCommand(composeCmd)
}
