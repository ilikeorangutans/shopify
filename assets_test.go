package shopify

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
	"time"
)

// AssertingURLBuilder returns a URLBuilder that asserts that the built URL ends with the given expected string.
func AssertingURLBuilder(t *testing.T, expected string) URLBuilder {
	return func(p ...string) string {

		joined := strings.Join(p, "/")
		assert.Equal(t, expected, joined)
		return joined
	}
}

func DummyRequestAndParse(result interface{}, err error) RequestAndParse {
	return func(req *http.Request, element string, f JSONResourceParser) (interface{}, error) {
		return result, err
	}
}

var testTheme = &Theme{
	CommonFields: CommonFields{ID: 12345},
}

func TestAssetsList(t *testing.T) {
	assets := &Assets{
		buildURL:        AssertingURLBuilder(t, "/admin/themes/12345/assets.json"),
		Theme:           testTheme,
		requestAndParse: DummyRequestAndParse([]*Asset{}, nil),
	}

	assets.List()
}

func TestAssetsGet(t *testing.T) {
	assets := &Assets{
		buildURL:        AssertingURLBuilder(t, "/admin/themes/12345/assets.json?asset[key]=templates/my-asset-key.liquid"),
		Theme:           testTheme,
		requestAndParse: DummyRequestAndParse(&Asset{DecodingComplete: make(chan bool)}, nil),
	}

	assets.Download("templates/my-asset-key.liquid")

}

func TestDownloadAll(t *testing.T) {
	assets := &Assets{
		buildURL:        AssertingURLBuilder(t, "/admin/themes/12345/assets.json?fields=key,value,attachment"),
		Theme:           testTheme,
		requestAndParse: DummyRequestAndParse([]*Asset{}, nil),
	}

	assets.DownloadAll()
}

func TestDecodeAssetsList(t *testing.T) {

	x, err := decodeAssetsList([]byte(AssetListJSON))

	assets := x.([]*Asset)
	assert.Nil(t, err)
	assert.Equal(t, 23, len(assets))
}

func TestDecodeAsset(t *testing.T) {
	x, err := decodeAsset([]byte(SingleAssetWithBase64Attachment))
	asset := x.(*Asset)
	select {
	case <-asset.DecodingComplete:
	case <-time.After(time.Duration(50 * time.Millisecond)):
		t.Fatal("Decoding should be complete by now; if this test fails it might indicate that the actual decoding operation was too slow. This is not necessarily a failure!")
	}

	assert.Nil(t, err)
	assert.Equal(t, "assets/arrow-dark.png", asset.Key)
	assert.True(t, asset.HasAttachment())
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

const SingleAssetWithBase64Attachment = `
{"key":"assets\/arrow-dark.png","public_url":"https:\/\/cdn.shopify.com\/s\/files\/1\/0761\/2111\/t\/1\/assets\/arrow-dark.png?4540024002816414779","attachment":"iVBORw0KGgoAAAANSUhEUgAAAAcAAAAECAYAAABCxiV9AAAAGXRFWHRTb2Z0\nd2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAyRpVFh0WE1MOmNvbS5hZG9i\nZS54bXAAAAAAADw\/eHBhY2tldCBiZWdpbj0i77u\/IiBpZD0iVzVNME1wQ2Vo\naUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6\nbnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuMC1jMDYxIDY0\nLjE0MDk0OSwgMjAxMC8xMi8wNy0xMDo1NzowMSAgICAgICAgIj4gPHJkZjpS\nREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJk\nZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIg\neG1sbnM6eG1wPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvIiB4bWxu\nczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1s\nbnM6c3RSZWY9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9S\nZXNvdXJjZVJlZiMiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9w\nIENTNS4xIE1hY2ludG9zaCIgeG1wTU06SW5zdGFuY2VJRD0ieG1wLmlpZDo1\nMUM5QzRCMDZGOTUxMUUxQTAzQ0U3RDQ4RjU1M0ZDQiIgeG1wTU06RG9jdW1l\nbnRJRD0ieG1wLmRpZDo1MUM5QzRCMTZGOTUxMUUxQTAzQ0U3RDQ4RjU1M0ZD\nQiI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAu\naWlkOjIyQTZFOUVCNkY5NTExRTFBMDNDRTdENDhGNTUzRkNCIiBzdFJlZjpk\nb2N1bWVudElEPSJ4bXAuZGlkOjIyQTZFOUVDNkY5NTExRTFBMDNDRTdENDhG\nNTUzRkNCIi8+IDwvcmRmOkRlc2NyaXB0aW9uPiA8L3JkZjpSREY+IDwveDp4\nbXBtZXRhPiA8P3hwYWNrZXQgZW5kPSJyIj8+21NY3QAAAChJREFUeNpiYGBg\naABihv\/\/\/6NgmDgDugJkCRQFKBJQlQwYEkAAEGAAgLIYbJBud\/kAAAAASUVO\nRK5CYII=\n","created_at":"2015-02-02T13:15:14-05:00","updated_at":"2015-02-02T13:15:14-05:00","content_type":"image\/png","size":950,"theme_id":9751085,"warnings":[]}
`
