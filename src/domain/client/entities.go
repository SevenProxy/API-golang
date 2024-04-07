package client

// 👤 Criando a instancia do usuário.
type PropsUser struct {
	ID 			int
	Name		string
	Avatar	string
}

// 🧩 Uma representação de um banco de dados. Você deve fazer a conexão com o banco de dados no dentro do infra em seguida chamado-lo aqui.
var db = []PropsUser{
	{ID: 123, Name: "proxy", Avatar: "https://app.com/upload/123123.jpg"},
}

func (pu *PropsUser) GetUserById(id int) (*PropsUser, error) {
	for _, user := range db {
		// ♟️ Código simples que irá realizar uma verificação se o usuário existe.
		if user.ID == id {
      return &PropsUser{ID: user.ID, Name: user.Name, Avatar: user.Avatar}, nil
		}
	}

	return nil, nil
}