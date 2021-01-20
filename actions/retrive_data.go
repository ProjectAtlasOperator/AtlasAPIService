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

	collection := client.Database("allPods").Collection("podInformation")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	var volumes []bson.M
	if err = cursor.All(ctx, &volumes); err != nil {
		log.Fatal(err)
	}

	for _, volume := range volumes {
		c.Set("id", volume["podInformations"])
	}

	//tableString := &strings.Builder{}
	//table := tablewriter.NewWriter(tableString)
	//
	//data := [][]string{
	//	[]string{"Pali", "The Good", "500"},
	//	[]string{"Gabi", "The Very very Bad Man", "288"},
	//	[]string{"C", "The Ugly", "120"},
	//	[]string{"D", "The Gopher", "800"},
	//}
	//
	//table = tablewriter.NewWriter(os.Stdout)
	//table.SetHeader([]string{"Name", "Sign", "Rating"})
	//
	//for _, v := range data {
	//	table.Append(v)
	//}
	//
	//f, err := os.Create("data.txt")
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer f.Close()
	//table.Render() // Send output
	//fmt.Println(tableString.String())
	//_, err2 := f.WriteString(tableString.String())
	//
	//if err2 != nil {
	//	log.Fatal(err2)
	//}
	//
	//fmt.Println("done")

	return c.Render(http.StatusOK, r.HTML("retrieve_data.html"))
}
