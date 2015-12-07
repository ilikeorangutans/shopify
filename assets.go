package shopify

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
)

type Assets struct {
	RemoteJSONResource
	Theme *Theme
}

// List downloads metadata for all assets associated with the Theme set on the instance.
func (a *Assets) List() ([]*Asset, error) {

	req, err := http.NewRequest("GET", a.BuildURL(a.themeBaseURL(), "assets.json"), nil)
	if err != nil {
		return nil, err
	}

	var assets []*Asset
	if err := a.RequestAndDecode(req, "assets", assets); err != nil {
		return nil, err
	}
	return assets, nil
}

type AttachmentRetrieval struct {
	Asset *Asset
	Error error
}

// DownloadAll downloads all assets including their attachments. This can cause large requests!
func (a *Assets) DownloadAll() ([]*Asset, error) {
	req, err := http.NewRequest("GET", a.BuildURL(a.themeBaseURL(), "assets.json?fields=key,value,attachment"), nil)
	if err != nil {
		return nil, err
	}
	var assets []*Asset
	if err := a.RequestAndDecode(req, "assets", assets); err != nil {
		return nil, err
	}
	return assets, nil
}

// Download downloads a single Asset identified by the given key with all its data.
func (a *Assets) Download(key string) (*Asset, error) {
	req, err := http.NewRequest("GET", a.BuildURL(a.themeBaseURL(), fmt.Sprintf("assets.json?asset[key]=%s", key)), nil)
	if err != nil {
		return nil, err
	}
	var asset *Asset
	if err := a.RequestAndDecode(req, "asset", asset); err != nil {
		return nil, err
	}
	asset.decodeAttachment()

	return asset, nil
}

func (a *Assets) themeBaseURL() string {
	return fmt.Sprintf("/admin/themes/%d", a.Theme.ID)
}

type Asset struct {
	Timestamps

	Key         string `json:"key"`
	ContentType string `json:"content_type"`
	PublicURL   string `json:"public_url"`
	Size        int    `json:"size"`
	ThemeID     int64  `json:"theme_id"`
	Value       string `json:"value"`
	// Attachment holds the binary attachment of this asset, if available. Note that you should check the
	// DecodingComplete channel on this asset to ensure decoding is complete.
	Attachment []byte `json:"-"`
	// EncodedAttachment holds a base64 encoded representation of the attachment.
	EncodedAttachment string `json:"attachment"`
	// DecodingComplete is a channel that blocks until decoding of this asset's attachment is complete.
	DecodingComplete chan bool `json:"-"`
	EncodingComplete chan bool `json:"-"`
}

func (a *Asset) HasAttachment() bool {
	return len(a.Attachment) > 0 || len(a.EncodedAttachment) > 0
}

func (a *Asset) String() string {
	return fmt.Sprintf("Asset{key: %s, content_type: %s, size: %d}", a.Key, a.ContentType, a.Size)
}

func (a *Asset) decodeAttachment() error {
	if len(a.EncodedAttachment) == 0 {
		close(a.DecodingComplete)
		return nil
	}
	b, err := base64.StdEncoding.DecodeString(a.EncodedAttachment)
	if err != nil {
		close(a.DecodingComplete)
		return err
	}
	if a.Size != len(b) {
		close(a.DecodingComplete)
		return fmt.Errorf("Attachment length does not match expected value, expected %d bytes but got %d", a.Size, len(b))
	}
	a.Attachment = b
	a.DecodingComplete <- true
	close(a.DecodingComplete)
	return nil
}

func decodeAssetsList(body []byte) (interface{}, error) {
	var assets []*Asset
	err := json.Unmarshal(body, &assets)
	if err != nil {
		return nil, err
	}

	for i := range assets {
		asset := assets[i]
		asset.DecodingComplete = make(chan bool)
		go asset.decodeAttachment()
	}

	return assets, nil
}

func decodeAsset(body []byte) (interface{}, error) {
	var asset *Asset
	err := json.Unmarshal(body, &asset)
	if err != nil {
		return nil, err
	}
	asset.DecodingComplete = make(chan bool)
	go asset.decodeAttachment()
	return asset, nil
}
