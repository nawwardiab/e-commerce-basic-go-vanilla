{{define "singleProduct.tpl"}}
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>{{with .Product}}{{.Title}}{{else}}Product Not Found{{end}}</title>
</head>
<body>
  <nav>
    <a href="/">Home</a> |
    <a href="/products">Products</a> |
    <a href="/cart">My Cart</a> |
    <a href="/logout">Logout</a>
  </nav>

  {{with .Product}}
  <div style="max-width:600px; margin:2em auto;">
    <img src="/static/{{.Img}}" alt="{{.Title}}" style="width:100%; max-width:300px;">
    <h1>{{.Title}}</h1>
    <p>By {{.Artist}} ({{.Year}}) • Genre: {{.Genre}}</p>
    <p><strong>Price:</strong> ${{printf "%d" .Price}}</p>

    <form method="post" action="/cart/add">
      <input type="hidden" name="product_id" value="{{.ID}}">

      <label>
        Quantity:
        <input type="number" name="quantity" value="1" min="1" required>
      </label>

      <button type="submit">Add to Cart</button>
    </form>
  </div>
  {{else}}
    <p>Sorry, we couldn’t find that product.</p>
  {{end}}
</body>
</html>
{{end}}
