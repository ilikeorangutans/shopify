package shopify

import (
	"log"
	"testing"
)

const listPaylod = `[{"channel_id":2,"created_at":"2015-04-28T17:38:50+02:00","id":36,"product_id":36,"shop_id":2,"updated_at":"2015-04-28T17:38:50+02:00","body_html":"","handle":"bam2","product_type":"","published":true,"published_at":"2015-04-28T17:38:00+02:00","tags":"","title":"Bam2","vendor":"Shop One","variants":[{"barcode":"","compare_at_price":null,"created_at":"2015-04-28T17:38:50+02:00","fulfillment_service":"manual","grams":0,"id":38,"inventory_management":null,"inventory_policy":"deny","option1":"Default Title","option2":null,"option3":null,"position":1,"price":"0.00","product_id":36,"requires_shipping":true,"sku":"","taxable":true,"title":"Default Title","updated_at":"2015-04-28T17:38:50+02:00","inventory_quantity":1,"old_inventory_quantity":1,"image_id":null,"weight":0.0,"weight_unit":"kg"}],"options":[{"id":36,"name":"Title","position":1,"product_id":36}],"images":[]},{"channel_id":2,"created_at":"2015-04-28T17:38:21+02:00","id":35,"product_id":35,"shop_id":2,"updated_at":"2015-04-28T17:38:21+02:00","body_html":"","handle":"bam-1","product_type":"","published":true,"published_at":"2015-04-28T17:38:00+02:00","tags":"","title":"Bam!","vendor":"Shop One","variants":[{"barcode":"","compare_at_price":null,"created_at":"2015-04-28T17:38:21+02:00","fulfillment_service":"manual","grams":0,"id":37,"inventory_management":null,"inventory_policy":"deny","option1":"Default Title","option2":null,"option3":null,"position":1,"price":"0.00","product_id":35,"requires_shipping":true,"sku":"","taxable":true,"title":"Default Title","updated_at":"2015-04-28T17:38:21+02:00","inventory_quantity":1,"old_inventory_quantity":1,"image_id":null,"weight":0.0,"weight_unit":"kg"}],"options":[{"id":35,"name":"Title","position":1,"product_id":35}],"images":[]},{"channel_id":2,"created_at":"2015-04-14T16:50:52+02:00","id":2,"product_id":2,"shop_id":2,"updated_at":"2015-04-14T16:50:52+02:00","body_html":null,"handle":"small-steel-pants","product_type":"expedite bleeding-edge mindshare","published":true,"published_at":"2015-04-14T16:50:46+02:00","tags":"","title":"Small Steel Pants","vendor":"Eichmann, Christiansen and Weimann","variants":[{"barcode":null,"compare_at_price":null,"created_at":"2015-04-14T16:50:46+02:00","fulfillment_service":"manual","grams":1800,"id":3,"inventory_management":"shopify","inventory_policy":"deny","option1":"Orchid","option2":null,"option3":null,"position":1,"price":"1245.99","product_id":2,"requires_shipping":true,"sku":"","taxable":true,"title":"Orchid","updated_at":"2015-04-14T16:50:46+02:00","inventory_quantity":934,"old_inventory_quantity":934,"image_id":null,"weight":1.8,"weight_unit":"kg"},{"barcode":null,"compare_at_price":null,"created_at":"2015-04-14T16:50:46+02:00","fulfillment_service":"manual","grams":1800,"id":4,"inventory_management":null,"inventory_policy":"deny","option1":"Olive","option2":null,"option3":null,"position":2,"price":"1245.99","product_id":2,"requires_shipping":true,"sku":"","taxable":true,"title":"Olive","updated_at":"2015-04-14T16:50:46+02:00","inventory_quantity":1,"old_inventory_quantity":1,"image_id":null,"weight":1.8,"weight_unit":"kg"}],"options":[{"id":2,"name":"Color or something","position":1,"product_id":2}],"images":[]}]`

func TestList(t *testing.T) {
	pp := &ProductPublications{}
	pps := pp.parseList([]byte(listPaylod))
	log.Println(pps)

	if len(pps) != 3 {
		t.Fatalf("Expect 2 publications but got %d\n", len(pps))
	}
}
