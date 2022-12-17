package src

import (
	"html/template"
	"path/filepath"
	"strings"

	"github.com/ktnyt/go-moji"
)

const FILE_PATH = "../templates"

func Template(fileName string, obj any) (string, error) {
	path := filepath.Join(FILE_PATH, fileName)

	templ, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}

	writer := new(strings.Builder)

	err = templ.Execute(writer, obj)
	if err != nil {
		return "", err
	}

	return writer.String(), nil
}

// 文字列をイイカンジに整形する
//
// - 全ての改行を消す (isDeleteReturnsがtrueの場合)
// - 全角英数字を半角英数字に変換する
// - 全角スペースを半角スペースに変換する
func Convert(text string, isDeleteReturns bool) string {
	if isDeleteReturns {
		text = strings.ReplaceAll(text, "\n", "")
	}

	return moji.Convert(
		moji.Convert(
			text,
			moji.ZS,
			moji.HS,
		),
		moji.ZE,
		moji.HE,
	)
}
