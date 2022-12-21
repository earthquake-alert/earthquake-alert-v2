package jma_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/earthquake-alert/erarthquake-alert-v2/src/jma"
	"github.com/stretchr/testify/require"
)

const TEST_DATA_PATH = "../../test_data/jma_xml/"

var TestEarthquakeActivity = []string{
	"32-35_09_01_191111_VXSE56.xml",
	"32-35_09_02_220316_VXSE56.xml",
}

var TestEarthquakeCount = []string{
	"32-35_03_01_100514_VXSE60.xml",
	"32-35_10_02_220510_VXSE60.xml",
}

var TestEarthquakeUpdate = []string{
	"32-35_06_09_100915_VXSE61.xml",
	"32-35_06_10_100915_VXSE61.xml",
	"32-35_07_09_100915_VXSE61.xml",
}

var TestEarthquakeData = []string{
	// 震度速報
	"32-35_01_01_100806_VXSE51.xml",
	"32-35_04_01_100831_VXSE51.xml",
	"32-35_04_02_100831_VXSE51.xml",
	"32-35_04_05_100831_VXSE51.xml",
	"32-35_07_02_100915_VXSE51.xml",
	"32-35_07_03_100915_VXSE51.xml",
	"32-35_07_04_100915_VXSE51.xml",
	"32-35_07_05_100915_VXSE51.xml",
	"32-35_08_01_100915_VXSE51.xml",
	"32-35_08_02_100915_VXSE51.xml",
	"32-35_08_03_100915_VXSE51.xml",
	"32-35_08_04_100915_VXSE51.xml",
	"32-35_08_05_100915_VXSE51.xml",
	"32-39_11_01_120615_VXSE51.xml",
	"32-35_10_01_220510_VXSE51.xml",
	// 震源に関する情報
	"32-35_01_02_100514_VXSE52.xml",
	"32-35_04_03_100831_VXSE52.xml",
	"32-35_06_01_100915_VXSE52.xml",
	"32-35_06_02_100915_VXSE52.xml",
	"32-35_08_06_100915_VXSE52.xml",
	"33_12_01_120615_VXSE52.xml",
	// 震源・震度に関する情報
	"32-35_01_03_100514_VXSE53.xml",
	"32-35_01_03_100806_VXSE53.xml",
	"32-35_04_04_100831_VXSE53.xml",
	"32-35_04_06_100831_VXSE53.xml",
	"32-35_06_03_100915_VXSE53.xml",
	"32-35_06_04_100915_VXSE53.xml",
	"32-35_06_05_100915_VXSE53.xml",
	"32-35_06_06_100915_VXSE53.xml",
	"32-35_07_06_100915_VXSE53.xml",
	"32-35_07_07_100915_VXSE53.xml",
	"32-35_08_07_100915_VXSE53.xml",
	"32-35_08_08_100915_VXSE53.xml",
	"32-39_05_01_100831_VXSE53.xml",
	"32-39_05_02_100831_VXSE53.xml",
	"32-39_05_04_100831_VXSE53.xml",
	"32-39_05_06_100831_VXSE53.xml",
	"32-39_05_07_100831_VXSE53.xml",
	"32-39_05_08_100831_VXSE53.xml",
	"32-39_05_12_100831_VXSE53.xml",
	"32-39_05_12_100915_VXSE53.xml",
	"32-39_11_05_120615_VXSE53.xml",
}

// 地震の活動状況等に関する情報
// パーステスト1
func TestParseEarthquakeActivity1(t *testing.T) {
	target := "32-35_09_01_191111_VXSE56.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquakeActivity(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "地震の活動状況等に関する情報",
			DateTime:         "2019-11-11T08:00:53Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "地震の活動状況等に関する情報")
		require.Equal(t, head.ReportDateTime, "2019-11-11T17:00:00+09:00")
		require.Equal(t, head.TargetDateTime, "2019-11-11T17:00:00+09:00")
		require.Equal(t, head.EventID, "20191111170000")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "地震の活動状況等に関する情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "伊豆東部の地震活動の見通しに関する情報を発表します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		require.Equal(t, body.Text, `伊豆東部の地震活動の見通しに関する情報（第１号）

１．見出し
本日（１１日）昼から東伊豆奈良本（ならもと）観測点で縮みのひずみ変
化が観測され始め、本日（１１日）昼からは体に感じない小さな地震が発生
し始めました。
このことから、伊豆東部の地下でマグマの貫入と上昇が始まったと思われ
ます。このため、今後、概ね４日、長くて１週間程度、地震活動が活発にな
るおそれがあります。

２．現状
伊豆東部の地下でマグマの貫入と上昇が始まったとみられ、これに伴い、
地殻変動と地震活動が観測されています。

（地殻変動の状況）
本日（１１日）昼から東伊豆奈良本（ならもと）観測点で観測されている
縮みのひずみ変化は、本日（１１日）１６時現在継続しています。２４時間
変化量（換算値）は４００ｎｓｔｒａｉｎ（ナノストレイン）となっていま
す。また、防災科学技術研究所及び気象庁が整備している周辺の傾斜計にも
同期した変化が現われています。
今回のひずみ変化量に近い同様の活動は○○○○年○○月で、その時には
震度１以上を観測する地震が○○回、震度３以上が○○回発生しました。

（地震活動の状況）
本日（１１日）昼に始まった地震活動は、本日（１１日）１６時現在、活
発な状態となっています。本日（１１日）１６時００分にはＭ５.５の地震
が発生して、伊東市大原で震度５弱を観測しました。
これまでに、最大震度５弱を観測する地震が１回、最大震度３が２回、最
大震度２が３回、最大震度１が５回発生しました。
震源はいずれも川奈東沖と川奈崎付近の深さ０から５ｋｍです。

（火山活動の状況）
噴火に直ちに結びつくような現象は観測されていません。

３．見通し
本日（１１日）１６時００分現在の観測データから予測される地震活動の
見通しは、以下のとおりです。
＜今回の地震活動の見通し＞
地震の規模と震度 ： 最大マグニチュード６程度
最大震度５弱から５強程度 ※
震度１以上の地震回数： ５００から９００回程度
活動期間 ： 概ね４日、長くて１週間程度
※地盤の状況等により、更に揺れが大きくなる場合があります。

４．防災上の留意事項
活動期間の予測は一回のマグマ上昇に基づくため、複数回の上昇が起きた
場合はさらに長引くことがあります。
マグマがさらに浅部へ上昇した場合、地震活動がさらに活発になることが
あります。

次の伊豆東部の地震活動の見通しに関する情報は、本日（１１日）１８時
頃に発表の予定です。
なお、見通しの内容を更新する場合や、活動に顕著な変化があった場合等
には、随時お知らせします。`)

		require.Nil(t, body.Comments)
	})
}

// 地震の活動状況等に関する情報
// パーステスト2（取り消し）
func TestParseEarthquakeActivity2(t *testing.T) {
	target := "32-35_09_02_220316_VXSE56.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquakeActivity(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "地震の活動状況等に関する情報",
			DateTime:         "2022-03-16T04:46:04Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "地震の活動状況等に関する情報")
		require.Equal(t, head.ReportDateTime, "2022-03-16T13:46:00+09:00")
		require.Equal(t, head.TargetDateTime, "2022-03-16T13:46:00+09:00")
		require.Equal(t, head.EventID, "20220316133000")
		require.Equal(t, head.InfoType, jma.Cancel)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "地震の活動状況等に関する情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "伊豆東部の地震活動の見通しに関する情報を発表します。 ")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		require.Equal(t, body.Text, "１６日１３時３０分に発表した「伊豆東部の地震活動の見通しに関する情報」は誤りですので取り消します。 ")

		require.Nil(t, body.Comments)
	})
}

// 地震回数に関する情報
// パーステスト1
func TestParseEarthquakeCount1(t *testing.T) {
	target := "32-35_03_01_100514_VXSE60.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquakeCount(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "地震回数に関する情報",
			DateTime:         "2008-08-26T03:00:15Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "地震回数に関する情報")
		require.Equal(t, head.ReportDateTime, "2008-08-26T12:00:00+09:00")
		require.Equal(t, head.TargetDateTime, "2008-08-26T12:00:00+09:00")
		require.Equal(t, head.EventID, "20080824150500")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "地震回数情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "地震回数に関する情報をお知らせします。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		t.Run("earthquake count", func(t *testing.T) {
			earthquake := body.EarthquakeCount

			require.Len(t, earthquake.Item, 5)
			require.Equal(t, earthquake.Item[0].Type, "地震回数")
			require.Equal(t, earthquake.Item[0].StartTime, "2008-08-24T15:00:00+09:00")
			require.Equal(t, earthquake.Item[0].EndTime, "2008-08-25T09:00:00+09:00")
			require.Equal(t, earthquake.Item[0].Number, 1587)
			require.Equal(t, earthquake.Item[0].FeltNumber, 1)
		})

		require.Equal(t, body.Text, "")
		require.Equal(t, body.NextAdvisory, "次の「地震回数に関する情報」は、２６日１８時００分頃に発表します。")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.FreeFormComment, `　８月２４日１５時過ぎから伊豆半島東方沖で地震が発生しています。この
付近で発生した地震については、震度３以上の場合は「震源・震度情報」で
発表しますが、震度２以下の場合は、「地震回数に関する情報」（本情報）
で地震回数をまとめて発表します。`)
	})
}

// 地震回数に関する情報
// パーステスト2 （取り消し）
func TestParseEarthquakeCount2(t *testing.T) {
	target := "32-35_10_02_220510_VXSE60.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquakeCount(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "地震回数に関する情報",
			DateTime:         "2022-05-10T09:49:07Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "地震回数に関する情報")
		require.Equal(t, head.ReportDateTime, "2022-05-10T18:49:00+09:00")
		require.Equal(t, head.TargetDateTime, "2022-05-10T18:49:00+09:00")
		require.Equal(t, head.EventID, "20220510173400")
		require.Equal(t, head.InfoType, jma.Cancel)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "地震回数情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "地震回数に関する情報を取り消します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		require.Nil(t, body.EarthquakeCount)

		require.Equal(t, body.Text, "先ほどの、地震回数に関する情報を取り消します。")
		require.Equal(t, body.NextAdvisory, "")

		require.Nil(t, body.Comments)
	})
}

// 顕著な地震の震源要素更新のお知らせ
// パーステスト1
func TestParseEarthquakeUpdate1(t *testing.T) {
	target := "32-35_06_09_100915_VXSE61.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquakeUpdate(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "顕著な地震の震源要素更新のお知らせ",
			DateTime:         "2008-06-14T03:30:00Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "顕著な地震の震源要素更新のお知らせ")
		require.Equal(t, head.ReportDateTime, "2008-06-14T12:30:00+09:00")
		require.Equal(t, head.TargetDateTime, "2008-06-14T12:30:00+09:00")
		require.Equal(t, head.EventID, "20080614084350")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震源要素更新のお知らせ")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "平成２０年　６月１４日１２時３０分をもって、地震の発生場所と規模を更新します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake

			require.Equal(t, earthquake.OriginTime, "2008-06-14T08:43:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2008-06-14T08:43:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "岩手県内陸南部")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "213")
			require.Len(t, earthquake.Hypocenter.Area.Coordinate, 2)
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Value, "+39.0+140.9-10000/")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Description, "北緯３９．０度　東経１４０．９度　深さ　１０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Type, "")

			require.Equal(t, earthquake.Magnitude.Condition, "")
			require.Equal(t, earthquake.Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ７．２")
			require.Equal(t, earthquake.Magnitude.Value, "7.2")
		})

		require.Equal(t, body.Text, "")
		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.FreeFormComment, "度単位の震源要素は、津波情報等を引き続き発表する場合に使用されます。")
	})
}

// 顕著な地震の震源要素更新のお知らせ
// パーステスト2 （取り消し）
func TestParseEarthquakeUpdate2(t *testing.T) {
	target := "32-35_06_10_100915_VXSE61.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquakeUpdate(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "顕著な地震の震源要素更新のお知らせ",
			DateTime:         "2008-06-14T03:35:00Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "顕著な地震の震源要素更新のお知らせ")
		require.Equal(t, head.ReportDateTime, "2008-06-14T12:35:00+09:00")
		require.Equal(t, head.TargetDateTime, "2008-06-14T12:35:00+09:00")
		require.Equal(t, head.EventID, "20080614084350")
		require.Equal(t, head.InfoType, jma.Cancel)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震源要素更新のお知らせ")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "顕著な地震の震源要素更新のお知らせを取り消します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		require.Nil(t, body.Earthquake)
		require.Equal(t, body.Text, "先ほどの、顕著な地震の震源要素更新のお知らせを取り消します。")
		require.Nil(t, body.Comments)
	})
}

// 顕著な地震の震源要素更新のお知らせ
// パーステスト3
func TestParseEarthquakeUpdate3(t *testing.T) {
	target := "32-35_07_09_100915_VXSE61.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquakeUpdate(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "顕著な地震の震源要素更新のお知らせ",
			DateTime:         "2009-08-10T23:55:00Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "顕著な地震の震源要素更新のお知らせ")
		require.Equal(t, head.ReportDateTime, "2009-08-11T08:55:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-08-11T06:45:00+09:00")
		require.Equal(t, head.EventID, "20090811050711")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震源要素更新のお知らせ")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "平成２１年　８月１１日０６時４５分をもって、地震の発生場所と規模を更新します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake

			require.Equal(t, earthquake.OriginTime, "2009-08-11T05:07:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2009-08-11T05:07:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "駿河湾")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "485")
			require.Len(t, earthquake.Hypocenter.Area.Coordinate, 2)
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Value, "+34.8+138.5-20000/")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Description, "北緯３４．８度　東経１３８．５度　深さ　２０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate[0].Type, "")

			require.Equal(t, earthquake.Magnitude.Condition, "")
			require.Equal(t, earthquake.Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ６．５")
			require.Equal(t, earthquake.Magnitude.Value, "6.5")
		})

		require.Equal(t, body.Text, "")
		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.FreeFormComment, "度単位の震源要素は、津波情報等を引き続き発表する場合に使用されます。")
	})
}

// 震度速報パーステスト
func TestParseEarthquakeReport1(t *testing.T) {
	target := "32-35_01_01_100806_VXSE51.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震度速報",
			DateTime:         "2009-08-10T20:09:11Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "震度速報")
		require.Equal(t, head.ReportDateTime, "2009-08-11T05:09:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-08-11T05:07:00+09:00")
		require.Equal(t, head.EventID, "20090811050711")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震度速報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "１１日０５時０７分ころ、地震による強い揺れを感じました。震度３以上が観測された地域をお知らせします。")

		// Information
		require.Len(t, head.Headline.Information, 1)
		require.Equal(t, head.Headline.Information[0].Type, "震度速報")

		require.Len(t, head.Headline.Information[0].Item, 4)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "震度６弱")

		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "地震情報／細分区域")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 2)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "静岡県中部")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "442")
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		require.Nil(t, body.Earthquake)

		t.Run("intensity", func(t *testing.T) {
			intensity := body.Intensity

			require.NotNil(t, intensity)

			require.Len(t, intensity.Observation.CodeDefine.Types, 2)
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].Value, "地震情報／都道府県等")
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].XPath, "Pref/Code")

			require.Equal(t, intensity.Observation.MaxInt, jma.Int6l)

			require.Len(t, intensity.Observation.Pref, 16)
			require.Equal(t, intensity.Observation.Pref[0].Name, "静岡県")
			require.Equal(t, intensity.Observation.Pref[0].Code, "22")
			require.Equal(t, intensity.Observation.Pref[0].MaxInt, jma.Int6l)

			require.Len(t, intensity.Observation.Pref[0].Area, 4)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Name, "静岡県中部")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Code, "442")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].MaxInt, jma.Int6l)
		})

		require.Equal(t, body.Text, "")
		require.Nil(t, body.Comments)

	})
}

// 震度速報パーステスト（東日本大震災）
func TestParseEarthquakeReport2(t *testing.T) {
	target := "32-39_11_01_120615_VXSE51.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震度速報",
			DateTime:         "2011-03-11T05:48:10Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "震度速報")
		require.Equal(t, head.ReportDateTime, "2011-03-11T14:48:00+09:00")
		require.Equal(t, head.TargetDateTime, "2011-03-11T14:46:00+09:00")
		require.Equal(t, head.EventID, "20110311144640")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震度速報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "１１日１４時４６分ころ、地震による強い揺れを感じました。震度３以上が観測された地域をお知らせします。")

		// Information
		require.Len(t, head.Headline.Information, 1)
		require.Equal(t, head.Headline.Information[0].Type, "震度速報")

		require.Len(t, head.Headline.Information[0].Item, 6)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "震度６強")

		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "地震情報／細分区域")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 2)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "宮城県北部")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "220")
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		require.Nil(t, body.Earthquake)

		t.Run("intensity", func(t *testing.T) {
			intensity := body.Intensity

			require.NotNil(t, intensity)

			require.Len(t, intensity.Observation.CodeDefine.Types, 2)
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].Value, "地震情報／都道府県等")
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].XPath, "Pref/Code")

			require.Equal(t, intensity.Observation.MaxInt, jma.Int6u)

			require.Len(t, intensity.Observation.Pref, 6)
			require.Equal(t, intensity.Observation.Pref[0].Name, "宮城県")
			require.Equal(t, intensity.Observation.Pref[0].Code, "04")
			require.Equal(t, intensity.Observation.Pref[0].MaxInt, jma.Int6u)

			require.Len(t, intensity.Observation.Pref[0].Area, 3)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Name, "宮城県北部")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Code, "220")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].MaxInt, jma.Int6u)
		})

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, "今後の情報に注意してください。")
		require.Equal(t, body.Comments.ForecastComment.Code, "0217")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Nil(t, body.Comments.VarComment)

	})
}

// 震度速報パーステスト（取り消し）
func TestParseEarthquakeReport3(t *testing.T) {
	target := "32-35_10_01_220510_VXSE51.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震度速報",
			DateTime:         "2022-05-10T09:10:20Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "震度速報")
		require.Equal(t, head.ReportDateTime, "2022-05-10T18:10:00+09:00")
		require.Equal(t, head.TargetDateTime, "2022-05-09T15:23:00+09:00")
		require.Equal(t, head.EventID, "20220509152315")
		require.Equal(t, head.InfoType, jma.Cancel)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震度速報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "震度速報を取り消します。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		require.Nil(t, body.Earthquake)
		require.Nil(t, body.Intensity)

		require.Equal(t, body.Text, "先ほどの、震度速報を取り消します。")
		require.Nil(t, body.Comments)

	})
}

// 震源に関する情報パーステスト
func TestParseEarthquakeEpicenter1(t *testing.T) {
	target := "32-35_01_02_100514_VXSE52.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震源に関する情報",
			DateTime:         "2009-10-01T04:48:03Z",
			Status:           jma.Training,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "震源に関する情報")
		require.Equal(t, head.ReportDateTime, "2009-10-01T13:48:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-10-01T13:48:00+09:00")
		require.Equal(t, head.EventID, "20091001134500")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震源速報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "　１日１３時４５分ころ、地震がありました。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake
			require.NotNil(t, earthquake)

			require.Equal(t, earthquake.OriginTime, "2009-10-01T13:45:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2009-10-01T13:45:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "駿河湾")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "485")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Description, "北緯３４．８度　東経１３８．５度　深さ　１０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Value, "+34.8+138.5-10000/")
			require.Equal(t, earthquake.Hypocenter.Area.NameFromMark, "")
			require.Nil(t, earthquake.Hypocenter.Area.MarkCode)
			require.Equal(t, earthquake.Hypocenter.Area.Direction, "")
			require.Nil(t, earthquake.Hypocenter.Area.Distance)

			require.Equal(t, earthquake.Hypocenter.Source, "")

			require.Equal(t, earthquake.Magnitude.Condition, "")
			require.Equal(t, earthquake.Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ６．６")
			require.Equal(t, earthquake.Magnitude.Value, "6.6")

		})

		require.Nil(t, body.Intensity)

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, "この地震による津波の心配はありません。")
		require.Equal(t, body.Comments.ForecastComment.Code, "0203")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Nil(t, body.Comments.VarComment)
	})
}

// 震源に関する情報パーステスト
func TestParseEarthquakeEpicenter2(t *testing.T) {
	target := "33_12_01_120615_VXSE52.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震源に関する情報",
			DateTime:         "2012-02-14T12:42:53Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "震源に関する情報")
		require.Equal(t, head.ReportDateTime, "2012-02-14T21:42:00+09:00")
		require.Equal(t, head.TargetDateTime, "2012-02-14T21:42:00+09:00")
		require.Equal(t, head.EventID, "20120214214013")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "")
		require.Equal(t, head.InfoKind, "震源速報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "１４日２１時４０分ころ、地震がありました。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake
			require.NotNil(t, earthquake)

			require.Equal(t, earthquake.OriginTime, "2012-02-14T21:40:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2012-02-14T21:40:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "岐阜県美濃中西部")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "432")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Description, "北緯３５．６度　東経１３６．６度　深さ　１０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Value, "+35.6+136.6-10000/")
			require.Equal(t, earthquake.Hypocenter.Area.NameFromMark, "")
			require.Nil(t, earthquake.Hypocenter.Area.MarkCode)
			require.Equal(t, earthquake.Hypocenter.Area.Direction, "")
			require.Nil(t, earthquake.Hypocenter.Area.Distance)

			require.Equal(t, earthquake.Hypocenter.Source, "")

			require.Equal(t, earthquake.Magnitude.Condition, "不明")
			require.Equal(t, earthquake.Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ８を超える巨大地震")
			require.Equal(t, earthquake.Magnitude.Value, "NaN")

		})

		require.Nil(t, body.Intensity)

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, `この地震による津波の心配はありません。
この地震について、緊急地震速報を発表しています。`)
		require.Equal(t, body.Comments.ForecastComment.Code, "0215 0241")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Nil(t, body.Comments.VarComment)
	})
}

// 震源・震度に関する情報のパーステスト
func TestParseEarthquakeDetails1(t *testing.T) {
	target := "32-35_01_03_100806_VXSE53.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震源・震度に関する情報",
			DateTime:         "2009-10-01T04:50:01Z",
			Status:           jma.Training,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "震源・震度情報")
		require.Equal(t, head.ReportDateTime, "2009-10-01T13:50:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-10-01T13:50:00+09:00")
		require.Equal(t, head.EventID, "20091001134500")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "地震情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "　１日１３時４５分ころ、地震がありました。各地の震度をお知らせします。")

		// Information
		require.Len(t, head.Headline.Information, 2)
		require.Equal(t, head.Headline.Information[0].Type, "震源・震度に関する情報（細分区域）")
		require.Len(t, head.Headline.Information[0].Item, 3)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "震度５弱")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "地震情報／細分区域")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 2)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "静岡県伊豆")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "440")
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake
			require.NotNil(t, earthquake)

			require.Equal(t, earthquake.OriginTime, "2009-10-01T13:45:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2009-10-01T13:45:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "駿河湾")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "485")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Description, "北緯３４．８度　東経１３８．５度　深さ　１０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Value, "+34.8+138.5-10000/")
			require.Equal(t, earthquake.Hypocenter.Area.NameFromMark, "")
			require.Nil(t, earthquake.Hypocenter.Area.MarkCode)
			require.Equal(t, earthquake.Hypocenter.Area.Direction, "")
			require.Nil(t, earthquake.Hypocenter.Area.Distance)

			require.Equal(t, earthquake.Hypocenter.Source, "")

			require.Equal(t, earthquake.Magnitude.Condition, "")
			require.Equal(t, earthquake.Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ５．９")
			require.Equal(t, earthquake.Magnitude.Value, "5.9")

		})

		t.Run("intensity", func(t *testing.T) {
			intensity := body.Intensity

			require.NotNil(t, intensity)

			require.Len(t, intensity.Observation.CodeDefine.Types, 4)
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].Value, "地震情報／都道府県等")
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].XPath, "Pref/Code")

			require.Equal(t, intensity.Observation.MaxInt, jma.Int5l)

			require.Len(t, intensity.Observation.Pref, 8)
			require.Equal(t, intensity.Observation.Pref[0].Name, "静岡県")
			require.Equal(t, intensity.Observation.Pref[0].Code, "22")
			require.Equal(t, intensity.Observation.Pref[0].MaxInt, jma.Int5l)

			require.Len(t, intensity.Observation.Pref[0].Area, 4)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Name, "静岡県伊豆")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Code, "440")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].MaxInt, jma.Int5l)

			require.Len(t, intensity.Observation.Pref[0].Area[0].City, 10)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].Name, "西伊豆町")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].Code, "2230600")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].MaxInt, jma.Int5l)
			require.Len(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation, 1)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Name, "西伊豆町仁科＊")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Code, "2230630")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Int, jma.Int5l)

		})

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, "この地震による津波の心配はありません。")
		require.Equal(t, body.Comments.ForecastComment.Code, "0203")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Equal(t, body.Comments.VarComment.Text, "＊印は気象庁以外の震度観測点についての情報です。")
		require.Equal(t, body.Comments.VarComment.Code, "0210")
		require.Equal(t, body.Comments.VarComment.CodeType, "固定付加文")
	})
}

// 震源・震度に関する情報のパーステスト（東日本大震災）
func TestParseEarthquakeDetails2(t *testing.T) {
	target := "32-39_11_05_120615_VXSE53.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震源・震度に関する情報",
			DateTime:         "2011-03-11T05:54:58Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "震源・震度情報")
		require.Equal(t, head.ReportDateTime, "2011-03-11T14:54:00+09:00")
		require.Equal(t, head.TargetDateTime, "2011-03-11T14:54:00+09:00")
		require.Equal(t, head.EventID, "20110311144640")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "地震情報")
		require.Equal(t, head.InfoKindVersion, "1.0_1")

		require.Equal(t, head.Headline.Text, "１１日１４時４６分ころ、地震がありました。")

		// Information
		require.Len(t, head.Headline.Information, 2)
		require.Equal(t, head.Headline.Information[0].Type, "震源・震度に関する情報（細分区域）")
		require.Len(t, head.Headline.Information[0].Item, 7)
		require.Equal(t, head.Headline.Information[0].Item[0].Kind.Name, "震度７")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.CodeType, "地震情報／細分区域")
		require.Len(t, head.Headline.Information[0].Item[0].Areas.Area, 1)
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Name, "宮城県北部")
		require.Equal(t, head.Headline.Information[0].Item[0].Areas.Area[0].Code, "220")
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake
			require.NotNil(t, earthquake)

			require.Equal(t, earthquake.OriginTime, "2011-03-11T14:46:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2011-03-11T14:46:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "三陸沖")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "288")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Datum, "日本測地系")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Description, "北緯３８．０度　東経１４２．９度　深さ　１０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Value, "+38.0+142.9-10000/")
			require.Equal(t, earthquake.Hypocenter.Area.NameFromMark, "牡鹿半島の東南東１３０ｋｍ付近")
			require.Equal(t, earthquake.Hypocenter.Area.MarkCode.Type, "震央補助")
			require.Equal(t, earthquake.Hypocenter.Area.MarkCode.Value, "202")
			require.Equal(t, earthquake.Hypocenter.Area.Direction, "東南東")
			require.Equal(t, earthquake.Hypocenter.Area.Distance.Unit, "km")
			require.Equal(t, earthquake.Hypocenter.Area.Distance.Value, "130")

			require.Equal(t, earthquake.Hypocenter.Source, "")

			require.Equal(t, earthquake.Magnitude.Condition, "不明")
			require.Equal(t, earthquake.Magnitude.Type, jma.JMAMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ８を超える巨大地震")
			require.Equal(t, earthquake.Magnitude.Value, "NaN")

		})

		t.Run("intensity", func(t *testing.T) {
			intensity := body.Intensity

			require.NotNil(t, intensity)

			require.Len(t, intensity.Observation.CodeDefine.Types, 4)
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].Value, "地震情報／都道府県等")
			require.Equal(t, intensity.Observation.CodeDefine.Types[0].XPath, "Pref/Code")

			require.Equal(t, intensity.Observation.MaxInt, jma.Int7)

			require.Len(t, intensity.Observation.Pref, 41)
			require.Equal(t, intensity.Observation.Pref[0].Name, "宮城県")
			require.Equal(t, intensity.Observation.Pref[0].Code, "04")
			require.Equal(t, intensity.Observation.Pref[0].MaxInt, jma.Int7)

			require.Len(t, intensity.Observation.Pref[0].Area, 3)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Name, "宮城県北部")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].Code, "220")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].MaxInt, jma.Int7)

			require.Len(t, intensity.Observation.Pref[0].Area[0].City, 8)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].Name, "栗原市")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].Code, "0421300")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].MaxInt, jma.Int7)
			require.Len(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation, 2)
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Name, "栗原市築館＊")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Code, "0421320")
			require.Equal(t, intensity.Observation.Pref[0].Area[0].City[0].IntensityStation[0].Int, jma.Int7)

		})

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, `津波警報等（大津波警報・津波警報あるいは津波注意報）を発表中です。
この地震について、緊急地震速報を発表しています。`)
		require.Equal(t, body.Comments.ForecastComment.Code, "0211 0241")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Equal(t, body.Comments.VarComment.Text, "＊印は気象庁以外の震度観測点についての情報です。")
		require.Equal(t, body.Comments.VarComment.Code, "0262")
		require.Equal(t, body.Comments.VarComment.CodeType, "固定付加文")
	})
}

// 震源・震度に関する情報のパーステスト（遠地地震）
func TestParseEarthquakeDetails3(t *testing.T) {
	target := "32-35_01_03_100514_VXSE53.xml"

	testPath := filepath.Join(TEST_DATA_PATH, target)
	row, err := os.ReadFile(testPath)
	require.NoError(t, err)

	ea, err := jma.ParseEarthquake(row)
	require.NoError(t, err)

	t.Run("control", func(t *testing.T) {
		control := ea.Control

		require.Equal(t, control, jma.JmaXmlControl{
			Title:            "震源・震度に関する情報",
			DateTime:         "2009-11-09T11:15:00Z",
			Status:           jma.Common,
			EditorialOffice:  "気象庁本庁",
			PublishingOffice: "気象庁",
		})
	})

	t.Run("head", func(t *testing.T) {
		head := ea.Head

		require.Equal(t, head.Title, "遠地地震に関する情報")
		require.Equal(t, head.ReportDateTime, "2009-11-09T20:15:00+09:00")
		require.Equal(t, head.TargetDateTime, "2009-11-09T20:15:00+09:00")
		require.Equal(t, head.EventID, "20091109194822")
		require.Equal(t, head.InfoType, jma.Publication)
		require.Equal(t, head.Serial, "1")
		require.Equal(t, head.InfoKind, "地震情報")
		require.Equal(t, head.InfoKindVersion, "1.0_0")

		require.Equal(t, head.Headline.Text, "　９日１９時４５分ころ、海外で規模の大きな地震がありました。")

		// Information
		require.Len(t, head.Headline.Information, 0)
	})

	t.Run("body", func(t *testing.T) {
		body := ea.Body

		t.Run("earthquake", func(t *testing.T) {
			earthquake := body.Earthquake
			require.NotNil(t, earthquake)

			require.Equal(t, earthquake.OriginTime, "2009-11-09T19:45:00+09:00")
			require.Equal(t, earthquake.ArrivalTime, "2009-11-09T19:45:00+09:00")

			require.Equal(t, earthquake.Hypocenter.Area.Name, "南太平洋")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Value, "950")
			require.Equal(t, earthquake.Hypocenter.Area.Code.Type, "震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Datum, "")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Description, "南緯１７．２度　東経１７８．６度　深さ５７０ｋｍ")
			require.Equal(t, earthquake.Hypocenter.Area.Coordinate.Value, "-17.2+178.6-570000/")
			require.Equal(t, earthquake.Hypocenter.Area.NameFromMark, "")
			require.Nil(t, earthquake.Hypocenter.Area.MarkCode)
			require.Equal(t, earthquake.Hypocenter.Area.Direction, "")
			require.Nil(t, earthquake.Hypocenter.Area.Distance)
			require.Equal(t, earthquake.Hypocenter.Area.DetailedName, "フィジー諸島")
			require.Equal(t, earthquake.Hypocenter.Area.DetailedCode.Type, "詳細震央地名")
			require.Equal(t, earthquake.Hypocenter.Area.DetailedCode.Value, "182")

			require.Equal(t, earthquake.Hypocenter.Source, "ＰＴＷＣ")

			require.Equal(t, earthquake.Magnitude.Condition, "")
			require.Equal(t, earthquake.Magnitude.Type, jma.MomentMagnitude)
			require.Equal(t, earthquake.Magnitude.Description, "Ｍ７．１")
			require.Equal(t, earthquake.Magnitude.Value, "7.1")

		})

		require.Nil(t, body.Intensity)

		require.Equal(t, body.Text, "")

		require.NotNil(t, body.Comments)
		require.Equal(t, body.Comments.ForecastComment.Text, "この地震による津波の心配はありません。")
		require.Equal(t, body.Comments.ForecastComment.Code, "0203")
		require.Equal(t, body.Comments.ForecastComment.CodeType, "固定付加文")
		require.Equal(t, body.Comments.FreeFormComment, "")
		require.Nil(t, body.Comments.VarComment)
	})
}
