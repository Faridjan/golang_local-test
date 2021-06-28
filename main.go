package main

import (
	"context"
	pb "country/proto/country"
	"log"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func main() {
	clientConn, err := grpc.Dial("localhost:50080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}

	client := pb.NewCountryServiceClient(clientConn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Process...
	//add(ctx, client)
	//addMany(ctx, client)

	getAll(ctx, client)
	//getById(ctx, client)

	//remove(ctx, client)

	log.Println("Success!")

}

/******************************************************************************************
							Getters
******************************************************************************************/
func getById(ctx context.Context, client pb.CountryServiceClient) {
	result, err := client.GetById(ctx, &pb.UUID{
		UUID: "",
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)
}

func getAll(ctx context.Context, client pb.CountryServiceClient) {
	result, err := client.GetAll(ctx, &pb.PaginationRequest{
		Page:   0,
		Limit:  -1,
		Offset: -1,
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result.Countries)
}

/******************************************************************************************
							Setters
******************************************************************************************/
func add(ctx context.Context, client pb.CountryServiceClient) {
	result, err := client.Add(ctx, &pb.CountryRequest{
		Code: "CODE",
		Name: &pb.Name{
			Ru: "COUNTRY",
			En: "COUNTRY",
		},
	})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println(result)
}

func addMany(ctx context.Context, client pb.CountryServiceClient) {
	for i := 0; i < 10; i++ {
		result, err := client.Add(ctx, &pb.CountryRequest{
			Code: strconv.Itoa(i),
			Name: &pb.Name{
				Ru: "COUNTRY " + strconv.Itoa(i),
				En: "COUNTRY " + strconv.Itoa(i),
			},
		})
		if err != nil {
			log.Println(err)
			return
		}

		log.Println(result)
	}
}

/******************************************************************************************
							Remove
******************************************************************************************/
func remove(ctx context.Context, client pb.CountryServiceClient) {
	_, err := client.Remove(ctx, &pb.UUID{
		UUID: "",
	})
	if err != nil {
		log.Println(err)
	}
}
