package actions

import (
	"context"
	"encoding/json"
	"github.com/gobuffalo/buffalo"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type podInformation map[string]interface{}

type PodInfo struct {
	PodInfoData string `bson:"PodInfoData " json:"PodInfoData"`
}

// PodInfoHander is a default handler to serve up
// a home page.
func PodInfoHander(c buffalo.Context) error {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://username:password@mongodb-service:27017")) // mongodb-service
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//Connection to MDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.Get("http://10.111.92.42:3000/api/podInformation/")

	if err != nil {
		panic(err.Error())
	}

	data, _ := ioutil.ReadAll(res.Body)

	var podInfo podInformation
	json.Unmarshal(data, &podInfo)
	res.Body.Close()

	collection := client.Database("pods").Collection("podInformation")
	_, err = collection.InsertOne(ctx, podInfo)
	if err != nil {
		log.Fatal(err)
	}

	return c.Render(http.StatusOK, r.JSON(podInfo))
}
