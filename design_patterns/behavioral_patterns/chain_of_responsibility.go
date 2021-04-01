package behavioral_patterns

import (
	"fmt"
	"regexp"
)

/* Avoid coupling the sender of a request to its receiver
by giving more than one object a chance to handle the request.
Chain the receiving objects and pass the request along the chain
until an object handles it
*/

type (
	Request struct {
		token, data string
	}

	Handler interface {
		SetNext(Handler)
		Run(Request)
	}

	AuthHandler struct {
		next Handler
	}
	ValidatorHandler struct {
		next Handler
	}
)

func (auth *AuthHandler) SetNext(handler Handler) {
	auth.next = handler
}

func (auth *AuthHandler) Run(req Request) {
	if req.token != "" {
		fmt.Println("Auth Passed!!")
		if auth.next != nil {
			auth.next.Run(req)
		}
	} else {
		fmt.Println("Auth failed!!")
	}
}

func (validator *ValidatorHandler) SetNext(handler Handler) {
	validator.next = handler
}

func (validator *ValidatorHandler) Run(req Request) {
	if regexp.MustCompile("^[0-9A-Za-z]+$").MatchString(req.data) {
		fmt.Println("Validator Passed!!")
		if validator.next != nil {
			validator.next.Run(req)
		}
	} else {
		fmt.Println("Validator failed!!")
	}
}

func ChainOfResponsibilityExample() {
	handler := new(AuthHandler)
	handler.SetNext(new(ValidatorHandler))

	fmt.Println("1.")
	handler.Run(Request{
		token: "abcdef",
		data:  "ABCDEFGH",
	})

	fmt.Println("2.")
	handler.Run(Request{
		token: "a",
		data:  "ABCDE---FGH",
	})
}
