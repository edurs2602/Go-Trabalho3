package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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

func printMenu() {
	fmt.Println("Selecione uma opção:")
	fmt.Println("1 - Listar Pets")
	fmt.Println("2 - Adicionar Pet")
	fmt.Println("3 - Atualizar Pet")
	fmt.Println("4 - Listar Clientes")
	fmt.Println("5 - Adicionar Cliente")
	fmt.Println("6 - Atualizar Cliente")
	fmt.Println("0 - Sair")
	fmt.Printf("Digite sua opção: ")
}

func userInput(input int) bool {
	switch input {
	case 1:
		listPets()
		return true
	case 2:
		addPet()
		return true
	case 3:
		updatePet()
		return true
	case 4:
		listClients()
		return true
	case 5:
		addClient()
		return true
	case 6:
		updateClient()
		return true
	case 0:
		fmt.Println("Saindo...")
		return false
	default:
		fmt.Println("Opção inválida!")
		return true
	}
}

func listPets() {
	resp, err := http.Get("http://localhost:8090/pets")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Erro ao obter a lista de pets. Código de status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&pets)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Lista de Pets:")
	for _, pet := range pets {
		fmt.Printf("Nome: %s, Dono: %s, Idade: %d, Peso: %.2f\n", pet.Name, pet.Owner, pet.Age, pet.Weight)
	}
}

func listClients() {
	resp, err := http.Get("http://localhost:8090/clients")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Erro ao obter a lista de clientes. Código de status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&clients)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Lista de Clientes:")
	for _, client := range clients {
		fmt.Printf("Nome: %s, Email: %s, Phone: %s, Adress: %s, Pets: %s\n", client.Name, client.Email, client.Phone, client.Address, client.Pets)
	}
}

func addPet() string {
	var pet Pet
	fmt.Println("Informe os detalhes do novo pet")
	fmt.Print("Nome: ")
	_, err := fmt.Scanln(&pet.Name)
	if err != nil {
		return ""
	}
	fmt.Print("Dono: ")
	_, err = fmt.Scanln(&pet.Owner)
	if err != nil {
		return ""
	}
	fmt.Print("Idade: ")
	_, err = fmt.Scanln(&pet.Age)
	if err != nil {
		return ""
	}
	fmt.Print("Peso: ")
	_, err = fmt.Scanln(&pet.Weight)
	if err != nil {
		return ""
	}

	petJSON, err := json.Marshal(pet)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8090/pets/add", "application/json", bytes.NewBuffer(petJSON))
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Novo pet adicionado com sucesso!")
	} else {
		fmt.Printf("Erro ao adicionar o pet. Código de status: %d\n", resp.StatusCode)
	}

	return pet.Name
}

func addClient() {
	var client Client
	var numPets int
	fmt.Println("Informe os detalhes do novo cliente")
	fmt.Print("Nome: ")
	_, err := fmt.Scanln(&client.Name)
	if err != nil {
		return
	}
	fmt.Print("Email: ")
	_, err = fmt.Scanln(&client.Email)
	if err != nil {
		return
	}
	fmt.Print("Phone: ")
	_, err = fmt.Scanln(&client.Phone)
	if err != nil {
		return
	}
	fmt.Print("Endereço: ")
	_, err = fmt.Scanln(&client.Address)
	if err != nil {
		return
	}
	fmt.Print("Quantidade de pets: ")
	_, err = fmt.Scanln(&numPets)
	if err != nil {
		return
	}

	for i := 0; i < numPets; i++ {
		name := addPet()
		client.Pets = append(client.Pets, name)
	}

	clientJSON, err := json.Marshal(client)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post("http://localhost:8090/clients/add", "application/json", bytes.NewBuffer(clientJSON))
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusCreated {
		fmt.Println("Novo cliente adicionado com sucesso!")
	} else {
		fmt.Printf("Erro ao adicionar o novo cliente. Código de status: %d\n", resp.StatusCode)
	}
}

func updatePet() {
	resp, err := http.Get("http://localhost:8090/pets")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Erro ao obter a lista de pets. Código de status: %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&pets)
	if err != nil {
		log.Fatal(err)
	}

	var name string
	fmt.Println("Digite o nome do pet que deseja atualizar")
	fmt.Print("Nome: ")
	if _, err := fmt.Scan(&name); err != nil {
		log.Fatal(err)
	}

	// Busca o pet pelo nome
	pet := findPetByName(name)
	if pet == nil {
		fmt.Printf("Pet '%s' não encontrado\n", name)
		return
	}

	var updatedPet Pet
	fmt.Println("Informe os novos detalhes do pet")
	fmt.Print("Nome: ")
	if _, err := fmt.Scan(&updatedPet.Name); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Dono: ")
	if _, err := fmt.Scan(&updatedPet.Owner); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Idade: ")
	if _, err := fmt.Scan(&updatedPet.Age); err != nil {
		log.Fatal(err)
	}
	fmt.Print("Peso: ")
	if _, err := fmt.Scan(&updatedPet.Weight); err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedPet)
	fmt.Println(pet)

	// Atualiza o pet no servidor
	petJSON, err := json.Marshal(updatedPet)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8090/pets/update/%s", pet.Name), bytes.NewBuffer(petJSON))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Pet atualizado com sucesso!")
	} else {
		fmt.Printf("Erro ao atualizar o pet. Código de status: %d\n", resp.StatusCode)
	}
}

func updateClient() {
	var name string
	fmt.Println("Digite o nome do cliente que deseja atualizar:")
	fmt.Print("Nome: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Println(err)
	}

	// Busca o cliente pelo nome
	pet := findClientByName(name)
	if pet == nil {
		fmt.Printf("Cliente '%s' não encontrado\n", name)
		return
	} else {
		fmt.Println("Cliente encontrado, prosseguindo com a atualização...")
	}

	var updatedClient Client
	var numPets int
	fmt.Println("Informe os detalhes do novo cliente")
	fmt.Print("Nome: ")
	_, err = fmt.Scanln(&updatedClient.Name)
	if err != nil {
		return
	}
	fmt.Print("Email: ")
	_, err = fmt.Scanln(&updatedClient.Email)
	if err != nil {
		return
	}
	fmt.Print("Phone: ")
	_, err = fmt.Scanln(&updatedClient.Phone)
	if err != nil {
		return
	}
	fmt.Print("Endereço: ")
	_, err = fmt.Scanln(&updatedClient.Address)
	if err != nil {
		return
	}
	fmt.Print("Quantidade de pets: ")
	_, err = fmt.Scanln(&numPets)
	if err != nil {
		return
	}

	for i := 0; i < numPets; i++ {
		name := addPet()
		updatedClient.Pets = append(updatedClient.Pets, name)
	}

	// Atualiza o cliente no servidor
	clientJSON, err := json.Marshal(updatedClient)
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("http://localhost:8090/client/update/%s", name), bytes.NewBuffer(clientJSON))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode == http.StatusOK {
		fmt.Println("Cliente atualizado com sucesso!")
	} else {
		fmt.Printf("Erro ao atualizar o cliente. Código de status: %d\n", resp.StatusCode)
	}
}

func findPetByName(name string) *Pet {
	for i := range pets {
		if pets[i].Name == name {
			return &pets[i]
		}
	}
	return nil
}

func findClientByName(name string) *Client {
	for i := range clients {
		if clients[i].Name == name {
			return &clients[i]
		}
	}
	return nil
}

func main() {
	var input int

	printMenu()

	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(err)
	}

	if !userInput(input) {
		os.Exit(0)
	} else {
		main()
	}
}
