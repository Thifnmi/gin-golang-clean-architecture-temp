package usecases

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/thifnmi/gin-golang-clean-architecture-temp/pkg/domain"
	"github.com/thifnmi/gin-golang-clean-architecture-temp/pkg/repositories"
)

type ExampleUsecase interface {
	GetAll(ctx context.Context, query *domain.ExampleQuery) (domain.ListExampleResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.ExampleResponse, error)
}

type exampleUsecase struct {
	exampleRepo repositories.ExampleRepository
}

func NewExampleUsecase(exampleRepo repositories.ExampleRepository) ExampleUsecase {
	return &exampleUsecase{
		exampleRepo,
	}
}

func (lu *exampleUsecase) GetAll(ctx context.Context, query *domain.ExampleQuery) (domain.ListExampleResponse, error) {
	examples, metaData, err := lu.exampleRepo.GetAll(ctx, query)
	if err != nil {
		fmt.Printf("err %v", err)
	}
	response := domain.ListExampleResponse{
		Success:   true,
		ErrorCode: 0,
		Message:   "List examples success",
		Data:      make([]domain.BaseExampleResponse, 0),
	}
	if metaData != nil {
		response.Metadata = domain.MetadataResponse{
			Total:        metaData.Total,
			CurrentPage:  metaData.CurrentPage,
			Page:         metaData.Page,
			HasPrevious:  metaData.HasPrevious,
			HasNext:      metaData.HasNext,
			PreviousPage: metaData.PreviousPage,
			NextPage:     metaData.NextPage,
		}
	}
	for _, log := range examples {
		// serviceUUID, _ := uuid.FromString(log.Service)
		baseLog := domain.BaseExampleResponse{
			Uuid:        log.Uuid,
			CreatedAt:   log.CreatedAt,
		}

		response.Data = append(response.Data, baseLog)
	}
	return response, err
}

func (lu *exampleUsecase) GetByID(ctx context.Context, id uuid.UUID) (domain.ExampleResponse, error) {
	example, err := lu.exampleRepo.GetByID(ctx, id)

	response := domain.ExampleResponse{
		Success:   true,
		ErrorCode: 0,
		Message:   "Get example success",
	}

	baseExample := domain.BaseExampleResponse{
		Uuid:        example.Uuid,
		CreatedAt:   example.CreatedAt,
	}
	response.Data = baseExample
	return response, err
}
