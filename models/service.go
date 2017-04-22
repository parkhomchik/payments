package payment

//Service сервисы
type Service struct{
	ID int
	Name string
	ShortName string
	Description string
	IsEnable bool
	IsAnonimusEnable bool
}

//Client клиенты приложения
type Client struct {
	ID int
	Name string
	Enable bool
}

//ServiceToClient какие сервисы доступны клиенту
type ServiceToClient struct{
	ServiceID int
	ClientID int
	Enable bool
}


/*
type Command struct {
	ID int
	Name string
}

type ServiceCommand struct {
	ID int
	ServiceID int
	CommandID int
	Step int
	Execute Execute
}

type Execute struct {
	ID int
	Name string
	Command string
}

type Parameter struct {
	ID int
	Name string
	Type Type		
}

type Type struct {
	ID int
	Name string
}*/