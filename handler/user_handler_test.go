package handler

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestCreateUser(t *testing.T) {
	repo := &MockUserHandler{}
	repo.On("CreateUser", mock.Anything).Return(nil)

	err := repo.CreateUser(nil)
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	repo.AssertExpectations(t)
}
