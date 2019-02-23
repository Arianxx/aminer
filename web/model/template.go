package model

var ListTemplate = `
query {{.Name}}($offset: int, $first: int
	{{ range $k, $v := .Variables }}
		,{{ $k }}: {{ $v }}
	{{ end }}
){
	data(func: {{ .Function }}, first: $first, offset: $offset)
	{{ if .Filters }}
		@filter(
		{{ range $i, $f := .Filters }}
			{{ if eq $i 0 }}
				{{ $f }}
			{{ end }}
			and {{ $f }}
		{{ end }}
		)
	{{ end }}
	{
		{{ .Want }}
	}
	info(func: {{ .Function }})
	{{ if .Filters }}
		@filter(
		{{ range $i, $f := .Filters }}
			{{ if eq $i 0 }}
				{{ $f }}
			{{ end }}
			and {{ $f }}
		{{ end }}
		)
	{{ end }}
	{
		num: count(uid)
	}
}
`

var CountTemplate = `
query {{.Name}}($offset: int, $first: int
	{{ range $k, $v := .Variables }}
		,{{ $k }}: {{ $v }}
	{{ end }}
){
	{{ range $k, $v := .Other }}
		{{ $v }}(func: {{ $.Function }})
		@filter( eq( {{ $.Name }}, "{{ $v }}")) {
			num: count(uid)
		}
	{{ end }}
}
`
