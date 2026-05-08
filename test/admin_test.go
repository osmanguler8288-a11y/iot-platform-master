package test

import (
	"fmt"
	"iot-platform-master/helper"
	"testing"
)

const adminServiceAdder = "http://127.0.0.1:14010"

func TestDeviceList(t *testing.T) {
	url := adminServiceAdder + "/device/list?page=1&size=20&name="
	rep, err := helper.HttpGet(url)
	if err != nil {
		t.Fatalf("请求失败: %v", err) // 去掉 args...:
	}
	fmt.Println("响应结果:", string(rep)) // 去掉 a...:
}
