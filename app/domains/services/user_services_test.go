package services_test

import (
	"echo_sprint_planner/app/domains/repositories"
	"echo_sprint_planner/app/domains/services"
	"testing"
)

func Test_userService_UserCreate(t *testing.T) {
	type args struct {
		name     string
		email    string
		isActive bool
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				name:     "test",
				email:    "test@test.com",
				isActive: true,
				password: "testtest",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockUserRepository := repositories.NewMockUserRepository()
			us := services.NewUserService(mockUserRepository)
			if err := us.UserCreate(tt.args.name, tt.args.email, tt.args.isActive, tt.args.password); (err != nil) != tt.wantErr {
				t.Errorf("userService.UserCreate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
