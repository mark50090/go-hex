package main

import (
	"go-hex/adapters"
	"go-hex/core"

	"go-hex/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// ✅ Connect to MongoDB
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()

	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	// if err != nil {
	// 	panic("failed to connect to MongoDB: " + err.Error())
	// }
	// fmt.Println("connect to MongoDB Success✅")

	// db := client.Database("orders_db")

	db := config.ConnectMongo()
	// ✅ Set up the core service and adapters
	orderRepo := adapters.NewMongoOrderRepository(db)
	orderService := core.NewOrderService(orderRepo)
	orderHandler := adapters.NewHttpOrderHandler(orderService)

	// ✅ Define routes
	app.Post("/order", orderHandler.CreateOrder)

	// ✅ Start the server
	app.Listen(":8000")
}
