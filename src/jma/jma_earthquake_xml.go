// 気象庁電文のXML定義
//
// 地震関係の以下4種類のXMLを定義しています。
// - 震度速報、震源に関する情報、震源・震度に関する情報
// - 地震の活動状況等に関する情報
// - 地震回数に関する情報
// - 顕著な地震の震源要素更新のお知らせ
package jma

type EarthquakeClass string

// 地震情報等におけるXML伝聞（NOT 緊急地震速報）
type EarthquakeJmaXml struct {
	Control JmaXmlControl `xml:"Control"`

	Head struct {
		JmaXmlHeader

		Headline EarthquakeHeadline `xml:"Headline"`
	} `xml:"Head"`

	Body EarthquakeBody `xml:"Body"`
}

// 地震の活動状況等に関する情報
type EarthquakeActivityJmaXml struct {
	Control JmaXmlControl `xml:"Control"`

	Head struct {
		JmaXmlHeader

		Headline EarthquakeHeadline `xml:"Headline"`
	} `xml:"Head"`

	Body EarthquakeActivityBody `xml:"Body"`
}

// 地震回数に関する情報
type EarthquakeCountInfoJmaXml struct {
	Control JmaXmlControl `xml:"Control"`

	Head struct {
		JmaXmlHeader

		Headline EarthquakeHeadline `xml:"Headline"`
	} `xml:"Head"`

	Body EarthquakeCountInfoBody `xml:"Body"`
}

// 顕著な地震の震源要素更新のお知らせ
type EarthquakeUpdateInfoJmaXml struct {
	Control JmaXmlControl `xml:"Control"`

	Head struct {
		JmaXmlHeader

		Headline EarthquakeHeadline `xml:"Headline"`
	} `xml:"Head"`

	Body EarthquakeUpdateBody `xml:"Body"`
}

// 地震情報等におけるHeadline
type EarthquakeHeadline struct {
	// 見出し分を自由文形式で掲載する
	Text string `xml:"Text"`

	// 見出し防災気象情報事項
	//
	// 震度速報では、@type が“ 震度速報” である本要素が１ 回出現する。情報形態
	// （Head/InfoType）が“取消”の場合は出現しない。
	// 地震情報(震源・震度に関する情報)では、@type が“震源・震度に関する情報（細分区域）”、
	// “震源・震度に関する情報（市町村等）”である本要素が各々１回ずつ出現する。なお、以下の
	// 場合は本要素は出現しない。
	// ・観測された震度が全て２以下だった場合
	// ・震度が観測されなかった場合（遠地地震の場合など）
	// ・情報形態（Head/InfoType）が“取消”の場合
	// 長周期地震動に関する観測情報では、@type が“長周期地震動に関する観測情報（細分
	// 区域）”である本要素が１回出現する。なお、以下の場合は本要素は出現しない。
	// ・情報形態（Head/InfoType）が“取消”の場合
	// 地震情報(震源に関する情報)、地震情報(地震の活動状況等に関する情報)、地震情報(地
	// 震回数に関する情報)、及び地震情報(顕著な地震の震源要素更新のお知らせ)には、本要素
	// は出現しない。
	Information []struct {
		Type string `xml:"type,attr"`

		// 個々の防災気象情報要素
		//
		// 本要素は、Information/@type の値や観測された最大震度により出現回数が決まる。
		// Information/@type が“震度速報”又は“震源・震度に関する情報(細分区域)”の場合は、観
		// 測された震度のうち、震度３以上の震度階級の数だけ本要素が出現する。
		// Information/@type が“長周期地震動に関する観測情報（細分区域）”の場合は、観測された
		// 長周期地震動階級の数だけ本要素が出現する。
		// Information/@type が“震源・震度に関する情報(市町村等)”の場合、当面は下表に示す震
		// 度階級の要素が出現する。また、基準となる震度以上と考えられるが情報発表時点で震度が
		// 入電していない市町村がある場合は、その旨を記載するための要素が追加される。当面は震
		// 度５弱を基準とし、震度５弱以上と考えられるが震度が入電していない市町村を、“震度５弱以
		// 上未入電”の要素に記載する。
		// 子要素として、Kind 及びAreas をもつ。
		Item []struct {
			// 防災気象情報要素
			//
			//   - 津波警報・注意報・予報の場合
			//     津波警報等の種類を記載する。子要素にName とCode をもつ。
			//
			//   - 沖合の津波観測に関する情報の場合
			//     子要素にName をもつ。
			Kind struct {
				// 防災気象情報要素名
				//
				// とりうる値：“震度７”、 “震度６強”、
				// “震度６弱”、 “震度５強”、 “震度５弱”、 “震度４”、 “震度３”、 “震度５弱以
				// 上未入電” 、“長周期地震動階級４”、“長周期地震動階級３” 、“長周期地震
				// 動階級２” 、“長周期地震動階級１”）
				//
				// 観測された震度、長周期地震動階級等を記載する。
				// Information/@type が“震源・震度に関する情報(市町村等)”の場合は、11-2(3)-1 の表に示
				// す震度等が記載の対象となる。
				Name string `xml:"Name"`
			} `xml:"Kind"`

			// 対象地域・地点
			//
			// 区域を記載する。Information/@type の値に応じて、@codeType が“地震情報／細分区域”
			// 又は“気象・地震・火山情報／市町村等”に設定される。
			// 子要素にArea をもつ。
			Areas struct {
				CodeType string `xml:"codeType,attr"`

				// 対象地域・地点
				//
				// 子要素にName とCode をもつ。
				Area []struct {
					// 	Areas/@codeType の値に応じて、Kind に記載されている震度、長周期地震動階級を観測し
					// た細分区域又は市町村等のいずれかを記載する（市町村等を記載する場合であって、Kind
					// の内容が“震度５弱以上未入電”の場合は、震度５弱以上未入電と推定される市町村等を記
					// 載する）。
					Name string `xml:"Name"`

					// 上記Name の内容に対応するコードを記載する。参照するコードはAreas/@codetype に記載
					// されている。
					Code string `xml:"Code"`
				} `xml:"Area"`
			} `xml:"Areas"`
		} `xml:"Item"`
	} `xml:"Information"`
}

// 震度速報、震源に関する情報、震源・震度に関する情報
type EarthquakeBody struct {
	// 震度（震度速報、震源・震度に関する情報）
	//
	// 震度に関する情報を記載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Intensity *struct {
		// 震度の観測
		//
		// 震度の観測に関する諸要素を記載する。
		Observation struct {
			CodeDefine CodeDefine `xml:"CodeDefine"`

			// 最大震度
			//
			// 本情報で発表する最大の震度を記載する。
			// 値：“3”/“4”/“5-”/“5+”/“6-”/“6+”/“7”
			MaxInt EarthquakeIntensity `xml:"MaxInt"`

			// 都道府県
			//
			// 都道府県毎の震度の観測状況を記載する。震度を観測した都道府県の数に応じて、本要
			// 素が複数出現する。
			// 子要素Name に都道府県名を記載し、対応するコードを子要素Code に記載する。対応する
			// コードは、「コード体系の定義」（Body/Intensity/Observation/CodeDefine）で定義されている。
			// 具体的なコードの値については、別途提供するコード表を参照。
			Pref []struct {
				Name string `xml:"Name"`
				Code string `xml:"Code"`

				// 最大震度（都道府県）
				//
				// 当該都道府県における最大震度を記載する。
				// 震度速報：“3”/“4”/“5-”/“5+”/“6-”/“6+”/“7”
				// 震源・震度に関する情報：“1”/“2”/“3”/“4”/“5-”/“5+”/“6-”/“6+”/“7”
				MaxInt EarthquakeIntensity `xml:"MaxInt"`

				// 情報の更新（都道府県）
				//
				// 地震情報の続報において、当該都道府県が新規に追加される場合は本要素を追加し、“追
				// 加”と記載する。また、当該都道府県の最大震度が更新された場合も本要素を追加し、“上方
				// 修正”と記載する。
				Revise string `xml:"Revise,omitempty"`

				// 地域
				//
				// 地域毎の震度の観測状況を記載する。震度を観測した地域の数に応じて、本要素が複数出
				// 現する。
				// 子要素Name に地域名を記載し、対応するコードを子要素Code に記載する。対応するコー
				// ドは、「コード体系の定義」（Body/Intensity/Observation/CodeDefine）で定義されている。具
				// 体的なコードの値については、別途提供するコード表を参照。
				Area []struct {
					Name string `xml:"Name"`
					Code string `xml:"Code"`

					// 最大震度（地域）
					//
					// 当該地域における最大震度を記載する。
					// 震度速報：“3”/“4”/“5-”/“5+”/“6-”/“6+”/“7”
					// 震源・震度に関する情報：“1”/“2”/“3”/“4”/“5-”/“5+”/“6-”/“6+”/“7”
					MaxInt EarthquakeIntensity `xml:"MaxInt"`

					// 情報の更新（地域）
					//
					// 地震情報の続報において、当該地域が新規に追加される場合は本要素を追加し、“追加”と
					// 記載する。また、当該地域の最大震度が更新された場合も本要素を追加し、“上方修正”と記
					// 載する。
					Revise string `xml:"Revise,omitempty"`

					// 市町村
					// 震源・震度に関する情報のみ（震度速報にはない）
					//
					// 市町村毎の震度の観測状況を記載する。震度を観測した市町村の数に応じて、本要素が複
					// 数出現する。
					// 子要素Name に市町村名を記載し、対応するコードを子要素Code に記載する。対応するコ
					// ードは、「コード体系の定義」（Body/Intensity/Observation/CodeDefine）で定義されている。
					// 具体的なコードの値については、別途提供するコード表を参照。
					// 当該市町村の中に、基準となる震度以上と考えられるが震度の値を入手していない震度観
					// 測点が存在し、当該市町村の最大震度が基準の震度未満（又は入電なし）の場合は、子要素
					// Condition を追加し、その旨を記載する。当面は震度５弱を基準とし、当該市町村の最大震度
					// が震度４以下（又は入電なし）の場合にCondition が出現し、ここに“震度５弱以上未入電”を
					// 記載する。
					City []struct {
						Name string `xml:"Name"`
						Code string `xml:"Code"`

						Condition string `xml:"Condition"`

						// 最大震度（市町村）
						//
						// 当該市町村における最大震度を記載する。当該市町村内に、基準となる震度以上（当面は
						// 震度５弱以上とする）と考えられるが震度の値を入手していない震度観測点のみしか存在しな
						// い場合、本要素は出現しない。
						// 値：“1”/“2”/“3”/“4”/“5-”/“5+”/“6-”/“6+”/“7”
						MaxInt EarthquakeIntensity `xml:"MaxInt"`

						// 情報の更新（市町村）
						//
						// 地震情報の続報において、当該地域が新規に追加される場合は本要素を追加し、“追加”と
						// 記載する。また、当該地域の最大震度が更新された場合も本要素を追加し、“上方修正”と記
						// 載する。
						Revise string `xml:"Revise,omitempty"`

						// 震度観測点
						//
						// 親要素City に記載した市町村に所属する震度観測点について、観測点毎の震度の観測状
						// 況を記載する。震度を観測した観測点の数に応じて、本要素が複数出現する。
						// 子要素Name に観測点名を記載し、対応するコードを子要素Code に記載する。対応するコ
						// ードは、「コード体系の定義」（Body/Intensity/Observation/CodeDefine）で定義されている。
						IntensityStation []struct {
							Name   string              `xml:"Name"`
							Code   string              `xml:"Code"`
							Int    EarthquakeIntensity `xml:"Int"`
							Revise string              `xml:"Revise,omitempty"`
						} `xml:"IntensityStation"`
					} `xml:"City"`
				} `xml:"Area"`
			} `xml:"Pref"`
		} `xml:"Observation"`
	} `xml:"Intensity,omitempty"`

	// 地震の諸要素（震源に関する情報、震源・震度に関する情報
	//
	// 地震の諸要素（発生日時、震央地名、震源要素、マグニチュード等）を記載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Earthquake *EarthquakeElement `xml:"Earthquake,omitempty"`

	// テキスト要素
	//
	// 自由文形式で追加的に情報を記載する必要がある場合等に、本要素を用いて記載する。例
	// えば、ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合に、取消しの概要等を本要素
	// に記載する。
	Text string `xml:"Text,omitempty"`

	// 付加文
	//
	// 情報の本文に加えて付加的な情報を記載する必要がある場合は、本要素以下に情報を記
	// 載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Comments *struct {
		// 固定付加文
		//
		// 津波や緊急地震速報に関する付加的な情報を、固定付加文の形式で子要素Text に、また、
		// 	対応するコードを子要素Code に記載する。具体的なコードの値については、別途提供するコ
		// ード表を参照。@codeType には“固定付加文”を記載する。
		ForecastComment *FixComment `xml:"ForecastComment,omitempty"`

		// 固定付加文（その他）
		//
		// その他の付加的な情報を、固定付加文の形式で子要素Text に、また、対応するコードを子
		// 要素Code に記載する。具体的なコードの値については、別途提供するコード表を参照。
		// @codeType には“固定付加文”を記載する。
		// 複数の固定付加文を記載する場合、Text においては改行し、Code においてはxs:list 型に
		// よりコードを併記する。
		VarComment *FixComment `xml:"VarComment,omitempty"`

		// 自由付加文
		//
		// その他の付加的な情報を、自由付加文の形式で記載する。
		FreeFormComment string `xml:"FreeFormComment,omitempty"`
	} `xml:"Comments,omitempty"`
}

// 地震の活動状況等に関する情報
type EarthquakeActivityBody struct {
	// 命名地震
	//
	// 顕著な被害を起こした地震について命名した場合は、その名称を記載する。さらに、英語に
	// よる名称がある場合は、@english を追加し、その名称を記載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Naming struct {
		English string `xml:"english,attr,omitempty"`
		Value   string `xml:",chardata"`
	} `xml:"Naming,omitempty"`

	// テキスト要素
	//
	// 自由文の形式で、情報の本文を記載する。
	// また、ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合に、取消しの概要等を本要
	// 素に記載する。
	Text string `xml:"Text"`

	// 付加文
	//
	// 情報の本文に加えて付加的な情報を記載する必要がある場合は、本要素以下に情報を記
	// 載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Comments *struct {
		// 自由付加文
		//
		// その他の付加的な情報を、自由付加文の形式で記載する。
		FreeFormComment string `xml:"FreeFormComment"`
	} `xml:"Comments,omitempty"`
}

// 地震回数に関する情報
type EarthquakeCountInfoBody struct {
	EarthquakeCount *struct {
		Item []struct {
			Type       string `xml:"type,attr"`
			StartTime  string `xml:"StartTime"`
			EndTime    string `xml:"EndTime"`
			Number     int    `xml:"Number"`
			FeltNumber int    `xml:"FeltNumber"`
		} `xml:"Item"`
	} `xml:"EarthquakeCount,omitempty"`

	// テキスト要素
	//
	// 自由文形式で追加的に情報を記載する必要がある場合等に、本要素を用いて記載する。例
	// えば、ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合に、取消しの概要等を本要素
	// に記載する。
	Text string `xml:"Text,omitempty"`

	// 次回発表予定
	//
	// 続報を発表する予定がある場合は、次回発表予定時刻に関する情報を記載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	NextAdvisory string `xml:"NextAdvisory,omitempty"`

	// 付加文
	//
	// 情報の本文に加えて付加的な情報を記載する必要がある場合は、本要素以下に情報を記
	// 載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Comments *struct {
		// 自由付加文
		//
		// その他の付加的な情報を、自由付加文の形式で記載する。
		FreeFormComment string `xml:"FreeFormComment"`
	} `xml:"Comments,omitempty"`
}

// 顕著な地震の震源要素更新のお知らせ
type EarthquakeUpdateBody struct {
	// 地震の諸要素
	//
	// 地震の諸要素（発生日時、震央地名、震源要素、マグニチュード等）を記載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Earthquake *struct {
		// 地震発生時刻
		//
		// 地震の発生した時刻を記載する。
		OriginTime string `xml:"OriginTime"`

		// 地震発現時刻
		//
		// 観測点で地震を検知した時刻（発現時刻）を記載する。
		ArrivalTime string `xml:"ArrivalTime"`

		// 地震の位置要素
		//
		// 地震の位置に関する要素（震央地名、震源要素等）を記載する。
		Hypocenter struct {
			// 震源位置
			//
			// 震源の位置に関する情報を記載する。
			Area struct {
				// 震央地名
				//
				// 	震央地名を記載する。また、これに対応するコードを、後に続く要素Code に記載し、その
				// @type にコード種別“震央地名”と記載する。具体的なコードの値については、別途提供するコ
				// ード表を参照。
				Name string `xml:"Name"`

				Code DetailCode `xml:"Code"`

				// 震源要素
				//
				// 更新後の震源要素を、ISO6709 の規格に従って記載する。日本測地系（度単位）と世界測地
				// 系（度分単位）で表現するため、本要素が２回出現する。
				// 日本測地系については@datum に“日本測地系”を記載し、世界測地系については@datum
				// に代わって@type が出現し、ここに“ 震源位置（ 度分） ” を記載する。また、いずれも
				// @description に文字列表現を記載する。本要素に記載する深さの値は、深さ700km より浅いと
				// ころでは世界測地系については1,000 メートル、日本測地系については10,000 メートルの単
				// 位で有効であり、@description における深さは、世界測地系については1km 単位で、日本測
				// 地系については1,000 メートルの位を四捨五入して10km 単位で表現する。
				Coordinate []struct {
					Value       string `xml:",chardata"`
					Type        string `xml:"type,attr"`
					Description string `xml:"description,attr"`
					Datum       string `xml:"datum,attr,omitempty"`
				} `xml:"Coordinate"`
			} `xml:"Area"`
		} `xml:"Hypocenter"`

		// マグニチュード
		//
		// 地震のマグニチュードの値を記載する。@type にはマグニチュードの種別を、@descripion に
		// は文字列表現を記載する。
		// また、マグニチュードが不明の場合やマグニチュードが8 を超える巨大地震と推定される場
		// 合は、これらの属性に加えて@condition が出現し、ここにマグニチュードが不明である旨を示
		// す固定値“不明”を記載する。マグニチュードの値には“NaN”を記載する。
		Magnitude Magnitude `xml:"Magnitude"`
	} `xml:"Earthquake,omitempty"`

	// テキスト要素
	//
	// 自由文形式で追加的に情報を記載する必要がある場合等に、本要素を用いて記載する。例
	// えば、ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合に、取消しの概要等を本要素
	// に記載する。
	Text string `xml:"Text,omitempty"`

	// 付加文
	//
	// 情報の本文に加えて付加的な情報を記載する必要がある場合は、本要素以下に情報を記
	// 載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Comments *struct {
		// 自由付加文
		//
		// その他の付加的な情報を、自由付加文の形式で記載する。
		FreeFormComment string `xml:"FreeFormComment"`
	} `xml:"Comments,omitempty"`
}
