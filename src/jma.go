package src

import (
	"net/http"
	"net/url"
	"time"
)

// HEADリクエストを送り、[Last-Modified](https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/Last-Modified)を引数と比較する
func ValidModified(url url.URL, m string) (bool, error) {
	res, err := http.Head(url.String())
	if err != nil {
		return false, err
	}

	lastModified := res.Header.Get("Last-Modified")

	return lastModified != m, nil
}

// [If-Modified-Since](https://developer.mozilla.org/ja/docs/Web/HTTP/Headers/If-Modified-Since)リクエストヘッダを付与した条件付きGETリクエストをします。
// リソースが変更されていない場合、ステータスは304で返されます。
func GetIfModifiedSince(url url.URL, d time.Time) (*http.Response, error) {
	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("If-Modified-Since", d.UTC().Format(http.TimeFormat))

	client := new(http.Client)
	return client.Do(req)
}
