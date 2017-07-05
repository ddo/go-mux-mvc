package product

import (
	"encoding/json"
	"errors"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/ddo/go-mux-mvc/db/mongodb"
	"github.com/ddo/go-mux-mvc/models/logger"
)

var (
	ErrInvalidName = errors.New("invalid product name")
	ErrInvalidID   = errors.New("invalid product id")
	ErrInvalidJSON = errors.New("invalid product json")
)

// Product .
type Product struct {
	ID         bson.ObjectId `json:"id" bson:"_id"`
	Name       string        `json:"name" bson:"name"`
	Status     string        `json:"status" bson:"status,omitempty"`
	CreateTime time.Time     `json:"create_time" bson:"create_time"`
}

func (p *Product) String() string {
	b, err := json.Marshal(p)
	if err != nil {
		err = ErrInvalidJSON
		logger.Log("ERR:", err)
		return err.Error()
	}

	return string(b)
}

// New .
func New(name string) (product *Product, err error) {
	logger.Log("name:", name)

	// verify product name
	if name == "" {
		err = ErrInvalidName
		logger.Log("ERR:", err)
		return
	}

	// create new product
	product = &Product{
		Name:       name,
		CreateTime: time.Now(),
	}

	// save product into db
	collection := mongodb.DB.C("product")
	err = collection.Insert(product)
	if err != nil {
		logger.Log("ERR collection.Insert:", err)
		return
	}

	logger.Log("product:", product)
	return
}

// Get .
func Get(id string) (product *Product, err error) {
	logger.Log("id:", id)

	// verify product id
	if !bson.IsObjectIdHex(id) {
		err = ErrInvalidID
		logger.Log("ERR:", err)
		return
	}

	// get product by id
	product = &Product{}
	collection := mongodb.DB.C("product")
	err = collection.FindId(bson.ObjectIdHex(id)).One(product)
	if err != nil {
		if err == mgo.ErrNotFound {
			err = nil
			product = nil
			logger.Log("not found")
			return
		}

		logger.Log("ERR collection.FindId:", err)
		return
	}

	logger.Log("product:", product)
	return
}
