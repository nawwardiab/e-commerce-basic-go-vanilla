<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Login</title>
  </head>
  <body>    
      <form method="post" action="/login">
        <label>
          Username:
          <input name="username" placeholder="username" required>
        </label>
        <br>
        <label>
          Password:
          <input type="password" name="password" placeholder="password" required>
        </label>
        <br>
        <button type="submit">Log in</button>
      </form>
      <p><span>Don't have an account yet? Register </span><a href="/register">here</a></p>
  </body>
</html>
