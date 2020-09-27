package cmd

import (
	"context"
	"epgu-generator/internal/artefact"
	"epgu-generator/internal/config"
	"epgu-generator/internal/consumer"
	"epgu-generator/internal/di"
	"epgu-generator/internal/registry"
	"epgu-generator/internal/replication"
	"github.com/spf13/cobra"
)

var converterCmd = &cobra.Command{
	Use:   "replication",
	Short: "replication",
	Long:  "replication",
	Run:   converterMain,
}

func init() {
	rootCmd.AddCommand(converterCmd)

	params := []config.Param{
		{Name: "registry", Value: "/tmp/registry.xlsx", Usage: "path to targets registry", ViperBind: "Registry"},
		{Name: "dir-template", Value: "/tmp/template", Usage: "path to templates dir", ViperBind: "Dir.Template"},
		{Name: "dir-artefact", Value: "/tmp/artefact", Usage: "path to artefact dir", ViperBind: "Dir.Artefact"},
		{Name: "dir-tmp", Value: "/tmp", Usage: "path to tmp dir", ViperBind: "Dir.Tmp"},
	}

	config.Apply(converterCmd, params)
}

func converterMain(cmd *cobra.Command, args []string) {
	ctx := context.Background()

	di := &di.Runner{
		Provide: map[string]interface{}{
			"config":                config.NewConfig,
			"registry.NewParser":    registry.NewParser,
			"artefact.NewTemplates": artefact.NewTemplates,
			"artefact.NewContent":   artefact.NewContent,
			"artefact.NewFolders":   artefact.NewFolders,
			"replication.Converter": replication.NewConverter,
			"consumer.Converter":    consumer.NewReplication,
		},
		Invoke: func(ctx context.Context, args []string) interface{} {
			return func(i consumer.Consumer) {
				i.Run(ctx, args)
			}
		},
	}

	di.Run(ctx, args)
}
