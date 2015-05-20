package shopify

import (
	"testing"
)

func TestDecodeTransactions(t *testing.T) {
	json := `[{"amount":"1058.00","authorization":"2008873","created_at":"2015-04-14T17:23:09+02:00","currency":"CAD","gateway":"bogus","id":2,"kind":"authorization","location_id":null,"message":null,"order_id":2,"parent_id":null,"status":"success","test":true,"user_id":null,"device_id":null,"receipt":{"auth":"123456"},"error_code":null,"source_name":"web","payment_details":{"avs_result_code":"M","credit_card_bin":"1","cvv_result_code":"M","credit_card_number":"•••• •••• •••• 1","credit_card_company":"Bogus"}}]`
	transactions := decodeTransactionsList([]byte(json))
	if len(transactions) != 1 {
		t.Fatalf("Expected 1 transaction but got %d\n", len(transactions))
	}
}
