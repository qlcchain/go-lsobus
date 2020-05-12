package commands

import (
	"encoding/hex"
	"errors"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/iixlabs/virtual-lsobus/common/util"

	"github.com/iixlabs/virtual-lsobus/services"
	ct "github.com/iixlabs/virtual-lsobus/services/context"

	"github.com/spf13/cobra"

	qlctypes "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/iixlabs/virtual-lsobus/log"
)

var (
	accountP       string
	cfgPathP       string
	chainEndPointP string
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd := &cobra.Command{
		Use:   "virtual-lsobus",
		Short: "virtual lsobus is a agent for MEF Sonata APIs",
		Long:  `virtual lsobus is a agent for MEF Sonata APIs`,
		Run: func(cmd *cobra.Command, args []string) {
			err := start()
			if err != nil {
				cmd.Println(err)
			}
		},
	}
	flag := rootCmd.PersistentFlags()
	flag.StringVar(&cfgPathP, "config", "", "config file")
	flag.StringVar(&accountP, "account", "", "sign with account for cdr data")
	flag.StringVarP(&chainEndPointP, "chainEndpoint", "", "ws://127.0.0.1:19736", "chain endpoint")
	if err := rootCmd.Execute(); err != nil {
		log.Root.Info(err)
		os.Exit(1)
	}
}

func start() error {
	var account *qlctypes.Account
	serviceContext := ct.NewServiceContext(cfgPathP)
	cm, err := serviceContext.ConfigManager()
	if err != nil {
		return err
	}
	cfg, err := cm.Config()
	if err != nil {
		return err
	}
	if chainEndPointP != "" {
		cfg.ChainUrl = chainEndPointP
	}
	s := util.ToIndentString(cfg)
	_ = ioutil.WriteFile(cm.ConfigFile, []byte(s), 0600)
	log.Root.Info("Run service id: ", serviceContext.Id())

	if len(accountP) > 0 {
		bytes, err := hex.DecodeString(accountP)
		if err != nil {
			return err
		}
		account = qlctypes.NewAccount(bytes)
	} else {
		return errors.New("must run with qlc account")
	}

	// save accounts to context
	serviceContext.SetAccount(account)
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
	log.Root.Info("virtual router closed successfully")
}
