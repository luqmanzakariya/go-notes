//go:build wireinject
// +build wireinject

package simple

import (
	"io"
	"os"

	"github.com/google/wire"
)

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(
		NewDatabaseRepository,
		NewDatabasePostgreSQL,
		NewDatabaseMongoDB,
	)

	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FoobarService {
	wire.Build(fooSet, barSet, NewFoobarService)
	return nil
}

// Versi salah 253-binding-iterface_test.go
// func InitializedHelloService() *HelloService {
// 	wire.Build(NewSayHelloService, NewSayHelloImpl)
// 	return nil
// }

// Versi Betul 253-binding-iterface_test.go
var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

func InitializedHelloService() *HelloService {
	wire.Build(helloSet, NewHelloService)
	return nil
}

// Versi 254-struct-provider_test.go
func InitializedFooBar() *FooBar {
	wire.Build(NewFoo, NewBar, wire.Struct(new(FooBar), "Foo", "Bar"))
	return nil
}

// Versi 255-binding_values.go
var fooValue = &Foo{}
var barValue = &Bar{}

func InitializedFooBarUsingValue() *FooBar {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBar), "*"))
	return nil
}

func InitializedReader() io.Reader {
	wire.Build(wire.InterfaceValue(new(io.Reader), os.Stdin))
	return nil
}

// Versi 256-struct-field-provider_test.go
func InitializedConfiguration() *Configuration {
	wire.Build(
		NewApplication,
		wire.FieldsOf(new(*Application), "Configuration"),
	)
	return nil
}

// Versi 257-cleanup-function_test.go
func InitializedConnection(name string) (*Connection, func()) {
	wire.Build(NewConnection, NewFile)
	return nil, nil
}
