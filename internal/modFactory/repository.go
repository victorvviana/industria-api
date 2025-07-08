package modfactory

type ProducaoRepository interface {
	CriarProduto(p Produto) (int, error)
	ListarProdutos(tipo string) ([]Produto, error)
	CriarOrdemProducao(op OrdemProducao) (int, error)
	AtualizarStatusOP(id int, status string) error
	BuscarOPPorID(id int) (*OrdemProducao, error)
	ListarOPs() ([]OrdemProducao, error)
}
