package dependencies

import (
	"context"
	"dataon-test/infrastructure/database"
	"dataon-test/repository"
	"dataon-test/service"
	"log"

	"github.com/joho/godotenv"
)


type Dependency struct {
	OrderService service.OrderService
}

func SetupDependencies() Dependency {
	ctx := context.Background()
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	db := database.InitMongoDB(ctx)

	orderRepo := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepo)
	return Dependency{
		OrderService: orderService,
	}
}