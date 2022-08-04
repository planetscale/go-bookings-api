package data

import (
	"github.com/pkg/errors"
)

type Hotel struct {
	Id          int64
	Name        string
	Address     string
	Stars       float32
	Description *string
}

func FetchHotels() ([]Hotel, error) {
	conn, err := GetDbConnection()
	if err != nil {
		return nil, errors.Wrap(err, "(FetchHotels) GetConnection")
	}

	query := "SELECT * FROM hotels"
	rows, err := conn.Query(query)
	if err != nil {
		return nil, errors.Wrap(err, "(FetchHotels) db.Query")
	}

	hotels := []Hotel{}

	for rows.Next() {
		var hotel Hotel
		err = rows.Scan(&hotel.Id, &hotel.Name, &hotel.Address, &hotel.Stars, &hotel.Description)
		if err != nil {
			return nil, errors.Wrap(err, "(FetchHotels) rows.Scan")
		}
		hotels = append(hotels, hotel)
	}

	return hotels, nil
}

func CreateHotel(hotel *Hotel) error {
	conn, err := GetDbConnection()
	if err != nil {
		return errors.Wrap(err, "(CreateHotel) GetConnection")
	}

	query := "INSERT INTO hotels (name, address, stars) VALUES (?, ?, ?)"

	result, err := conn.Exec(query, hotel.Name, hotel.Address, hotel.Stars)
	if err != nil {
		return errors.Wrap(err, "(CreateHotel) conn.Exec")
	}

	id, err := result.LastInsertId()
	if err != nil {
		return errors.Wrap(err, "(CreateHotel) result.LastInsertId")
	}

	hotel.Id = id
	return nil
}

func UpdateHotel(hotelId int64, hotel *Hotel) error {
	conn, err := GetDbConnection()
	if err != nil {
		return errors.Wrap(err, "(UpdateHotel) GetConnection")
	}

	query := "UPDATE hotels SET name = ?, address = ?, stars = ? WHERE id = ?"

	_, err = conn.Exec(query, hotel.Name, hotel.Address, hotel.Stars, hotelId)
	if err != nil {
		return errors.Wrap(err, "(UpdateHotel) conn.Exec")
	}

	return nil
}

func DeleteHotel(hotelId int64) error {
	conn, err := GetDbConnection()
	if err != nil {
		return errors.Wrap(err, "(DeleteHotel) GetConnection")
	}

	query := "DELETE FROM hotels WHERE id = ?"

	_, err = conn.Exec(query, hotelId)
	if err != nil {
		return errors.Wrap(err, "(DeleteHotel) conn.Exec")
	}

	return nil
}
