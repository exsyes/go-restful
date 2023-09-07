//go:build wireinject
// +build wireinject

package simple

import (
	"github.com/google/wire"
	"io"
	"os"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseMongoDB,
		NewDatabasePostgreSQL,
		NewDatabaseRepository,
	)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

var HelloSet = wire.NewSet(NewSayHelloImpl, wire.Bind(new(SayHello), new(*SayHelloImpl)))

func InitializedHelloService() *HelloService {
	wire.Build(HelloSet, NewSayHelloService)
	return nil
}

var CustomerSet = wire.NewSet(NewCustomerImpl, wire.Bind(new(Customer), new(*CustomerImpl)))

func InitializedCustomerService() *CustomerService {
	wire.Build(CustomerSet, NewCustomerService)
	return nil
}

var FooBarSet = wire.NewSet(
	NewFoo, NewBar)

func InitializedFooBar() *FooBar {
	wire.Build(FooBarSet, wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedFooBarValue() *FooBar {
	wire.Build(wire.Value(&Foo{}), wire.Value(&Bar{}), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}
