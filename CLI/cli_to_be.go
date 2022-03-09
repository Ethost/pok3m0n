package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type labResponse struct {
	LabID     string `json:"labID"`
	StatusLab string `json:"statusLab"`
}

func LabToBE(labID string, action string) {
	var i labResponse

	client := resty.New()
	_, err := client.R().
		SetResult(&i).
		Get("http://127.0.0.1:8080/labs/" + labID + "/" + action)
		// Explore response object
	if err != nil {
		fmt.Println("Erreur:\n", err)
	}
	fmt.Println("Response BE:")
	fmt.Println("\t-> Lab ID:", i.LabID)
	fmt.Println("\t-> Status Lab:", i.StatusLab)

}

type userResponse struct {
	SignRequest string `json:"signRequest"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

func UserToBE(signRequest string, username string, password string) {
	var i userResponse

	client := resty.New()
	_, err := client.R().
		SetResult(&i).
		Get("http://127.0.0.1:8080/users/" + signRequest + "/" + username + "/" + password)

	if err != nil {
		fmt.Println("Erreur:\n", err)
	}

	fmt.Println("Response BE:")
	fmt.Println("\t-> signRequest:", i.SignRequest)
	fmt.Println("\t-> username:", i.Username)
	fmt.Println("\t-> password:", i.Password)
}
