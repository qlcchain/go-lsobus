package commands

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/qlcchain/qlc-go-sdk/pkg/util"
	"gopkg.in/validator.v2"

	"github.com/qlcchain/go-lsobus/config"

	"github.com/qlcchain/go-lsobus/services"
	ct "github.com/qlcchain/go-lsobus/services/context"

	"github.com/spf13/cobra"

	"github.com/qlcchain/go-lsobus/log"
)

var (
	cfgPathP      string
	configParamsP string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := &cobra.Command{
		Use:   "lsobus",
		Short: "LSOBUS is a agent for MEF Sonata APIs",
		Long:  `LSOBUS is a agent for MEF Sonata APIs`,
		Run: func(cmd *cobra.Command, args []string) {
			err := start()
			if err != nil {
				cmd.Println(err)
			}
		},
	}
	flag := rootCmd.PersistentFlags()
	flag.StringVar(&cfgPathP, "config", "", "config file")
	flag.StringVar(&configParamsP, "configParams", "", "parameters that can be override")
	if err := rootCmd.Execute(); err != nil {
		log.Root.Info(err)
		os.Exit(1)
	}
}

func start() error {
	serviceContext := ct.NewServiceContext(cfgPathP)
	cm, err := serviceContext.ConfigManager(func(cm *config.CfgManager) error {
		if len(configParamsP) > 0 {
			cfg, _ := cm.Config()
			params := strings.Split(configParamsP, ";")

			if len(params) > 0 {
				_, err := cm.PatchParams(params, cfg)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	cfg, err := cm.Config()
	if err != nil {
		return err
	}

	if err = validator.Validate(cfg); err != nil {
		return err
	}

	log.Root.Debug(util.ToIndentString(cfg))
	log.Root.Info("Run service id: ", serviceContext.Id())

	// start all services by chain context
	err = serviceContext.Init(func() error {
		return services.RegisterServices(serviceContext)
	})
	if err != nil {
		return err
	}
	err = serviceContext.Start()

	if err != nil {
		return err
	}
	trapSignal()
	return nil
}

func trapSignal() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	serviceContext := ct.NewServiceContext(cfgPathP)
	err := serviceContext.Stop()
	if err != nil {
		log.Root.Info(err)
	}
	time.Sleep(1 * time.Second)
	log.Root.Info("LSOBUS closed successfully")
}
