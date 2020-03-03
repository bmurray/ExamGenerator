package webserver

const question_templ = `
<!DOCTYPE html>
<html>
<head>
<title>Questions ( {{ .Seed }} )</title>
<style type="text/css">
.question {
	break-inside: avoid;
}
.question ol {
	list-style-type: upper-alpha;
}
</style>
</head>
<body>
{{range $i, $q := .Questions}}
<div class="question">
<h3>{{ oneIndex $i }}: {{$q.Question}}</h3>
<ol>
{{range $q.Answers}}
<li>{{.Answer}}</li>
{{ end }}
{{if $q.AllOfTheAbove}}
<li>All of the above</li>
{{end}}
{{if $q.NoneOfTheAbove}}
<li>None of the above</li>
{{end}}
</ol>
</div>
{{ else }}
No Questions
{{end}}
</body>
</html>
`
const answer_templ = `
<!DOCTYPE html>
<html>
<head>
<title>Answers {{ .Seed }}</title>
</head>
<body>
{{range $i, $q := .Questions}}
<div class="question">
{{ oneIndex $i }}: {{range $q.CorrectAnswers }}{{.}}{{end}}
{{ else }}
No Questions
{{end}}
</body>
</html>
`
