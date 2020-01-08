package main

import (
	pb "DetailPageService/proto"
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"reflect"
	"sync"
	"time"
)

type Authentication struct {
	Login string
	Password string
}

func (a *Authentication) GetRequestMetadata(context.Context, ...string) (map[string]string , error){
	return map[string]string{
		"login": a.Login,
		"password": a.Password,
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires transport security
func (a *Authentication) RequireTransportSecurity() bool {
	return true
}


func makingConnection() (*grpc.ClientConn, error){
	// creating creds for grpc client connection
	//creds, err := credentials.NewClientTLSFromFile("cert/server.crt", "")
	//if err != nil {
	//	log.Fatalf("could not load tls cert: %s", err)
	//}
	//
	//auth := Authentication{
	//	Login:    "nayan",
	//	Password: "makasare",
	//}

	//conn , err := grpc.Dial("192.168.1.143:7775", grpc.WithTransportCredentials(creds), grpc.WithPerRPCCredentials(&auth))

	conn , err := grpc.Dial("192.168.1.143:7775", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	return conn, err
}


func main(){

	// stress testing run the rpc for how many times
	run_test_index := 1;

	var wg sync.WaitGroup
	conn, err := makingConnection()
	if err != nil {
		log.Fatal(err)
	}
	c :=  pb.NewDetailPageServiceClient(conn)


	wg.Add(run_test_index)
	for i:=0 ; i < run_test_index ; i++ {
		//go CreateSchedule(c, &wg)
		go GetDetails(c, &wg)
		//go GetScheduler(c, &wg)
		//go UpdateSchedule(c, &wg)
	}
	wg.Wait()






	//rabbitMq()


}



type nullawareStrDecoder struct{}

func (nullawareStrDecoder) DecodeValue(dctx bsoncodec.DecodeContext, vr bsonrw.ValueReader, val reflect.Value) error {
	if !val.CanSet() || val.Kind() != reflect.String {
		return errors.New("bad type or not settable")
	}
	var str string
	var err error
	switch vr.Type() {
	case bsontype.String:
		if str, err = vr.ReadString(); err != nil {
			return err
		}
	case bsontype.Null: // THIS IS THE MISSING PIECE TO HANDLE NULL!
		if err = vr.ReadNull(); err != nil {
			return err
		}
	default:
		return fmt.Errorf("cannot decode %v into a string type", vr.Type())
	}

	val.SetString(str)
	return nil
}



func getMongoCollection(dbName, collectionName, mongoHost string)  *mongo.Collection {

	// Register custom codecs for protobuf Timestamp and wrapper types
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	mongoClient, err :=  mongo.Connect(ctx, options.Client().ApplyURI(mongoHost), options.Client().SetRegistry(bson.NewRegistryBuilder().
		RegisterDecoder(reflect.TypeOf(""), nullawareStrDecoder{}).
		Build(),))

	if err != nil {
		log.Println("Error while making collection obj ")
		log.Fatal(err)
	}
	return mongoClient.Database(dbName).Collection(collectionName)
}

const (
	defaultHost = "mongodb://nayan:tlwn722n@cluster0-shard-00-00-8aov2.mongodb.net:27017,cluster0-shard-00-01-8aov2.mongodb.net:27017,cluster0-shard-00-02-8aov2.mongodb.net:27017/test?ssl=true&replicaSet=Cluster0-shard-0&authSource=admin&retryWrites=true&w=majority"
	//developmentMongoHost = "mongodb://192.168.1.9:27017"
	schedularMongoHost = "mongodb://192.168.1.143:27017"
	schedularRedisHost = "192.168.1.143:6379"
)

func GetDetails(c pb.DetailPageServiceClient, wg *sync.WaitGroup){


	tileCollection := getMongoCollection("test", "cwmovies", defaultHost);

	cur, err := tileCollection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(context.Background()){
		resp, err := c.GetDetailInfo(context.Background(), &pb.TileInfoRequest{TileId:cur.Current.Lookup("ref_id").StringValue()})
		if err != nil {
			log.Fatal(err)
		}
		log.Println(resp.String())
	}
	cur.Close(context.Background())
	wg.Done()
}