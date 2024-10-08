package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	"gitlab.com/ltp2-c-megalodonte/ps-backend-augusto-cesar-joaoguilherme-costa/internal/db"
	"gitlab.com/ltp2-c-megalodonte/ps-backend-augusto-cesar-joaoguilherme-costa/internal/product"
	"gitlab.com/ltp2-c-megalodonte/ps-backend-augusto-cesar-joaoguilherme-costa/internal/user"
)

// Definir a variável store globalmente
var store = sessions.NewCookieStore([]byte("something-very-secret"))

func main() {
	// Conecta ao banco de dados
	if err := db.ConnectDB(); err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v\n", err)
	}

	if db.IsDBConnected() {
		fmt.Println("Conexão com o banco de dados estabelecida com sucesso!")
	} else {
		fmt.Println("Erro ao verificar a conexão com o banco de dados.")
		return
	}

	// Configure as rotas com o middleware CORS
	http.HandleFunc("/user/register", CORS(user.RegisterUser))
	http.HandleFunc("/user/login", CORS(user.AuthenticateUser(store)))
	http.HandleFunc("/newproduct/create", CORS(product.RegisterProduct))
	http.HandleFunc("/search", CORS(product.SearchProduct(store)))
	http.HandleFunc("/product/delete", CORS(AuthenticatedMiddleware(store, product.DeleteProduct(store))))
	http.HandleFunc("/product/update", CORS(AuthenticatedMiddleware(store, product.UpdateProduct(store))))

	// Inicie o servidor
	http.ListenAndServe(":8080", nil)

	/* Configurar rotas com suporte a CORS e sessions
	http.HandleFunc("/user/register", CORS(user.RegisterUser))
	http.HandleFunc("/user/login", CORS(user.AuthenticateUser(store)))
	http.HandleFunc("/newproduct/create", CORS(product.RegisterProduct))
	http.HandleFunc("/search", CORS(product.SearchProduct(store)))
	http.HandleFunc("/product/delete", CORS(AuthenticatedMiddleware(store, product.DeleteProduct(store))))
	http.HandleFunc("/product/update", CORS(AuthenticatedMiddleware(store, product.UpdateProduct(store))))

	*/
	// Iniciar o servidor
	fmt.Println("Listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Printf("Erro ao iniciar o servidor: %v\n", err)
	}
}

// CORS middleware para permitir requisições de diferentes origens
func CORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next(w, r)
	}
}

// Middleware para verificar se o usuário está autenticado
func AuthenticatedMiddleware(store *sessions.CookieStore, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		next(w, r)
	}
}
