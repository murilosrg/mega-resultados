package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Config struct to database
type Config struct {
	URL  string
	Name string
}

const defaultURL = "mongodb://root:root@localhost:27017"
const defaultName = "megaresultados"

//Database is mongo database implementation
type Database struct {
	conn *mongo.Database
}

//New create a new connection in database
func New(config Config) *Database {
	url := config.URL
	name := config.Name

	if url == "" {
		url = defaultURL
	}

	if name == "" {
		name = defaultName
	}

	opts := options.Client().ApplyURI(url)
	conn, err := mongo.NewClient(opts)

	if err != nil {
		log.Fatalln("error on connect database", err)
	}

	if err := conn.Connect(context.Background()); err != nil {
		log.Fatalln("error on connection", err)
	}

	return &Database{
		conn: conn.Database(name),
	}
}

//InsertMany is an insert for many documents in the database
func (d Database) InsertMany(collection string, documents []interface{}) (interface{}, error) {
	res, err := d.conn.Collection(collection).InsertMany(nil, documents)

	if err != nil {
		log.Fatalln("error to bulk insert", err)
		return nil, err
	}

	return res.InsertedIDs, nil
}
