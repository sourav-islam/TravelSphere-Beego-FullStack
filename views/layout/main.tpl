<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>TravelSphere</title>
    <link rel="stylesheet" href="/static/css/style.css">
</head>
<body>
<nav class="navbar">
    <a class="navbar-brand" href="/">TravelSphere</a>
    <div class="nav-links">
        <a href="/" class="{{if eq .ActiveMenu "home"}}active{{end}}">Home</a>
        <a href="/countries" class="{{if eq .ActiveMenu "countries"}}active{{end}}">Countries</a>
        {{if .IsLoggedIn}}
        <a href="/wishlist" class="{{if eq .ActiveMenu "wishlist"}}active{{end}}">Wishlist</a>
        <a href="/dashboard" class="{{if eq .ActiveMenu "dashboard"}}active{{end}}">Dashboard</a>
        {{end}}
    </div>
    <div class="nav-auth">
        {{if .IsLoggedIn}}
        <span>Hi, {{.CurrentUser}}</span>
        <a href="/logout" class="btn-logout">Logout</a>
        {{else}}
        <form action="/login" method="post" class="login-form">
            <input type="text" name="username" placeholder="Username" required />
            <button type="submit">Login</button>
        </form>
        {{end}}
    </div>
</nav>

<main class="container">
    {{.LayoutContent}}
</main>

<footer class="footer">
    <p>TravelSphere &copy; 2025 — Discover. Explore. Wishlist.</p>
</footer>
</body>
</html>