package postgresql

import (
	"connect/internal/models"
	"context"
	"time"
)

// GetBusinessInfoById get the business information
func (p *DbPostgres) GetMyBusinessInfoById(businessId int) *models.BusinessAccount {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select 
			bus_id, bus_name, bus_type, email, founded
		from
			business_account
		where 
			bus_id = $1
		`
	row := p.Db.QueryRowContext(ctx, query, businessId)

	var business models.BusinessAccount

	err := row.Scan(
		&business.BusinessID,
		&business.BusinessName,
		&business.BusinessType,
		&business.Email,
		&business.Founded,
	)

	if err != nil {
		p.Error.Println(err)
		return &business
	}

	return &business
}

// CreateNewDeal
func (p *DbPostgres) CreateNewDeal(deal *models.Deal) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := p.Db.BeginTx(ctx, nil)
	if err != nil {
		p.Error.Fatal(err)
	}
	defer tx.Rollback()

	var dealId int

	row := tx.QueryRowContext(ctx, `
		insert into active (bus_id, active, sold)
			values ($1, $2, $3)
		returning deal_id
		`,
		deal.BusinessID,
		deal.IsActive.DealIsActive,
		deal.IsActive.Sold,
	)

	err = row.Scan(&dealId)

	query := `
		insert into new_deal
			(deal_id, bus_id, bus_type, pro_name, pro_description, created_at, price)
		values
			($1, $2, $3, $4, $5, $6, $7)
	`
	_, err = tx.ExecContext(ctx, query,
		dealId, deal.BusinessID,
		deal.BusinessType, deal.ProductName,
		deal.DealDescription, time.Now(),
		deal.Price,
	)

	if err != nil {
		p.Error.Println(err)
	}

	err = tx.Commit()
	if err != nil {
		p.Error.Println(err)
		return false
	}

	return true
}