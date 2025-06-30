<!doctype html>
<html>
  <head>
    <meta charset="utf-8">
    <title>Login</title>
  </head>
  <body>    
      <form method="post" action="/register">
        <label>
          Username:
          <input name="username" placeholder="username" required>
        </label>
        <br>
         <label>
          Email:
          <input name="email" placeholder="email" required>
        </label>
        <br>
        <label>
          Password:
          <input type="password" name="password" placeholder="password" required>
        </label>
        <br>
        <br>
        <label>
          Repeat Your Password:
          <input type="password" name="repeatedPassword" placeholder="repeat password" required>
        </label>
        <br>
        <button type="submit">Sign up</button>
      </form>
  </body>
</html>
