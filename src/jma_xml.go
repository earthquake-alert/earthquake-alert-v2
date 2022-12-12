package src

type Status int
type InfoType int

const (
	Common Status = iota
	Training
	Test
	StatusUnknown
)

const (
	Publication InfoType = iota
	Correction
	Cancel
	InfoTypeUnknown
)

// 管理部
// 本情報の配信に関連する情報を記載する。
type JmaXmlControl struct {
	// 情報名称
	// 本要素は、「独立した情報単位」判別のキーとしても用いられる（（ⅲ）共通別紙ア．「地震火
	// 山関連XML 電文の「独立した情報単位」の運用」参照）。
	Title string `xml:"Title"`

	// 発表時刻
	// 気象庁システムからの発信時刻を記載する。この値は秒値まで有効である
	// TODO: ISO 8601 でエンコードできるかどうか
	DateTime string `xml:"DateTime"`

	// 運用種別
	// - 通常の運用で発表する情報: 通常
	// - 事前に日時を定めて行う業務訓練等で発表する情報: 訓練
	// - 定期または臨時に伝聞疎通確認等を目的として発表する緊急地震速報の配信テスト伝聞: 試験
	Status string `xml:"Status"`

	// 編集官署名
	// 	本要素は、「独立した情報単位」判別のキーとしても用いられるが、地震・津波に関連する情
	// 報、南海トラフ地震に関連する情報及び地震・津波に関するお知らせについては、システム障
	// 害発生等により一連の情報であっても編集官署が切り替わる場合があることに留意が必要で
	// ある。地震・津波に関連する情報等のこうした取扱については、（ⅲ）共通別紙ア．「地震火山
	// 関連XML 電文の「独立した情報単位」の運用」を参照すること。
	EditorialOffice string `xml:"EditorialOffice"`

	// 発表官署名
	PublishingOffice string `xml:"PublishingOffice"`
}

// ヘッダ部
type JmaXmlHeader struct {
	// 課題
	// 	情報の標題を記載する。
	// 震源・震度に関する情報において、近地地震の場合には“震源・震度情報”、遠地地震の場
	// 合には“遠地地震に関する情報”と記載する。
	// 津波警報・注意報・予報については、発表する情報に含まれる津波予報等の種類の総和表
	// 現を記載する。なお、津波警報・注意報を全解除し、全ての津波予報区等で津波予報（若干
	// の海面変動）又は津波なしとなる場合は、事例に示すとおり“津波予報”と記載する。
	// 各地の満潮時刻と津波到達予想時刻を発表する津波情報については“各地の満潮時刻・
	// 津波到達予想時刻に関する情報”を、津波の観測値を発表する津波情報については“津波観
	// 測に関する情報”を記載する。両者をひとつの津波情報電文で発表する場合は、本要素の中
	// に二つの標題を半角スペースで区切って併記する。
	// 南海トラフ地震に関連する情報においては、情報名称（Control/Title）が”南海トラフ地震臨
	// 時情報” の場合は、” 南海トラフ地震臨時情報” に続けて情報種別番号名
	// （Body/EarthquakeInfo/InfoSerial/Name）の内容を付記する（例：”南海トラフ地震臨時情報
	// （巨大地震警戒）”）。また、情報名称（Control/Title）が”南海トラフ地震関連解説情報”の場
	// 合は、”南海トラフ地震関連解説情報”と標記し、情報番号（Head/Serial）に値が記載されてい
	// る場合に限り、一連の情報番号を付記する（例：”南海トラフ地震関連解説情報（第○号）”）。
	// 火山に関連する情報においては、火山名と情報の種別を記載する。
	//
	// - 事例１（津波注意報と津波予報を発表する場合）
	//   - 津波注意報・津波予報
	// - 事例２（津波注意報を全解除し、津波予報（若干の海面変動）が残る場合）
	//   - 津波予報
	// - 事例３（津波注意報を全解除し、全ての津波予報区で津波なしとなる場合）
	//   - 津波予報
	// - 事例４（大津波警報、津波警報、津波注意報、津波予報を発表する場合）
	//   - 大津波警報・津波警報・津波注意報・津波予報
	Title string `xml:"Title"`

	// 発表時刻
	// 発表官署が本情報を発表した時刻を記載する。
	// 緊急地震速報（警報）、緊急地震速報（地震動予報）、緊急地震速報（予報）、及び緊急地
	// 震速報の配信テスト電文については秒値まで、その他の地震・津波・南海トラフ地震・火山に
	// 関連する情報については、分値まで有効である。
	ReportDateTime string `xml:"ReportDateTime"`

	// 基点時刻
	// 	情報の内容が発現・発効する基点時刻を記載する。
	// 震度速報については最初に地震波を自動検知した観測点における地震波の検知時刻を、
	// 地震情報（顕著な地震の震源要素更新のお知らせ）については震源要素を切り替えた時刻を、
	// 津波の観測値を発表する津波情報、沖合の津波観測に関する情報については津波の観測
	// 状況を確定した時刻を記載する。火山現象に関する海上警報については火山活動の観測時
	// 刻、噴火に関する火山観測報、噴火速報、推定噴煙流向報については報じる現象の発現時
	// 刻、降灰予報については情報の対象となる時間帯の基点時刻を記載する。その他の地震・津
	// 波・火山に関連する情報については、ヘッダ部の発表時刻（Head/ReportDateTime）の値を記
	// 載する。
	// なお、緊急地震速報（警報）、緊急地震速報（地震動予報）、緊急地震速報（予報）、及び緊
	// 急地震速報の配信テスト電文については秒値まで、その他の地震・津波・南海トラフ地震・火
	// 山に関連する情報については、分値まで有効である。ただし、噴火に関する火山観測報、噴
	// 火速報、推定噴煙流向報については、基本的に分値まで有効であるが、TargetDTDubious が
	// 出現する場合は、それで示すあいまいさに応じた単位までが有効、発現時刻が不明の場合に
	// は~~xsi:nil=“true”属性値により空要素となる~~ nilとなる。
	TargetDateTime string `xml:"TargetDateTime"`

	// 基点時刻の曖昧さ
	// "頃", "年頃", "月頃", "日頃", "時頃", "分頃", "秒頃"
	// Optional
	//
	// 噴火に関する火山観測報、噴火速報、推定噴煙流向報で用いる場合があり、報じる現象の
	// 発現時刻にあいまいさがある場合に記載する。
	// 例えば“日頃”のときは年月日までが有効となる。具体的な精度の有効な範囲は、内容部の
	// EventDateTime 及びEventDateTimeUTC の@significant に記載する。
	TargetDTDubious string `xml:"TargetDTDubious,omitempty"`

	// 失効時刻
	// 津波警報・注意報・予報の電文及び降灰予報の電文において情報の失効時刻を記載する。
	// 津波警報・注意報・予報の電文については、津波予報（若干の海面変動）のみ発表の場合
	// 	や、津波警報・注意報解除後に津波予報（若干の海面変動）のみが残る場合に、その失効時
	// 刻を記載する。
	// 降灰予報については、それぞれの情報における失効時刻を記載し、降灰予報（定時）は基
	// 点時刻から18 時間後、降灰予報（速報）は基点時刻から1 時間後、降灰予報（詳細）は基点
	// 時刻から概ね6 時間後となる。
	// 存在しない場合はnil
	// Optional
	ValidDateTim string `xml:"ValidDateTim,omitempty"`

	// 識別情報
	// 地震・津波に関連する情報については、ある特定の地震を識別するための地震識別番号
	// （14 桁の数字）を記載する。津波に関連する情報では、当該警報等に寄与している地震の地
	// 震識別番号を記載するため、１つの電文に複数の地震識別番号が出現する場合もある。詳細
	// については、（ⅲ）共通別紙イ．「地震・津波に関連する情報のEventID 要素の運用」を参照。
	// 南海トラフ地震に関連する情報については、任意の識別番号（14 桁の数字）を記載する。詳
	// 細については、（ⅲ）共通別紙エ．「南海トラフ地震に関連する情報におけるEventID 要素及
	// びSerial 要素の運用」を参照。
	// 火山に関連する情報については、３桁の火山番号を記載する。ただし、噴火に関する火山
	// 観測報及び噴火速報、推定噴煙流向報については、ReportDateTime と火山番号を“_”で連
	// 結して記載する。
	// 地震・津波に関するお知らせや火山に関するお知らせについては、情報発表日時分（14 桁
	// の数字）を記載する。
	EventID string `xml:"EventID"`

	// 情報形態
	// 情報を発表する場合は“発表”を、「独立した情報単位」において直前の時点で発表されて
	// いるControl/DateTime の最も新しい電文を訂正する場合は“訂正”を、「独立した情報単位」
	// 全体を取り消す場合は“取消”を記載する。取消電文の運用については、（ⅲ）共通別紙ウ．
	// 「取消電文の運用」を参照。
	InfoType string `xml:"InfoType"`

	// 情報番号
	// 続報を発表し、内容を更新する情報については、情報番号を記載する。続報を発表する度
	// に情報番号を更新するが、取消報の場合は、番号は更新しない。訂正報の場合は訂正する
	// 直近の情報の情報番号を記載する。
	// 南海トラフ地震に関連する情報については、続報を発表する情報で情報番号を記載する。
	// 詳細については、（ⅲ）共通別紙エ．「南海トラフ地震に関連する情報におけるEventID 要素
	// 及びSerial 要素の運用」を参照。
	Serial string `xml:"Serial"`

	// スキーマの運用種別情報
	InfoKind string `xml:"InfoKind"`

	// スキーマの運用種別情報のバージョン番
	InfoKindVersion string `xml:"InfoKindVersion"`
}
