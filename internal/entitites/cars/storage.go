package cars

import (
	"effective_mobile_tech_task/logger"
	"effective_mobile_tech_task/pkg/database"
	"errors"
)

func saveCars(request []ApiResponse) error {
	var (
		tx    = database.GetDB().Begin()
		owner People
		err   error
	)

	for _, req := range request {
		var car Car
		car.RegNum = req.RegNum
		car.Mark = req.Mark
		car.Model = req.Model
		car.Year = req.Year

		owner, err = getPeopleByNameAndSurname(req.Owner.Name, req.Owner.Surname)
		if err != nil {
			tx.Rollback()
			return err
		}

		if owner.ID == 0 {
			owner.Name = req.Owner.Name
			owner.Surname = req.Owner.Surname
			err = tx.Create(&owner).Error
			if err != nil {
				logger.Error.Printf("saveCars func create people query error: %s \n", err.Error())
				return errors.New("error on saving owner data")
			}
		}

		car.OwnerID = owner.ID

		err = tx.Create(&car).Error
		if err != nil {
			logger.Error.Printf("saveCars func create car query error: %s \n", err.Error())
			return errors.New("error on saving car data")
		}
	}

	tx.Commit()

	return nil
}

func getPeopleByNameAndSurname(name, surname string) (resp People, err error) {
	db := database.GetDB()

	err = db.Find(&resp, "name = ? and surname = ?", name, surname).Error
	if err != nil {
		logger.Error.Printf("getPeopleByNameAndSurname func query error: \n", err.Error())
		return People{}, errors.New("something went wrong")
	}

	return
}

func updateCar(car *Car) error {
	err := database.GetDB().Updates(&car).Error
	if err != nil {
		logger.Error.Printf("updateCar func query error: %s", err.Error())
		return errors.New("error on updating car info")
	}
	return nil
}

func deleteCarByID(carID int) error {
	car := Car{ID: carID}
	err := database.GetDB().Delete(&car).Error
	if err != nil {
		logger.Error.Printf("deleteCarByID func query error: %s", err.Error())
		return errors.New("error on deleting car")
	}
	return nil
}

func getCars(filter GetCarsFilter) (cars []Response, totalRows int64, page, pageLimit int, err error) {
	query := database.GetDB().Table("cars c").Where("c.deleted_at is null").
		Joins(`LEFT JOIN people p on c.owner_id = p.id`)

	// countQuery to build separate query to count total rows. For some reason using one query repeats the joins!!!
	countQuery := database.GetDB().Table("cars c").Where("c.deleted_at is null").
		Joins(`LEFT JOIN people p on c.owner_id = p.id`)

	if filter.RegNum != nil {
		query = query.Where("c.reg_num = ?", *filter.RegNum)
		countQuery = countQuery.Where("c.reg_num = ?", *filter.RegNum)
	}

	if filter.Mark != nil {
		query = query.Where("c.mark = ?", *filter.Mark)
		countQuery = countQuery.Where("c.mark = ?", *filter.Mark)
	}

	if filter.Model != nil {
		query = query.Where("c.model = ?", *filter.Model)
		countQuery = countQuery.Where("c.model = ?", *filter.Model)
	}

	if filter.Year != nil {
		query = query.Where("c.year = ?", *filter.Year)
		countQuery = countQuery.Where("c.year = ?", *filter.Year)
	}

	if filter.OwnerID != nil {
		query = query.Where("c.owner_id = ?", *filter.OwnerID)
		countQuery = countQuery.Where("c.owner_id = ?", *filter.OwnerID)
	}

	//pagination
	if filter.Page != nil {
		page = *filter.Page
	} else {
		page = 1
	}
	if filter.PageLimit != nil {
		pageLimit = *filter.PageLimit
	} else {
		pageLimit = 15
	}

	var countList []int64
	err = countQuery.Select("c.id").Scan(&countList).Error
	if err != nil {
		logger.Error.Println("getCars func count err:", err.Error())
		return []Response{}, 0, 0, 0, errors.New("something went wrong")
	}
	totalRows = int64(len(countList))

	selectQuery := `c.id 					   					   as id,
					c.reg_num 				   					   as reg_num,
					c.mark 					   					   as mark,
					c.model					   					   as model,
					c.year 					   					   as year,
					c.owner_id				   					   as owner_id,
					p.surname || ' ' || p.name 					   as owner,
					to_char(c.created_at, 'DD.MM.YYYY HH24:MI:SS') as created_at`

	err = query.Select(selectQuery).Offset(pageLimit * (page - 1)).Limit(pageLimit).Scan(&cars).Error
	if err != nil {
		logger.Error.Printf("getCars func select query error: %s \n", err.Error())
		return []Response{}, 0, 0, 0, errors.New("something went wrong")
	}

	if cars == nil {
		cars = []Response{}
	}

	return
}
