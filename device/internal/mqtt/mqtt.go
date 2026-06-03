package mqtt

import (
	"fmt"
	"iot-platform-master/models"
	"log"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var topic = "/sys/#"
var MC mqtt.Client

func NewMqttServer(broker, clientID, username, password string) {
	opt := mqtt.NewClientOptions().AddBroker(broker).SetClientID(clientID).
		SetUsername(username).SetPassword(password)

	// 回调
	opt.SetDefaultPublishHandler(publishHandler)

	MC = mqtt.NewClient(opt)

	// 连接
	if token := MC.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// 订阅主题
	if token := MC.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	log.Println("[MQTT] connected and subscribed to", topic)

	// 阻塞，保持 MQTT 连接存活（MQTT 库内部有 keepalive 机制）
	select {}
}
func publishHandler(client mqtt.Client, message mqtt.Message) {
	fmt.Printf("MESSAGE : %s\n", message.Payload())
	fmt.Printf("TOPIC : %s\n", message.Topic())

	topicArray := strings.Split(strings.TrimPrefix(message.Topic(), "/"), "/")
	if len(topicArray) >= 4 {
		if topicArray[3] == "ping" {
			err := models.UpdateDeviceOnlineTime(topicArray[1], topicArray[2])
			if err != nil {
				log.Printf("[DB ERROR] : %v\n", err)
			}
		}
	}
}
