<html>
<body>
<h1>Admin</h1>
<table>
{{range .Pages}}
<tr><td>{{.ID}}</td><td><a href="{{.Edit}}">{{.Name}}</a></td></tr>
{{end}}
</table>

<a href="/new/info">New info page</a><br />
<a href="/new/playlist">New playlist</a>
</body>
</html>
