package helpers

import (
	"encoding/json"
	"fmt"

	"github.com/tidwall/pretty"
)

// รับ data เพื่อ print เป็น json
func ConvertToJson(v ...interface{}) {
	bytes, _ := json.MarshalIndent(v, "", "\t")
	fmt.Println(string(pretty.Color(pretty.PrettyOptions(bytes, pretty.DefaultOptions), nil)))
}
