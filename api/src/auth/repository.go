// Repository providing CRUD operations of User Account
package auth

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"scul"
	"time"
)

type AccountRepository struct {
	Db *scul.DataB
}

type Account struct {
	Id          int64
	Email       string
	CompanyName string
	Password    string
	Created     time.Time
}

const (
	ACC_INSERT                    = "INSERT INTO USER_ACCOUNT (EMAIL, COMPANY_NAME, PASSWORD, CREATED) values ($1, $2, $3, $4) RETURNING ID"
	ACC_SELECT_EMAIL              = "SELECT ID, PASSWORD, COMPANY_NAME, CREATED FROM USER_ACCOUNT WHERE EMAIL = $1"
	ACC_SELECT_EMAIL_AND_PASSWORD = "SELECT ID, COMPANY_NAME, CREATED FROM USER_ACCOUNT WHERE EMAIL = $1 AND PASSWORD = $2"
	ACC_SELECT_CREDENTIAL         = "SELECT ID, PASSWORD, COMPANY_NAME, CREATED FROM USER_ACCOUNT"
)

func (repo *AccountRepository) Create(email string, companyName string, password string) (*Account, error) {
	created := time.Now()
	var id int64
	err := repo.Db.QueryRow(ACC_INSERT, 4, email, companyName, password, created, &id)
	if err != nil {
		return nil, err
	}
	return &Account{id, email, companyName, password, created}, nil
}

func (repo *AccountRepository) FindEmail(email string) (*Account, error) {
	var created time.Time
	var company_name, password string
	var id int64
	err := repo.Db.QueryRow(ACC_SELECT_EMAIL, 1, email, &id, &password, &company_name, &created)
	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	default:
		return &Account{id, email, company_name, password, created}, nil
	}
}

func (repo *AccountRepository) FindEmailAndPassword(email string, password string) (*Account, error) {
	var created time.Time
	var company_name string
	var id int64
	err := repo.Db.QueryRow(ACC_SELECT_EMAIL_AND_PASSWORD, 2, email, password, &id, &company_name, &created)
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	default:
		return &Account{id, email, company_name, password, created}, nil
	}
}

func (acc *Account) Tostring() string {
	return fmt.Sprintf("Account [email='%v',company='%v',created='%v', password='%v...']", acc.Email, acc.CompanyName, acc.Created, acc.Password[0:1])
}
