<html>
<head>
</head>
<body>
<form action="/upload" method="POST" enctype="multipart/form-data">
<input type="file" name="file" />
<input type="hidden" name="token" value="{{.}}" />
<input type="submit" value="upload" />
</form>
</body>
</html>