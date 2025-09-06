package main

import (
	"fmt"
	"context"
	"os"
	"os/signal"
	"github.com/Mussab2003/Golangs-basics/web-app/application"
)

func main(){
	app := application.New()
	
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)

	if err != nil {
		fmt.Println("Failed to start application:", err)
	}
	
}

