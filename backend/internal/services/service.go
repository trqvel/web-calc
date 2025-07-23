package services

import (
	"fmt"

	"github.com/Knetic/govaluate"
	"github.com/google/uuid"
)

type Service interface {
	CreateCalculation(expression string) (Calculation, error)
	ReadAllCalculations() ([]Calculation, error)
	GetCalculationByID(id string) (Calculation, error)
	UpdateCalculation(id, expression string) (Calculation, error)
	DeleteCalculation(id string) error
}

type svc struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &svc{repo: r}
}

func (s *svc) calculateExpression(expression string) (string, error) {
	expr, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return "", err
	}

	res, err := expr.Evaluate(nil)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", res), err
}

func (s *svc) CreateCalculation(expression string) (Calculation, error) {
	res, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calc := Calculation{
		ID:         uuid.NewString(),
		Expression: expression,
		Result:     res,
	}

	if err := s.repo.CreateCalculation(calc); err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

func (s *svc) ReadAllCalculations() ([]Calculation, error) {
	return s.repo.ReadAllCalculations()
}

func (s *svc) GetCalculationByID(id string) (Calculation, error) {
	return s.repo.GetCalculationByID(id)
}

func (s *svc) UpdateCalculation(id, expression string) (Calculation, error) {
	calc, err := s.repo.GetCalculationByID(id)
	if err != nil {
		return Calculation{}, err
	}

	res, err := s.calculateExpression(expression)
	if err != nil {
		return Calculation{}, err
	}

	calc.Expression = expression
	calc.Result = res

	if err := s.repo.UpdateCalculation(calc); err != nil {
		return Calculation{}, err
	}

	return calc, nil
}

func (s *svc) DeleteCalculation(id string) error {
	return s.repo.DeleteCalculation(id)
}
