package shopify

import (
	"testing"
)

func TestDecodeOrders(t *testing.T) {
	jsonPayload := `[{"buyer_accepts_marketing":false,"cancel_reason":null,"cancelled_at":null,"cart_token":"7c9907c604010b2b6d8f0bb7b910547f","checkout_token":"433f781dcbe3a35d4f68b88926ea238b","closed_at":null,"confirmed":true,"created_at":"2015-04-14T17:23:09+02:00","currency":"CAD","device_id":null,"email":"foo@test.com","financial_status":"authorized","gateway":"bogus","id":2,"landing_site":"\/","location_id":null,"name":"#1002","note":"","number":2,"processed_at":"2015-04-14T17:23:09+02:00","reference":null,"referring_site":"","source_identifier":null,"source_url":null,"subtotal_price":"1000.00","taxes_included":false,"test":true,"token":"7a464fdd09f283fb4f9d4f18076de69d","total_discounts":"0.00","total_line_items_price":"1000.00","total_price":"1058.00","total_price_usd":"849.28","total_tax":"50.00","total_weight":0,"updated_at":"2015-04-22T22:41:25+02:00","user_id":null,"browser_ip":"127.0.0.1","landing_site_ref":null,"order_number":1002,"discount_codes":[],"note_attributes":[],"processing_method":"direct","source":"checkout_next","checkout_id":2,"source_name":"web","fulfillment_status":null,"tax_lines":[{"price":"50.00","rate":0.05,"title":"GST"}],"tags":"","line_items":[{"fulfillment_service":"manual","fulfillment_status":null,"gift_card":false,"grams":0,"id":2,"price":"1000.00","product_id":null,"quantity":1,"requires_shipping":true,"sku":"","taxable":true,"title":"asdfasdf","variant_id":null,"variant_title":"","vendor":"Shop One","name":"asdfasdf","variant_inventory_management":null,"properties":[],"product_exists":false,"fulfillable_quantity":1,"total_discount":"0.00","tax_lines":[{"price":"50.00","rate":0.05,"title":"GST"}]}],"shipping_lines":[{"code":"Standard Shipping","price":"8.00","source":"shopify","title":"Standard Shipping","tax_lines":[{"price":"0.00","rate":0.05,"title":"GST"}]}],"billing_address":{"address1":"jaldf","address2":"asdfj","city":"asdf","company":"jkadls","country":"Canada","first_name":"asdf","last_name":"asfd","latitude":45.416311,"longitude":-75.68683,"phone":"","province":"Manitoba","zip":"a1a1a1","name":"asdf asfd","country_code":"CA","province_code":"MB"},"shipping_address":{"address1":"jaldf","address2":"asdfj","city":"asdf","company":"jkadls","country":"Canada","first_name":"asdf","last_name":"asfd","latitude":45.416311,"longitude":-75.68683,"phone":"","province":"Manitoba","zip":"a1a1a1","name":"asdf asfd","country_code":"CA","province_code":"MB"},"fulfillments":[],"client_details":{"accept_language":"en-US,en;q=0.8,de;q=0.6","browser_height":778,"browser_ip":"127.0.0.1","browser_width":1440,"session_hash":"1884410f63ba1c96d27d62e1873a3129c4cbb3151f505387fadabf7dbad987cf","user_agent":"Mozilla\/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/41.0.2272.118 Safari\/537.36"},"refunds":[],"payment_details":{"avs_result_code":"M","credit_card_bin":"1","cvv_result_code":"M","credit_card_number":"•••• •••• •••• 1","credit_card_company":"Bogus"},"customer":{"accepts_marketing":false,"created_at":"2015-04-14T17:22:58+02:00","email":"foo@test.com","first_name":"asdf","id":2,"last_name":"asfd","last_order_id":null,"multipass_identifier":null,"note":null,"orders_count":0,"state":"disabled","tax_exempt":false,"total_spent":"0.00","updated_at":"2015-04-14T17:23:10+02:00","verified_email":true,"tags":"","last_order_name":null,"default_address":{"address1":"jaldf","address2":"asdfj","city":"asdf","company":"jkadls","country":"Canada","first_name":"asdf","id":2,"last_name":"asfd","phone":"","province":"Manitoba","zip":"a1a1a1","name":"asdf asfd","province_code":"MB","country_code":"CA","country_name":"Canada","default":true}}},{"buyer_accepts_marketing":false,"cancel_reason":null,"cancelled_at":null,"cart_token":"6a8054eb2d33faa4ce35b526fd20962c","checkout_token":"2428761efe1497c79ed66dcbfe5a8dd5","closed_at":null,"confirmed":true,"created_at":"2015-04-14T17:20:52+02:00","currency":"CAD","device_id":null,"email":"test@test.com","financial_status":"authorized","gateway":"bogus","id":1,"landing_site":"\/","location_id":null,"name":"#1001","note":"","number":1,"processed_at":"2015-04-14T17:20:52+02:00","reference":null,"referring_site":"","source_identifier":null,"source_url":null,"subtotal_price":"1000.00","taxes_included":false,"test":true,"token":"4fe68a12860f2ac60cf7a5b04cf71c47","total_discounts":"0.00","total_line_items_price":"1000.00","total_price":"1138.00","total_price_usd":"913.50","total_tax":"130.00","total_weight":0,"updated_at":"2015-04-22T22:41:25+02:00","user_id":null,"browser_ip":"127.0.0.1","landing_site_ref":null,"order_number":1001,"discount_codes":[],"note_attributes":[],"processing_method":"direct","source":"checkout_next","checkout_id":1,"source_name":"web","fulfillment_status":null,"tax_lines":[{"price":"130.00","rate":0.13,"title":"HST"}],"tags":"","line_items":[{"fulfillment_service":"manual","fulfillment_status":null,"gift_card":false,"grams":0,"id":1,"price":"1000.00","product_id":null,"quantity":1,"requires_shipping":true,"sku":"","taxable":true,"title":"asdfasdf","variant_id":null,"variant_title":"","vendor":"Shop One","name":"asdfasdf","variant_inventory_management":null,"properties":[],"product_exists":false,"fulfillable_quantity":1,"total_discount":"0.00","tax_lines":[{"price":"130.00","rate":0.13,"title":"HST"}]}],"shipping_lines":[{"code":"Standard Shipping","price":"8.00","source":"shopify","title":"Standard Shipping","tax_lines":[{"price":"0.00","rate":0.13,"title":"HST"}]}],"billing_address":{"address1":"asdf","address2":"12","city":"asdf","company":"asdf","country":"Canada","first_name":"asdaf","last_name":"asdf","latitude":45.416311,"longitude":-75.68683,"phone":"","province":"Ontario","zip":"a1a1a1","name":"asdaf asdf","country_code":"CA","province_code":"ON"},"shipping_address":{"address1":"asdf","address2":"12","city":"asdf","company":"asdf","country":"Canada","first_name":"asdaf","last_name":"asdf","latitude":45.416311,"longitude":-75.68683,"phone":"","province":"Ontario","zip":"a1a1a1","name":"asdaf asdf","country_code":"CA","province_code":"ON"},"fulfillments":[],"client_details":{"accept_language":"en-US,en;q=0.8,de;q=0.6","browser_height":778,"browser_ip":"127.0.0.1","browser_width":1440,"session_hash":"1884410f63ba1c96d27d62e1873a3129c4cbb3151f505387fadabf7dbad987cf","user_agent":"Mozilla\/5.0 (Macintosh; Intel Mac OS X 10_10_3) AppleWebKit\/537.36 (KHTML, like Gecko) Chrome\/41.0.2272.118 Safari\/537.36"},"refunds":[],"payment_details":{"avs_result_code":"M","credit_card_bin":"1","cvv_result_code":"M","credit_card_number":"•••• •••• •••• 1","credit_card_company":"Bogus"},"customer":{"accepts_marketing":false,"created_at":"2015-04-14T17:20:36+02:00","email":"test@test.com","first_name":"asdaf","id":1,"last_name":"asdf","last_order_id":null,"multipass_identifier":null,"note":null,"orders_count":0,"state":"disabled","tax_exempt":false,"total_spent":"0.00","updated_at":"2015-04-14T17:20:53+02:00","verified_email":true,"tags":"","last_order_name":null,"default_address":{"address1":"asdf","address2":"12","city":"asdf","company":"asdf","country":"Canada","first_name":"asdaf","id":1,"last_name":"asdf","phone":"","province":"Ontario","zip":"a1a1a1","name":"asdaf asdf","province_code":"ON","country_code":"CA","country_name":"Canada","default":true}}}]`

	orders := decodeOrdersList([]byte(jsonPayload))

	if len(orders) != 2 {
		t.Fatalf("Expected but got %d orders\n", len(orders))
	}
}