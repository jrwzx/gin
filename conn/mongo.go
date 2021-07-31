package conn
import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	_ "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)
var dbGameMain *mongo.Database
var dbGameConfig *mongo.Database
var ctx = context.TODO()
func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("godotenv err:", err)
		os.Exit(2)
	}
	host := os.Getenv("MONGO_HOST")
	envGameMain := os.Getenv("DB_GAME_MAIN")
	envGameConfig := os.Getenv("DB_GAME_CONFIG")

	clientOptions := options.Client().ApplyURI(host)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("connect err:", err)
		os.Exit(2)
	}
	dbGameMain = client.Database(envGameMain)
	dbGameConfig = client.Database(envGameConfig)
}

func GetMongoDB(name string) *mongo.Database {
	switch name {
		case "GAME_CONFIG":
			return dbGameConfig
		default:
			return dbGameMain
	}

}