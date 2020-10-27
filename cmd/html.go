package main

const dirTemp = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>panand</title>
</head>
<body>
    <ul>
		{{if .Url}} <a href="{{.Url}}">../</a> {{end}}
        {{ range $index, $element := .N }}
            <li> <a href="{{$.C}}/{{$element}}">{{$element}}/</a> </li>
        {{ end }}
		{{ range $index, $element := .F }}
			<li> <a href="{{$.C}}/{{$element}}">{{$element}}</a> </li>
		{{ end }}
    </ul>
</body>
</html>
`

type pathInfo struct {
	Url     string
	C       string
	IsEmpty bool
	N       []string
	F       []string
}
