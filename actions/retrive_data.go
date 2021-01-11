package actions

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"time"

	"github.com/gobuffalo/buffalo"
)

// HomeHandler is a default handler to serve up
// a home page.
func RetrieveData(c buffalo.Context) error {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://username:password@mongodb-service:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	//Connection to MDB
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	collection := client.Database("pods").Collection("podInformation")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var volumes []bson.M
	if err = cursor.All(ctx, &volumes); err != nil {
		log.Fatal(err)
	}
	for _, volume := range volumes {
		c.Set("podName", volume["podName"])
		c.Set("namespace", volume["namespace"])
		c.Set("hostIp", volume["hostIp"])
		c.Set("podIP", volume["podIP"])
		c.Set("startTime", volume["startTime"])
		c.Set("volumePodName", volume["volumePodName"])
		c.Set("volumeName", volume["volumeName"])
		c.Set("volumeMount", volume["volumeMount"])
		c.Set("cpuPodName", volume["cpuPodName"])
		c.Set("cpuUsage", volume["cpuUsage"])
		c.Set("memoryUsage", volume["memoryUsage"])
		c.Set("imageName", volume["imageName"])
		c.Set("mountPath", volume["mountPath"])
		c.Set("configMapName", volume["configMapName"])
		c.Set("NodeName", volume["NodeName"])
		c.Set("nodeMemory", volume["nodeMemory"])
		c.Set("mdbPort", volume["mdbPort"])
		c.Set("mExpressPort", volume["mExpressPort"])
		c.Set("cpsPort", volume["cpsPort"])
	}

	return c.Render(http.StatusOK, r.HTML("retrieve_data.html"))
}
