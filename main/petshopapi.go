package main

import (
	_ "Petshop/docs"
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi"
	httpSwagger "github.com/swaggo/http-swagger"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Pet struct {
	Name   string  `json:"name"`
	Owner  string  `json:"owner"`
	Age    int     `json:"age"`
	Weight float64 `json:"weight"`
}

var pets []Pet

type Client struct {
	Name    string   `json:"name"`
	Email   string   `json:"email"`
	Phone   string   `json:"phone"`
	Address string   `json:"address"`
	Pets    []string `json:"pets"`
}

var clients []Client

type DataStore struct {
	Pets    []Pet
	Clients []Client
}

var dataStore DataStore
var dbFileName = "database.json"

func loadDataFromDB() {
	data, err := ioutil.ReadFile(dbFileName)
	if err != nil {
		log.Println("Database file not found. Creating a new one.")
		dataStore.Pets = []Pet{}
		dataStore.Clients = []Client{}
		return
	}

	// Verifica se o arquivo não está vazio antes de decodificar
	if len(data) == 0 {
		dataStore.Pets = []Pet{}
		dataStore.Clients = []Client{}
		return
	}

	err = json.Unmarshal(data, &dataStore)
	if err != nil {
		log.Fatal(err)
	}
}

func saveDataToDB() {
	data, err := json.Marshal(dataStore)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(dbFileName, data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	loadDataFromDB()

	if len(dataStore.Pets) == 0 {
		// Adiciona dados iniciais apenas se o arquivo estiver vazio
		pets = []Pet{
			{"Fofão", "Pablo", 2, 12.5},
			{"Safira", "Tereza", 4, 8.2},
			{"Xulipa", "Luis", 1, 5.7},
		}

		clients = []Client{
			{"Pablo Santos", "pablo.santos@mail.com", "123456789", "IMD/UFRN", []string{"Fofão"}},
			{"Tereza Aparecida", "tereza_aparecida@example.com", "987654321", "DIMAP/UFRN", []string{"Safira"}},
			{"Luis Eduardo", "luis-eduard@domain.com", "987654321", "NPITI/UFRN", []string{"Xulipa"}},
		}

		dataStore.Pets = pets
		dataStore.Clients = clients
		saveDataToDB()
	}
}

// GetPetsHandler retorna a lista de pets.
// @Summary Lista os pets.
// @Description Retorna a lista de pets.
// @Tags Pets
// @Produce json
// @Success 200 {array} Pet
// @Router /pets [get]
func getPets(res http.ResponseWriter, req *http.Request) {
	if resBody, err := json.Marshal(dataStore.Pets); err != nil {
		log.Fatal(err)
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		_, err := res.Write(resBody)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// AddPetHandler adiciona um novo pet.
// @Summary Adiciona um novo pet.
// @Description Adiciona um novo pet.
// @Tags Pets
// @Accept json
// @Param pet body Pet true "Novo pet a ser adicionado"
// @Success 201 "Pet adicionado com sucesso"
// @Router /pets/add [post]
func addPet(res http.ResponseWriter, req *http.Request) {
	contents, _ := io.ReadAll(req.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(req.Body)

	var p Pet
	err := json.Unmarshal(contents, &p)
	if err != nil {
		log.Fatal(err)
	} else {
		dataStore.Pets = append(dataStore.Pets, p)
		saveDataToDB()
		res.WriteHeader(http.StatusCreated)
	}
}

// UpdatePetHandler atualiza os detalhes de um pet existente.
// @Summary Atualiza os detalhes de um pet.
// @Description Atualiza os detalhes de um pet.
// @Tags Pets
// @Accept json
// @Param name path string true "Nome do pet a ser atualizado"
// @Param pet body Pet true "Detalhes atualizados do pet"
// @Success 200 "Pet atualizado com sucesso"
// @Router /pets/update/{name} [put]
func updatePet(res http.ResponseWriter, req *http.Request) {
	urlPathElements := strings.Split(req.URL.Path, "/")
	name := urlPathElements[3] // Update URL path element index corrected
	var updatedPet Pet
	contents, _ := io.ReadAll(req.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(req.Body)
	err := json.Unmarshal(contents, &updatedPet)
	if err != nil {
		log.Fatal(err)
	}

	for i, p := range dataStore.Pets {
		if p.Name == name {
			dataStore.Pets[i] = updatedPet
			saveDataToDB() // Saving after update
			break
		}
	}

	res.WriteHeader(http.StatusOK)
}

// GetClientsHandler retorna a lista de clientes.
// @Summary Lista os clientes.
// @Description Retorna a lista de clientes.
// @Tags Clients
// @Produce json
// @Success 200 {array} Client
// @Router /clients [get]
func getClients(res http.ResponseWriter, req *http.Request) {
	if resBody, err := json.Marshal(dataStore.Clients); err != nil {
		log.Fatal(err)
	} else {
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		_, err := res.Write(resBody)
		if err != nil {
			return
		}
	}
}

// AddClientHandler adiciona um novo cliente.
// @Summary Adiciona um novo cliente.
// @Description Adiciona um novo cliente.
// @Tags Clients
// @Accept json
// @Param client body Client true "Novo cliente a ser adicionado"
// @Success 201 "Cliente adicionado com sucesso"
// @Router /clients/add [post]
func addClient(res http.ResponseWriter, req *http.Request) {
	contents, _ := io.ReadAll(req.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(req.Body)

	var c Client
	err := json.Unmarshal(contents, &c)
	if err != nil {
		log.Fatal(err)
	} else {
		dataStore.Clients = append(dataStore.Clients, c)
		saveDataToDB()
		res.WriteHeader(http.StatusCreated)
	}
}

// UpdateClientHandler atualiza os detalhes de um cliente existente.
// @Summary Atualiza os detalhes de um cliente.
// @Description Atualiza os detalhes de um cliente.
// @Tags Clients
// @Accept json
// @Param name path string true "Nome do cliente a ser atualizado"
// @Param client body Client true "Detalhes atualizados do cliente"
// @Success 200 "Cliente atualizado com sucesso"
// @Router /clients/update/{name} [put]
func updateClient(res http.ResponseWriter, req *http.Request) {
	urlPathElements := strings.Split(req.URL.Path, "/")
	name := urlPathElements[3] // Update URL path element index corrected
	var updatedClient Client
	contents, _ := io.ReadAll(req.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(req.Body)
	err := json.Unmarshal(contents, &updatedClient)
	if err != nil {
		log.Fatal(err)
	}

	for i, c := range dataStore.Clients {
		if c.Name == name {
			dataStore.Clients[i] = updatedClient
			saveDataToDB() // Saving after update
			break
		}
	}

	res.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("A API ESTÁ RODANDO NA PORTA 8090")
	loadDataFromDB()
	r := chi.NewRouter()

	r.Get("/pets", getPets)
	r.Post("/pets/add", addPet)
	r.Put("/pets/update/", updatePet)

	r.Get("/clients", getClients)
	r.Post("/clients/add", addClient)
	r.Put("/clients/update/", updateClient)

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8090/swagger/doc.json"), //The url pointing to API definition
	))

	server := http.Server{
		Addr:         "localhost:8090",
		Handler:      r,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
