package auth

import (
	"net/http"
)

var (
	adminUser = "admin"
	adminPass = "password"
)

// AdminMiddleware ensures only authenticated users can access admin routes
func AdminMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != adminUser || password != adminPass {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// LoginHandler handles admin login
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == adminUser && password == adminPass {
			// Redirect to admin dashboard on successful login
			http.Redirect(w, r, "/admin/dashboard", http.StatusSeeOther)
			return
		}

		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Serve the login form
	tmpl := `
    <h1>Login</h1>
    <form method="POST">
        <label>Username:</label>
        <input type="text" name="username" required><br>
        <label>Password:</label>
        <input type="password" name="password" required><br>
        <button type="submit">Login</button>
    </form>
    `
	w.Write([]byte(tmpl))
}
