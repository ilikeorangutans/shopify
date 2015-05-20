package shopify

import (
	"testing"
)

func TestDecode(t *testing.T) {
	jsonToDecode := `[
		{"credential1":null,"email":"test@test.com","handle":"custom","id":1,"include_pending_stock":false,"name":"custom","service_name":"custom","inventory_management":false,"tracking_support":true,"provider_id":4,"credential2_exists":false},
		{"credential1":"credential","email":"test@test.com","handle":"test","id":2,"include_pending_stock":false,"name":"TEST","requires_shipping_method":false,"service_name":"Test","inventory_management":false,"tracking_support":false,"provider_id":null,"credential2_exists":false}]`

	services := decodeServicesJSON([]byte(jsonToDecode))

	t.Logf("Found %d ", len(services))

	if len(services) != 2 {
		t.Fail()
	}
}
