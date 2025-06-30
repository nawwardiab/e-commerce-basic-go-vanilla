<!doctype html>
<html>
  <body>
    {{ if .Logged }}
        <a href="/products">Products</a>
        <a href="/logout">Logout</a>
    {{ end }}
  </body>
</html>
