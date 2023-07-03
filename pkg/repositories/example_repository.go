package repositories

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/thifnmi/gin-golang-clean-architecture-temp/pkg/domain"
	base_query "github.com/thifnmi/gin-golang-clean-architecture-temp/pkg/interfaces/http/query"

	// "go.elastic.co/apm/v2"
	"gorm.io/gorm"
)

type ExampleRepository interface {
	GetAll(ctx context.Context, query *domain.ExampleQuery) ([]domain.BaseExampleResponse, *domain.MetadataResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (domain.Example, error)
}

type exampleRepository struct {
	db *gorm.DB
}

func NewExampleRepository(db *gorm.DB) ExampleRepository {
	return &exampleRepository{
		db,
	}
}

func (lr *exampleRepository) GetAll(ctx context.Context, query *domain.ExampleQuery) ([]domain.BaseExampleResponse, *domain.MetadataResponse, error) {
	var (
		logs     []domain.Example
		baseExample []domain.BaseExampleResponse
		metaData domain.MetadataResponse
	)
	q := lr.db.Model(domain.Example{})

	metaData, err := base_query.GetPaginatedResults(ctx, lr.db, &logs, q, 1, 20)
	if err != nil {
		fmt.Sprintf("error %v when query metadata %v", err, query)
	}
	// span, ctx = apm.StartSpan(ctx, "GetAll Repository", "DB MySQL query results")
	// defer span.End()
	result := q.Limit(20).Offset(10).Find(&baseExample)
	if result.Error != nil {
		fmt.Sprintf("query all with filter %v have err %v", query, result.Error)
		return nil, nil, result.Error
	}
	return baseExample, &metaData, nil
}

func (lr *exampleRepository) GetByID(ctx context.Context, id uuid.UUID) (domain.Example, error) {
	var example domain.Example
	lr.db.Where("uuid = ?", id).First(&example)
	return example, nil
}
