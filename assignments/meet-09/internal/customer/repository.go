package customer


import (
	"context"
	"database/sql"
)

type Repository struct{ db *sql.DB }

func NewRepository(db *sql.DB) *Repository { return &Repository{db: db} }

func (repo *Repository) GetAll(ctx context.Context) ([]Customer, error) {
	rows, err := repo.db.QueryContext(ctx, `
		SELECT 
			customer_id, 
			company_name,
			COALESCE(contact_title, ''),
			COALESCE(address, ''),
			COALESCE(city, ''),
			COALESCE(region, ''),
			COALESCE(postal_code, ''),
			COALESCE(country, ''),
			COALESCE(phone, ''),
			COALESCE(fax, '')
		FROM public.customers ORDER BY customer_id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	values := make([]Customer, 0)
	for rows.Next() {
		var value Customer
		err := rows.Scan(
			&value.CustomerID, 
			&value.CompanyName, 
			&value.ContanctTitle,
			&value.Address,
			&value.City,
			&value.Region,
			&value.PostalCode,
			&value.Country,
			&value.Phone,
			&value.Fax);
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	return values, rows.Err()
}


func (repo *Repository) GetByID(ctx context.Context, id string) (Customer, error) {
	var value Customer
	err := repo.db.QueryRowContext(ctx, `
		SELECT 
			customer_id, 
			company_name,
			COALESCE(contact_title, ''),
			COALESCE(address, ''),
			COALESCE(city, ''),
			COALESCE(region, ''),
			COALESCE(postal_code, ''),
			COALESCE(country, ''),
			COALESCE(phone, ''),
			COALESCE(fax, '')
		FROM public.customers
		WHERE customer_id=$1`, id).Scan(
			&value.CustomerID, 
			&value.CompanyName, 
			&value.ContanctTitle,
			&value.Address,
			&value.City,
			&value.Region,
			&value.PostalCode,
			&value.Country,
			&value.Phone,
			&value.Fax);
	return value, err
}

func (repo *Repository) Create(ctx context.Context, value Customer) error {
	_, err := repo.db.ExecContext(ctx, `
		INSERT INTO public.customers (
			customer_id, 
			company_name,
			contact_title,
			address, 
			city, 
			region, 
			postal_code, 
			country, 
			phone,
			fax )
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`, 
			value.CustomerID, 
			value.CompanyName, 
			value.ContanctTitle, 
			value.Address, 
			value.City, 
			value.Region, 
			value.PostalCode,
			value.Country,
			value.Phone,
			value.Fax)
	return err
}

func (repo *Repository) Update(ctx context.Context, value Customer) error {
	result, err := repo.db.ExecContext(ctx, `
		UPDATE public.customers 
		SET 
			customer_id = $1, 
			company_name = $2,
			contact_title = $3,
			address = $4, 
			city = $5, 
			region = $6, 
			postal_code = $7, 
			country = $8, 
			phone = $9,
			fax = $10 	
		WHERE 
			customer_id=$1`,
			value.CustomerID, 
			value.CompanyName, 
			value.ContanctTitle, 
			value.Address, 
			value.City, 
			value.Region, 
			value.PostalCode,
			value.Country,
			value.Phone,
			value.Fax)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (repo *Repository) Delete(ctx context.Context, id string) error {
	result, err := repo.db.ExecContext(ctx, `
		DELETE FROM public.customers 
		WHERE customer_id=$1`, id)
	if err != nil {
		return err
	}
	count, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return sql.ErrNoRows
	}
	return nil
}
 
