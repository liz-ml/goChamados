package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// Chamado representa a estrutura do reporte
type Chamado struct {
	ID          int       `json:"id"`
	Titulo      string    `json:"titulo"`
	Descricao   string    `json:"descricao"`
	Status      string    `json:"status"` // ex: "Aberto", "Em Andamento", "Resolvido"
	CriadoEm    time.Time `json:"criado_em"`
}

// Simulando um banco de dados em memória por enquanto
var chamados = []Chamado{ // sim é possível criar um slice de um struct
	{
		ID:        1,
		Titulo:    "Erro de autenticação",
		Descricao: "Usuário não consegue logar via SSO",
		Status:    "Aberto",
		CriadoEm:  time.Now(),
	},
}

// Handler para listar todos os chamados (GET)
func getChamadosHandler(w http.ResponseWriter, r *http.Request) { //
	// Define o cabeçalho como JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Serializa o slice/lista de chamados para JSON e envia na resposta
	json.NewEncoder(w).Encode(chamados)} // Encode pega a variável chamados e converte p json
	//NewEnconder indica para onde deve ser enviado o JSON

// Handler para criar um novo chamado (POST)
func createChamadoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var novoChamado Chamado
	// Decodifica o JSON do corpo da requisição para a struct
	err := json.NewDecoder(r.Body).Decode(&novoChamado)
	if err != nil {
		http.Error(w, "JSON inválido", http.StatusBadRequest)
		return
	}

	novoChamado.ID = len(chamados) + 1
	novoChamado.CriadoEm = time.Now()
	novoChamado.Status = "Aberto"

	chamados = append(chamados, novoChamado)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(novoChamado)
}

func main() {
	// Rotas
	http.HandleFunc("/chamados", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getChamadosHandler(w, r)
		case http.MethodPost:
			createChamadoHandler(w, r)
		default:
			http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("🚀 Servidor Go rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}