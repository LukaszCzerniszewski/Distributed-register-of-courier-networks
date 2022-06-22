package main



import (

	"encoding/json"

	"fmt"

	"io/ioutil"

	"log"

	"net/http"

	"os"

	"path/filepath"

	"strings"



	"github.com/gorilla/mux"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"

	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"

)



// Struktury potrzebne do utworzenia obiektu z JSONa

type User struct {

	ID         string      `json:"ID"`

	Attributes []ParcelType `json:"ParcelType"`

}



type attribute struct {

	Name  string `json:"name"`

	Value string `json:"value"`

}



type UserID struct {

	ID string `json:"ID"`

}



type ParcelID struct {

	ID string `json:"ID"`

}



type Kurier struct {

	IDPaczki string `json:"ID"`

	IDKuriera string `json:"ID"`

}



type ParcelType struct {

	ID           string `json:"id"`

	Destination  string `json:"destination"`

	Product_List string `json:"product_list"`

	Consignor    string `json:"consignor"`

	Localization string `json:"localization"`

	Track        string `json:"Track"`

}



// Zmienna globalna do przzechowywania uchwytu do smart contract

var contract_aa *gateway.Contract



// Rejestracja RegisterParcel

func RegisterParcel(w http.ResponseWriter, r *http.Request) {

	var ParcelObject ParcelType


	// odczyt danych wejsciowych

	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Printf("BODY: %v", string(reqBody))

	// Parsowanie JSON do obiektu

	err := json.Unmarshal(reqBody, &ParcelObject)

	if err != nil {

		log.Printf("ERROR: %v", err.Error())

	}
	//print("ParcelObject = ", string(ParcelObject.ID))
	// Parsowanie obiektu do JSON

	//attr_JSON, err := json.Marshal(ParcelObject)

	if err != nil {

		log.Printf("11 failed to create JSON: %v", err)

	}



	//UserKey := "1234567890"

	log.Printf("%v", ParcelObject)

	// Wyslanie danych do rejestru, RegisterUser to nazwa funkcji, dalej sa parametry
	
	result, err := contract_aa.SubmitTransaction("RegisterParcel",  ParcelObject.Destination, ParcelObject.Product_List,ParcelObject.Consignor)

	if strings.Contains(string(result), "Error") {

		log.Printf("0 Failed to Submit transaction: %v", string(result))

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(string(result))

	}

	if err != nil {

		log.Printf("1 Failed to Submit transaction: %v, %v", err.Error(), result)

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())

	}

	// Wyslanie odpowiedzi, jezeli wszystko zakonczylo sie sukcesem

	w.WriteHeader(200)

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	enc.Encode(string(result))

}






func GetTrace(w http.ResponseWriter, r *http.Request) {

	var ID string


	// odczyt danych wejsciowych

	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Printf("BODY: %v", string(reqBody))

	// Parsowanie JSON do obiektu

	err := json.Unmarshal(reqBody, &ID)

	if err != nil {

		log.Printf("ERROR: %v", err.Error())

	}
	
	//print("ParcelObject = ", string(ParcelObject.ID))
	// Parsowanie obiektu do JSON

	//attr_JSON, err := json.Marshal(ParcelObject)

	if err != nil {

		log.Printf("11 failed to create JSON: %v", err)

	}



	//UserKey := "1234567890"

	log.Printf("%v", ID)
	

	// Wyslanie danych do rejestru, RegisterUser to nazwa funkcji, dalej sa parametry
	
	result, err := contract_aa.SubmitTransaction("GetTrace",  string(ID))

	if strings.Contains(string(result), "Error") {

		log.Printf("0 Failed to Submit transaction: %v", string(result))

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(string(result))

	}

	if err != nil {

		log.Printf("1 Failed to Submit transaction: %v, %v", err.Error(), result)

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())

	}

	// Wyslanie odpowiedzi, jezeli wszystko zakonczylo sie sukcesem

	w.WriteHeader(200)

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	enc.Encode(string(result))

}



func GetParcel(w http.ResponseWriter, r *http.Request) {

	var ID string


	// odczyt danych wejsciowych

	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Printf("BODY: %v", string(reqBody))

	// Parsowanie JSON do obiektu

	err := json.Unmarshal(reqBody, &ID)

	if err != nil {

		log.Printf("ERROR: %v", err.Error())

	}
	
	//print("ParcelObject = ", string(ParcelObject.ID))
	// Parsowanie obiektu do JSON

	//attr_JSON, err := json.Marshal(ParcelObject)

	if err != nil {

		log.Printf("11 failed to create JSON: %v", err)

	}



	//UserKey := "1234567890"

	log.Printf("%v", ID)
	

	// Wyslanie danych do rejestru, RegisterUser to nazwa funkcji, dalej sa parametry
	
	result, err := contract_aa.SubmitTransaction("GetParcel",  string(ID))

	if strings.Contains(string(result), "Error") {

		log.Printf("0 Failed to Submit transaction: %v", string(result))

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(string(result))

	}

	if err != nil {

		log.Printf("1 Failed to Submit transaction: %v, %v", err.Error(), result)

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())

	}

	// Wyslanie odpowiedzi, jezeli wszystko zakonczylo sie sukcesem

	w.WriteHeader(200)

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	enc.Encode(string(result))

}


func SortingO1(w http.ResponseWriter, r *http.Request) {

	var ID string


	// odczyt danych wejsciowych

	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Printf("BODY: %v", string(reqBody))

	// Parsowanie JSON do obiektu

	err := json.Unmarshal(reqBody, &ID)

	if err != nil {

		log.Printf("ERROR: %v", err.Error())

	}
	
	//print("ParcelObject = ", string(ParcelObject.ID))
	// Parsowanie obiektu do JSON

	//attr_JSON, err := json.Marshal(ParcelObject)

	if err != nil {

		log.Printf("11 failed to create JSON: %v", err)

	}



	//UserKey := "1234567890"

	log.Printf("%v", ID)
	

	// Wyslanie danych do rejestru, RegisterUser to nazwa funkcji, dalej sa parametry
	
	result, err := contract_aa.SubmitTransaction("SortingO1",  string(ID))

	

	if strings.Contains(string(result), "Error") {

		log.Printf("0 Failed to Submit transaction: %v", string(result))

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(string(result))

	}

	if err != nil {

		log.Printf("1 Failed to Submit transaction: %v, %v", err.Error(), result)

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())

	}

	// Wyslanie odpowiedzi, jezeli wszystko zakonczylo sie sukcesem

	w.WriteHeader(200)

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	enc.Encode(string(result))


}

func SortingO2(w http.ResponseWriter, r *http.Request) {

	var ID string


	// odczyt danych wejsciowych

	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Printf("BODY: %v", string(reqBody))

	// Parsowanie JSON do obiektu

	err := json.Unmarshal(reqBody, &ID)

	if err != nil {

		log.Printf("ERROR: %v", err.Error())

	}
	
	//print("ParcelObject = ", string(ParcelObject.ID))
	// Parsowanie obiektu do JSON

	//attr_JSON, err := json.Marshal(ParcelObject)

	if err != nil {

		log.Printf("11 failed to create JSON: %v", err)

	}

	//UserKey := "1234567890"

	log.Printf("%v", ID)
	

	// Wyslanie danych do rejestru, RegisterUser to nazwa funkcji, dalej sa parametry
	
	result, err := contract_aa.SubmitTransaction("SortingO2",  string(ID))

	if strings.Contains(string(result), "Error") {

		log.Printf("0 Failed to Submit transaction: %v", string(result))

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(string(result))

	}

	if err != nil {

		log.Printf("1 Failed to Submit transaction: %v, %v", err.Error(), result)

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())

	}

	// Wyslanie odpowiedzi, jezeli wszystko zakonczylo sie sukcesem

	w.WriteHeader(200)

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	enc.Encode(string(result))


}


func BranchO1(w http.ResponseWriter, r *http.Request) {

	var ID string


	// odczyt danych wejsciowych

	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Printf("BODY: %v", string(reqBody))

	// Parsowanie JSON do obiektu

	err := json.Unmarshal(reqBody, &ID)

	if err != nil {

		log.Printf("ERROR: %v", err.Error())

	}
	
	//print("ParcelObject = ", string(ParcelObject.ID))
	// Parsowanie obiektu do JSON

	//attr_JSON, err := json.Marshal(ParcelObject)

	if err != nil {

		log.Printf("11 failed to create JSON: %v", err)

	}

	//UserKey := "1234567890"

	log.Printf("%v", ID)
	

	// Wyslanie danych do rejestru, RegisterUser to nazwa funkcji, dalej sa parametry
	
	result, err := contract_aa.SubmitTransaction("BranchO1",  string(ID))

	if strings.Contains(string(result), "Error") {

		log.Printf("0 Failed to Submit transaction: %v", string(result))

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(string(result))

	}

	if err != nil {

		log.Printf("1 Failed to Submit transaction: %v, %v", err.Error(), result)

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())

	}

	// Wyslanie odpowiedzi, jezeli wszystko zakonczylo sie sukcesem

	w.WriteHeader(200)

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	enc.Encode(string(result))


}

func BranchO2(w http.ResponseWriter, r *http.Request) {

	var ID string


	// odczyt danych wejsciowych

	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Printf("BODY: %v", string(reqBody))

	// Parsowanie JSON do obiektu

	err := json.Unmarshal(reqBody, &ID)

	if err != nil {

		log.Printf("ERROR: %v", err.Error())

	}
	
	//print("ParcelObject = ", string(ParcelObject.ID))
	// Parsowanie obiektu do JSON

	//attr_JSON, err := json.Marshal(ParcelObject)

	if err != nil {

		log.Printf("11 failed to create JSON: %v", err)

	}

	//UserKey := "1234567890"

	log.Printf("%v", ID)
	

	// Wyslanie danych do rejestru, RegisterUser to nazwa funkcji, dalej sa parametry
	
	result, err := contract_aa.SubmitTransaction("BranchO2",  string(ID))

	if strings.Contains(string(result), "Error") {

		log.Printf("0 Failed to Submit transaction: %v", string(result))

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(string(result))

	}

	if err != nil {

		log.Printf("1 Failed to Submit transaction: %v, %v", err.Error(), result)

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())

	}

	// Wyslanie odpowiedzi, jezeli wszystko zakonczylo sie sukcesem

	w.WriteHeader(200)

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	enc.Encode(string(result))


}


func GiveToCourier(w http.ResponseWriter, r *http.Request) {

	var ParcelObject ParcelType


	// odczyt danych wejsciowych

	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Printf("BODY: %v", string(reqBody))

	// Parsowanie JSON do obiektu

	err := json.Unmarshal(reqBody, &ParcelObject)

	if err != nil {

		log.Printf("ERROR: %v", err.Error())

	}
	//print("ParcelObject = ", string(ParcelObject.ID))
	// Parsowanie obiektu do JSON

	//attr_JSON, err := json.Marshal(ParcelObject)

	if err != nil {

		log.Printf("11 failed to create JSON: %v", err)

	}



	//UserKey := "1234567890"

	log.Printf("%v", ParcelObject)

	// Wyslanie danych do rejestru, RegisterUser to nazwa funkcji, dalej sa parametry
	
	result, err := contract_aa.SubmitTransaction("GiveToCourier",  ParcelObject.ID, ParcelObject.Destination)

	if strings.Contains(string(result), "Error") {

		log.Printf("0 Failed to Submit transaction: %v", string(result))

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(string(result))

	}

	if err != nil {

		log.Printf("1 Failed to Submit transaction: %v, %v", err.Error(), result)

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())

	}

	// Wyslanie odpowiedzi, jezeli wszystko zakonczylo sie sukcesem

	w.WriteHeader(200)

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	enc.Encode(string(result))

}


func Delivered(w http.ResponseWriter, r *http.Request) {

	var ID string


	// odczyt danych wejsciowych

	reqBody, _ := ioutil.ReadAll(r.Body)

	log.Printf("BODY: %v", string(reqBody))

	// Parsowanie JSON do obiektu

	err := json.Unmarshal(reqBody, &ID)

	if err != nil {

		log.Printf("ERROR: %v", err.Error())

	}
	
	//print("ParcelObject = ", string(ParcelObject.ID))
	// Parsowanie obiektu do JSON

	//attr_JSON, err := json.Marshal(ParcelObject)

	if err != nil {

		log.Printf("11 failed to create JSON: %v", err)

	}

	//UserKey := "1234567890"

	log.Printf("%v", ID)
	

	// Wyslanie danych do rejestru, RegisterUser to nazwa funkcji, dalej sa parametry
	
	result, err := contract_aa.SubmitTransaction("Delivered",  string(ID))

	if strings.Contains(string(result), "Error") {

		log.Printf("0 Failed to Submit transaction: %v", string(result))

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(string(result))

	}

	if err != nil {

		log.Printf("1 Failed to Submit transaction: %v, %v", err.Error(), result)

		http.ResponseWriter.WriteHeader(w, http.StatusInternalServerError)

		json.NewEncoder(w).Encode(err.Error())

	}

	// Wyslanie odpowiedzi, jezeli wszystko zakonczylo sie sukcesem

	w.WriteHeader(200)

	enc := json.NewEncoder(w)

	enc.SetEscapeHTML(false)

	enc.Encode(string(result))


}

// Utworzenie serwera REST oraz zdefiniowanie jego funkcji

func handleRequests() {

	// Create router

	log.Printf("Start Listen on :8080")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/RegisterParcel", RegisterParcel).Methods("POST")

	//myRouter.HandleFunc("/GetUser", GetUser).Methods("POST")

	myRouter.HandleFunc("/GetTrace", GetTrace).Methods("POST")

	myRouter.HandleFunc("/GetParcel", GetParcel).Methods("POST")
	
	myRouter.HandleFunc("/SortingO1", SortingO1).Methods("POST")

	myRouter.HandleFunc("/SortingO2", SortingO2).Methods("POST")
	
	myRouter.HandleFunc("/BranchO1", BranchO1).Methods("POST")
	
	myRouter.HandleFunc("/BranchO2", BranchO2).Methods("POST")
	
	myRouter.HandleFunc("/GiveToCourier", GiveToCourier).Methods("POST")

	myRouter.HandleFunc("/Delivered", Delivered).Methods("POST")

	

	log.Printf("%v", http.ListenAndServe(":8080", myRouter))



}



// Main function

func main() {

	//print("Test Main 1 \n")

	// Utworzenie wallet, jezeli nie istnieje

	log.Println("============ Hyperledger Fabric v2.2 TEST APP ============")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")

	if err != nil {

		log.Fatalf("Error setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)

	}

	//print("Test Main 2 \n")

	wallet, err := gateway.NewFileSystemWallet("wallet")

	if err != nil {

		log.Fatalf("4 Failed to create wallet: %v", err)

	}

	

	//print("Test Main 3 \n")



	if !wallet.Exists("Org1MSP") {

		err = populateWallet(wallet)

		if err != nil {

			log.Fatalf("5 Failed to populate wallet contents: %v", err)

		}

	}

	

	//print("Test Main 4 \n")



	// Plik YAML z konfiguracja sieci

	ccpPath := "/home/fabric/go/src/WAT/network/organizations/peerOrganizations/org1.wat.net/connection-org1.yaml"



	// Laczenie sie do HLF

	gw, err := gateway.Connect(

		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),

		gateway.WithIdentity(wallet, "Org1MSP"),

	)

	

	//print("Test Main 5 \n")

	if err != nil {

		log.Fatalf("6 Failed to connect to gateway: %v", err)

	}

	

	//print("Test Main 6 \n")

	defer gw.Close()

	

	//print("Test gw.Close 1 \n")

	// Wskazanie kanalu

	//!!!!!!!!!!!!!!!ZMiana !!!!!!!!!1

	// STare network_aa, err := gw.GetNetwork("attrauth")

	//NOwe

	

	//print("Test gw.Close 2 \n")

	

	network_aa, err := gw.GetNetwork("example")

	

	//print("Test gw.Close 3 \n")

	if err != nil {

		log.Fatalf("7 Failed to get network: %v", err)

	}

	//print("Test gw.Close 4 \n")

	// Wskazanie Smart Contract

	//Stare contract_aa = network_aa.GetContract("AttrAuth")

	//print("2 print \n")

	contract_aa = network_aa.GetContract("Example")

	//print(string("Wywolanie funkcji handleRequests() \n"))

	handleRequests()

}



// Tworzenie wallet

func populateWallet(wallet *gateway.Wallet) error {

	log.Println("============ Populating wallet ============")

	//print("Test populateWallet 1 \n")

	credPath := "/home/fabric/go/src/WAT/network/organizations/peerOrganizations/org1.wat.net/users/User1@org1.wat.net/msp"



	certPath := filepath.Join(credPath, "signcerts", "User1@org1.wat.net-cert.pem")

	// read the certificate pem

	//print("Test populateWallet 2 \n")

	cert, err := ioutil.ReadFile(filepath.Clean(certPath))

	print("cert = ", cert,"\n")

	if err != nil {

		return err

	}



	keyDir := filepath.Join(credPath, "keystore")

	files, err := ioutil.ReadDir(keyDir)

	//print("Test populateWallet 3 \n")

	if err != nil {

		return err

	}

	if len(files) != 1 {

		return fmt.Errorf("keystore folder should have contain one file")

	}

	//print("Test populateWallet 4 \n")

	keyPath := filepath.Join(keyDir, files[0].Name())

	key, err := ioutil.ReadFile(filepath.Clean(keyPath))

	if err != nil {

		return err

	}

	//identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	//print("Test populateWallet 6 \n")

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	//print("Test populateWallet 7 \n")

	return wallet.Put("Org1MSP", identity)

}
