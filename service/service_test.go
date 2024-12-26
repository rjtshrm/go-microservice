package service

import "testing"

func Test_returnHelloFromMicroService(t *testing.T) {
	if returnHelloFromMicroService() != "Hello From Microservice!" {
		t.Errorf("Expected returnHelloFromMicroService() to return 'Hello From Microservice!', got %v", returnHelloFromMicroService())
	}
}
