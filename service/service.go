package service

import (
	"context"
	"dataon-test/model"
	"dataon-test/repository"
	"time"
)

type OrderService interface {
	Order(ctx context.Context, payload *model.OrderRequest) (model.OrderResult, error)
	FindByID(ctx context.Context, id string) (*model.Orders, error)
	List(ctx context.Context) ([]*model.Orders, error)
	Update(ctx context.Context, payload *model.OrderUpdateRequest)  error
}

type orderService struct {
	orderRepo  repository.OrderRepositoryImpl
}

func NewOrderService(orderRepo repository.OrderRepositoryImpl) OrderService{
	return &orderService{
		orderRepo: orderRepo,
	}
}


func (f orderService) Order(ctx context.Context, payload *model.OrderRequest)  (model.OrderResult, error) {
	mappingOrderRequest := model.NewOrderMapping(*payload)
	err := f.orderRepo.InsertOrder(ctx, mappingOrderRequest)
	if err != nil {
		return model.OrderResult{}, err
	}
	// TODO
	// mapping with master data and calculate total order
	resp := model.NewMappingOrderResponse(*payload)
	
	return resp, nil
}

func (f orderService) FindByID(ctx context.Context, id string) (*model.Orders, error) {
	resp, err := f.orderRepo.FindByID(ctx, id)
	if err != nil {
		return nil, nil
	}

	// TODO
	// mapping with master data and calculate total order
	return resp, nil
}


func (f orderService) List(ctx context.Context) ([]*model.Orders, error)  {
	resp, err := f.orderRepo.List(ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (f orderService) Update(ctx context.Context, payload *model.OrderUpdateRequest)  error  {
	menuOrder := model.MenuOrder{}
	menusOrder := []model.MenuOrder{}
	for _, v := range payload.MenuMakanan {
		menuOrder.IdMakanan = v.IdMakanan
		menuOrder.Qty = v.Qty
		menusOrder = append(menusOrder, menuOrder)
	}

	newPayload := model.OrderUpdate{
		NoTable: payload.NoTable,
		MenuOrder: menusOrder,
		UpdatedAt: time.Now(),
	}
	err := f.orderRepo.Update(ctx, payload.Id, newPayload)
	if err != nil {
		return err
	}
	return nil
}