package db

import (
	"github.com/google/uuid"
)

type CompanyCreateTup struct {
	Id          uuid.UUID
	Name        string
	Description string
	Employees   int
	Registered  bool
	Type        string
}

func (c *CompanyDbC) CreateCompany(company *CompanyCreateTup) error {
	_, err := c.Pg.Exec(
		`INSERT INTO companies (
            id,
            name,
            description,
            employees,
            registered,
            type
        ) VALUES (
            $1,
            $2,
            $3,
            $4,
            $5,
            $6
        )`,
		company.Id,
		company.Name,
		company.Description,
		company.Employees,
		company.Registered,
		company.Type,
	)

	return err
}

func (c *CompanyDbC) DeleteCompany(id uuid.UUID) error {
	_, err := c.Pg.Exec("DELETE FROM companies WHERE id = $1", id)
	if err != nil {
		return err
	}
	return nil
}

type CompanyInfoTup struct {
	Id          uuid.UUID
	Name        string
	Description string
	Employees   int
	Registered  bool
	Type        string
}

func (c *CompanyDbC) GetCompanyById(id uuid.UUID) (*CompanyInfoTup, error) {
	company := &CompanyInfoTup{}
	err := c.Pg.QueryRow("SELECT id, name, description, employees, registered, type FROM companies WHERE id=$1", id).Scan(&company.Id, &company.Name, &company.Description, &company.Employees, &company.Registered, &company.Type)
	if err != nil {
		return nil, err
	}

	return company, nil
}

type CompanyUpdateTup struct {
	Id          uuid.UUID
	Name        string
	Description string
	Employees   int
	Registered  bool
	Type        string
}

func (c *CompanyDbC) UpdateCompany(company CompanyUpdateTup) error {
	_, err := c.Pg.Exec(
		`UPDATE companies SET
            name=$2,
            description=$3,
            employees=$4,
            registered=$5,
            type=$6
        WHERE id=$1`,
		company.Id,
		company.Name,
		company.Description,
		company.Employees,
		company.Registered,
		company.Type,
	)

	return err
}
