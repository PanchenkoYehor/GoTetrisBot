package lib

const help_html = `<html lang="en">
<head>
<meta charset="UTF-8" http-equiv="refresh" content="0.15">
<title>Tetris Bot by Yehor</title>
</head>
<body>
<h1>Count of destroyed lines: {{.Lines}}</h1>
<ul>
{{range $index, $element := .Pitch}}
<li>{{$element}}</li>
{{end}}
</ul>
</body>
</html>`
