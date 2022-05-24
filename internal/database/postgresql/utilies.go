package postgresql

import (
	"connect/internal/models"
	"context"
	"time"
)

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
