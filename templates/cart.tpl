{{define "cart.tpl"}}
<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <title>Your Cart</title>
</head>
<body>
  <nav>
    <a href="/">Home</a> |
    <a href="/products">Products</a> |
    <a href="/logout">Logout</a>
  </nav>

  <h1>Your Shopping Cart</h1>
  {{if .CartItems}}
  <table border="1" cellpadding="8" cellspacing="0">
    <thead>
      <tr>
        <th>Cover</th><th>Title</th><th>Price</th><th>Qty</th><th>Remove</th>
      </tr>
    </thead>
    <tbody>
    {{range .CartItems}}
      <tr>
        <td><img src="{{.Product.Img}}" alt="{{.Product.Title}}" style="height:100px;"></td>
        <td>{{.Product.Title}}</td>
        <td>${{printf "%d" .Product.Price}}</td>
        <td>{{.Quantity}}</td>
        <td>
          <form method="post" action="/cart/remove">
            <input type="hidden" name="product_id" value="{{.ProductID}}">
            <button type="submit">×</button>
          </form>
        </td>
      </tr>
    {{end}}
    </tbody>
  </table>
  <button><a href="/checkout">Proceed to Checkout</a></button>
  {{else}}
    <p>Your cart is empty.</p>
  {{end}}
</body>
</html>
{{end}}
