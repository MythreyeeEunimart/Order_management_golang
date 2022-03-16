package services

import (
	"context"
	"fmt"
	"net/http"
	"order-management/models"
	"order-management/responses"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "golang.org/x/tools/go/analysis/passes/nilfunc"
)


func CreateOrderStatusHistory (c *fiber.Ctx) error {

	var orderStatus models.Orders

	err := c.BodyParser(&orderStatus); 
	if err != nil {
        return c.JSON(responses.OrderResponse{	
			Status: http.StatusBadRequest, 
			Message: "error", 
			Data: &fiber.Map{"data": err.Error()}})
    }

	errStatus := mgm.CollectionByName("order_status").Create(&orderStatus)

	if errStatus != nil{
		return c.JSON(responses.OrderResponse{
            Status: http.StatusInternalServerError, 
            Message: "Unable to create order", 
            Data: &fiber.Map{"data": err.Error()}})
	}

	return c.JSON(responses.OrderResponse{
        Status: http.StatusCreated,     
        Message: "Order Status History got Created Successfully", 
        Data: &fiber.Map{"data": orderStatus}})
}

func UpdateOrderStatusHistory (c *fiber.Ctx) error{
	var orderStatus models.Orders 

	collection := mgm.CollectionByName("order_status")

	object := bson.M{
		"account_id" : c.Query("account_id"),
	}
	fmt.Println(object)

	err := collection.FindOne(context.Background(), object)

    errStatus := mgm.CollectionByName("order_status").Update(&orderStatus)
	if errStatus!= nil{
		return c.JSON(responses.OrderResponse{
			Status :  http.StatusCreated,
			Message : "FAILED_TO_UPDATE_ORDER",
			Data : &fiber.Map{"data" : err.Err()}})

	}

	return c.JSON(responses.OrderResponse{
		Status: http.StatusCreated,
		Message: "Order Status History Updated Successfully",
		Data : &fiber.Map{"data" : orderStatus}})


	}

  func GetOrderStatusHistory (c *fiber.Ctx) error{

	var orderStatus []models.Orders
    fmt.Println("1")

	object := bson.M{
		"account_id" : c.Query("account_id"),
		"order_id" : c.Query("order_id"),
	}
	findOptions := options.Find()
	findOptions.SetSort(bson.D{{"channel_id", -1}})

	err:= mgm.CollectionByName("order_status").SimpleFind(&orderStatus,object,findOptions)
	// .SetSort(Orders{}.cd) *FindOneOptions
	if err!= nil{
		return c.JSON(responses.OrderResponse{
			Status : http.StatusCreated,
			Message :"Failed to fetch data",
			Data : &fiber.Map{"data": err.Error()}})
            
	}
	 return c.JSON(responses.OrderResponse{
		Status : http.StatusCreated,
		Message :"Data fetched successfully",
		Data : &fiber.Map{"data": orderStatus}})
  
  }



