package main

import (
	"fmt"
	"context"
	"time"
	"math/rand"
)

func openConnection(done chan bool){
	if rand.Intn(100) > 50 {
		fmt.Println("OOPS! Hanging Connection")
		time.Sleep(10000 + time.Hour)
	}else{
		fmt.Println("Connection successful")
	}
	done <- true
}

func openConnectionWithContext(){
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()

	done := make(chan bool)
	go openConnection(done)

	select {
		case <- done:
			fmt.Println("Connection Successfull")
		case <- ctx.Done():
			fmt.Println("Connection timeout")
		}
}

func main(){
	openConnectionWithContext()
}