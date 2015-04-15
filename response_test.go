package shopify

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"
)

type res struct {
}

func TestFoo(t *testing.T) {

	var d map[string]json.RawMessage

	decoder := json.NewDecoder(strings.NewReader("{\"channels\":[ {\"id\": 3, \"shop_id\": 2} ]}"))

	decoder.Decode(&d)

	println(d["channels"])
	println(len(d))

	if d["channels"] != nil {
		var ch []*Channel
		json.Unmarshal(d["channels"], &ch)
		println("length: ", len(ch))
		first := ch[0]
		fmt.Printf("id %d shop %d \n\n", first.Id, first.ShopId)

	}

}
