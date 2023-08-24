package inmemory

import (
	"context"
	"testing"
	"todoapp/internal/models"
	"todoapp/internal/repository/inmemory/mocks"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func TestInMemoryUserRepository_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	type args struct {
		ctx  context.Context
		user *models.User
	}

	mockEmailCache := mocks.NewMockCache(ctrl)
	mockIDCache := mocks.NewMockCache(ctrl)

	tests := []struct {
		name           string
		emailCacheMock func()
		idCacheMock    func()
		args           args
		wantErr        bool
	}{
		{
			name: "Create user success",
			emailCacheMock: func() {
				mockEmailCache.EXPECT().Contains("test@example.com").Return(false)
				mockEmailCache.EXPECT().Put("test@example.com", &models.User{
					Email: "test@example.com",
				})
			},
			idCacheMock: func() {
				mockIDCache.EXPECT().Contains("1").Return(false)
				mockIDCache.EXPECT().Put("1", &models.User{
					Email: "test@example.com",
				})
			},
			args: args{
				ctx: context.TODO(),
				user: &models.User{
					ID:    uuid.New(),
					Email: "test@example.com",
				},
			},
			wantErr: false,
		},
		{
			name:           "Create user, user is nil",
			emailCacheMock: func() {},
			idCacheMock:    func() {},
			args: args{
				ctx:  context.TODO(),
				user: nil,
			},
			wantErr: true,
		},
		{
			name: "Create user, user already exists",
			emailCacheMock: func() {
				mockEmailCache.EXPECT().Contains("test@example.com").Return(true)
			},
			idCacheMock: func() {},
			args: args{
				ctx: context.TODO(),
				user: &models.User{
					ID:    uuid.New(),
					Email: "test@example.com",
				},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.emailCacheMock()
			tt.idCacheMock()
			i := &InMemoryUserRepository{
				emailCache: mockEmailCache,
				idCache:    mockIDCache,
			}
			if err := i.CreateUser(tt.args.ctx, tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
