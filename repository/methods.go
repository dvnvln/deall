package repo

import (
	"context"
	"log"

	"github.com/dvnvln/deallscrud/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repository) Disconnect() error {
	return r.client.Disconnect(context.Background())
}

func (r *Repository) Get(ctx context.Context) ([]model.User, error) {
	projection := bson.M{
		"password": 0,
	}
	log.Println("getting user data")
	cursor, err := r.db.Collection(model.UserCollection).Find(ctx, bson.M{}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var result []model.User = nil
	err = cursor.All(ctx, &result)
	return result, err
}

func (r *Repository) GetByUserID(ctx context.Context, userID string) ([]model.User, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}
	log.Println("getting user data: ", objID)
	projection := bson.M{
		"password": 0,
	}
	cursor, err := r.db.Collection(model.UserCollection).Find(ctx, bson.M{model.BsonObjID: objID}, options.Find().SetProjection(projection))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	var result []model.User = nil
	err = cursor.All(ctx, &result)
	return result, err
}

func (r *Repository) Add(ctx context.Context, user model.User) error {
	_, err := r.db.Collection(model.UserCollection).InsertOne(ctx, user)
	if err != nil {
		return err
	}
	log.Println("data added")
	return nil
}

func (r *Repository) Update(ctx context.Context, userID string, user model.User) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	log.Println("getting user data: ", objID)
	user.ID = objID
	_, err = r.db.Collection(model.UserCollection).UpdateOne(ctx, bson.M{model.BsonObjID: objID}, bson.M{"$set": user})
	return err
}

func (r *Repository) Delete(ctx context.Context, userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}
	_, err = r.db.Collection(model.UserCollection).DeleteOne(ctx, bson.M{model.BsonObjID: objID})
	return err
}

func (r *Repository) Login(ctx context.Context, user model.UserReqBody) (model.User, error) {
	log.Print("repo here..")
	uname := user.Username
	var result model.User = model.User{}
	err := r.db.Collection(model.UserCollection).FindOne(ctx, bson.M{model.BsonUname: uname}).Decode(&result)
	log.Print("repo done here..")
	return result, err
}
