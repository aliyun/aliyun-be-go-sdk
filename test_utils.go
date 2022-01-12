package be

import (
	"encoding/json"
	"fmt"
)

func PrintResult(i interface{}) {
	iByte, _ := json.MarshalIndent(i, "", "\t")
	fmt.Println(string(iByte))
}
