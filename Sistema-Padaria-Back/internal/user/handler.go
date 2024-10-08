package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"gitlab.com/ltp2-c-megalodonte/ps-backend-augusto-cesar-joaoguilherme-costa/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// User estrutura de dados do usuário
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password"`
	Token     string             `bson:"token" json:"token"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

// RegisterUser handler para o registro do usuário
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Payload Inválido", http.StatusBadRequest)
		return
	}

	// Verificar campos obrigatórios
	if user.Email == "" || user.Password == "" {
		http.Error(w, "É necessário preencher todos os campos", http.StatusBadRequest)
		return
	}

	collection := db.GetCollection("SistemaPadaria", "usuarios")

	// Verificar se o email já existe
	count, err := collection.CountDocuments(context.TODO(), bson.M{"email": user.Email})
	if err != nil {
		http.Error(w, "Falha ao checar o email", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Esse e-mail já está sendo utilizado", http.StatusConflict)
		return
	}

	// Hash da senha antes de salvar no banco de dados
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}
	user.Password = string(hashedPassword)

	// Gerar um token para o usuário
	user.Token, err = GenerateToken()
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	_, err = collection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, fmt.Sprintf("Falha ao criar o usuário %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// LoginRequest estrutura para o login request
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// AuthenticateUser handler para o login do usuário
func AuthenticateUser(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
			return
		}

		var loginReq LoginRequest
		err := json.NewDecoder(r.Body).Decode(&loginReq)
		if err != nil {
			http.Error(w, "Payload Inválido", http.StatusBadRequest)
			return
		}

		fmt.Printf("Login request received for email: %s\n", loginReq.Email)

		collection := db.GetCollection("SistemaPadaria", "usuarios")
		var user User
		err = collection.FindOne(context.TODO(), bson.M{"email": loginReq.Email}).Decode(&user)
		if err != nil {
			fmt.Printf("Error finding user: %v\n", err)
			http.Error(w, "Usuário não encontrado", http.StatusUnauthorized)
			return
		}

		fmt.Printf("User found: %+v\n", user)

		// Compare a senha fornecida com o hash armazenado
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
		if err != nil {
			fmt.Printf("Invalid password for user: %s\n", loginReq.Email)
			http.Error(w, "Senha inválida", http.StatusUnauthorized)
			return
		}

		// Armazenar o token na session
		session, _ := store.Get(r, "session-name")
		session.Values["token"] = user.Token
		err = session.Save(r, w)
		if err != nil {
			http.Error(w, "Falha ao salvar sessão", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Logado com sucesso",
			"token":   user.Token,
		})
	}
}
