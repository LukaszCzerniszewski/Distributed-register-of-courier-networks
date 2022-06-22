package main



import (

	"bytes"

	"encoding/json"

	"fmt"

	"io"

	"log"

	"net/http"

	"strings"
	 
	"regexp"

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


type ParcelID struct {

	ID string `json:"ID"`

}

type Kurier struct {

	IDPaczki string `json:"ID"`

	IDKuriera string `json:"ID"`

}


func main() {

	
	// ======> 1 Krok towrzymy nowa przesylke <===================
	var numerPrzesylki = RegisterParcel("Destination" , "Product_List" , "Consignor")
	print("Nowo utworzona  przesylka ma numer = ",numerPrzesylki,"\n")

	// ======> 2 Krok pobranie obecnej trasy przesylki  <===================
	
	var trasaPrzesylki = GetTrace(numerPrzesylki)
	print("Przesyka obecnie pokonala trase = ",trasaPrzesylki,"\n")
	

	// ======> 3 Krok pobranie danych przesylki  <===================
	
	var lokalizacjaPrzesylki = GetParcel(numerPrzesylki)
	print("Przesylka = ",lokalizacjaPrzesylki,"\n")
	
	// ======> 4 Krok proba wyslania przesylki do sortowni firmy Org2MSP  <===================
	
	var wynikSortownia1  = SortingO1(numerPrzesylki)
	print("Proba przeniesienia paczki do sortowni Org2MSP ", wynikSortownia1 ,"\n")

	//======> 5 Krok proba wyslania przesylki do sortowni firmy Org1MSP  <===================
	
	var wynikSortownia2  = SortingO2(numerPrzesylki)
	print("Proba przeniesienia paczki do sortowni Org1MSP ", wynikSortownia2 ,"\n")
	


	// ======> 6 Krok proba wyslania przesylki do sortowni firmy Org2MSP  <===================
	
	var wynikBranch1  = BranchO2(numerPrzesylki)
	print("Proba przeniesienia paczki do odzialu Org2MSP ", wynikBranch1 ,"\n")

	//======> 7 Krok proba wyslania przesylki do sortowni firmy Org1MSP  <===================
	
	var wynikBranch2  = BranchO1(numerPrzesylki)
	print("Proba przeniesienia paczki do odzialu  Org1MSP ", wynikBranch2 ,"\n")
	
	//======> 8 Krok przekazanie paczki do kuriera numer 1  <===================
	
	var numerKuriera = "1"
	var wynikKurier  = GiveToCourier(numerPrzesylki, numerKuriera)
	print("Proba dostarczenia przesylki ", wynikKurier ,"\n")

	
	// ======> 9 Krok dostarczenie przesylki  <===================
	
	var dostarczeniePrzesylki = Delivered(numerPrzesylki)
	print("Przesyka dotarla do celu = ",dostarczeniePrzesylki,"\n")
		
	
		

	// ======> 10 Krok pobranie obecnej trasy przesylki  <===================
	
	trasaPrzesylki = GetTrace(numerPrzesylki)
	print("Przesyka obecnie pokonala trase = ",trasaPrzesylki,"\n")
	

	fmt.Println("Terminating the application...")

}


func Delivered(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")

	jsonValue, _ := json.Marshal(numerPrzesylki)

	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/Delivered", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("Blad")
}

func GiveToCourier(numerPrzesylki string, numerKuriera string) string{
	
	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")
	//UserObject := Kurier{numerPrzesylki, numerKuriera}
	//UserObject := Kurier{"fwfawf", "wafawf"}
	UserObject := ParcelType{numerPrzesylki,numerKuriera,"null"  ,"null","null","null"}
	jsonValue, _ := json.Marshal(UserObject)

	log.Printf(string(jsonValue))

	//print(UserObject.IDKuriera,"      ", UserObject.IDPaczki, "Koniec \n")
	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/GiveToCourier", "application/json", bytes.NewBuffer(jsonValue))



	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("NULL")
}


func BranchO1(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")

	jsonValue, _ := json.Marshal(numerPrzesylki)

	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/BranchO1", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("Blad")
}
func BranchO2(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")

	jsonValue, _ := json.Marshal(numerPrzesylki)

	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/BranchO2", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("Blad")
}


func SortingO2(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")

	jsonValue, _ := json.Marshal(numerPrzesylki)

	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/SortingO2", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("Blad")
}

func SortingO1(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")

	jsonValue, _ := json.Marshal(numerPrzesylki)

	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/SortingO1", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("Blad")
}



func GetTrace(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")


	
	jsonValue, _ := json.Marshal(numerPrzesylki)




	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/GetTrace", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("NULL")
}

func RegisterParcel(Destination string, Product_List string, Consignor string) string{
	//var attrib []ParcelType

	//attrib = append(attrib, ParcelType{"id",string(Destination),string(Product_List),string(Consignor),"localization","Track"})

	UserObject := ParcelType{"id",Destination, Product_List ,Consignor,"localization","Track"}

	jsonValue, _ := json.Marshal(UserObject)

	log.Printf(string(jsonValue))


	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/RegisterParcel", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("NULL")
}

func GetParcel(numerPrzesylki string) string{
	//hx := hex.EncodeToString([]byte(numerPrzesylki))

	re, err := regexp.Compile(`[^\w]`)
	if err != nil {
		log.Fatal(err)
	}
	numerPrzesylki = re.ReplaceAllString(numerPrzesylki, "")


	
	jsonValue, _ := json.Marshal(numerPrzesylki)




	// Wyslanie danych do serwera i obsluga odpowiedzi -- Rjestracja przesyki

	response, err := http.Post("http://127.0.0.1:8080/GetParcel", "application/json", bytes.NewBuffer(jsonValue))

	if err != nil {

		fmt.Printf("The HTTP request failed with error %s\n", err)

	} else {

		data, _ := io.ReadAll(response.Body)

		body := strings.Split(string(data), "\n")

		fmt.Println(body[0])
		
		return body[0]
		
}
return string("NULL")
}