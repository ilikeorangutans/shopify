package shopify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProductsList(t *testing.T) {
	products := &Products{
		RemoteJSONResource: &ShopifyRemoteJSONResource{
			URLBuilder:     &ShopifyAdminURLBuilder{},
			RemoteResource: NewTestRemote(t).ReturnsBody([]byte(productsListJSON)).ExpectsURL("/admin/products?limit=50"),
		},
	}
	var paginator Paginator

	prods, err := products.List(paginator)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(prods))
}

func TestProductsGet(t *testing.T) {
	products := &Products{
		RemoteJSONResource: &ShopifyRemoteJSONResource{
			URLBuilder:     &ShopifyAdminURLBuilder{},
			RemoteResource: NewTestRemote(t).ReturnsBody([]byte(singleProductJSON)).ExpectsURL("/admin/products/632910392"),
		},
	}

	prod, err := products.Get(632910392)

	assert.Nil(t, err)
	assert.Equal(t, 632910392, prod.ID)
	assert.Equal(t, "IPod Nano - 8GB", prod.Title)
	assert.Equal(t, 4, len(prod.Variants))
}

const productsListJSON = `
{
  "products":[
    {
      "id": 632910392,
      "title": "IPod Nano - 8GB",
      "body_html": "<p>It's the small iPod with one very big idea: Video. Now the world's most popular music player, available in 4GB and 8GB models, lets you enjoy TV shows, movies, video podcasts, and more. The larger, brighter display means amazing picture quality. In six eye-catching colors, iPod nano is stunning all around. And with models starting at just $149, little speaks volumes.<\/p>",
      "vendor": "Apple",
      "product_type": "Cult Products",
      "created_at": "2015-09-02T14:50:32-04:00",
      "handle": "ipod-nano",
      "updated_at": "2015-09-02T14:50:32-04:00",
      "published_at": "2007-12-31T19:00:00-05:00",
      "template_suffix": null,
      "published_scope": "web",
      "tags": "Emotive, Flash Memory, MP3, Music",
      "variants": [
        {
          "id": 808950810,
          "product_id": 632910392,
          "title": "Pink",
          "sku": "IPOD2008PINK",
          "position": 1,
          "grams": 200,
          "inventory_policy": "continue",
          "fulfillment_service": "manual",
          "inventory_management": "shopify",
          "price": "199.00",
          "compare_at_price": null,
          "option1": "Pink",
          "option2": null,
          "option3": null,
          "created_at": "2015-09-02T14:50:32-04:00",
          "updated_at": "2015-09-02T14:50:32-04:00",
          "taxable": true,
          "requires_shipping": true,
          "barcode": "1234_pink",
          "inventory_quantity": 10,
          "old_inventory_quantity": 10,
          "image_id": 562641783,
          "weight": 0.2,
          "weight_unit": "kg"
        },
        {
          "id": 49148385,
          "product_id": 632910392,
          "title": "Red",
          "sku": "IPOD2008RED",
          "position": 2,
          "grams": 200,
          "inventory_policy": "continue",
          "fulfillment_service": "manual",
          "inventory_management": "shopify",
          "price": "199.00",
          "compare_at_price": null,
          "option1": "Red",
          "option2": null,
          "option3": null,
          "created_at": "2015-09-02T14:50:32-04:00",
          "updated_at": "2015-09-02T14:50:32-04:00",
          "taxable": true,
          "requires_shipping": true,
          "barcode": "1234_red",
          "inventory_quantity": 20,
          "old_inventory_quantity": 20,
          "image_id": null,
          "weight": 0.2,
          "weight_unit": "kg"
        },
        {
          "id": 39072856,
          "product_id": 632910392,
          "title": "Green",
          "sku": "IPOD2008GREEN",
          "position": 3,
          "grams": 200,
          "inventory_policy": "continue",
          "fulfillment_service": "manual",
          "inventory_management": "shopify",
          "price": "199.00",
          "compare_at_price": null,
          "option1": "Green",
          "option2": null,
          "option3": null,
          "created_at": "2015-09-02T14:50:32-04:00",
          "updated_at": "2015-09-02T14:50:32-04:00",
          "taxable": true,
          "requires_shipping": true,
          "barcode": "1234_green",
          "inventory_quantity": 30,
          "old_inventory_quantity": 30,
          "image_id": null,
          "weight": 0.2,
          "weight_unit": "kg"
        },
        {
          "id": 457924702,
          "product_id": 632910392,
          "title": "Black",
          "sku": "IPOD2008BLACK",
          "position": 4,
          "grams": 200,
          "inventory_policy": "continue",
          "fulfillment_service": "manual",
          "inventory_management": "shopify",
          "price": "199.00",
          "compare_at_price": null,
          "option1": "Black",
          "option2": null,
          "option3": null,
          "created_at": "2015-09-02T14:50:32-04:00",
          "updated_at": "2015-09-02T14:50:32-04:00",
          "taxable": true,
          "requires_shipping": true,
          "barcode": "1234_black",
          "inventory_quantity": 40,
          "old_inventory_quantity": 40,
          "image_id": null,
          "weight": 0.2,
          "weight_unit": "kg"
        }
      ],
      "options": [
        {
          "id": 594680422,
          "product_id": 632910392,
          "name": "Color",
          "position": 1,
          "values": [
            "Pink",
            "Red",
            "Green",
            "Black"
          ]
        }
      ],
      "images": [
        {
          "id": 850703190,
          "product_id": 632910392,
          "position": 1,
          "created_at": "2015-09-02T14:50:32-04:00",
          "updated_at": "2015-09-02T14:50:32-04:00",
          "src": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano.png?v=1441219832",
          "variant_ids": [
          ]
        },
        {
          "id": 562641783,
          "product_id": 632910392,
          "position": 2,
          "created_at": "2015-09-02T14:50:32-04:00",
          "updated_at": "2015-09-02T14:50:32-04:00",
          "src": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano-2.png?v=1441219832",
          "variant_ids": [
            808950810
          ]
        }
      ],
      "image": {
        "id": 850703190,
        "product_id": 632910392,
        "position": 1,
        "created_at": "2015-09-02T14:50:32-04:00",
        "updated_at": "2015-09-02T14:50:32-04:00",
        "src": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano.png?v=1441219832",
        "variant_ids": [
        ]
      }
    },
    {
      "id": 921728736,
      "title": "IPod Touch 8GB",
      "body_html": "<p>The iPod Touch has the iPhone's multi-touch interface, with a physical home button off the touch screen. The home screen has a list of buttons for the available applications.<\/p>",
      "vendor": "Apple",
      "product_type": "Cult Products",
      "created_at": "2015-09-02T14:50:32-04:00",
      "handle": "ipod-touch",
      "updated_at": "2015-09-02T14:50:32-04:00",
      "published_at": "2008-09-25T20:00:00-04:00",
      "template_suffix": null,
      "published_scope": "global",
      "tags": "",
      "variants": [
        {
          "id": 447654529,
          "product_id": 921728736,
          "title": "Black",
          "sku": "IPOD2009BLACK",
          "position": 1,
          "grams": 200,
          "inventory_policy": "continue",
          "fulfillment_service": "manual",
          "inventory_management": "shopify",
          "price": "199.00",
          "compare_at_price": null,
          "option1": "Black",
          "option2": null,
          "option3": null,
          "created_at": "2015-09-02T14:50:32-04:00",
          "updated_at": "2015-09-02T14:50:32-04:00",
          "taxable": true,
          "requires_shipping": true,
          "barcode": "1234_black",
          "inventory_quantity": 13,
          "old_inventory_quantity": 13,
          "image_id": null,
          "weight": 0.2,
          "weight_unit": "kg"
        }
      ],
      "options": [
        {
          "id": 891236591,
          "product_id": 921728736,
          "name": "Title",
          "position": 1,
          "values": [
            "Black"
          ]
        }
      ],
      "images": [
      ],
      "image": null
    }
  ]
}
`

const singleProductJSON = `
{
  "product": {
    "id": 632910392,
    "title": "IPod Nano - 8GB",
    "body_html": "<p>It's the small iPod with one very big idea: Video. Now the world's most popular music player, available in 4GB and 8GB models, lets you enjoy TV shows, movies, video podcasts, and more. The larger, brighter display means amazing picture quality. In six eye-catching colors, iPod nano is stunning all around. And with models starting at just $149, little speaks volumes.<\/p>",
    "vendor": "Apple",
    "product_type": "Cult Products",
    "created_at": "2015-09-02T14:50:32-04:00",
    "handle": "ipod-nano",
    "updated_at": "2015-09-02T14:50:32-04:00",
    "published_at": "2007-12-31T19:00:00-05:00",
    "template_suffix": null,
    "published_scope": "web",
    "tags": "Emotive, Flash Memory, MP3, Music",
    "variants": [
      {
        "id": 808950810,
        "product_id": 632910392,
        "title": "Pink",
        "sku": "IPOD2008PINK",
        "position": 1,
        "grams": 200,
        "inventory_policy": "continue",
        "fulfillment_service": "manual",
        "inventory_management": "shopify",
        "price": "199.00",
        "compare_at_price": null,
        "option1": "Pink",
        "option2": null,
        "option3": null,
        "created_at": "2015-09-02T14:50:32-04:00",
        "updated_at": "2015-09-02T14:50:32-04:00",
        "taxable": true,
        "requires_shipping": true,
        "barcode": "1234_pink",
        "inventory_quantity": 10,
        "old_inventory_quantity": 10,
        "image_id": 562641783,
        "weight": 0.2,
        "weight_unit": "kg"
      },
      {
        "id": 49148385,
        "product_id": 632910392,
        "title": "Red",
        "sku": "IPOD2008RED",
        "position": 2,
        "grams": 200,
        "inventory_policy": "continue",
        "fulfillment_service": "manual",
        "inventory_management": "shopify",
        "price": "199.00",
        "compare_at_price": null,
        "option1": "Red",
        "option2": null,
        "option3": null,
        "created_at": "2015-09-02T14:50:32-04:00",
        "updated_at": "2015-09-02T14:50:32-04:00",
        "taxable": true,
        "requires_shipping": true,
        "barcode": "1234_red",
        "inventory_quantity": 20,
        "old_inventory_quantity": 20,
        "image_id": null,
        "weight": 0.2,
        "weight_unit": "kg"
      },
      {
        "id": 39072856,
        "product_id": 632910392,
        "title": "Green",
        "sku": "IPOD2008GREEN",
        "position": 3,
        "grams": 200,
        "inventory_policy": "continue",
        "fulfillment_service": "manual",
        "inventory_management": "shopify",
        "price": "199.00",
        "compare_at_price": null,
        "option1": "Green",
        "option2": null,
        "option3": null,
        "created_at": "2015-09-02T14:50:32-04:00",
        "updated_at": "2015-09-02T14:50:32-04:00",
        "taxable": true,
        "requires_shipping": true,
        "barcode": "1234_green",
        "inventory_quantity": 30,
        "old_inventory_quantity": 30,
        "image_id": null,
        "weight": 0.2,
        "weight_unit": "kg"
      },
      {
        "id": 457924702,
        "product_id": 632910392,
        "title": "Black",
        "sku": "IPOD2008BLACK",
        "position": 4,
        "grams": 200,
        "inventory_policy": "continue",
        "fulfillment_service": "manual",
        "inventory_management": "shopify",
        "price": "199.00",
        "compare_at_price": null,
        "option1": "Black",
        "option2": null,
        "option3": null,
        "created_at": "2015-09-02T14:50:32-04:00",
        "updated_at": "2015-09-02T14:50:32-04:00",
        "taxable": true,
        "requires_shipping": true,
        "barcode": "1234_black",
        "inventory_quantity": 40,
        "old_inventory_quantity": 40,
        "image_id": null,
        "weight": 0.2,
        "weight_unit": "kg"
      }
    ],
    "options": [
      {
        "id": 594680422,
        "product_id": 632910392,
        "name": "Color",
        "position": 1,
        "values": [
          "Pink",
          "Red",
          "Green",
          "Black"
        ]
      }
    ],
    "images": [
      {
        "id": 850703190,
        "product_id": 632910392,
        "position": 1,
        "created_at": "2015-09-02T14:50:32-04:00",
        "updated_at": "2015-09-02T14:50:32-04:00",
        "src": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano.png?v=1441219832",
        "variant_ids": [
        ]
      },
      {
        "id": 562641783,
        "product_id": 632910392,
        "position": 2,
        "created_at": "2015-09-02T14:50:32-04:00",
        "updated_at": "2015-09-02T14:50:32-04:00",
        "src": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano-2.png?v=1441219832",
        "variant_ids": [
          808950810
        ]
      }
    ],
    "image": {
      "id": 850703190,
      "product_id": 632910392,
      "position": 1,
      "created_at": "2015-09-02T14:50:32-04:00",
      "updated_at": "2015-09-02T14:50:32-04:00",
      "src": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/products\/ipod-nano.png?v=1441219832",
      "variant_ids": [
      ]
    }
  }
}
`
