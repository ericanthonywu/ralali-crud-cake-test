package Repository

import (
	"database/sql"
	"ralali-crud-cake-test/Model"
	"ralali-crud-cake-test/Model/Database"
)

type (
	ICakeRepositoryHandler interface {
		CheckExists(id uint64) (exists bool, err error)
		GetCake(pagination uint) (data []Database.Cakes, err error)
		GetCakeById(id uint64) (data Database.Cakes, err error)
		AddCake(data Model.CakeRequestDto) (err error)
		UpdateCake(id uint64, data Model.CakeRequestDto) (err error)
		DeleteCake(id uint64) (err error)
	}

	CakeRepositoryHandler struct {
		DB *sql.DB
	}
)

func CakeRepositoryControllerProvider(DB *sql.DB) *CakeRepositoryHandler {
	return &CakeRepositoryHandler{
		DB: DB,
	}
}

func (h *CakeRepositoryHandler) CheckExists(id uint64) (exists bool, err error) {
	if err = h.DB.QueryRow("select exists(select 1 from ralali_crud_cake_test.cakes where id = ? and deleted_at is null)", id).Scan(&exists); err != nil {
		return false, err
	}

	return exists, nil
}

func (h *CakeRepositoryHandler) GetCake(pagination uint) (data []Database.Cakes, err error) {
	var (
		result *sql.Rows
	)

	if result, err = h.DB.Query("select id, title, description, rating, image, created_at, COALESCE(updated_at, '') from ralali_crud_cake_test.cakes where deleted_at IS NULL limit 10 offset ?", (pagination-1)*10); err != nil {
		return nil, err
	}

	for result.Next() {
		var cake Database.Cakes
		err = result.Scan(&cake.Id, &cake.Title, &cake.Description, &cake.Rating, &cake.Image, &cake.CreatedAt, &cake.UpdatedAt)
		if err != nil {
			return nil, err
		}

		data = append(data, cake)
	}

	return data, nil
}

func (h *CakeRepositoryHandler) GetCakeById(id uint64) (data Database.Cakes, err error) {
	if err = h.DB.QueryRow("select id, title, description, rating, image, created_at, COALESCE(updated_at, '') from ralali_crud_cake_test.cakes where id = ?", id).
		Scan(
			&data.Id,
			&data.Title,
			&data.Description,
			&data.Rating,
			&data.Image,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
		return data, err
	}

	return data, nil
}

func (h *CakeRepositoryHandler) AddCake(data Model.CakeRequestDto) (err error) {
	if _, err = h.DB.Exec("insert into ralali_crud_cake_test.cakes (title, description, rating, image) VALUES (?,?,?,?)",
		data.Title,
		data.Description,
		data.Rating,
		data.Image,
	); err != nil {
		return err
	}
	return nil
}

func (h *CakeRepositoryHandler) UpdateCake(id uint64, data Model.CakeRequestDto) (err error) {
	if _, err = h.DB.Exec("update ralali_crud_cake_test.cakes set title = ?, description = ?, rating = ?, image = ?, updated_at = current_timestamp where id = ?",
		data.Title,
		data.Description,
		data.Rating,
		data.Image,
		id,
	); err != nil {
		return err
	}
	return nil
}

func (h *CakeRepositoryHandler) DeleteCake(id uint64) (err error) {
	if _, err = h.DB.Exec("update ralali_crud_cake_test.cakes set deleted_at = current_timestamp where id = ?", id); err != nil {
		return err
	}
	return nil
}
