// Copyright © 2018 Guillaume Jacquet
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

// labelsCmd represents the labels command
var labelsCmd = &cobra.Command{
	Use:   "labels",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := viper.GetString("token")

		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)

		client := github.NewClient(tc)

		query := args[0]

		owner, _ := cmd.Flags().GetString("owner")
		repositoryName, _ := cmd.Flags().GetString("repository")
		var repositoryID int64
		if repositoryName != "" {
			repository, _, err := client.Repositories.Get(ctx, owner, repositoryName)
			if err != nil {
				panic(err)
			}
			repositoryID = *repository.ID
		}

		sort, _ := cmd.Flags().GetString("sort")
		order, _ := cmd.Flags().GetString("order")

		results, _, err := client.Search.Labels(ctx, repositoryID, query, &github.SearchOptions{Sort: sort, Order: order})
		if err != nil {
			panic(err)
		}

		fmt.Println(*results)
	},
}

func init() {
	searchCmd.AddCommand(labelsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// labelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	labelsCmd.Flags().String("owner", "", "Owner")
	labelsCmd.Flags().String("repository", "", "Repository")
	labelsCmd.Flags().String("sort", "", "The sort field. Can be one of created or updated.")
	labelsCmd.Flags().String("order", "", "The sort order if the sort parameter is provided. Can be one of asc or desc.")
}
