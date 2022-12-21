package src

import (
	"os"
	"path/filepath"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var C *Config

type Config struct {
	// MySQLの設定
	DatabaseConfig *mysql.Config

	// 公開先の設定ファイル名
	PublishConfigName string

	// 公開先のファイルパス
	PublishConfigPath string

	PublishConfig *PublishConfig

	// 気象庁XML電文のURL
	JmaXmlUrl string

	// これを設定するとWebサーバーはBasic認証をします
	// Webサーバーを外部に公開する際には必ず設定してください
	AuthenticationUser string
	AuthenticationPw   string
}

// 公開先の共通設定項目
type PublishConfigClientCommon struct {
	ClientName string `yaml:"name"`

	// 送信する最低震度
	// 指定しない場合、全ての地震を送信する
	MinInt jma.EarthquakeIntensity `yaml:"min_int,omitempty"`

	// 送信する対象都道府県
	// これを指定すると、この地域で震度1以上の地震を観測したときのみ送信します。
	// 指定しないと、全ての都道府県の情報を送信します。
	Areas []string `yaml:"areas,omitempty"`

	// 津波情報を送信するか
	// デフォルトfalse
	IsTsunami bool `yaml:"is_tsunami,omitempty"`
}

// 公開先の設定など
type PublishConfig struct {
	Twitter []struct {
		// TODO: Twitter API v2のトークンどうなんだろう…？
		Token string `yaml:"token"`

		PublishConfigClientCommon `yaml:",inline"`
	} `yaml:"twitter,omitempty"`

	Discord []struct {
		// DiscordWebhook URL
		WebhookURL string `yaml:"webhook_url"`

		PublishConfigClientCommon `yaml:",inline"`
	} `yaml:"discord,omitempty"`

	Slack []struct {
		// SlackWebhook URL
		WebhookURL string `yaml:"webhook_url"`
		// 送信するSlackチャンネル
		Channel string `yaml:"channel"`

		PublishConfigClientCommon `yaml:",inline"`
	} `yaml:"slack,omitempty"`

	LineNotify []struct {
		// LINE Notify トークン
		Token string `yaml:"token"`

		PublishConfigClientCommon `yaml:",inline"`
	} `yaml:"line_notify,omitempty"`
}

var LocalConfig = &Config{
	DatabaseConfig: &mysql.Config{
		DBName:               "earthquake-alert",
		User:                 "docker",
		Passwd:               "docker",
		Addr:                 "localhost:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	},
	PublishConfigName: "publish_config.yaml",
	PublishConfigPath: ".",
	PublishConfig:     nil,

	JmaXmlUrl: "",

	AuthenticationUser: "user",
	AuthenticationPw:   "password",
}
var TestConfig = &Config{
	DatabaseConfig: &mysql.Config{
		DBName:               "earthquake-alert-test",
		User:                 "docker",
		Passwd:               "docker",
		Addr:                 "localhost:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	},
	PublishConfigName: "publish_config_test.yaml",
	PublishConfigPath: ".",
	PublishConfig:     nil,

	JmaXmlUrl: "",

	AuthenticationUser: "",
	AuthenticationPw:   "",
}
var ProdConfig = &Config{
	DatabaseConfig: &mysql.Config{
		DBName:               "earthquake-alert",
		User:                 "docker",
		Passwd:               "docker",
		Addr:                 "localhost:3306",
		Net:                  "tcp",
		ParseTime:            true,
		AllowNativePasswords: true,
	},
	PublishConfigName: "publish_config.yaml",
	PublishConfigPath: ".",
	PublishConfig:     nil,

	JmaXmlUrl: "",

	AuthenticationUser: "",
	AuthenticationPw:   "",
}

// 設定初期化
func InitConfig(mode string) error {
	switch mode {
	case "test":
		C = TestConfig
	case "local":
		C = LocalConfig
	case "prod":
		C = ProdConfig
	default:
		C = TestConfig
	}

	return InitPublishConfig()
}

// 公開先設定をyamlファイルから読み取りパースする
func InitPublishConfig() error {
	publishConfigPath := filepath.Join(C.PublishConfigPath, C.PublishConfigName)
	_, err := os.Stat(publishConfigPath)
	if err != nil {
		return err
	}
	data, err := os.ReadFile(publishConfigPath)
	if err != nil {
		return err
	}
	publishConfig := new(PublishConfig)
	err = yaml.Unmarshal(data, publishConfig)
	if err != nil {
		return err
	}
	C.PublishConfig = publishConfig
	return nil
}
