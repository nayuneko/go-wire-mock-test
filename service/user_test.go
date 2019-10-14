package service

import (
	"testing"

	"github.com/pkg/errors"

	"github.com/nayuneko/go-wire-mock-test/entity"

	"github.com/nayuneko/go-wire-mock-test/mock_repository"

	"github.com/golang/mock/gomock"
)

func Test_UserService_CreateEntry(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock_repository.NewMockUserRepository(ctrl)
	entryRepo := mock_repository.NewMockEntryRepository(ctrl)

	t.Run("success", func(t *testing.T) {

		id := int64(1)
		userRepo.
			EXPECT().
			FindByID(gomock.Eq(id)).
			Return(&entity.User{ID: id}, nil)
		entryRepo.
			EXPECT().
			CreateEntry(gomock.Eq(&entity.Entry{UserID: id})).
			Return(nil, nil)

		s := NewUserService(userRepo, entryRepo)

		err := s.CreateEntry(id)

		if err != nil {
			t.Fatalf("err is not nil. err = %#v", err)
		}
	})

	t.Run("failed FindByID", func(t *testing.T) {

		id := int64(101)
		userRepo.
			EXPECT().
			FindByID(gomock.Eq(id)).
			Return(nil, errors.New("failed FindByID"))

		s := NewUserService(userRepo, entryRepo)

		err := s.CreateEntry(id)

		if err == nil {
			t.Fatal("err is nil.")
		}
		want := "failed FindByID"
		if got := err.Error(); got != want {
			t.Fatalf("mismatch error message. got = %#v, want = %#v", got, want)
		}
	})
	t.Run("failed CreateEntry", func(t *testing.T) {

		id := int64(102)
		userRepo.
			EXPECT().
			FindByID(gomock.Eq(id)).
			Return(&entity.User{ID: id}, nil)
		entryRepo.
			EXPECT().
			CreateEntry(gomock.Eq(&entity.Entry{UserID: id})).
			Return(nil, errors.New("failed CreateEntry"))

		s := NewUserService(userRepo, entryRepo)

		err := s.CreateEntry(id)

		if err == nil {
			t.Fatal("err is nil.")
		}
		want := "failed CreateEntry"
		if got := err.Error(); got != want {
			t.Fatalf("mismatch error message. got = %#v, want = %#v", got, want)
		}
	})
}
