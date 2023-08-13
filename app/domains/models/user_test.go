package models_test

import (
	"echo_sprint_planner/app/domains/models"
	"echo_sprint_planner/app/utils"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestUser_CreateValidate(t *testing.T) {
	now := time.Now()
	user_id := uuid.New()
	password := "password"

	// 6文字以下のパスワード
	short_password, err := utils.MakeRandomString(5)
	if err != nil {
		t.Fatal(err)
	}
	// 51文字以上のパスワード
	long_password, err := utils.MakeRandomString(51)
	if err != nil {
		t.Fatal(err)
	}

	type fields struct {
		ID       *uuid.UUID
		Name     string
		Email    string
		IsActive bool
		Password *string
		CreateAt *time.Time
		UpdateAt *time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				ID:       &user_id,
				Name:     "test",
				Email:    "test@test.com",
				IsActive: true,
				Password: &password,
				CreateAt: &now,
				UpdateAt: &now,
			},
			wantErr: false,
		},
		{
			name: "異常系:名前が空",
			fields: fields{
				ID:       &user_id,
				Name:     "",
				Email:    "test@test.com",
				IsActive: true,
				Password: &password,
				CreateAt: &now,
				UpdateAt: &now,
			},
			wantErr: true,
		},
		{
			name: "異常系:名前が3文字未満",
			fields: fields{
				ID:       &user_id,
				Name:     "te",
				Email:    "test@test.com",
				IsActive: true,
				Password: &password,
				CreateAt: &now,
				UpdateAt: &now,
			},
			wantErr: true,
		},
		{
			name: "異常系:名前が51文字以上",
			fields: fields{
				ID:       &user_id,
				Name:     "testtesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttesttest",
				Email:    "test@test.com",
				IsActive: true,
				Password: &password,
				CreateAt: &now,
				UpdateAt: &now,
			},
			wantErr: true,
		},
		{
			name: "異常系:メールアドレスが空",
			fields: fields{
				ID:       &user_id,
				Name:     "test",
				Email:    "",
				IsActive: true,
				Password: &password,
				CreateAt: &now,
				UpdateAt: &now,
			},
			wantErr: true,
		},
		{
			name: "異常系:メールアドレスが不正",
			fields: fields{
				ID:       &user_id,
				Name:     "test",
				Email:    "testtest.com",
				IsActive: true,
				Password: &password,
				CreateAt: &now,
				UpdateAt: &now,
			},
			wantErr: true,
		},
		{
			name: "異常系:パスワードが空",
			fields: fields{
				ID:       &user_id,
				Name:     "test",
				Email:    "test@test.com",
				IsActive: true,
				Password: nil,
				CreateAt: &now,
				UpdateAt: &now,
			},
			wantErr: true,
		},
		{
			name: "異常系:パスワードが6文字未満",
			fields: fields{
				ID:       &user_id,
				Name:     "test",
				Email:    "test@test.com",
				IsActive: true,
				Password: &short_password,
				CreateAt: &now,
				UpdateAt: &now,
			},
			wantErr: true,
		},
		{
			name: "異常系:パスワードが51文字以上",
			fields: fields{
				ID:       &user_id,
				Name:     "test",
				Email:    "test@test.com",
				IsActive: true,
				Password: &long_password,
				CreateAt: &now,
				UpdateAt: &now,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &models.User{
				ID:       tt.fields.ID,
				Name:     tt.fields.Name,
				Email:    tt.fields.Email,
				IsActive: tt.fields.IsActive,
				Password: tt.fields.Password,
				CreateAt: tt.fields.CreateAt,
				UpdateAt: tt.fields.UpdateAt,
			}
			if err := u.CreateValidate(); (err != nil) != tt.wantErr {
				t.Errorf("User.CreateValidate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
