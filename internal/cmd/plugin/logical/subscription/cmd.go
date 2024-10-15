/*
Copyright The CloudNativePG Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package subscription

import (
	"github.com/spf13/cobra"

	"github.com/cloudnative-pg/cloudnative-pg/internal/cmd/plugin"
	"github.com/cloudnative-pg/cloudnative-pg/internal/cmd/plugin/logical/subscription/create"
	"github.com/cloudnative-pg/cloudnative-pg/internal/cmd/plugin/logical/subscription/drop"
	"github.com/cloudnative-pg/cloudnative-pg/internal/cmd/plugin/logical/subscription/syncsequences"
)

// NewCmd initializes the subscription command
func NewCmd() *cobra.Command {
	subscriptionCmd := &cobra.Command{
		Use:     "subscription",
		Short:   "Logical subscription management commands",
		GroupID: plugin.GroupIDDatabase,
	}
	subscriptionCmd.AddCommand(create.NewCmd())
	subscriptionCmd.AddCommand(drop.NewCmd())
	subscriptionCmd.AddCommand(syncsequences.NewCmd())

	return subscriptionCmd
}
