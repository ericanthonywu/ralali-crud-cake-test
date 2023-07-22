package Tests

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/mock"
	"ralali-crud-cake-test/Model"
	"ralali-crud-cake-test/Model/Database"
	"ralali-crud-cake-test/Repository"
	"ralali-crud-cake-test/Services"
	"regexp"
	"testing"
)

type MockedService struct {
	Services.ICakeServicesHandler
	mock.Mock
}

type MockedRepository struct {
	Repository.ICakeRepositoryHandler
	mock.Mock
}

func TestGetCake(t *testing.T) {
	repo := new(MockedRepository)
	service := new(MockedService)

	db, dbMock, _ := sqlmock.New()
	defer db.Close()

	dbMock.ExpectQuery(regexp.QuoteMeta("select id, title, description, rating, image, created_at, COALESCE(updated_at, '') from cakes where deleted_at IS NULL limit 10 offset ?")).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "title", "description", "rating", "image", "created_at", "updated_at"}).
				AddRow("1", "title", "desc", float64(7), "a", "", ""),
		)

	repo.On("GetCake", mock.Anything).Return([]Database.Cakes{
		{
			Id:          0,
			Title:       "",
			Description: "",
			Rating:      0,
			Image:       "",
			CreatedAt:   "",
			UpdatedAt:   "",
			DeletedAt:   "",
		},
	}, nil)

	Services.CakeServicesControllerProvider(Repository.CakeRepositoryControllerProvider(db)).GetCake(1)

	service.AssertExpectations(t)
}

func TestGetCakeById(t *testing.T) {
	repo := new(MockedRepository)
	service := new(MockedService)

	db, dbMock, _ := sqlmock.New()
	defer db.Close()

	dbMock.ExpectQuery(regexp.QuoteMeta("select id, title, description, rating, image, created_at, COALESCE(updated_at, '') from cakes where id = ?")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "title", "description", "rating", "imiage", "created_at", "updated_at"}))

	repo.On("GetCakeById", mock.Anything).Return(nil, nil)

	Services.CakeServicesControllerProvider(Repository.CakeRepositoryControllerProvider(db)).GetCakeById(1)

	service.AssertExpectations(t)
}

func TestAddCake(t *testing.T) {
	repo := new(MockedRepository)
	service := new(MockedService)

	db, dbMock, _ := sqlmock.New()
	defer db.Close()

	var (
		title       = "title"
		description = "description"
		rating      = float64(7)
		image       = "image"
	)

	dbMock.ExpectExec(regexp.QuoteMeta("insert into cakes (title, description, rating, image) VALUES (?,?,?,?)")).
		WithArgs(title, description, rating, image)

	repo.On("AddCake", mock.Anything).Return(nil)

	Services.CakeServicesControllerProvider(Repository.CakeRepositoryControllerProvider(db)).AddCake(Model.CakeRequestDto{
		Title:       title,
		Description: description,
		Rating:      rating,
		Image:       image,
	})

	service.AssertExpectations(t)
}

func TestUpdateCake(t *testing.T) {
	repo := new(MockedRepository)
	service := new(MockedService)

	db, dbMock, _ := sqlmock.New()
	defer db.Close()

	var (
		title       = "title"
		description = "description"
		rating      = float64(7)
		image       = "image"
		id          = uint64(1)
	)

	dbMock.ExpectExec(regexp.QuoteMeta("update cakes set title = ?, description = ?, rating = ?, image = ?, updated_at = current_timestamp where id = ?")).
		WithArgs(title, description, rating, image, id)

	repo.On("UpdateCake", mock.Anything).Return(nil)

	Services.CakeServicesControllerProvider(Repository.CakeRepositoryControllerProvider(db)).UpdateCake(id, Model.CakeRequestDto{
		Title:       title,
		Description: description,
		Rating:      rating,
		Image:       image,
	})

	service.AssertExpectations(t)
}

func TestDeleteCake(t *testing.T) {
	repo := new(MockedRepository)
	service := new(MockedService)

	db, dbMock, _ := sqlmock.New()
	defer db.Close()

	var (
		id = uint64(1)
	)

	dbMock.ExpectExec(regexp.QuoteMeta("update cakes set deleted_at = current_timestamp where id = ?")).
		WithArgs(id)

	repo.On("DeleteCake", mock.Anything).Return(nil)

	Services.CakeServicesControllerProvider(Repository.CakeRepositoryControllerProvider(db)).DeleteCake(id)

	service.AssertExpectations(t)
}
