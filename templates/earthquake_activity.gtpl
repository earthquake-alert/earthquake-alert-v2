{{- /* 地震の活動状況等に関する情報のテンプレート */ -}}
{{- convert .Parsed.Head.Headline.Text false}}

{{ convert .Parsed.Body.Text false -}}
