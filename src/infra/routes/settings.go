package routes

type PropsRoute struct {
	Name string
	Fun func()
}

// ParamsRoutes representa os parâmetros de rota; O nome precisa ser igual a o do arquivo, para que o for dentro do arquivo routes.go funcione.
var ParamsRoutes = []PropsRoute{
	{Name: "getuser", Fun: GetUserRoute},
}