package usecases

import (
	"microservice/databases/postgresql/tables"
	"microservice/internals/module_a/adapters/inputs"
	"microservice/internals/module_a/adapters/outputs"
	mocksRepository "microservice/internals/module_a/repositories/mocks"
	"testing"

	"github.com/stretchr/testify/suite"
)

type UsecaseTestSuite struct {
	suite.Suite
	mockRepository *mocksRepository.IRepository
	usecase        IUsecase
}

func TestServiceTestSuite(t *testing.T) {
	suite.Run(t, new(UsecaseTestSuite))
}

func (suite *UsecaseTestSuite) SetupTest() {
	suite.mockRepository = mocksRepository.NewIRepository(suite.T())
	suite.usecase = NewUsecase(suite.mockRepository)
}

func (suite *UsecaseTestSuite) TestGet() {
	//Input

	//mock
	suite.mockRepository.EXPECT().Get().Return([]tables.Users{
		{
			ID:   1,
			Name: "A",
		},
	}, nil)

	//expected
	expected := []outputs.GetOutput{
		{
			ID:   1,
			Name: "A",
		},
	}
	//function test
	actual, err := suite.usecase.Get()

	suite.Require().NoError(err)
	suite.Equal(expected, actual)
}

func (suite *UsecaseTestSuite) TestAdd() {
	//Input
	input := inputs.AddInput{
		Name: "B",
	}
	fields := "name"

	//expected
	suite.mockRepository.EXPECT().Add(input).Return(nil)
	suite.mockRepository.EXPECT().Log(fields, input.Name, 1, "ADD").Return(nil)

	//function test
	err := suite.usecase.Add(input)

	suite.Require().NoError(err)
}

func (suite *UsecaseTestSuite) TestUpdate() {
	//Input
	input := inputs.UpdateInput{
		ID: 1,
	}
	fields := "id"

	//expected
	suite.mockRepository.EXPECT().Update(input).Return(nil)
	suite.mockRepository.EXPECT().Log(fields, input.ID, input.ID, "UPDATE").Return(nil)

	//function test
	err := suite.usecase.Update(input)

	suite.Require().NoError(err)
}

func (suite *UsecaseTestSuite) TestDelete() {
	//Input
	input := inputs.DeleteInput{
		ID: 1,
	}
	fields := "id"

	//expected
	suite.mockRepository.EXPECT().Delete(input).Return(nil)
	suite.mockRepository.EXPECT().Log(fields, input.ID, input.ID, "DELETE").Return(nil)

	//function test
	err := suite.usecase.Delete(input)

	suite.Require().NoError(err)
}
