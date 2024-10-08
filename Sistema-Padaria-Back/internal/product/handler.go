package product

import (
	"context"
	"encoding/json"
	"fmt"

	"log"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"gitlab.com/ltp2-c-megalodonte/ps-backend-augusto-cesar-joaoguilherme-costa/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Product estrutura de dados do produto
type Product struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string             `bson:"name" json:"name"`
	Quantity  int                `bson:"quantity" json:"quantity"`
	Price     float32            `bson:"price" json:"price"`
	Token     string             `bson:"token" json:"token"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}

func SearchProduct(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("SearchProduct called")

		if r.Method != http.MethodGet {
			log.Println("Método de request inválido:", r.Method)
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		// Pegar o parâmetro de busca da query string
		searchQuery := r.URL.Query().Get("search")
		if searchQuery == "" {
			http.Error(w, "Parâmetro de busca é necessário", http.StatusBadRequest)
			return
		}

		/* Obter o token do cabeçalho da requisição
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := authHeader[len("Bearer "):]
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		*/
		collection := db.GetCollection("SistemaPadaria", "produtos")
		filter := bson.M{
			"name": bson.M{"$regex": searchQuery, "$options": "i"},
			//"token": token,
		}
		findOptions := options.Find()

		cursor, err := collection.Find(context.TODO(), filter, findOptions)
		if err != nil {
			log.Println("Erro ao pesquisar produtos no banco de dados:", err)
			http.Error(w, "Erro ao pesquisar produtos", http.StatusInternalServerError)
			return
		}
		defer cursor.Close(context.TODO())

		var products []Product
		for cursor.Next(context.TODO()) {
			var product Product
			if err := cursor.Decode(&product); err != nil {
				log.Println("Erro ao decodificar produto:", err)
				http.Error(w, "Erro ao decodificar produto", http.StatusInternalServerError)
				return
			}
			products = append(products, product)
		}

		if err := cursor.Err(); err != nil {
			log.Println("Erro durante iteração no cursor:", err)
			http.Error(w, "Erro ao buscar produtos", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(products); err != nil {
			http.Error(w, "Erro ao codificar resposta", http.StatusInternalServerError)
		}
	}
}

// RegisterProduct handler para o registro do produto
func RegisterProduct(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	var product Product
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		http.Error(w, "Payload Inválido", http.StatusBadRequest)
		return
	}

	// Verifique os dados recebidos
	fmt.Printf("Dados recebidos: %+v\n", product)

	// Verificar campos obrigatórios
	if product.Name == "" || product.Price == 0 || product.Quantity == 0 {
		http.Error(w, "É necessário preencher todos os campos", http.StatusBadRequest)
		return
	}

	collection := db.GetCollection("SistemaPadaria", "produtos")

	// Verificar se o produto já existe
	count, err := collection.CountDocuments(context.TODO(), bson.M{"name": product.Name})
	if err != nil {
		http.Error(w, "Falha ao checar produto", http.StatusInternalServerError)
		return
	}
	if count > 0 {
		http.Error(w, "Esse produto já existe", http.StatusConflict)
		return
	}

	product.ID = primitive.NewObjectID()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	_, err = collection.InsertOne(context.TODO(), product)
	if err != nil {
		http.Error(w, fmt.Sprintf("Falha ao criar o produto %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

// DeleteProduct handler para deletar o produto
func DeleteProduct(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
			return
		}

		var request struct {
			ID string `json:"id"`
		}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			log.Println("Error decoding request payload:", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Converting string ID to ObjectID
		objID, err := primitive.ObjectIDFromHex(request.ID)
		if err != nil {
			log.Println("Invalid ID format:", err)
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}

		collection := db.GetCollection("SistemaPadaria", "produtos")
		filter := bson.M{"_id": objID}

		result, err := collection.DeleteOne(context.TODO(), filter)
		if err != nil {
			log.Println("Erro ao deletar o produto do banco de dados:", err)
			http.Error(w, "Falha ao deletar o produto do banco de dados", http.StatusInternalServerError)
			return
		}

		if result.DeletedCount == 0 {
			http.Error(w, "Nenhum produto encontrado com o ID fornecido", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Produto deletado com sucesso"})
	}
}

func UpdateProduct(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("UpdateProduct called")

		if r.Method != http.MethodPut {
			log.Println("Método de request inválido:", r.Method)
			http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
			return
		}

		var product Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			log.Println("Erro ao decodificar o payload:", err)
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if product.ID.IsZero() {
			log.Println("ID do produto não fornecido")
			http.Error(w, "ID do produto é necessário", http.StatusBadRequest)
			return
		}

		// Obter o token do cabeçalho da requisição
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := authHeader[len("Bearer "):]
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Atualizar os campos do produto
		product.Token = token
		product.UpdatedAt = time.Now()

		log.Println("Produto para ser atualizado:", product)

		collection := db.GetCollection("SistemaPadaria", "produtos")
		filter := bson.M{"_id": product.ID}
		update := bson.M{
			"$set": bson.M{
				"name":       product.Name,
				"quantity":   product.Quantity,
				"price":      product.Price,
				"updated_at": product.UpdatedAt,
			},
		}

		result, err := collection.UpdateOne(context.TODO(), filter, update)
		if err != nil {
			log.Println("Erro ao atualizar o produto no banco de dados:", err)
			http.Error(w, "Falha ao atualizar o produto", http.StatusInternalServerError)
			return
		}

		if result.MatchedCount == 0 {
			http.Error(w, "Nenhum produto encontrado com o ID fornecido", http.StatusNotFound)
			return
		}

		log.Println("Produto atualizado com sucesso")

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(product)
	}
}
