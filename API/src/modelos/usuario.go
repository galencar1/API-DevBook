package modelos

import "time"

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     uint64    `json:"nome,omitempty"`
	Nick     uint64    `json:"nick,omitempty"`
	Email    uint64    `json:"email,omitempty"`
	Senha    uint64    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}
