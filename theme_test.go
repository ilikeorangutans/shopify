package shopify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecodeThemesList(t *testing.T) {
	themes, err := decodeThemesList([]byte(ThemesListJSON))

	assert.Nil(t, err)
	assert.Equal(t, 3, len(themes.([]*Theme)))
}

func TestDecodeTheme(t *testing.T) {
	x, err := decodeTheme([]byte(ThemeJSON))
	theme := x.(*Theme)
	assert.Nil(t, err)
	assert.NotNil(t, theme)
	assert.Equal(t, 828155753, theme.ID)
	t.Log(theme)
}

const ThemesListJSON = `
[
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
`

const ThemeJSON = `
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
`
