package main

import (
	"log"

	"github.com/trqvel/web-calc/backend/internal/db"
	"github.com/trqvel/web-calc/backend/internal/handlers"
	"github.com/trqvel/web-calc/backend/internal/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	calcRepo := services.NewRepository(database)
	calcSvc := services.NewService(calcRepo)
	calcHandlers := handlers.NewHandler(calcSvc)

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/calculations", calcHandlers.GetCalculations)
	e.POST("/calculations", calcHandlers.PostCalculations)
	e.PATCH("/calculations/:id", calcHandlers.PatchCalculations)
	e.DELETE("/calculations/:id", calcHandlers.DeleteCalculations)

	e.Start("localhost:8080")
}
