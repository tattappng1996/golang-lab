package repository

import (
	"encoding/json"
	"fmt"
	"golang-lab/dto"
	"io/ioutil"
	"os"
)

// ListUsers ..
func (r *repository) ListUsers() (dto.TableUsers, error) {
	byteValue, err := ioutil.ReadAll(r.db)
	if err != nil {
		fmt.Println(err)
	}

	tableUsers := dto.TableUsers{}
	if err := json.Unmarshal(byteValue, &tableUsers); err != nil {
		fmt.Println(err)
	}

	return tableUsers, nil
}

// CreateUser ..
func (r *repository) CreateUser(user dto.User) error {
	byteValue, err := ioutil.ReadAll(r.db)
	if err != nil {
		return err
	}

	tableUsers := dto.TableUsers{}
	if err := json.Unmarshal(byteValue, &tableUsers); err != nil {
		return err
	}

	tableUsers.Users = append(tableUsers.Users, user)

	file, err := json.MarshalIndent(tableUsers, "", " ")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile("save.json", file, os.ModePerm); err != nil {
		return err
	}

	return nil
}
