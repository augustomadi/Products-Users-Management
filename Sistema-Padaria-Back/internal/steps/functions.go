package steps

/*
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"gitlab.com/ltp2-c-megalodonte/ps-backend-celeste-salvioni/internal/db"

	"go.mongodb.org/mongo-driver/bson"
)

// Creates a Step
func CreateStep(req *http.Request) error {
	var step Step

	err := json.NewDecoder(req.Body).Decode(&step)
	if err != nil {
		panic(fmt.Sprintf("Deu pau no decode da response: %v", err))
	}

	collection := db.ClientDB.Database("oniondb").Collection("response")

	_, err = collection.InsertOne(context.TODO(), step)
	if err != nil {
		return err
	}
	return nil
}

// shows all Steps in database
func GetSteps() []Step {
	var steps []Step
	collection := db.ClientDB.Database("dancedb").Collection("steps")
	result, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		panic(err)
	}

	for result.Next(context.TODO()) {
		var step Step

		err := result.Decode(&step)
		if err != nil {
			panic(fmt.Sprintf("Deu pau no decode da pergunta: %v", err))
		}

		steps = append(steps, step)
	}

	return steps
}
*/
