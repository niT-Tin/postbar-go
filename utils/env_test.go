package utils

import (
	"fmt"
	"testing"
)

func TestGetMongoENV(t *testing.T) {
	host, port, err := GetMongoENV()
	if err != nil || len(host) == 0 || len(port) == 0 {
		t.Errorf("error encountered getting envs: %v", err)
	}
	fmt.Println("host：", host, "port：", port)
}
