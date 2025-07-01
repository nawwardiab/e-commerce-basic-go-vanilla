{{define "products.tpl"}}
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Available Products</title>
  <style>
    ul { list-style:none; padding:0; display:flex; flex-wrap:wrap; }
    li { border:1px solid #ddd; padding:1em; margin:0.5em; width:200px; }
    img { max-width:100%; height:auto; }
    .actions { margin-top:0.5em; display:flex; gap:0.5em; }
    .actions form { display:inline; }
  </style>
</head>
<body>
  <nav>
    <a href="/">Home</a> |
    <a href="/cart">My Cart</a> |
    <a href="/logout">Logout</a>
  </nav>

  <h1>Available Products</h1>

  {{if .Products}}
    <ul>
    {{range .Products}}
      <li>
        <a href="/products/{{.ID}}">
          <img src="/static/{{.Img}}" alt="{{.Title}}">
        </a>

        <h2>{{.Title}}</h2>
        <p>By {{.Artist}} ({{.Year}})</p>
        <p>Genre: {{.Genre}}</p>
        <p><strong>${{printf "%d" .Price}}.00</strong></p>

        <div class="actions">
          <a href="/products/{{.ID}}">View Details</a>
          <form method="post" action="/cart/add">
            <input type="hidden" name="product_id" value="{{.ID}}">
            <input type="hidden" name="quantity" value="1">
            <button type="submit">Add to Cart</button>
          </form>
        </div>
      </li>
    {{end}}
    </ul>
  {{else}}
    <p>No products available at the moment.</p>
  {{end}}
</body>
</html>
{{end}}
