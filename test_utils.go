package be

import (
	"encoding/json"
	"fmt"
)

func PrintResult(i interface{}) {
	iByte, _ := json.Marshal(i)
	fmt.Println(string(iByte))
}
