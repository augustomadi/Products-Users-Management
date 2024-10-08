package api

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"gitlab.com/ltp2-c-megalodonte/ps-backend-augusto-cesar-joaoguilherme-costa/internal/product"
	"gitlab.com/ltp2-c-megalodonte/ps-backend-augusto-cesar-joaoguilherme-costa/internal/user"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))

type Router struct{}

func RunServer() {
	s := &http.Server{
		Addr:         ":8080",
		Handler:      CORS(Router{}), // Aplicando o middleware CORS
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
	}

	log.Fatal(s.ListenAndServe())
}

func (Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("Requested URL: %s", req.URL.Path)
	switch req.URL.Path {
	case "/user/register":
		user.RegisterUser(res, req)
	case "/user/login":
		user.AuthenticateUser(store)(res, req)
	case "/search":
		product.SearchProduct(store)(res, req)
	case "/product/create":
		product.RegisterProduct(res, req)
	case "/product/delete":
		AuthenticatedMiddleware(store, product.DeleteProduct(store))(res, req)
	case "/product/update":
		AuthenticatedMiddleware(store, product.UpdateProduct(store))(res, req)
	case "/test":
		TestRoute(res, req)
	default:
		log.Printf("404 Not Found: %s", req.URL.Path)
		http.NotFound(res, req)
	}
}

// CORS middleware para permitir requisições de diferentes origens
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Middleware para verificar se o usuário está autenticado
func AuthenticatedMiddleware(store *sessions.CookieStore, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		token, ok := session.Values["token"].(string)
		if !ok || token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}

// TestRoute é uma rota de teste simples
func TestRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Write([]byte("Test route reached successfully!"))
}
