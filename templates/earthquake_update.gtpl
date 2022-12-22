{{- /* 顕著な地震の震源要素更新のお知らせのテンプレート */ -}}
{{- convert .Parsed.Head.Headline.Text true -}}

■ 更新前震源地: {{- .GetOldEpicenter -}}
■ 更新前規模: {{- .GetOldMagnitude -}}

□ 更新後震源地: {{- .NewLatLonStr -}}
□ 更新後規模: {{- .NewMagnitude -}}

{{- if .Parsed.Body.Comments -}}
    {{ printf "\n" }}{{- convert .Parsed.Body.Comments.FreeFormComment true -}}
{{- end -}}
