package shopify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecodeAssets(t *testing.T) {

	x, err := decodeAssetsList([]byte(AssetListJSON))

	assets := x.([]*Asset)
	assert.Nil(t, err)
	assert.Equal(t, 23, len(assets))
}

const AssetListJSON = `
[
    {
      "key": "assets\/bg-body-green.gif",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/bg-body-green.gif?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "image\/gif",
      "size": 1542,
      "theme_id": 828155753
    },
    {
      "key": "assets\/bg-body-orange.gif",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/bg-body-orange.gif?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "image\/gif",
      "size": 1548,
      "theme_id": 828155753
    },
    {
      "key": "assets\/bg-body-pink.gif",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/bg-body-pink.gif?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "image\/gif",
      "size": 1562,
      "theme_id": 828155753
    },
    {
      "key": "assets\/bg-body.gif",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/bg-body.gif?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "image\/gif",
      "size": 1571,
      "theme_id": 828155753
    },
    {
      "key": "assets\/bg-content.gif",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/bg-content.gif?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "image\/gif",
      "size": 134,
      "theme_id": 828155753
    },
    {
      "key": "assets\/bg-footer.gif",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/bg-footer.gif?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "image\/gif",
      "size": 1434,
      "theme_id": 828155753
    },
    {
      "key": "assets\/bg-main.gif",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/bg-main.gif?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "image\/gif",
      "size": 297,
      "theme_id": 828155753
    },
    {
      "key": "assets\/bg-sidebar.gif",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/bg-sidebar.gif?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "image\/gif",
      "size": 124,
      "theme_id": 828155753
    },
    {
      "key": "assets\/shop.css",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/shop.css?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "text\/css",
      "size": 14058,
      "theme_id": 828155753
    },
    {
      "key": "assets\/shop.css.liquid",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/shop.css.liquid?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "text\/x-liquid",
      "size": 14675,
      "theme_id": 828155753
    },
    {
      "key": "assets\/shop.js",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/shop.js?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "application\/javascript",
      "size": 348,
      "theme_id": 828155753
    },
    {
      "key": "assets\/sidebar-devider.gif",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/sidebar-devider.gif?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "image\/gif",
      "size": 1016,
      "theme_id": 828155753
    },
    {
      "key": "assets\/sidebar-menu.jpg",
      "public_url": "https:\/\/cdn.shopify.com\/s\/files\/1\/0006\/9093\/3842\/t\/1\/assets\/sidebar-menu.jpg?7364251289859092031",
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "image\/jpeg",
      "size": 1609,
      "theme_id": 828155753
    },
    {
      "key": "config\/settings_data.json",
      "public_url": null,
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "application\/json",
      "size": 4570,
      "theme_id": 828155753
    },
    {
      "key": "config\/settings_schema.json",
      "public_url": null,
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "application\/json",
      "size": 4570,
      "theme_id": 828155753
    },
    {
      "key": "layout\/theme.liquid",
      "public_url": null,
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "text\/x-liquid",
      "size": 3252,
      "theme_id": 828155753
    },
    {
      "key": "templates\/article.liquid",
      "public_url": null,
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "text\/x-liquid",
      "size": 2486,
      "theme_id": 828155753
    },
    {
      "key": "templates\/blog.liquid",
      "public_url": null,
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "text\/x-liquid",
      "size": 786,
      "theme_id": 828155753
    },
    {
      "key": "templates\/cart.liquid",
      "public_url": null,
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "text\/x-liquid",
      "size": 2047,
      "theme_id": 828155753
    },
    {
      "key": "templates\/collection.liquid",
      "public_url": null,
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "text\/x-liquid",
      "size": 946,
      "theme_id": 828155753
    },
    {
      "key": "templates\/index.liquid",
      "public_url": null,
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "text\/x-liquid",
      "size": 1068,
      "theme_id": 828155753
    },
    {
      "key": "templates\/page.liquid",
      "public_url": null,
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "text\/x-liquid",
      "size": 147,
      "theme_id": 828155753
    },
    {
      "key": "templates\/product.liquid",
      "public_url": null,
      "created_at": "2010-07-12T15:31:50-04:00",
      "updated_at": "2010-07-12T15:31:50-04:00",
      "content_type": "text\/x-liquid",
      "size": 2796,
      "theme_id": 828155753
    }
  ]
`
