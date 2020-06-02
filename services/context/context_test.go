package context

import (
	"errors"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	pkg "github.com/qlcchain/qlc-go-sdk/pkg/types"

	"github.com/qlcchain/go-lsobus/common"
	"github.com/qlcchain/go-lsobus/config"
)

type testService struct {
	Id int
}

func (*testService) Init() error {
	panic("implement me")
}

func (*testService) Start() error {
	return nil
}

func (*testService) Stop() error {
	return nil
}

func (*testService) Status() int32 {
	panic("implement me")
}

type waitService struct {
	common.ServiceLifecycle
}

func (w *waitService) Init() error {
	if !w.PreInit() {
		return errors.New("pre init fail")
	}
	defer w.PostInit()
	return nil
}

func (w *waitService) Start() error {
	if !w.PreStart() {
		return errors.New("pre init fail")
	}
	defer w.PostStart()

	time.Sleep(time.Duration(3) * time.Second)
	return nil
}

func (w *waitService) Stop() error {
	if !w.PreStop() {
		return errors.New("pre init fail")
	}
	defer w.PostStop()
	return nil
}

func (w *waitService) Status() int32 {
	return w.State()
}

func Test_serviceContainer(t *testing.T) {
	sc := newServiceContainer()
	serv1 := &testService{Id: 1}
	t.Logf("serv1 %p", serv1)
	err := sc.Register(LogService, serv1)
	if err != nil {
		t.Fatal(err)
	}
	err = sc.Register(LogService, serv1)
	if err == nil {
		t.Fatal(err)
	}

	if service, err := sc.Get(LogService); service == nil || err != nil {
		t.Fatal(err)
	} else {
		t.Logf("%p, %p", service, serv1)
	}

	if b := sc.HasService(LogService); !b {
		t.Fatal("can not find ledger service")
	}

	serv2 := &testService{Id: 2}
	t.Logf("serv2 %p", serv2)
	err = sc.Register("TestService", serv2)
	if err != nil {
		t.Fatal(err)
	}

	sc.Iter(func(name string, service common.Service) error {
		t.Logf("%s: %p", name, service)
		return nil
	})

	sc.ReverseIter(func(name string, service common.Service) error {
		t.Logf("%s: %p", name, service)
		return nil
	})
	var i int
	sc.IterWithPredicate(func(name string, service common.Service) error {
		t.Logf("IterWithPredicate ==>%s: %p", name, service)
		i++
		return nil
	}, func(name string) bool {
		return name != "TestService"
	})

	if i != 1 {
		t.Fatal("invalid IterWithPredicate ", i)
	}

	err = sc.UnRegister(LogService)
	if err != nil {
		t.Fatal(err)
	}

	sc.Iter(func(name string, service common.Service) error {
		t.Logf("%s: %p", name, service)
		return nil
	})

	if _, err := sc.Get(LogService); err == nil {
		t.Fatal("shouldn't find log service")
	}
}

func TestNewServiceContextWithDefaultConfig(t *testing.T) {
	ctx := NewServiceContext("")
	defer func() {
		_ = os.Remove(config.DefaultDataDir())
	}()
	if ctx.cfgFile != path.Join(config.DefaultDataDir(), config.VirtualLSOBus) {
		t.Fatal("use default config error")
	}
	ac := account()
	ctx.SetAccount(ac)
	ac1 := ctx.Account()
	if ac1.String() != ac.String() {
		t.Fatal("set account error")
	}
	err := ctx.Init(func() error {
		return ctx.Register("waitService", &waitService{})
	})
	if err != nil {
		t.Fatal(err)
	}
	if !ctx.HasService("waitService") {
		t.Fatal("search waitService error")
	}
	if err := ctx.UnRegister("waitService"); err != nil {
		t.Fatal(err)
	}
	if ctx.HasService("waitService") {
		t.Fatal("waitService is already delete")
	}
	_ = ctx.Register("waitService", &waitService{})
	_, err = ctx.Service("waitService")
	if err != nil {
		t.Fatal(err)
	}
	if err = ctx.Start(); err != nil {
		t.Fatal(err)
	}
	if err = ctx.Stop(); err != nil {
		t.Fatal(err)
	}
	if ctx.Status() != 6 {
		t.Fatal("context status error")
	}
}

func TestNewServiceContext(t *testing.T) {
	cfgFile1 := filepath.Join(config.TestDataDir(), "config1", "test1.json")
	cfgFile2 := filepath.Join(config.TestDataDir(), "config2", "test2.json")
	t.Log(filepath.Dir(cfgFile2), filepath.Base(cfgFile2))
	cm := config.NewCfgManagerWithFile(cfgFile2)
	cfg, err := cm.Load()
	if err != nil {
		t.Fatal(err)
	}
	cfg.DataDir = filepath.Dir(cfgFile2)
	err = cm.Save()
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.Remove(cfgFile1)
		_ = os.Remove(cfgFile2)
	}()

	c1 := NewServiceContext(cfgFile1)
	c2 := NewServiceContext(cfgFile2)
	if c1 == nil || c2 == nil {
		t.Fatal("failed to create context")
	} else {
		if c1.Id() == c2.Id() {
			t.Fatal("invalid c1 and c2")
		} else {
			t.Log(c1.Id(), c2.Id())
		}
	}

	c3 := NewServiceContext(cfgFile1)
	if c1 != c3 {
		t.Fatalf("invalid instance expect: %p,act :%p", c1, c3)
	}

	cfg1, err := c1.Config()
	if err != nil {
		t.Fatal(err)
	}
	cfg2, err := c2.Config()
	if err != nil {
		t.Fatal(err)
	}

	if cfg1.DataDir != filepath.Dir(cfgFile1) {
		t.Fatalf(cfg1.DataDir, filepath.Dir(cfgFile1))
	}
	if cfg2.DataDir != filepath.Dir(cfgFile2) {
		t.Fatalf(cfg2.DataDir, filepath.Dir(cfgFile2))
	}

	eb1 := c1.EventBus()
	eb2 := c2.EventBus()
	eb3 := c3.EventBus()
	if eb1 == eb2 {
		t.Fatal("eb1 shouldn't same as eb2")
	}

	if eb1 != eb3 {
		t.Fatal("eb1 shouldn same as eb3")
	}
}

func TestSerivceContext_ConfigManager(t *testing.T) {
	cfgFile := filepath.Join(config.TestDataDir(), "config2", "test.json")
	ctx := NewServiceContext(cfgFile)
	err := ctx.Init(func() error {
		return ctx.Register("waitService", &waitService{})
	})
	if err != nil {
		t.Fatal(err)
	}
	if cfg, err := ctx.ConfigManager(func(cm *config.CfgManager) error {
		if cfg, err := cm.Load(); err != nil {
			return err
		} else {
			cfg.LogLevel = "info"
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("1, ctx==>%p,cfg==>%p", ctx, cfg)
	}
	ctx2 := NewServiceContext(cfgFile)
	if cfg2, err := ctx2.Config(); err != nil {
		t.Fatal(err)
	} else {
		t.Logf("2, ctx==>%p,cfg==>%p", ctx2, cfg2)
		if cfg2.LogLevel != "info" {
			t.Fatal("invalid loglevel")
		}
	}

}

func account() *pkg.Account {
	seed, _ := pkg.NewSeed()
	_, priv, _ := pkg.KeypairFromSeed(seed.String(), 0)
	return pkg.NewAccount(priv)
}
