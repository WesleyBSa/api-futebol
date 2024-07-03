package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/WesleyBSa/api-tabela/models"
)

func GetTabela(w http.ResponseWriter, r *http.Request) {
	tabela := []models.Tabela{
		{Posicao: 1, Time: "Flamengo", Pontos: 27, Jogos: 13, Vitorias: 8, Empates: 3, Derrotas: 2, GolsPro: 22, GolsContra: 12, SaldoGols: 10},
		{Posicao: 2, Time: "Palmeiras", Pontos: 26, Jogos: 13, Vitorias: 8, Empates: 2, Derrotas: 3, GolsPro: 18, GolsContra: 9, SaldoGols: 9},
		{Posicao: 3, Time: "Botafogo", Pontos: 24, Jogos: 13, Vitorias: 7, Empates: 3, Derrotas: 3, GolsPro: 21, GolsContra: 13, SaldoGols: 8},
		{Posicao: 4, Time: "Bahia", Pontos: 24, Jogos: 13, Vitorias: 7, Empates: 3, Derrotas: 3, GolsPro: 21, GolsContra: 16, SaldoGols: 5},
		{Posicao: 5, Time: "Athletico-PR", Pontos: 22, Jogos: 13, Vitorias: 6, Empates: 4, Derrotas: 3, GolsPro: 16, GolsContra: 10, SaldoGols: 6},
		{Posicao: 6, Time: "São Paulo", Pontos: 21, Jogos: 13, Vitorias: 6, Empates: 3, Derrotas: 4, GolsPro: 20, GolsContra: 15, SaldoGols: 5},
		{Posicao: 7, Time: "Cruzeiro", Pontos: 20, Jogos: 12, Vitorias: 6, Empates: 2, Derrotas: 4, GolsPro: 16, GolsContra: 16, SaldoGols: 0},
		{Posicao: 8, Time: "Fortaleza", Pontos: 20, Jogos: 12, Vitorias: 5, Empates: 5, Derrotas: 2, GolsPro: 13, GolsContra: 12, SaldoGols: 1},
		{Posicao: 9, Time: "Bragantino", Pontos: 19, Jogos: 13, Vitorias: 5, Empates: 4, Derrotas: 4, GolsPro: 17, GolsContra: 15, SaldoGols: 2},
		{Posicao: 10, Time: "Internacional", Pontos: 18, Jogos: 11, Vitorias: 5, Empates: 3, Derrotas: 3, GolsPro: 10, GolsContra: 8, SaldoGols: 2},
		{Posicao: 11, Time: "Atlético-MG", Pontos: 18, Jogos: 12, Vitorias: 4, Empates: 6, Derrotas: 2, GolsPro: 18, GolsContra: 16, SaldoGols: 2},
		{Posicao: 12, Time: "Juventude", Pontos: 16, Jogos: 12, Vitorias: 4, Empates: 4, Derrotas: 4, GolsPro: 15, GolsContra: 17, SaldoGols: -2},
		{Posicao: 13, Time: "Criciúma", Pontos: 13, Jogos: 11, Vitorias: 3, Empates: 4, Derrotas: 4, GolsPro: 18, GolsContra: 19, SaldoGols: -1},
		{Posicao: 14, Time: "Cuiabá", Pontos: 13, Jogos: 13, Vitorias: 3, Empates: 4, Derrotas: 6, GolsPro: 14, GolsContra: 17, SaldoGols: -3},
		{Posicao: 15, Time: "EC Vitória", Pontos: 12, Jogos: 13, Vitorias: 3, Empates: 3, Derrotas: 7, GolsPro: 14, GolsContra: 20, SaldoGols: -6},
		{Posicao: 16, Time: "Vasco da Gama", Pontos: 11, Jogos: 13, Vitorias: 3, Empates: 2, Derrotas: 8, GolsPro: 13, GolsContra: 25, SaldoGols: -12},
		{Posicao: 17, Time: "Atlético-GO", Pontos: 11, Jogos: 13, Vitorias: 2, Empates: 5, Derrotas: 6, GolsPro: 11, GolsContra: 16, SaldoGols: -5},
		{Posicao: 18, Time: "Grêmio", Pontos: 10, Jogos: 11, Vitorias: 3, Empates: 1, Derrotas: 7, GolsPro: 8, GolsContra: 12, SaldoGols: -4},
		{Posicao: 19, Time: "Corinthians", Pontos: 9, Jogos: 13, Vitorias: 1, Empates: 6, Derrotas: 6, GolsPro: 9, GolsContra: 15, SaldoGols: -6},
		{Posicao: 20, Time: "Fluminense", Pontos: 6, Jogos: 13, Vitorias: 1, Empates: 3, Derrotas: 9, GolsPro: 10, GolsContra: 21, SaldoGols: -11},
	}

	query := r.URL.Query().Get("time")
	if query != "" {
		var filteredTabela []models.Tabela
		for _, t := range tabela {
			// Converte ambos para minúsculas para fazer a comparação case-insensitive
			if strings.Contains(strings.ToLower(t.Time), strings.ToLower(query)) {
				filteredTabela = append(filteredTabela, t)
			}
		}
		tabela = filteredTabela
	}

	// Define o cabeçalho da resposta como JSON
	w.Header().Set("Content-Type", "application/json")

	// Codifica e envia a tabela como JSON na resposta
	json.NewEncoder(w).Encode(tabela)
}
