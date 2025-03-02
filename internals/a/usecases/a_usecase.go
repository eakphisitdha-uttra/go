package usecases

import (
	"microservice/internals/a/adapters/inputs"
	"microservice/internals/a/adapters/outputs"
	"microservice/internals/a/repositories"
	"microservice/logs"
)

type IUsecase interface {
	Get() ([]outputs.GetOutput, error)
	Add(input inputs.AddInput) error
	Update(input inputs.UpdateInput) error
	Delete(input inputs.DeleteInput) error
}

type Usecase struct {
	repo repositories.IRepository
}

func NewUsecase(repo repositories.IRepository) IUsecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) Get() ([]outputs.GetOutput, error) {
	data, err := u.repo.Get()
	if err != nil {
		logs.Error(err.Error())
		return nil, err
	}
	//
	// put your business logic
	//
	output := []outputs.GetOutput{}
	for _, i := range data {
		var res outputs.GetOutput
		//
		// put your business logic
		//
		res.ID = i.ID
		res.Name = i.Name
		//
		output = append(output, res)
	}

	return output, nil
}

func (s *Usecase) Add(input inputs.AddInput) error {
	//
	// put your business logic
	//
	err := s.repo.Add(input)
	if err != nil {
		return err
	}

	err = s.repo.Log("name", input.Name, 1, "ADD")
	if err != nil {
		return err
	}

	return nil
}

func (s *Usecase) Update(input inputs.UpdateInput) error {
	//
	// put your business logic
	//
	err := s.repo.Update(input)
	if err != nil {
		return err
	}
	err = s.repo.Log("id", input.ID, input.ID, "UPDATE")
	if err != nil {
		return err
	}

	return nil
}

func (s *Usecase) Delete(input inputs.DeleteInput) error {
	//
	// put your business logic
	//
	err := s.repo.Delete(input)
	if err != nil {
		return err
	}

	err = s.repo.Log("id", input.ID, input.ID, "DELETE")
	if err != nil {
		return err
	}
	return nil
}
