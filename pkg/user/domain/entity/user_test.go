// FILEPATH: /Users/yudai/apps/grpc-practice/pkg/user/domain/entity/user_test.go
package entity

import (
	"testing"
	"time"

	"github.com/yudai2929/grpc-practice/pkg/libs/password"
)

func TestIsEqualPassword(t *testing.T) {
	// Arrange
	pass := "testPassword"
	hashedPass, _ := password.Hash(pass)
	user := &User{
		ID:        "1",
		Name:      "Test User",
		Email:     "test@example.com",
		Password:  hashedPass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// Act and Assert
	if !user.IsEqualPassword(pass) {
		t.Errorf("IsEqualPassword() failed, expected %v, got %v", true, false)
	}

	// Act and Assert
	if user.IsEqualPassword("wrongPassword") {
		t.Errorf("IsEqualPassword() failed, expected %v, got %v", false, true)
	}
}
