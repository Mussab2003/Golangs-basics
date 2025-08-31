package main

import (
	"fmt"
	"context"
	"github.com/Mussab2003/Golangs-basics/web-app/application"
)

func main(){
	app := application.New()
	err := app.Start(context.TODO())

	if err != nil {
		fmt.Println("Failed to start application:", err)
	}
	
}

