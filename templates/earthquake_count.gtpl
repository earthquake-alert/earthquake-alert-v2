{{- /* 地震回数に関する情報のテンプレート */}}
{{- convert .Parsed.Head.Headline.Text true}}
{{- $c := .TempCountText }}
{{- if $c }}
  {{ $c.FormattedStartTime }}から{{ $c.FormattedEndTime }}までの間に、{{ $c.Number }}回
    {{- if ne $c.FeltNumber -1 -}}
    （うち有感地震{{ $c.FeltNumber }}回）
    {{- end -}}
    の地震が発生しています。
{{- end }}
{{- if ne .Parsed.Body.NextAdvisory "" }}
{{ convert .Parsed.Body.NextAdvisory false }}
{{- end }}
{{- if ne .Parsed.Body.Text "" }}
{{ convert .Parsed.Body.Text true }}
{{- end }}
{{- if .Parsed.Body.Comments }}
{{ convert .Parsed.Body.Comments.FreeFormComment true }}
{{- end -}}
