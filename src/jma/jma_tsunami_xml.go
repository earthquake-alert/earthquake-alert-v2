// 気象庁電文のXML定義
//
// 津波関係のXMLの定義をしています。
package jma

// 津波警報・注意報・予報、津波情報、沖合の津波観測に関する情報
type TsunamiJmaXml struct {
	Control JmaXmlControl `xml:"Control"`

	Head struct {
		JmaXmlHeader

		Headline TsunamiHeadline `xml:"Headline"`
	} `xml:"Head"`

	Body TsunamiBody `xml:"Body"`
}

// 津波に関連する情報におけるHeadline
type TsunamiHeadline struct {
	// 見出し分を自由文形式で掲載する
	Text string `xml:"Text"`

	// 地震火山関連XML 電文では、情報によって本要素の運用が異なる。このため、以下のとお
	// り個別に解説する。
	//
	// 本要素は津波警報・注意報・予報、沖合の津波観測に関する情報のみに出現し、津波情
	// 報には出現しない。
	// 津波警報・注意報・予報においては、津波予報（若干の海面変動）のみ発表する場合、津
	// 波警報・注意報を全解除する場合、津波警報発表後に切り替わり津波注意報のみとなる場合、
	// 又は情報形態（Head/InfoType）が“取消”の場合を除き、本要素が出現する。
	// 沖合の津波観測に関する情報においては、大津波警報・津波警報に相当する観測値が含
	// まれない場合、又は情報形態（Head/InfoType）が“取消”の場合を除き、本要素が出現する。
	// 本要素が出現する場合、津波警報・注意報・予報においては、@type が“津波予報領域表
	// 現”となり、子要素としてItem をもち、沖合の津波観測に関する情報においては、@type が“沖
	// 合の津波観測に関する情報”となり、子要素としてItem をもつ。
	Information []struct {
		Type string `xml:"type,attr"`

		// 個々の防災気象情報要素 1~2回
		Item []struct {
			// 防災気象情報要素
			//
			//   - 津波警報・注意報・予報の場合
			//     津波警報等の種類を記載する。子要素にName とCode をもつ。
			//
			//   - 沖合の津波観測に関する情報の場合
			//     子要素にName をもつ。
			Kind struct {
				// b横災気象情報要素名
				//
				// - 津波警報・注意報・予報の場合
				// 	津波警報等の名称を記載する。
				// - 沖合の津波観測に関する情報の場合
				// 	本要素の値は“沖合の津波観測に関する情報”となる。
				Name string `xml:"Name"`

				// 防災気象情報要素コード
				// - 津波警報・注意報・予報の場合
				// 	上記Name の内容に対応するコード（“警報等情報要素／津波警報・注意報・予報”）を記載する。
				// - 沖合の津波観測に関する情報の場合
				// 	本要素は出現しない。
				Code string `xml:"Code,omitempty"`
			} `xml:"Kind"`

			// 対象地域・地点
			//
			//   - 津波警報・注意報・予報の場合
			//     津波警報等の対象となる津波予報区、津波予報区結合表現、又は領域表現を記載する。
			//     子要素にArea をもつ。
			//
			//   - 沖合の津波観測に関する情報の場合
			//     大津波警報・津波警報に相当する高い津波が観測された沖合の潮位観測点を記載する。
			//     子要素にArea をもつ。
			Areas struct {
				CodeType string `xml:"codeType,attr"`

				// 対象地域・地点
				//
				//   - 津波警報・注意報・予報の場合
				//     Kind の内容に対応する津波警報等の対象となる、津波予報区、津波予報区結合表現、又
				//     は領域表現の数と同数出現する。
				//     子要素にName とCode をもつ。
				//
				//   - 沖合の津波観測に関する情報の場合
				//     大津波警報・津波警報に相当する高い津波が観測された沖合の潮位観測点又は観測点名
				//     称を簡略化した表現（複数の観測点で同じ表現となる場合は1 回だけ記載する。）の数と同数
				//     出現する。
				//     子要素にName とCode をもつ。
				Area []struct {
					// 対象地域・地点名称
					//
					// - 津波警報・注意報・予報の場合
					// 	津波予報区、津波予報区結合表現、又は領域表現を記載する。
					//
					// - 沖合の津波観測に関する情報の場合
					// 	沖合の潮位観測点又は観測点名称を簡略化した表現（複数の観測点で同じ表現となる場
					// 合は1 回だけ記載する。）を記載する。
					Name string `xml:"Name"`

					// 対象地域・地点コード
					//
					// - 津波警報・注意報・予報の場合
					// 	上記Name の内容に対応するコード（“津波予報区”）を記載する。
					//
					// - 沖合の津波観測に関する情報の場合
					// 	上記Name の内容に対応するコード（“潮位観測点”）を記載する。“潮位観測点”コード表に
					// 	は、各観測点を示すコードと、観測点名称を簡略化した表現（複数の観測点を代表する地点
					// 	として抜粋して用いられる観測点名）を示すコードが含まれており、簡略化した観測点名称に
					// 	対しては、その名称に対応するコードを記載する。この簡略化した観測点名称は、「ヘッダ部」
					// 	（Head）に記載する場合のみ使用し、「内容部」（Body）では使用しない。このヘッダ部に出現
					// 	する簡略化した観測点名称は、電文の内容を簡潔に伝えることを目的としたものであり、実際
					// 	にどの観測点で観測したかを知るためには、内容部を参照することを想定している。
					Code string `xml:"Code"`
				} `xml:"Area"`
			} `xml:"Areas"`
		} `xml:"Item"`
	} `xml:"Information"`
}

// 内容部
// 本情報の量的な詳細内容を記載する。
type TsunamiBody struct {
	// 津波に関連する情報を記載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Tsunami struct {
		// 津波の観測値
		//
		// 津波が観測された場合、本要素に津波の観測に関する情報を記載する。
		Observation *struct {
			// コード体系の定義
			//
			// 「津波の観測」（Body/Tsunami/Observation）以下で使用するコード体系を定義する。使用
			// するコードの種類に応じて子要素Type が出現し、ここにコード種別を記載する。さらに、Type
			// の@xpath として、定義したコードを使用する要素の相対的な出現位置を記載する。
			CodeDefine CodeDefine `xml:"CodeDefine"`

			// 津波の観測値（津波予報区毎）
			//
			// 津波予報区毎に津波の観測値を記載する。津波を観測した津波予報区の数に応じて、本
			// 要素が複数出現する。
			Item []struct {
				// 津波予報区
				//
				// 当該津波予報区について、その名称を子要素Name に、対応するコードを子要素Code に記
				// 載する。対応するコードは、「コード体系の定義」（Body/Tsunami/Observation/CodeDefine）
				// で定義されている。具体的なコードの値については、別途提供するコード表を参照。
				Area struct {
					Name string `xml:"Name"`
					Code string `xml:"Code"`
				} `xml:"Area"`

				// 潮位観測点
				//
				// 潮位観測点毎に津波の観測値を記載する。津波を観測した潮位観測点の数に応じて、本
				// 要素が複数出現する。
				// 潮位観測点の名称を子要素Name に、対応するコードを子要素Code に記載する。対応する
				// コードは、「コード体系の定義」（Body/Tsunami/Observation/CodeDefine）で定義されている。
				// 具体的なコードの値については、別途提供するコード表を参照。
				Station []struct {
					Name string `xml:"Name"`
					Code string `xml:"Code"`

					Sensor string `xml:"Sensor,omitempty"`

					// 津波の第１波（観測値）
					//
					// 観測した津波の第１波について、子要素ArrivalTime に観測時刻を、子要素Initial に極性を
					// 記載する。
					// 津波の最大波を観測したものの第１波を観測できなかった場合は、子要素ArrivalTime 及び
					// 子要素Initial に代わって子要素Condition が出現し、ここに“第１波識別不能”と記載する。
					// 続報において、新たに本要素が出現する場合は子要素Revise に“追加”を、既出であった
					// 本要素の内容が更新される場合は“更新”を記載する。
					FirstHeight struct {
						ArrivalTime string `xml:"ArrivalTime,omitempty"`
						Initial     string `xml:"Initial,omitempty"`
						Condition   string `xml:"Condition,omitempty"`
						Revise      string `xml:"Revise,omitempty"`
					} `xml:"FirstHeight"`

					// 津波の最大波（観測値）
					//
					// 観測したこれまでの最大波について、子要素DateTime に観測時刻を、子要素
					// jmx_eb:TsunamiHeight に観測した津波の高さを記載する。
					// 子要素jmx_eb:TsunamiHeight の@type に“これまでの最大波の高さ”、@unit に津波の高さの
					// 単位である“m”、@description に文字列表現を記載する。また、これまでの最大波の高さが測
					// 定範囲を超え、「～以上」と表現する場合は、事例に示すとおり@description に記載する。水位
					// が上昇中の場合は、子要素jmx_eb:TsunamiHeight に@condition が出現し、“上昇中”を記載
					// する。
					// 津波注意報の予報区（警報・注意報を解除した予報区も含む）において、観測されたこれま
					// での最大波が非常に小さい場合は、子要素jmx_eb:TsunamiHeight に代わって子要素
					// Condition が出現し、ここに“微弱”と記載する。また、津波警報以上の津波予報区において、
					// 観測されたこれまでの最大波の高さが予想される高さに比べて十分小さい場合は、子要素
					// DateTime 及び子要素jmx_eb:TsunamiHeight に代わって子要素Condition が出現し、ここに
					// “観測中”と記載する。
					// これまでの最大波の高さが大津波警報の基準を超え、追加あるいは更新された場合は、子
					// 要素Condition を追加し、ここに“重要”と記載する。 続報において、新たに本要素が出現す
					// る場合は子要素Revise に“追加”を、既出であった本要素の内容が更新される場合は“更新”
					// を記載する。
					MaxHeight struct {
						DateTime      string `xml:"DateTime,omitempty"`
						ArrivalTime   string `xml:"ArrivalTime,omitempty"`
						Condition     string `xml:"Condition,omitempty"`
						Revise        string `xml:"Revise,omitempty"`
						TsunamiHeight *struct {
							Value       string `xml:",chardata"`
							Type        string `xml:"type,attr"`
							Unit        string `xml:"unit,attr"`
							Condition   string `xml:"condition,attr"`
							Description string `xml:"description,attr"`
						} `xml:"TsunamiHeight,omitempty"`
					} `xml:"MaxHeight"`
				} `xml:"Station"`
			} `xml:"Item"`
		} `xml:"Observation,omitempty"`

		// 津波の予測値
		//
		// 津波警報・注意報・予報に関する情報を本要素に記載する。
		Forecast *struct {
			// コード体系の定義
			//
			// 「津波の予測」（Body/Tsunami/Forecast）以下で使用するコード体系を定義する。使用する
			// コードの種類に応じて子要素Type が出現し、ここにコード種別を記載する。さらに、Type の
			// @xpath として、定義したコードを使用する要素の相対的な出現位置を記載する。
			CodeDefine CodeDefine `xml:"CodeDefine"`

			// 津波の予測値（津波予報区毎）
			//
			// 本情報で津波警報・注意報や津波予報（若干の海面変動）を発表している津波予報区及び
			// 津波警報・注意報を解除した津波予報区について、発表状況を記載する。記載する津波予報
			// 区の数に応じて、本要素が複数出現する。
			Item []struct {
				// 津波予報区
				//
				// 対象となる津波予報区の名称を子要素Name に、対応するコードを子要素Code に記載する。
				// 対応するコードは、「コード体系の定義」（Body/Tsunami/Forecast/CodeDefine）で定義されて
				// いる。具体的なコードの値については、別途提供するコード表を参照。
				Area struct {
					Name string `xml:"Name"`
					Code string `xml:"Code"`
				} `xml:"Area"`

				// 津波警報等の種類
				//
				// 本情報による、当該津波予報区の津波警報等の発表状況を子要素Kind に記載する。また、
				// 発表状況の状態遷移を表すために、一つ前の情報による発表状況を子要素LastKind に記載
				// する。さらに、各要素の子要素Name 及びCode に、その名称と対応するコードを記載する。対
				// 応するコードは、「コード体系の定義」（Body/Tsunami/Forecast/CodeDefine）で定義されてい
				// る。具体的なコードの値については、別途提供するコード表を参照。
				// 大津波警報については、第1 報で大津波警報となる津波予報区および切り替え（更新報）で
				// 新たに大津波警報となる津波予報区においては”大津波警報：発表”、大津波警報を継続す
				// る津波予報区においては”大津波警報”を記載する。
				Category struct {
					Kind *struct {
						Name string `xml:"Name"`
						Code string `xml:"Code"`
					} `xml:"Kind"`

					LastKind struct {
						Name string `xml:"Name"`
						Code string `xml:"Code"`
					} `xml:"LastKind"`
				} `xml:"Category"`

				// 当該津波予報区への第１波の到達予想時刻を、子要素ArrivalTime に記載する。
				// 本情報の発表時点において、第１波の到達予想時刻までに時間的な猶予が無い場合は、
				// 子要素Condition を追加し、“ただちに津波来襲と予測”を記載する。また、既に第１波が到達
				// したと推測される場合、当該津波予報区内の潮位観測点で第１波が観測された場合は、
				// ArrivalTime に代わって子要素Condition が出現し、それぞれ、“津波到達中と推測”、“第１波
				// の到達を確認”を記載する。
				// 続報において、新たに本要素が出現する場合は子要素Revise に“追加”を、既出であった
				// 本要素の内容が更新される場合は“更新”を記載する。
				// また、津波警報・注意報を解除する又は津波予報（若干の海面変動）を発表している津
				// 波予報区については、本要素は出現しない。
				FirstHeight *struct {
					ArrivalTime string `xml:"ArrivalTime,omitempty"`
					Condition   string `xml:"Condition,omitempty"`
					Revise      string `xml:"Revise"`
				} `xml:"FirstHeight,omitempty"`

				// 当該津波予報区に対して予想される津波の高さを子要素jmx_eb:TsunamiHeight にメートル
				// 単位で記載する。jmx_eb:TsunamiHeight の@type に“津波の高さ”を、@unit に“m”を記載する。
				// また、@description に文字列表現を記載する。発表する津波の高さのとりうる値を下表に示す。
				// jmx_eb:TsunamiHeight に記載する値はxs:float 型とし、「～未満」又は「～超」の表現は、事例
				// に示すとおり@description に記載する。
				// マグニチュードが8 を超える巨大地震と推定されるなど、地震規模推定の不確定性が大きい
				// 場合は、これらの属性に加えて@condition が出現し、ここに津波の高さが不明である旨を示す
				// 固定値“不明”を記載する。津波の高さの値には“NaN”を記載する。また、@description に津
				// 波の高さに関する定性的表現を記載する。発表する定性的表現のとりうる値を下表に示す。
				// 定性的表現がない津波注意報や津波予報の場合は、@description は空属性となる。
				// 大津波警報の津波予報区に対して、予想される津波の高さが最初に数値で発表された場合
				// や、大津波警報の中で予想される津波の高さが上方修正された場合は、子要素Condition を
				// 追加し、ここに”重要”と記載する。
				// 続報において、新たに本要素が出現する場合は子要素Revise に“追加”を、既出であった
				// 本要素の内容が更新される場合は“更新”を記載する。
				// また、津波が減衰して、いずれかの津波予報区で津波警報・注意報等の種類を引き下げる
				// 場合（解除、津波予報（若干の海面変動）への切り替えを含む）は、津波警報・注意報を解除
				// した又は津波予報（若干の海面変動）を発表している全ての津波予報区について本要素
				// は出現しない。
				MaxHeight *struct {
					Condition     string `xml:"Condition,omitempty"`
					TsunamiHeight struct {
						Value       string `xml:",chardata"`
						Type        string `xml:"type,attr"`
						Unit        string `xml:"unit,attr"`
						Description string `xml:"description,attr"`
						Condition   string `xml:"condition,attr"`
					} `xml:"TsunamiHeight"`
					Revise string `xml:"Revise,omitempty"`
				} `xml:"MaxHeight,omitempty"`

				// 潮位観測点
				//
				// 対象となる潮位観測点の名称を子要素Name に、対応するコードを子要素Code に記載する。
				// 対応するコードは、「コード体系の定義」（Body/Tsunami/Forecast/CodeDefine）で定義されて
				// いる。具体的なコードの値については、別途提供するコード表を参照。
				// また、当該観測点での満潮時刻を子要素HighTideDateTime に、津波の到達予想時刻を子
				// 要素FirstHeight に記載する。津波警報・注意報を解除した又は津波予報（若干の海面変動）
				// を発表している津波予報区について、本要素は出現しない。
				Station []struct {
					Name             string `xml:"Name"`
					Code             string `xml:"Code"`
					HighTideDateTime string `xml:"HighTideDateTime"`

					// 津波の到達予想時刻（潮位観測点）
					//
					// 当該潮位観測点への第１波の到達予想時刻を、子要素ArrivalTime に記載する。
					// 本情報の発表時点において、既に第１波が到達したと推測される場合や当該潮位観測点で
					// 第１波が観測された場合は、ArrivalTime に代わって子要素Condition を追加し、それぞれ、
					// “津波到達中と推測”、“第１波の到達を確認”を記載する。
					// 続報において、新たに本要素が出現する場合は子要素Revise に“追加”を、既出であった
					// 本要素の内容が更新される場合は“更新”を記載する。
					FirstHeight struct {
						ArrivalTime string `xml:"ArrivalTime,omitempty"`
						Condition   string `xml:"Condition,omitempty"`
						Revise      string `xml:"Revise"`
					} `xml:"FirstHeight"`
				} `xml:"Station,omitempty"`
			} `xml:"Item"`
		} `xml:"Forecast,omitempty"`

		// 津波の推定値（沖合の津波観測に関する情報のみ）
		//
		// 沖合の潮位観測点で観測された津波の情報に基づき、津波が到達すると推定される沿岸地
		// 域について、津波の推定値に関する情報を記載する。
		Estimation *struct {
			// コード体系の定義
			//
			// 「津波の推定」（Body/Tsunami/Estimation）以下で使用するコード体系を定義する。使用す
			// るコードの種類に応じて子要素Type が出現し、ここにコード種別を記載する。さらに、Type の
			// @xpath として、定義したコードを使用する要素の相対的な出現位置を記載する。
			CodeDefine CodeDefine `xml:"CodeDefine"`

			// 津波の推定値（沿岸地域毎）
			//
			// 沿岸地域毎に推定される津波の到達時刻、高さ等の情報を記載する。推定値を発表する沿
			// 岸地域の数に応じて、本要素が複数出現する。
			Item []struct {
				// 沿岸地域
				//
				// 対象となる沿岸地域の名称を子要素Name に、対応するコードを子要素Code に記載する。
				// 対応するコードは、「コード体系の定義」（Body/Tsunami/Estimation/CodeDefine）で定義され
				// ている。具体的なコードの値については、別途提供するコード表を参照。
				Area struct {
					Name string `xml:"Name"`
					Code string `xml:"Code"`
				} `xml:"Area"`

				// 津波到達時刻（推定値）
				//
				// 当該沿岸地域に第１波が到達すると推定される時刻を子要素ArrivalTime に記載する。
				// 沖合の潮位観測点による観測値から当該沿岸地域への津波到達予想時刻を推定し、推定
				// 時刻よりも早く沿岸地域に津波が到達している可能性がある場合は、子要素Condition を追加
				// し、“早いところでは既に津波到達と推定”と記載する。
				// 続報において、新たに本要素が出現する場合は子要素Revise に“追加”を、既出であった
				// 本要素の内容が更新される場合は“更新”を記載する。
				FirstHeight struct {
					Condition   string `xml:"Condition"`
					ArrivalTime string `xml:"ArrivalTime"`
				} `xml:"FirstHeight"`

				// 津波の高さ（推定値）
				//
				// 沖合の潮位観測点によるこれまでの最大波の観測値から、当該沿岸地域に到達すると推定
				// される時刻を子要素DateTime に、津波の高さを子要素jmx_eb:TsunamiHeight に記載する。
				// 子要素jmx_eb:TsunamiHeight の@type に“津波の高さ”、@unit に津波の高さの単位である
				// “m”、@description に文字列表現を記載する。発表する津波の高さのとりうる値を下表に示す。
				// jmx_eb:TsunamiHeight に記載する値はxs:float 型とし、「～超」の表現は、事例に示すとおり
				// @description に記載する。
				// マグニチュードが8 を超える巨大地震と推定されるなど、地震規模推定の不確定性が大きい
				// 場合は、これらの属性に加えて@condition が出現し、ここに津波の高さが不明である旨を示す
				// 固定値“不明”を記載する。津波の高さの値には“NaN”を記載する。また、@description に津
				// 波の高さに関する定性的表現を記載する。発表する定性的表現のとりうる値を下表に示す。
				// 定性的表現がない津波注意報の場合は、@description は空属性となる。
				// 津波警報以上の沿岸地域に対して推定される津波の高さが、予想される高さに比べて十分
				// 小さい場合は、子要素DateTime 及び子要素jmx_eb:TsunamiHeight に代わって子要素
				// Condition が出現し、ここに“推定中”と記載する（予想される高さが定性的表現で発表されて
				// いる場合を除く）。
				// 推定される津波の高さが大津波警報の基準を超え、追加あるいは更新された場合（定性的
				// 表現から数値表現に変更された場合も含む）は、子要素Condition を追加し、ここに“重要”と
				// 記載する。
				// 続報において、新たに本要素が出現する場合は子要素Revise に“追加”を、既出であった
				// 本要素の内容が更新される場合は“更新”を記載する。
				MaxHeight struct {
					DateTime      string `xml:"DateTime,omitempty"`
					TsunamiHeight *struct {
						Type        string `xml:"type,attr"`
						Unit        string `xml:"unit,attr"`
						Description string `xml:"description,attr"`
						Condition   string `xml:"condition,attr"`
						Value       string `xml:",chardata"`
					} `xml:"TsunamiHeight,omitempty"`
					Revise    string `xml:"Revise,omitempty"`
					Condition string `xml:"Condition,omitempty"`
				} `xml:"MaxHeight"`

				// 潮位観測点
				//
				// 潮位観測点毎に津波の観測値を記載する。津波を観測した潮位観測点の数に応じて、本
				// 要素が複数出現する。
				// 潮位観測点の名称を子要素Name に、対応するコードを子要素Code に記載する。対応する
				// コードは、「コード体系の定義」（Body/Tsunami/Observation/CodeDefine）で定義されている。
				// 具体的なコードの値については、別途提供するコード表を参照。
				Station []struct {
					Name string `xml:"Name"`
					Code string `xml:"Code"`

					Sensor string `xml:"Sensor,omitempty"`

					// 津波の第１波（観測値）
					//
					// 観測した津波の第１波について、子要素ArrivalTime に観測時刻を、子要素Initial に極性を
					// 記載する。
					// 津波の最大波を観測したものの第１波を観測できなかった場合は、子要素ArrivalTime 及び
					// 子要素Initial に代わって子要素Condition が出現し、ここに“第１波識別不能”と記載する。
					// 続報において、新たに本要素が出現する場合は子要素Revise に“追加”を、既出であった
					// 本要素の内容が更新される場合は“更新”を記載する。
					FirstHeight struct {
						ArrivalTime string `xml:"ArrivalTime,omitempty"`
						Initial     string `xml:"Initial,omitempty"`
						Condition   string `xml:"Condition,omitempty"`
						Revise      string `xml:"Revise,omitempty"`
					} `xml:"FirstHeight"`

					// 津波の最大波（観測値）
					//
					// 観測したこれまでの最大波について、子要素DateTime に観測時刻を、子要素
					// jmx_eb:TsunamiHeight に観測した津波の高さを記載する。
					// 子要素jmx_eb:TsunamiHeight の@type に“これまでの最大波の高さ”、@unit に津波の高さの
					// 単位である“m”、@description に文字列表現を記載する。また、これまでの最大波の高さが測
					// 定範囲を超え、「～以上」と表現する場合は、事例に示すとおり@description に記載する。水位
					// が上昇中の場合は、子要素jmx_eb:TsunamiHeight に@condition が出現し、“上昇中”を記載
					// する。
					// 津波注意報の予報区（警報・注意報を解除した予報区も含む）において、観測されたこれま
					// での最大波が非常に小さい場合は、子要素jmx_eb:TsunamiHeight に代わって子要素
					// Condition が出現し、ここに“微弱”と記載する。また、津波警報以上の津波予報区において、
					// 観測されたこれまでの最大波の高さが予想される高さに比べて十分小さい場合は、子要素
					// DateTime 及び子要素jmx_eb:TsunamiHeight に代わって子要素Condition が出現し、ここに
					// “観測中”と記載する。
					// これまでの最大波の高さが大津波警報の基準を超え、追加あるいは更新された場合は、子
					// 要素Condition を追加し、ここに“重要”と記載する。 続報において、新たに本要素が出現す
					// る場合は子要素Revise に“追加”を、既出であった本要素の内容が更新される場合は“更新”
					// を記載する。
					MaxHeight struct {
						ArrivalTime   string `xml:"ArrivalTime,omitempty"`
						Condition     string `xml:"Condition,omitempty"`
						Revise        string `xml:"Revise,omitempty"`
						TsunamiHeight *struct {
							Value       string `xml:",chardata"`
							Type        string `xml:"type,attr"`
							Unit        string `xml:"unit,attr"`
							Condition   string `xml:"condition,attr"`
							Description string `xml:"description,attr"`
						} `xml:"TsunamiHeight,omitempty"`
					} `xml:"MaxHeight"`
				} `xml:"Station"`
			} `xml:"Item"`
		} `xml:"Estimation,omitempty"`
	} `xml:"Tsunami,omitempty"`

	// 地震の諸要素
	//
	// 地震の諸要素（発生日時、震央地名、震源要素、マグニチュード等）を記載する。複数の地
	// 震が原因で本情報を発表する場合は、地震毎に本要素を記載する。
	// ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Earthquake []EarthquakeElement `xml:"Earthquake"`

	// テキスト要素
	//
	// 自由文形式で追加的に情報を記載する必要がある場合等に、本要素を用いて記載する。
	// 例えば、ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合に、取消しの概要等を本要素
	// に記載する。
	Text string `xml:"Text,omitempty"`

	// 付加文
	//
	// 情報の本文に加えて付加的な情報を記載する必要がある場合は、本要素以下に情報を記
	// 載する。ヘッダ部の「情報形態」（Head/InfoType）が“取消”の場合、本要素は出現しない。
	Comments struct {
		// 固定付加文
		//
		// 付加的な情報を、固定付加文の形式で子要素「Text」に、また、対応するコードを子要素
		// Code に記載する。具体的なコードの値については、別途提供するコード表を参照。
		// @codeType には“固定付加文”を記載する。
		// 複数の固定付加文を記載する場合、Text においては改行して空行を挿入し、Code におい
		// てはxs:list 型によりコードを併記する。
		WarningComment FixComment `xml:"WarningComment"`

		// 自由付加文
		// その他の付加的な情報を、自由付加文の形式で記載する。
		FreeFormComment string `xml:"FreeFormComment"`
	} `xml:"Comments,omitempty"`
}
