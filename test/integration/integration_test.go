package integration

import (
	"context"
	"log/slog"
	"testing"
	"time"

	"github.com/ZhdanovichVlad/service-podof/internal/entity"
	"github.com/ZhdanovichVlad/service-podof/internal/repository/postgres"
	"github.com/ZhdanovichVlad/service-podof/internal/service"
	"github.com/ZhdanovichVlad/service-podof/internal/service/mocks"
	"github.com/ZhdanovichVlad/service-podof/test/integration/test_repository"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)


const (
	getProductPath = "SELECT id, type, date_time, reception_id FROM products"
)

func TestCreatePveAndAddProducts(t *testing.T) {
	ctrl := gomock.NewController(t)

	db := test_repository.SetupTestDB(t)
	logger := slog.Default()
	repo := postgres.NewStorage(db, logger)

	tokenGen := mocks.NewMocktokenGenerator(ctrl)

	productList := []entity.Product{}

	service := service.NewService(repo, logger, tokenGen)

	ctx := context.Background()

	rows, err := db.Query(ctx, getProductPath)
	assert.NoError(t, err)
	defer rows.Close()

	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Type, &product.DateTime, &product.ReceptionID)
		assert.NoError(t, err)
		productList = append(productList, product)
	}

	assert.Equal(t, len(productList), 0)


	pvzId := uuid.New()

	pvz := &entity.Pvz{
		Id: pvzId,
		RegistrationDate: time.Now(),
		City: "Москва",
	}

	_, err = service.CreatePvz(ctx, pvz)
	assert.NoError(t, err)

	receptionRequest := &entity.Reception{
		DateTime: time.Now(),
		PvzID: pvzId,
	}
	reception, err := service.CreateReception(ctx, receptionRequest)
	assert.NoError(t, err)

	product := &entity.Product{
		ID: uuid.New(),
		Type: "Футболка",
		DateTime: time.Now(),
		ReceptionID: reception.Id,
	}

	for i := 0; i < 50; i++ {
		_, err = service.CreateProduct(ctx, product, pvzId)
		assert.NoError(t, err)
		time.Sleep(10 * time.Millisecond)
	}

	rows, err = db.Query(ctx, getProductPath)
	assert.NoError(t, err)
	defer rows.Close()

	newProductList := []entity.Product{}
	for rows.Next() {
		var product entity.Product
		err = rows.Scan(&product.ID, &product.Type, &product.DateTime, &product.ReceptionID)
		assert.NoError(t, err)
		newProductList = append(newProductList, product)
	}


	service.CloseReception(ctx, reception.Id)

	assert.Equal(t, 50, len(newProductList))

}
