// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"github.com/labring/lvscare/care"
	"github.com/labring/lvscare/utils"
	"github.com/spf13/cobra"
)

// careCmd represents the care command
var careCmd = &cobra.Command{
	Use:   "care",
	Short: "A lightweight LVS baby care, support ipvs health check.",
	RunE: func(cmd *cobra.Command, args []string) error {
		utils.Config(care.LVS.Logger)
		return care.LVS.VsAndRsCare()
	},
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if err := care.SetTargetIP(); err != nil {
			return err
		}
		// if target ip is empty, skip sync router
		if care.LVS.TargetIP == nil {
			return nil
		}
		return care.LVS.SyncRouter()
	},
}

func init() {
	rootCmd.AddCommand(careCmd)
	careCmd.Flags().IPVar(&care.LVS.TargetIP, "ip", nil, "target ip")
	careCmd.Flags().BoolVar(&care.LVS.RunOnce, "run-once", false, "is run once mode")
	careCmd.Flags().StringVar(&care.LVS.VirtualServer, "vs", "", "virtual server like 10.54.0.2:6443")
	careCmd.Flags().StringSliceVar(&care.LVS.RealServer, "rs", []string{}, "real server like 192.168.0.2:6443")
	careCmd.Flags().StringVar(&care.LVS.Logger, "logger", "INFO", "logger level: DEBG/INFO")
	careCmd.Flags().BoolVar(&care.LVS.Clean, "clean", false, "before run clean ipvs rules")
	careCmd.Flags().BoolVarP(&care.LVS.Test, "test", "t", true, "test mode")
	careCmd.Flags().StringVar(&care.LVS.HealthPath, "health-path", "/healthz", "health check path")
	careCmd.Flags().StringVar(&care.LVS.HealthSchem, "health-schem", "https", "health check schem")
	careCmd.Flags().Int32Var(&care.LVS.Interval, "interval", 5, "health check interval, unit is sec.")
}
