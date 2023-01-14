package storage

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
	"skillbox/internal/flags"
	"skillbox/internal/models"
	"strconv"
	"time"
)

type UserRepo struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewUserRepo(ctx *context.Context, DBName string) *UserRepo {
	userData := flags.GetData(ctx)

	opts := options.Client()
	opts.SetAuth(options.Credential{
		AuthMechanism: "SCRAM-SHA-256",
		Username:      "thatguy",
		Password:      "pass12345",
	})
	clientDB, err := mongo.NewClient(opts.ApplyURI(fmt.Sprintf("mongodb://%s", userData.DB)))
	if err != nil {
		panic(err)
	}
	ctxDatabaseClient, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = clientDB.Connect(ctxDatabaseClient)
	if err != nil {
		fmt.Println("error creating db connection: ", err)
		os.Exit(1)
	}
	err = clientDB.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	collectionDB := clientDB.Database(DBName).Collection("users")

	mod := mongo.IndexModel{
		Keys: bson.M{
			"id": 1,
		}, Options: options.Index().SetUnique(true),
	}

	_, err = collectionDB.Indexes().CreateOne(*ctx, mod)
	if err != nil {
		fmt.Println("Indexes().CreateOne() ERROR:", err)
		os.Exit(1)
	}

	return &UserRepo{
		client:     clientDB,
		collection: collectionDB,
	}
}

func (r *UserRepo) Get(id string) *models.User {
	var user *models.User
	filter := bson.M{"id": id}
	opts := options.FindOne().SetProjection(bson.D{{"_id", 0}})
	err := r.collection.FindOne(context.TODO(), filter, opts).Decode(&user)
	if err != nil {
		log.Printf("can't find user %s: %v", id, err)
		return &models.User{}
	}
	return user
}

func (r *UserRepo) Save(u *models.User) (*models.User, error) {
	var lastUser *models.User
	err := r.collection.FindOne(context.Background(), bson.M{}, options.FindOne().SetSort(bson.M{"$natural": -1})).Decode(&lastUser)
	if lastUser == nil {
		u.ID = "0"
	} else {
		lastID, _ := strconv.Atoi(lastUser.ID)
		u.ID = strconv.Itoa(lastID + 1)
	}

	if u.Friends == nil {
		u.Friends = make([]*models.User, 0)
	}

	insertResult, err := r.collection.InsertOne(context.TODO(), u)
	if err != nil {
		return &models.User{}, err
	}
	if insertResult.InsertedID == nil {
		return &models.User{}, errors.New(fmt.Sprintf("user not saved: %v", u.ID))
	}

	user := r.Get(u.ID)

	return user, nil
}

func (r *UserRepo) AllUsers() []models.User {
	opts := options.Find()
	filter := bson.M{}

	var results []models.User

	cur, err := r.collection.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Printf("can't find users: %v", err)
	}

	for cur.Next(context.TODO()) {
		var elem models.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Printf("can't decode user: %v\n", err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Printf("can't find users: %v\n", err)
	}

	err = cur.Close(context.TODO())
	if err != nil {
		log.Printf("can't find users: %v\n", err)
	}

	return results
}

func (r *UserRepo) Delete(id string) error {
	filter := bson.M{"id": id}
	deleteResult, err := r.collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		error500 := errors.New(fmt.Sprintf("can't delete user %s: %v", id, err))
		return error500
	}
	if deleteResult.DeletedCount == 0 {
		error404 := errors.New(fmt.Sprintf("user %v not found", id))
		return error404
	}
	return nil
}

func (r *UserRepo) Update(u *models.User) (*models.User, error) {
	updatedUser := r.Get(u.ID)

	if u.Name != "" {
		updatedUser.Name = u.Name
	}
	if u.Age != 0 {
		updatedUser.Age = u.Age
	}
	if u.Friends != nil {
		updatedUser.Friends = u.Friends
	}

	filter := bson.D{{"id", u.ID}}
	update := bson.D{
		{"$set", bson.D{
			{"name", updatedUser.Name},
			{"age", updatedUser.Age},
			{"friends", updatedUser.Friends},
		}},
	}

	updateResult, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if updateResult.MatchedCount == 0 {
		error404 := errors.New(fmt.Sprintf("user %v not found", u.ID))
		return nil, error404
	}

	if updateResult.ModifiedCount == 0 {
		error404 := errors.New(fmt.Sprintf("user %v not found", u.ID))
		return nil, error404
	}

	updatedUserFromDB := r.Get(u.ID)
	return updatedUserFromDB, nil
}

func (r *UserRepo) MakeFriend(source, target *models.User) (*models.User, error) {
	filter := bson.D{{"id", target.ID}}

	update := bson.D{
		{"$addToSet", bson.D{
			{"friends", source},
		}},
	}

	updateResult, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if updateResult.MatchedCount == 0 {
		error404 := errors.New(fmt.Sprintf("user %v not found", target.ID))
		return &models.User{}, error404
		// TODO: replace with NIL
	}

	if updateResult.ModifiedCount == 0 {
		error500 := errors.New(fmt.Sprintf("user %v was not updated", target.ID))
		return &models.User{}, error500
		// TODO: replace with NIL
	}

	updatedUser := r.Get(target.ID)
	return updatedUser, nil
}

func (r *UserRepo) DeleteFriend(source, target *models.User) (*models.User, error) {
	filter := bson.D{{"id", source.ID}}

	update := bson.D{
		{"$pull", bson.D{
			{"friends", target},
		}},
	}

	updateResult, err := r.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	if updateResult.MatchedCount == 0 {
		error404 := errors.New(fmt.Sprintf("user %v not found", source.ID))
		return &models.User{}, error404
		// TODO: replace with NIL
	}

	if updateResult.ModifiedCount == 0 {
		error500 := errors.New(fmt.Sprintf("user %v was not updated", source.ID))
		return &models.User{}, error500
		// TODO: replace with NIL
	}

	updatedUser := r.Get(source.ID)
	return updatedUser, nil
}

func (r *UserRepo) GetFriends(id string) ([]models.User, error) {
	filter := bson.D{{"id", id}}
	var result models.User

	err := r.collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	var friends []models.User
	for _, friend := range result.Friends {
		friends = append(friends, *friend)
	}

	return friends, nil
}
