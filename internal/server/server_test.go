package server

import (
	"context"
	"os"
	"testing"
	"wildfire/internal/clients/schema"
)

type MockNamesClient struct {
	GetRandomNameResponse *schema.GetRandomNameResponse
	GetRandomNameError    error
}

func (m *MockNamesClient) GetRandomName(ctx context.Context) (*schema.GetRandomNameResponse, error) {
	if m.GetRandomNameError != nil {
		return nil, m.GetRandomNameError
	}

	if m.GetRandomNameResponse != nil {
		return m.GetRandomNameResponse, nil
	}

	return &schema.GetRandomNameResponse{
		FirstName: "John",
		LastName:  "Doe",
	}, nil
}

type MockJokesClient struct {
	GetJokeResponse *schema.GetJokeResponse
	GetJokeError    error
}

func (m *MockJokesClient) GetJoke(ctx context.Context, firstName *string, lastName *string) (*schema.GetJokeResponse, error) {
	if m.GetJokeError != nil {
		return nil, m.GetJokeError
	}

	if m.GetJokeResponse != nil {
		return m.GetJokeResponse, nil
	}

	return &schema.GetJokeResponse{
		Value: schema.GetJokeResponseValue{
			Joke: "This is a joke",
		},
	}, nil
}

type serverTest struct {
	JokesClient *MockJokesClient
	NamesClient *MockNamesClient
}

var ServerTest serverTest = serverTest{}

func setup(ctx context.Context) {
	ServerTest.JokesClient = &MockJokesClient{}
	ServerTest.NamesClient = &MockNamesClient{}
}

// TODO:
func cleanup(ctx context.Context) {
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	setup(ctx)
	code := m.Run()
	cleanup(ctx)
	os.Exit(code)
}
