package modfactory

import (
	"errors"
	"time"
)

type ProducaoService interface {
	CadastrarProduto(p Produto) (int, error)
	ListarProdutos(tipo string) ([]Produto, error)
	CriarOrdemProducao(input CriarOPInput) (int, error)
	IniciarOP(id int) error
	FinalizarOP(id int) error
}

type producaoService struct {
	repo ProducaoRepository
}

func NewProducaoService(r ProducaoRepository) ProducaoService {
	return &producaoService{repo: r}
}

func (s *producaoService) CadastrarProduto(p Produto) (int, error) {
	if p.Nome == "" || p.Tipo == "" {
		return 0, errors.New("dados incompletos")
	}
	return s.repo.CriarProduto(p)
}

func (s *producaoService) ListarProdutos(tipo string) ([]Produto, error) {
	return s.repo.ListarProdutos(tipo)
}

func (s *producaoService) CriarOrdemProducao(input CriarOPInput) (int, error) {
	if input.Quantidade <= 0 {
		return 0, errors.New("quantidade inválida")
	}
	op := OrdemProducao{
		ProdutoID:  input.ProdutoID,
		Quantidade: input.Quantidade,
		Status:     "planejada",
		CriadoEm:   time.Now(),
	}
	return s.repo.CriarOrdemProducao(op)
}

func (s *producaoService) IniciarOP(id int) error {
	return s.repo.AtualizarStatusOP(id, "em_producao")
}

func (s *producaoService) FinalizarOP(id int) error {
	return s.repo.AtualizarStatusOP(id, "concluida")
}
