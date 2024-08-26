package main

import (
	"fmt"
	"github.com/bytedance/sonic"
	"time"
)

type YourSchema struct {
	ID int
	T  time.Time `json:"birthday" form:"birthday" time_format:"2006-01-02 15:04:05" time_utc:"0"`
}

func main() {
	var data YourSchema
	data.ID = 1
	data.T = time.Now()
	// Marshal
	output, err := sonic.Marshal(&data)
	fmt.Println(string(output), err)
}
