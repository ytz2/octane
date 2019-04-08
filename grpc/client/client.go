package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	pb "octane/grpc/lotto"

	"google.golang.org/grpc"
)

const (
	address = "localhost:50551"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLottoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	defer cancel()

	var id int64 = 13790
	var counter = 1
	for {
		r := &pb.GetUsersRequest{
			PageSize: 100,
			Id:       fmt.Sprint(id),
		}
		ret, err := c.GetUsers(ctx, r)
		if err != nil {
			fmt.Printf("Error %s", err.Error())
			break
		}
		id, err = strconv.ParseInt(ret.NextId, 10, 32)
		fmt.Printf("Page %d\n\n:", counter)
		//id++
		counter++
		for _, u := range ret.Users {
			fmt.Print(*u)
		}
		fmt.Println()
	}
	/*start := time.Now()

	var loop = 10000
	var st = 3790
	for i := st; i < loop+st; i++ {
		_, err = c.AddUser(ctx, &pb.AddUserRequest{Name: "Yanhua"})
		log.Println("Adding user Yanhua")
		if err != nil {
			log.Fatalf("could not add user: %v", err)
		} else {
			u, err := c.GetUser(ctx, &pb.GetUserRequest{Id: fmt.Sprint(i)})
			if err != nil {
				log.Fatalf("could not get user: %v", err)
			}
			log.Printf("Get user %v after adding", u)
			_, err = c.UpdateUser(ctx, &pb.UpdateUserRequest{Id: fmt.Sprint(u.Id), Name: "Danni"})
			if err != nil {
				log.Fatalf("could not get user: %v", err)
			}
			u, err = c.GetUser(ctx, &pb.GetUserRequest{Id: fmt.Sprint(i)})
			log.Printf("GOt user %v", u)
			_, err = c.DeleteUser(ctx, &pb.DeleteUserRequest{Id: fmt.Sprint(i)})
		}
	}

	diff := time.Now().Sub(start).Seconds()
	rps := float64(loop) / diff
	fmt.Printf("RPS = %v / second to run.\n", rps)*/
}
