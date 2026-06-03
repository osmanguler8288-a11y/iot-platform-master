package test

import (
	"fmt"
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func TestMqtt(t *testing.T) {
	opt := mqtt.NewClientOptions().AddBroker("tcp://127.0.0.1:1883")
	opt.SetUsername("admin").SetPassword("123456")
	//
	opt.SetDefaultPublishHandler(func(client mqtt.Client, message mqtt.Message) {
		fmt.Printf("message:%s\n", message.Payload())
		fmt.Printf("TOPIC:%s\n", message.Topic())
	})
	c := mqtt.NewClient(opt)
	//
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
	//
	if token := c.Subscribe("/sys/1/device_key/receive", 0, nil); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
	//
	if token := c.Publish("/sys/1/device_key/ping", 0, false, "hello"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
	//
	time.Sleep(10 * time.Second)
	//取消订阅
	if token := c.Unsubscribe("topic/#"); token.Wait() && token.Error() != nil {
		t.Fatal(token.Error())
	}
	//
	c.Disconnect(250)
}
