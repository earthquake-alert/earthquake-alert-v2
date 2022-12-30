{{- /* 顕著な地震の震源要素更新のお知らせのテンプレート */ -}}
{{- convert .Parsed.Head.Headline.Text true -}}
{{- printf "\n" -}}
{{- if .Parsed.Body.Earthquake -}}
    {{- printf "\n" -}}■{{- printf " " -}}更新前震源地:{{- printf " " -}}{{- .GetOldEpicenter -}}
    {{- printf "\n" -}}■{{- printf " " -}}更新前規模:{{- printf " " -}}{{- .GetOldMagnitude -}}
    {{- printf "\n" -}}↓
    {{- printf "\n" -}}□{{- printf " " -}}更新後震源地:{{- printf " " -}}{{- .NewLatLonStr -}}
    {{- printf "\n" -}}□{{- printf " " -}}更新後規模:{{- printf " " -}}{{- .NewMagnitude -}}
{{- else if .Parsed.Body.Text -}}
    {{- printf "\n" -}}{{- .Parsed.Body.Text -}}
{{- else -}}
    {{- printf "\n" -}}
{{- end -}}
{{- if .Parsed.Body.Comments -}}
    {{- printf "\n" -}}
    {{- printf "\n" -}}{{- convert .Parsed.Body.Comments.FreeFormComment true -}}
{{- end -}}
