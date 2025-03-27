package utils_test

import (
	"testing"

	pb_user "github.com/agrotention/pb-user"
	"github.com/agrotention/service-user/utils"
)

// TestValidatePassword menguji berbagai skenario validasi password.
func TestValidatePassword(t *testing.T) {
	tests := []struct {
		Name      string
		Data      *pb_user.CreateUserRequest
		WantError bool
	}{
		{
			Name: "Valid password",
			Data: &pb_user.CreateUserRequest{
				Email:     "example@email.com",
				Username:  "johndoe",
				Password:  "Secret1234",
				FirstName: "John",
				LastName:  "Doe",
			},
			WantError: false,
		},
		{
			Name: "Password too short",
			Data: &pb_user.CreateUserRequest{
				Email:     "test@email.com",
				Username:  "shortuser",
				Password:  "abc",
				FirstName: "Short",
				LastName:  "User",
			},
			WantError: true,
		},
		{
			Name: "Password missing number",
			Data: &pb_user.CreateUserRequest{
				Email:     "user@email.com",
				Username:  "nonumber",
				Password:  "PasswordOnly",
				FirstName: "No",
				LastName:  "Number",
			},
			WantError: true,
		},
		{
			Name: "Password missing uppercase",
			Data: &pb_user.CreateUserRequest{
				Email:     "lower@email.com",
				Username:  "lowercase",
				Password:  "alllower123",
				FirstName: "Lower",
				LastName:  "Case",
			},
			WantError: true,
		},
		{
			Name: "Password missing lowercase",
			Data: &pb_user.CreateUserRequest{
				Email:     "upper@email.com",
				Username:  "uppercase",
				Password:  "ALLUPPER123",
				FirstName: "Upper",
				LastName:  "Case",
			},
			WantError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			validator := utils.NewValidator()
			errors := validator.Validate(tt.Data)
			gotError := len(errors) > 0
			if gotError != tt.WantError {
				t.Errorf("Test %q failed: expected error = %v, got = %v", tt.Name, tt.WantError, gotError)
			}
		})
	}
}
