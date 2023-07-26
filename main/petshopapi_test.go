package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPets(t *testing.T) {
	loadDataFromDB()

	req, err := http.NewRequest("GET", "/pets", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getPets)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v; got %v", http.StatusOK, rr.Code)
	}

	// Verifica se a resposta contém os dados corretos dos pets
	var petsResponse []Pet
	if err := json.Unmarshal(rr.Body.Bytes(), &petsResponse); err != nil {
		t.Fatal(err)
	}

	// Verifica se a quantidade de pets retornada é igual à quantidade de pets esperada
	expectedNumPets := len(dataStore.Pets)
	if len(petsResponse) != expectedNumPets {
		t.Errorf("expected %v pets; got %v", expectedNumPets, len(petsResponse))
	}
}

func TestAddPet(t *testing.T) {
	loadDataFromDB()

	// Cria um novo pet para adicionar
	newPet := Pet{
		Name:   "Biro",
		Owner:  "Paulo",
		Age:    3,
		Weight: 9.7,
	}

	// Converte o novo pet para JSON
	petJSON, err := json.Marshal(newPet)
	if err != nil {
		t.Fatal(err)
	}

	// Cria uma requisição POST para adicionar o pet
	req, err := http.NewRequest("POST", "/pets/add", bytes.NewBuffer(petJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addPet)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status code %v; got %v", http.StatusCreated, rr.Code)
	}

	// Verifica se o pet foi adicionado corretamente
	loadDataFromDB()
	found := false
	for _, pet := range dataStore.Pets {
		if pet.Name == newPet.Name {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("added pet not found in the list of pets")
	}
}

func TestUpdatePet(t *testing.T) {
	loadDataFromDB()

	// Cria um pet existente com dados atualizados
	updatedPet := Pet{
		Name:   "Fofão",
		Owner:  "Marcelo",
		Age:    3,
		Weight: 15.2,
	}

	// Converte o pet atualizado para JSON
	petJSON, err := json.Marshal(updatedPet)
	if err != nil {
		t.Fatal(err)
	}

	// Cria uma requisição POST para atualizar o pet com o nome correto na URL
	req, err := http.NewRequest("POST", "/pets/update/Fofão", bytes.NewBuffer(petJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updatePet)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v; got %v", http.StatusOK, rr.Code)
	}

	// Verifica se o pet foi atualizado corretamente
	loadDataFromDB()
	found := false
	for _, pet := range dataStore.Pets {
		if pet.Name == updatedPet.Name {
			// Verifica se os dados foram atualizados corretamente
			if pet.Owner != updatedPet.Owner {
				t.Errorf("expected owner %v; got %v", updatedPet.Owner, pet.Owner)
			}
			if pet.Age != updatedPet.Age {
				t.Errorf("expected age %v; got %v", updatedPet.Age, pet.Age)
			}
			if pet.Weight != updatedPet.Weight {
				t.Errorf("expected weight %v; got %v", updatedPet.Weight, pet.Weight)
			}
			found = true
			break
		}
	}

	if !found {
		t.Errorf("updated pet not found in the list of pets")
	}
}

func TestGetClients(t *testing.T) {
	loadDataFromDB()

	req, err := http.NewRequest("GET", "/clients", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getClients)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v; got %v", http.StatusOK, rr.Code)
	}

	// Verifica se a resposta contém os dados corretos dos clientes
	var clientsResponse []Client
	if err := json.Unmarshal(rr.Body.Bytes(), &clientsResponse); err != nil {
		t.Fatal(err)
	}

	// Verifica se a quantidade de clientes retornada é igual à quantidade de clientes esperada
	expectedNumClients := len(dataStore.Clients)
	if len(clientsResponse) != expectedNumClients {
		t.Errorf("expected %v clients; got %v", expectedNumClients, len(clientsResponse))
	}
}

func TestAddClient(t *testing.T) {
	loadDataFromDB()

	// Cria um novo cliente para adicionar
	newClient := Client{
		Name:    "Paulo Silva",
		Email:   "paulo.silva@example.com",
		Phone:   "9876543210",
		Address: "ICE/UFRN",
		Pets:    []string{"Biro"},
	}

	// Converte o novo cliente para JSON
	clientJSON, err := json.Marshal(newClient)
	if err != nil {
		t.Fatal(err)
	}

	// Cria uma requisição POST para adicionar o cliente
	req, err := http.NewRequest("POST", "/clients/add", bytes.NewBuffer(clientJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(addClient)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status code %v; got %v", http.StatusCreated, rr.Code)
	}

	// Verifica se o cliente foi adicionado corretamente
	loadDataFromDB()
	found := false
	for _, client := range dataStore.Clients {
		if client.Name == newClient.Name {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("added client not found in the list of clients")
	}
}

func TestUpdateClient(t *testing.T) {
	loadDataFromDB()

	// Cria um cliente existente com dados atualizados
	updatedClient := Client{
		Name:    "Luis Eduardo",
		Email:   "luiseduardo@mail.com",
		Phone:   "84993212183",
		Address: "IMD/UFRN",
		Pets:    []string{"Xulipa", "Ralf"},
	}

	// Converte o cliente atualizado para JSON
	clientJSON, err := json.Marshal(updatedClient)
	if err != nil {
		t.Fatal(err)
	}

	// Cria uma requisição POST para atualizar o cliente com o nome correto na URL
	req, err := http.NewRequest("POST", "/clients/update/Luis Eduardo", bytes.NewBuffer(clientJSON))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(updateClient)

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status code %v; got %v", http.StatusOK, rr.Code)
	}

	// Verifica se o cliente foi atualizado corretamente
	loadDataFromDB()
	found := false
	for _, client := range dataStore.Clients {
		if client.Name == updatedClient.Name {
			// Verifica se os dados foram atualizados corretamente
			if client.Email != updatedClient.Email {
				t.Errorf("expected email %v; got %v", updatedClient.Email, client.Email)
			}
			if client.Phone != updatedClient.Phone {
				t.Errorf("expected phone %v; got %v", updatedClient.Phone, client.Phone)
			}
			if client.Address != updatedClient.Address {
				t.Errorf("expected address %v; got %v", updatedClient.Address, client.Address)
			}
			found = true
			break
		}
	}

	if !found {
		t.Errorf("updated client not found in the list of clients")
	}
}
