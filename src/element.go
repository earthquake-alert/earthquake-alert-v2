package src

import (
	"context"
	"database/sql"
	"time"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
)

type JmaElement interface {
	// DBに格納したり、DBから関連の地震・津波情報を取得してくるメソッド
	Assembly(ctx context.Context, db *sql.DB) error

	// 震度分布図など様々な画像を生成するためのメソッド
	//
	// 並列で生成したいので`Assembly`と分けている
	SetImages() error

	// タイトルを返す
	GetTitle() string

	// 対象の基点時刻を返す
	GetTargetDate() (time.Time, error)

	// 運用種別を返す
	GetInfoType() jma.InfoType

	// 本文を返す
	GetText() (string, error)

	// 画像を返す
	GetImages() []string

	// EventIDを返す
	GetEventId() ([]int, error)
}
