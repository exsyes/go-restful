package simple

type Configuration struct {
	Name string
}

type Application struct {
	Configuration *Configuration
}

func NewApplication() *Application {
	return &Application{Configuration: &Configuration{
		Name: "SUHARDI",
	}}
}
