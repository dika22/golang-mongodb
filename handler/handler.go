package handler

import (
	"dataon-test/model"
	"dataon-test/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type HandlerImpl interface {
   Order(c *fiber.Ctx) error
   FindByID(c *fiber.Ctx) error
   List(c *fiber.Ctx) error
   Update(c *fiber.Ctx) error
}

type orderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(hs service.OrderService) HandlerImpl {
	return &orderHandler{
		orderService: hs,
	}
}

func (h orderHandler) Order(c *fiber.Ctx) error  {
	ctx := c.UserContext()
	request := new(model.OrderRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	if err := c.BodyParser(request); err != nil {
		fmt.Println("error body parser")
		return err
	}

	res, err := h.orderService.Order(ctx, request)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(res)
}

func (h orderHandler) FindByID(c *fiber.Ctx) error {
	ctx := c.UserContext()
	id := c.Params("id")

	res, err := h.orderService.FindByID(ctx, id)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(res)
}

func (h orderHandler) List(c *fiber.Ctx) error {
	ctx := c.UserContext()
	res, err := h.orderService.List(ctx)
	if  err != nil {
		return err
	}

	return c.Status(200).JSON(res)
}

func (h orderHandler) Update(c *fiber.Ctx) error {
	ctx := c.UserContext()
	request := new(model.OrderUpdateRequest)
	if err := c.BodyParser(request); err != nil {
		return err
	}

	if err := c.BodyParser(request); err != nil {
		fmt.Println("error body parser")
		return err
	}

	err := h.orderService.Update(ctx, request)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(request)
}