package command

import "context"

type (
	// Command represents the intention to execute an operation on our system that modifies its state
	//
	// Command is a DTO (Data Transfer Object) which represents the action to be performed
	Command interface {
		// CommandId method that implements all commands(Data Transfer Object)
		// returns a string representing the type of command to be performed
		CommandId() string
	}
	// Bus will be in charge of searching among the registered Command Handlers
	// and executing the function of such Handler when it receives a Command as parameter in its Handle method
	Bus[V Command] interface {
		// Dispatch method that implements the CommandBus that looks for the registered Command with its Handler
		// and executes the function
		//
		// If the register no searching it is returns an error
		Dispatch(ctx context.Context, v V) error
	}

	// Handler will be in charge of performing the action we are looking for,
	// simply returning an error if the process is not correct
	//
	// Maps the DTO values to value objects in our domain and invokes the use case
	Handler[V Command] interface {
		// Handler represents the action that you want to perform by means of the Command
		Handler(ctx context.Context, v V) error
	}
)
