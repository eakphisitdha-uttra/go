package repositories

import (
	"context"
	"database/sql"
	"microservice/databases/postgresql/tables"
	"microservice/internals/module_a/adapters/inputs"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockery --with-expecter --name "IRepository" --output $PWD/mocks
type IRepository interface {
	Get() ([]tables.Users, error)
	Add(input inputs.AddInput) error
	Update(input inputs.UpdateInput) error
	Delete(input inputs.DeleteInput) error
	Log(field string, values interface{}, id int, action string) error
}

type Repository struct {
	pg *sql.DB
	mg *mongo.Client
}

func NewRepository(pg *sql.DB, mg *mongo.Client) IRepository {
	return &Repository{pg: pg, mg: mg}
}

func (r *Repository) Get() ([]tables.Users, error) {
	//
	//SQL logic
	//
	data := []tables.Users{}
	return data, nil
}

func (r *Repository) Add(input inputs.AddInput) error {
	//
	//SQL logic
	//

	return nil
}

func (r *Repository) Update(input inputs.UpdateInput) error {
	//
	//SQL logic
	//

	return nil
}

func (r *Repository) Delete(input inputs.DeleteInput) error {
	//
	//SQL logic
	//

	return nil
}

func (r *Repository) Log(field string, value interface{}, id int, action string) error {
	//
	// logic
	//
	// create log in MongoDB
	doc := bson.D{{Key: "time", Value: time.Now().Format("2-Jan-06 03:04PM")}, {Key: field, Value: value}, {Key: "user", Value: id}, {Key: "action", Value: action}}

	_, err := r.mg.Database("your_database_name").Collection("logs").InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	return nil
}
