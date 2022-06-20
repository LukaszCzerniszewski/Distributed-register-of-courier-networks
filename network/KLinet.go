package main



import (

	"bytes"

	"encoding/json"

	"fmt"

	"io"

	"log"

	"net/http"

	"strings"
	 
	

)



// Struktury do utworzenia obiektu JSON

type User struct {

	ID         string      `json:"ID"`

	Attributes []ParcelType `json:"ParcelType"`

}

type attribute struct {

	Name  string `json:"name"`

	Value string `json:"value"`

}





type ParcelType struct {

	ID           string `json:"id"`

	Destination  string `json:"destination"`

	Product_List string `json:"product_list"`

	Consignor    string `json:"consignor"`

	Localization string `json:"localization"`

	Track        string `json:"Track"`

}



func main() {

	// Tworzenie obiektu JSON

	var attrib []ParcelType

	attrib = append(attrib, ParcelType{"id","Destination","Product_list","Consignor","localization","Track"})



	UserObject := User{ID: "127@org1", Attributes: attrib}

	jsonValue, _ := json.Marshal(UserObject)

	log.Printf(string(jsonValue))



	// Wyslanie danych do serwera i obsluga odpowiedzi

	response, err := http.Post("http://127.0.0.1:8080/RegisterUser", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])

		log.Printf("%v", response.Status)

	}

	fmt.Println("Terminating the application...")

}
