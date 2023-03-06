package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	start := time.Now();
	ctx := context.Background();
	userID := 10;
	value, err := fetchUserData(ctx, userID);

	if err != nil {
		log.Fatal(err);
	}

	fmt.Println("result : ", value);

	fmt.Println("it took : ", time.Since(start));


}

type Response struct {
	value int
	err error
}

func fetchUserData(ctx context.Context, userID int) (int, error){

	ctx, cancel := context.WithTimeout(ctx, (time.Millisecond * 200))
	defer cancel();

	responseChannel := make(chan Response);

	go func(){
		value, err := fetchThirdPartyStuffThatCanBeSlow();
		responseChannel <- Response{
			value: value,
			err : err,
		}
	}();

	for {
		select {
		case <- ctx.Done():
			return 0, fmt.Errorf("fetching data from 3rd party took too long");
		case response := <- responseChannel:
			return response.value, response.err;
		}
	}
}

func fetchThirdPartyStuffThatCanBeSlow() (int, error) {

	time.Sleep(time.Millisecond * 100);

	return 42, nil
}