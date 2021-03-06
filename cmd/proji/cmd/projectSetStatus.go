package cmd

import (
	"fmt"
	"strings"

	"github.com/nikoksr/proji/pkg/helper"
	"github.com/spf13/cobra"
)

var projectSetStatusCmd = &cobra.Command{
	Use:   "status STATUS PROJECT-ID",
	Short: "Set a new status",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("missing status or project-ID")
		}

		status := strings.ToLower(args[0])
		projectID, err := helper.StrToUInt(args[1])
		if err != nil {
			return err
		}

		err = setStatus(projectID, status)
		if err != nil {
			fmt.Printf("> Setting status '%s' for project %d failed: %v\n", status, projectID, err)
			return err
		}
		fmt.Printf("> Status '%s' was successfully set for project %d\n", status, projectID)
		return nil
	},
}

func init() {
	projectSetCmd.AddCommand(projectSetStatusCmd)
}

func setStatus(projectID uint, statusTitle string) error {
	// Load and validate status
	statusID, err := projiEnv.Svc.LoadStatusID(statusTitle)
	if err != nil {
		return err
	}
	// Validate project
	_, err = projiEnv.Svc.LoadProject(projectID)
	if err != nil {
		return err
	}
	return projiEnv.Svc.UpdateProjectStatus(projectID, statusID)
}
