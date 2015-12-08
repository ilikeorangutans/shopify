package shopify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestThemesList(t *testing.T) {
	themes := &Themes{
		RemoteJSONResource: &ShopifyRemoteJSONResource{
			URLBuilder: &ShopifyAdminURLBuilder{},
			RemoteResource: &TestRemoteResource{
				body: []byte(ThemesListJSON),
			},
		},
	}

	th, err := themes.List()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(th))
}

func TestThemesGet(t *testing.T) {
	themes := &Themes{
		RemoteJSONResource: &ShopifyRemoteJSONResource{
			URLBuilder: &ShopifyAdminURLBuilder{},
			RemoteResource: &TestRemoteResource{
				body: []byte(ThemeJSON),
			},
		},
	}

	th, err := themes.Get(828155753)
	assert.Nil(t, err)
	assert.Equal(t, 828155753, th.ID)
}

const ThemesListJSON = `
{
  "themes":[
    {
      "id": 828155753,
      "name": "Comfort",
      "created_at": "2015-09-02T14:52:56-04:00",
      "updated_at": "2015-09-02T14:52:56-04:00",
      "role": "main",
      "theme_store_id": null,
      "previewable": true,
      "processing": false
    },
    {
      "id": 976877075,
      "name": "Speed",
      "created_at": "2015-09-02T14:52:56-04:00",
      "updated_at": "2015-09-02T14:52:56-04:00",
      "role": "mobile",
      "theme_store_id": null,
      "previewable": true,
      "processing": false
    },
    {
      "id": 752253240,
      "name": "Sandbox",
      "created_at": "2015-09-02T14:52:56-04:00",
      "updated_at": "2015-09-02T14:52:56-04:00",
      "role": "unpublished",
      "theme_store_id": null,
      "previewable": true,
      "processing": false
    }
  ]
}
`

const ThemeJSON = `
{ "theme":
  {
    "id": 828155753,
    "name": "Comfort",
    "created_at": "2015-09-02T14:52:56-04:00",
    "updated_at": "2015-09-02T14:52:56-04:00",
    "role": "main",
    "theme_store_id": null,
    "previewable": true,
    "processing": false
  }
}
`
