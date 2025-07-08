package modfactory

import "time"

type Produto struct {
	ID       int
	Nome     string
	Codigo   string
	Tipo     string // "materia_prima" ou "produto_acabado"
	Unidade  string
	Estoque  float64
	Ativo    bool
	CriadoEm time.Time
}

// Ordem de Produção
type OrdemProducao struct {
	ID         int
	ProdutoID  int
	Quantidade float64
	Status     string // planejada, em_producao, concluida, cancelada
	DataInicio *time.Time
	DataFim    *time.Time
	CriadoEm   time.Time
}

// Entrada para criar uma OP
type CriarOPInput struct {
	ProdutoID  int
	Quantidade float64
}
