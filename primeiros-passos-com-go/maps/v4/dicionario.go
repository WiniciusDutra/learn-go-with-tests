package main

const (
	// ErrNaoEncontrado significa que a definição não pôde ser encontrada para determinada palavra
	ErrNaoEncontrado = ErrDicionario("não foi possível encontrar a palavra que você procura")

	// ErrPalavraExistente significa que você está tentando adicionar uma palavra que já existe
	ErrPalavraExistente = ErrDicionario("não é possível adicionar a palavra pois ela já existe")
)

// ErrDicionario são erros que podem acontecer quando se interage com o dicionário
type ErrDicionario string

func (e ErrDicionario) Error() string {
	return string(e)
}

// Dicionario armazena definições de palavras
type Dicionario map[string]string

// Busca encontra uma palavra no dicionário
func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}

	return definicao, nil
}

// Adiciona insere uma palavra e definição no dicionário
func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraExistente
	default:
		return err

	}

	return nil
}
