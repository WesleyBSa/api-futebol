package models

type Tabela struct {
	Posicao    int    `json:"posicao"`
	Time       string `json:"time"`
	Pontos     int    `json:"pontos"`
	Jogos      int    `json:"jogos"`
	Vitorias   int    `json:"vitorias"`
	Empates    int    `json:"empates"`
	Derrotas   int    `json:"derrotas"`
	GolsPro    int    `json:"gols_pro"`
	GolsContra int    `json:"gols_contra"`
	SaldoGols  int    `json:"saldo_gols"`
}
