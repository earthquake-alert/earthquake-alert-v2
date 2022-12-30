{{- /* 地震回数に関する情報のテンプレート */ -}}
{{- convert .Parsed.Head.Headline.Text true -}}

{{- $c := .TempCountText -}}
{{- if $c -}}
    {{- printf "\n" -}}{{- $c.FormattedStartTime -}}から{{- $c.FormattedEndTime -}}までの間に、{{- $c.Number -}}回
    {{- if ne $c.FeltNumber -1 -}}
        （うち有感地震{{- $c.FeltNumber -}}回）
    {{- end -}}
    の地震が発生しています。
{{- end -}}

{{- if ne .Parsed.Body.NextAdvisory "" -}}
    {{- printf "\n" -}}{{ convert .Parsed.Body.NextAdvisory false -}}
{{- end -}}

{{- if ne .Parsed.Body.Text "" -}}
    {{- printf "\n" -}}{{ convert .Parsed.Body.Text true -}}
{{- end -}}

{{- if .Parsed.Body.Comments -}}
    {{- printf "\n" -}}{{- convert .Parsed.Body.Comments.FreeFormComment true -}}
{{- end -}}
