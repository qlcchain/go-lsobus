/*
 * Copyright (c) 2019 QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package event

import (
	"errors"
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/qlcchain/go-lsobus/common"

	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
	bus := New()
	if bus == nil {
		t.Log("New EventBus not created!")
		t.Fail()
	}
}

func TestSubscribe(t *testing.T) {
	bus := NewEventBus(runtime.NumCPU())

	if _, err := bus.Subscribe("test", func() {}); err != nil {
		t.Fail()
	}

	if _, err := bus.Subscribe("test", 2); err == nil {
		t.Fail()
	}
}

func TestSubscribeSync(t *testing.T) {
	bus := NewEventBus(runtime.NumCPU())

	counter := int64(0)
	topic := common.TopicType("test")
	if _, err := bus.SubscribeSync(topic, func() {
		atomic.AddInt64(&counter, 1)
		t.Log("sub1")
	}); err != nil {
		t.Fail()
	}
	if _, err := bus.Subscribe(topic, func() {
		t.Log("sub2")
	}); err != nil {
		t.Fail()
	}

	bus.Publish(topic)

	if counter != int64(1) {
		t.Fatal("invalid ", counter)
	}

	_ = bus.Close()
}

func TestUnsubscribe(t *testing.T) {
	bus := NewEventBus(runtime.NumCPU())

	handler := func() {}

	id, _ := bus.Subscribe("test", handler)

	if err := bus.Unsubscribe("test", id); err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if err := bus.Unsubscribe("unexisted", "xxx"); err == nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestUnsubscribe2(t *testing.T) {
	bus := SimpleEventBus()

	handler := func() {}

	id, _ := bus.Subscribe("test", handler)

	t.Log(bus.(*DefaultEventBus).handlers.Len())
	if value, ok := bus.(*DefaultEventBus).handlers.GetStringKey("test"); ok {
		t.Log(value.(*eventHandlers).Size())
	}

	if err := bus.Unsubscribe("test", id); err != nil {
		fmt.Println(err)
		t.Fail()
	}
	t.Log(bus.(*DefaultEventBus).handlers.Len())
	if value, ok := bus.(*DefaultEventBus).handlers.GetStringKey("test"); ok {
		t.Log(value.(*eventHandlers).Size())
	}
	if err := bus.Unsubscribe("unexisted", "xxx"); err == nil {
		fmt.Println(err)
		t.Fail()
	}
}

func TestClose(t *testing.T) {
	bus := NewEventBus(runtime.NumCPU())

	handler := func() {}

	_, _ = bus.Subscribe("test", handler)

	original, ok := bus.(*DefaultEventBus)
	if !ok {
		t.Fatal("Could not cast message bus to its original type")
	}

	if original.handlers.Len() == 0 {
		t.Fatal("Did not subscribed handler to topic")
	}

	bus.CloseTopic("test")

	if original.handlers.Len() != 0 {
		t.Fatal("Did not unsubscribed handlers from topic")
	}
}

func TestPublish(t *testing.T) {
	bus := NewEventBus(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	first := false
	second := false

	_, _ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		first = v
	})

	_, _ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		second = v
	})

	bus.Publish("topic", true)

	wg.Wait()

	if first == false || second == false {
		t.Fatal(first, second)
	}
}

func TestHandleError(t *testing.T) {
	bus := NewEventBus(runtime.NumCPU())
	_, _ = bus.Subscribe("topic", func(out chan<- error) {
		out <- errors.New("I do throw error")
	})

	out := make(chan error)
	defer close(out)

	bus.Publish("topic", out)

	if <-out == nil {
		t.Fail()
	}
}

func TestHasCallback(t *testing.T) {
	bus := New()
	_, err := bus.Subscribe("topic", func() {})
	if err != nil {
		t.Fatal(err)
	}
	if bus.HasCallback("topic_topic") {
		t.Fail()
	}
	if !bus.HasCallback("topic") {
		t.Fail()
	}
}

func TestGetEventBus(t *testing.T) {
	eb0 := SimpleEventBus()
	eb1 := GetEventBus("")
	if eb0 != eb1 {
		t.Fatal("invalid default eb")
	}

	id1 := "111111"
	eb2 := GetEventBus(id1)
	eb3 := GetEventBus(id1)

	if eb2 != eb3 {
		t.Fatal("invalid eb of same id")
	}

	id2 := "222222"
	eb4 := GetEventBus(id2)
	if eb3 == eb4 {
		t.Fatal("invalid eb of diff ids")
	}
}

func TestEventSubscribe(t *testing.T) {
	bus := NewEventBus(runtime.NumCPU())
	topic := common.TopicType("test")

	counter := int64(0)
	_, _ = bus.Subscribe(topic, func(i int64) {
		fmt.Println("sub1", i, atomic.AddInt64(&counter, 1))
	})

	_, _ = bus.Subscribe(topic, func(i int64) {
		time.Sleep(time.Second)
		fmt.Println("sub2", i, atomic.AddInt64(&counter, 1))
	})

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			fmt.Println("publish: ", i)
			bus.Publish(topic, int64(i))
		}
	}()
	wg.Wait()
	_ = bus.Close()

	//if atomic.LoadInt64(&counter) != 10 {
	//	t.Fatal("invalid sub", atomic.LoadInt64(&counter))
	//}
	t.Log("result", atomic.LoadInt64(&counter))
}

type foo struct {
	id string
}

func (f *foo) Test(arg int) {
	fmt.Printf("foo:%s, args=%d\n", f.id, arg)
}

func TestFooSubscribe(t *testing.T) {
	bus := NewEventBus(runtime.NumCPU())
	topic := common.TopicType("test")
	foo := &foo{id: uuid.New().String()}

	wg := sync.WaitGroup{}
	wg.Add(1)
	i := 1
	go func(i int) {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			bus.Publish(topic, i)
		}
	}(i)
	id, err := bus.Subscribe(topic, foo.Test)
	t.Log(id)
	if err != nil {
		t.Fatal(err)
	}
	wg.Wait()
	if flag := bus.HasCallback(topic); !flag {
		t.Fatal()
	}
	if err = bus.Unsubscribe(topic, id); err != nil {
		t.Fatal(err)
	}
	if flag := bus.HasCallback(topic); flag {
		t.Fatal()
	}
	_ = bus.Close()
}
