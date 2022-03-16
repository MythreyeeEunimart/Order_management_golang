package config

import (
	"fmt"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connectdb() {
	err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI("mongodb+srv://root:root@cluster0.h2ivf.mongodb.net/Orders_golang?retryWrites=true&w=majority"))
	fmt.Println("connected to the database............")
	fmt.Println(err)
 }