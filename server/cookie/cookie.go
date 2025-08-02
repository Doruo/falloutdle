package cookie

import "net/http"

// NewCookie creates and returns a new pointer to a http.Cookie instance.
func NewCookie(name string, value string) *http.Cookie {
	return &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
}

// SetCookie creates and set a http.Cookie to the client.
func SetCookie(w http.ResponseWriter, name string, value string) {
	http.SetCookie(w, NewCookie(name, value))
}
