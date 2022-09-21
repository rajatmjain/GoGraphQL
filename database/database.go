package database

import (
	"context"
	"log"
	"time"

	"github.com/rajatjain007/GoGraphQL/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct{
	client *mongo.Client
}

func Connect() *DB{
	client,err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if(err!=nil){
		log.Fatal(err)
	}
	ctx,cancel := context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	return &DB{
		client: client,
	}
}

func (db* DB) Save(input *model.NewDog) *model.Dog{
	collection := db.client.Database("test").Collection("dogs")
	ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
	defer cancel()
	res,err := collection.InsertOne(ctx,input)
	if(err!=nil){
		log.Fatal(err)
	}
	return &model.Dog{
		ID: res.InsertedID.(primitive.ObjectID).Hex(),
		Name: input.Name,
		IsGoodBoy: input.IsGoodBoy,
	}
}

func (db* DB) FindDogByID(ID string) *model.Dog{
	collection := db.client.Database("test").Collection("dogs")
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if(err!=nil){
		log.Fatal(err)
	}
	filter := bson.M{"_id":ObjectID}
	ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
	defer cancel()
	res := collection.FindOne(ctx,filter)
	dog := model.Dog{}
	res.Decode(&dog)
	return &dog	
}

func (db* DB) FindAllDogs() []*model.Dog{
	collection := db.client.Database("test").Collection("dogs")
	ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
	defer cancel()
	cur,err := collection.Find(ctx,bson.D{})
	if(err!=nil){
		log.Fatal(err)
	}
	var dogs []*model.Dog
	for cur.Next(ctx){
		var dog *model.Dog
		err := cur.Decode(&dog)
		if(err!=nil){
			log.Fatal(err)
		}
		dogs = append(dogs, dog)
	}
	return dogs
}