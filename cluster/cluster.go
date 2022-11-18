// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
//
// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cluster

import (
	"github.com/temporalio/tctl-kit/pkg/output"
	"github.com/temporalio/temporal-cli/common"
	"github.com/urfave/cli/v2"
)

func NewClusterCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "health",
			Usage: "Check health of frontend service",
			Action: func(c *cli.Context) error {
				return HealthCheck(c)
			},
		},
		{
			Name:      "describe",
			Usage:     "Show information about the cluster",
			ArgsUsage: " ",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     output.FlagOutput,
					Aliases:  common.FlagOutputAlias,
					Usage:    output.UsageText,
					Value:    string(output.Table),
					Category: common.CategoryFormatting,
				},
				&cli.StringFlag{
					Name:     output.FlagFields,
					Usage:    "customize fields to print. Set to 'long' to automatically print more of main fields",
					Category: common.CategoryFormatting,
				},
			},
			Action: func(c *cli.Context) error {
				return DescribeCluster(c)
			},
		},
		{
			Name:      "system",
			Usage:     "Show information about the system and capabilities",
			ArgsUsage: " ",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     output.FlagOutput,
					Aliases:  common.FlagOutputAlias,
					Usage:    output.UsageText,
					Value:    string(output.Table),
					Category: common.CategoryFormatting,
				},
				&cli.StringFlag{
					Name:     output.FlagFields,
					Usage:    "customize fields to print. Set to 'long' to automatically print more of main fields",
					Category: common.CategoryFormatting,
				},
			},
			Action: func(c *cli.Context) error {
				return DescribeSystem(c)
			},
		},
		{
			Name:      "upsert",
			Usage:     "Add or update a remote cluster",
			ArgsUsage: " ",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     common.FlagClusterAddress,
					Usage:    "Frontend address of the remote cluster",
					Category: common.CategoryMain,
				},
				&cli.BoolFlag{
					Name:     common.FlagClusterEnableConnection,
					Usage:    "Enable cross cluster connection",
					Category: common.CategoryMain,
				},
			},
			Action: func(c *cli.Context) error {
				return UpsertCluster(c)
			},
		},
		{
			Name:      "list",
			Usage:     "List all remote clusters",
			ArgsUsage: " ",
			Flags:     common.FlagsForPaginationAndRendering,
			Action: func(c *cli.Context) error {
				return ListClusters(c)
			},
		},
		{
			Name:      "remove",
			Usage:     "Remove a remote cluster",
			ArgsUsage: " ",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:     common.FlagName,
					Usage:    "Frontend address of the remote cluster",
					Required: true,
					Category: common.CategoryMain,
				},
			},
			Action: func(c *cli.Context) error {
				return RemoveCluster(c)
			},
		},
	}
}
