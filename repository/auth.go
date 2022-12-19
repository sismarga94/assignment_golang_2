package repository

import (
	"assignment2/dto"
	"database/sql"
	"fmt"
)

type AuthRepository struct {
	DB *sql.DB
}

func (s *AuthRepository) GetUser(username string) (res dto.RegisterDto, err error) {
	sql := "select username, password, firstname, lastname from users where username = '" + username + "'"
	err = s.DB.QueryRow(sql).Scan(
		&res.Username,
		&res.Password,
		&res.Firstname,
		&res.Lastname,
	)
	if err != nil {
		fmt.Printf("get user : %v\n", err.Error())
		return res, err
	}
	fmt.Printf("user : %v", res)
	return res, nil
}

func (s *AuthRepository) InsertUser(req dto.RegisterDto) (err error) {
	sql := fmt.Sprintf("insert into users (username,password,firstname,lastname) values('%v','%v','%v','%v')", req.Username, req.Password, req.Firstname, req.Lastname)
	_, err = s.DB.Exec(sql)
	if err != nil {
		return err
	}
	return nil
}
