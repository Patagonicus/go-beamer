<html>
<body>
<h1>Edit playlist "{{.Playlist.Name}}"</h1>
<form action="/edit/playlist/{{.Playlist.ID}}/save" method="post">
  Name: <input type="text" name="name" value="{{.Playlist.Name}}" /><br />
  <input type="submit" value="Submit" />
</form>

Current entries:
<table>
{{range $i, $e := .Entries}}
<tr><td>{{$i}}</td><td>{{$e.ID}}</td><td>{{$e.Name}}</td></tr>
{{end}}
</table><br />

Available pages:
<table>
{{range .Pages}}
<tr><td>{{.ID}}</td><td>{{.Name}}</td><td><a href="/edit/playlist/add/{{$.Playlist.ID}}/{{.ID}}">Add</a></td></tr>
{{end}}
</table>
</body>
</html>
