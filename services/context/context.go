package context

import (
	"errors"
	"fmt"
	"path"
	"sync"

	"github.com/iixlabs/virtual-lsobus/common/event"

	"github.com/iixlabs/virtual-lsobus/common"
	"github.com/iixlabs/virtual-lsobus/log"

	"github.com/cornelk/hashmap"

	qlctypes "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/iixlabs/virtual-lsobus/config"
)

var cache = hashmap.New(10)

const (
	LogService      = "logService"
	ContractService = "contractService"
)

type Option func(cm *config.CfgManager) error

type ServiceContext struct {
	common.ServiceLifecycle
	services *serviceContainer
	cm       *config.CfgManager
	cfgFile  string
	chainId  string
	locker   sync.RWMutex
	account  *qlctypes.Account
}

func NewServiceContext(cfgFile string) *ServiceContext {
	var dataDir string
	if len(cfgFile) == 0 {
		dataDir = config.DefaultDataDir()
		cfgFile = path.Join(dataDir, config.VirtualLSOBus)
	} else {
		cm := config.NewCfgManagerWithFile(cfgFile)
		dataDir, _ = cm.ParseDataDir()
	}
	id := qlctypes.HashData([]byte(dataDir)).String()
	if v, ok := cache.GetStringKey(id); ok {
		return v.(*ServiceContext)
	} else {
		sr := &ServiceContext{
			services: newServiceContainer(),
			cfgFile:  cfgFile,
			chainId:  id,
		}
		cache.Set(id, sr)
		return sr
	}
}

func (sc *ServiceContext) EventBus() event.EventBus {
	return event.GetEventBus(sc.Id())
}

func (sc *ServiceContext) SetAccount(account *qlctypes.Account) {
	sc.locker.Lock()
	defer sc.locker.Unlock()
	sc.account = account
}

func (sc *ServiceContext) Account() *qlctypes.Account {
	sc.locker.RLock()
	defer sc.locker.RUnlock()
	return sc.account
}

func (sc *ServiceContext) Id() string {
	return sc.chainId
}

func (sc *ServiceContext) ConfigFile() string {
	return sc.cfgFile
}

func (sc *ServiceContext) Init(fn func() error) error {
	if !sc.PreInit() {
		return errors.New("pre init fail")
	}
	defer sc.PostInit()

	if fn != nil {
		err := fn()
		if err != nil {
			return err
		}
	}

	sc.services.IterWithPredicate(func(name string, service common.Service) error {
		err := service.Init()
		if err != nil {
			return err
		}
		log.Root.Infof("%s init successfully", name)
		return nil
	}, func(name string) bool {
		return name != LogService
	})

	return nil
}

func (sc *ServiceContext) Start() error {
	if !sc.PreStart() {
		return errors.New("pre start fail")
	}
	defer sc.PostStart()
	sc.services.Iter(func(name string, service common.Service) error {
		return nil
	})
	sc.services.Iter(func(name string, service common.Service) error {
		err := service.Start()
		if err != nil {
			return fmt.Errorf("%s, %s", name, err)
		}
		log.Root.Infof("%s start successfully", name)
		return nil
	})

	return nil
}

func (sc *ServiceContext) Stop() error {
	if !sc.PreStop() {
		return errors.New("pre stop fail")
	}
	defer sc.PostStop()

	sc.services.ReverseIter(func(name string, service common.Service) error {
		err := service.Stop()
		if err != nil {
			return err
		}
		fmt.Printf("%s stop successfully.\n", name)
		return nil
	})

	return nil
}

func (sc *ServiceContext) Status() int32 {
	return sc.State()
}

func (sc *ServiceContext) Register(name string, service common.Service) error {
	return sc.services.Register(name, service)
}

func (sc *ServiceContext) HasService(name string) bool {
	return sc.services.HasService(name)
}

func (sc *ServiceContext) UnRegister(name string) error {
	return sc.services.UnRegister(name)
}

func (sc *ServiceContext) AllServices() ([]common.Service, error) {
	var services []common.Service
	sc.services.Iter(func(name string, service common.Service) error {
		services = append(services, service)
		return nil
	})
	return services, nil
}

func (sc *ServiceContext) Service(name string) (common.Service, error) {
	return sc.services.Get(name)
}

func (sc *ServiceContext) ConfigManager(opts ...Option) (*config.CfgManager, error) {
	sc.locker.Lock()
	defer sc.locker.Unlock()
	if sc.cm == nil {
		sc.cm = config.NewCfgManagerWithFile(sc.cfgFile)
		_, err := sc.cm.Load()
		if err != nil {
			return nil, err
		}
	}

	for _, opt := range opts {
		_ = opt(sc.cm)
	}

	return sc.cm, nil
}

func (sc *ServiceContext) Config() (*config.Config, error) {
	cm, err := sc.ConfigManager()
	if err != nil {
		return nil, err
	}
	return cm.Config()
}

type serviceContainer struct {
	locker   sync.RWMutex
	services map[string]common.Service
	names    []string
}

func newServiceContainer() *serviceContainer {
	return &serviceContainer{
		locker:   sync.RWMutex{},
		services: make(map[string]common.Service),
		names:    []string{},
	}
}

func (sc *serviceContainer) Register(name string, s common.Service) error {
	sc.locker.Lock()
	defer sc.locker.Unlock()

	if _, ok := sc.services[name]; ok {
		return fmt.Errorf("service[%s] already exist", name)
	} else {
		sc.services[name] = s
		sc.names = append(sc.names, name)
		return nil
	}
}

func (sc *serviceContainer) UnRegister(name string) error {
	sc.locker.Lock()
	defer sc.locker.Unlock()

	if v, ok := sc.services[name]; ok {
		_ = v.Stop()
		delete(sc.services, name)
		for idx, n := range sc.names {
			if n == name {
				sc.names = append(sc.names[:idx], sc.names[idx+1:]...)
				break
			}
		}
		return nil
	} else {
		return fmt.Errorf("service[%s] not exist", name)
	}
}

func (sc *serviceContainer) Get(name string) (common.Service, error) {
	sc.locker.RLock()
	defer sc.locker.RUnlock()

	if v, ok := sc.services[name]; ok {
		return v, nil
	} else {
		return nil, fmt.Errorf("service[%s] not exist", name)
	}
}

func (sc *serviceContainer) HasService(name string) bool {
	sc.locker.RLock()
	defer sc.locker.RUnlock()

	if _, ok := sc.services[name]; ok {
		return true
	}

	return false
}

func (sc *serviceContainer) Iter(fn func(name string, service common.Service) error) {
	sc.IterWithPredicate(fn, func(name string) bool {
		return true
	})
}

func (sc *serviceContainer) IterWithPredicate(fn func(name string, service common.Service) error,
	predicate func(name string) bool) {
	sc.locker.RLock()
	defer sc.locker.RUnlock()
	for idx := range sc.names {
		name := sc.names[idx]
		if service, ok := sc.services[name]; ok && predicate(name) {
			err := fn(name, service)
			if err != nil {
				break
			}
		}
	}
}

func (sc *serviceContainer) ReverseIter(fn func(name string, service common.Service) error) {
	sc.locker.RLock()
	defer sc.locker.RUnlock()

	for i := len(sc.names) - 1; i >= 0; i-- {
		name := sc.names[i]
		if service, ok := sc.services[name]; ok {
			err := fn(name, service)
			if err != nil {
				break
			}
		}
	}
}
