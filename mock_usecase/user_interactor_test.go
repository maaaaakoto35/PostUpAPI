package mock_usecase

import (
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/maaaaakoto35/PostUpAPI/domain"
)

func TestUserByID(t testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	var expected domain.User
	var err error
	id := 1

	mockSample := NewMockUserRepository(ctrl)
	mockSample.EXPECT().FindByID(id).Return(expected, err)

	// mockを利用してusecase.UserByIDをテスト

}
