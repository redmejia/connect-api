package postgresql

import (
	"connect/internal/models"
	"connect/internal/utils"
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

// GetAuthInfo
func (p *DbPostgres) GetAuthInfo(email string) *models.LogIn {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select 
			bus_id, 
			email, 
			password 
		from 
			login 
		where email = $1
	`
	row := p.Db.QueryRowContext(ctx, query, email)

	var business models.LogIn
	err := row.Scan(
		&business.BusinessID,
		&business.Email,
		&business.Password,
	)

	if err != nil {
		p.Error.Println(err)
		return &business
	}

	return &business

}

// UpdateProfile
func (p *DbPostgres) UpdateProfile(business *models.BusinessAccount) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		update business_account 
			set bus_name = $1, bus_type = $2,  email = $3, founded = $4
		where 
			bus_id = $5
	`

	_, err := p.Db.ExecContext(
		ctx,
		query,
		business.BusinessName,
		business.BusinessType,
		business.Email,
		business.Founded,
		business.BusinessID,
	)

	if err != nil {
		p.Error.Println(err)
		return
	}
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

// GetDealsByType all by type
func (p *DbPostgres) GetDealsByType(businessType string) *[]models.Deal {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select 
			nd.deal_id, nd.bus_id, nd.bus_type, 
			nd.pro_name, nd.pro_description, nd.created_at, nd.price,
			a.deal_id, a.bus_id, a.active, a.sold
		from
			new_deal as nd
		join 
			active as a
		on 
			nd.deal_id = a.deal_id
		where 
			nd.bus_type = $1;
		`

	var dealsType []models.Deal

	rows, err := p.Db.QueryContext(ctx, query, businessType)
	if err != nil {
		p.Error.Println(err)
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var deal models.Deal
		err := rows.Scan(
			&deal.DealID,
			&deal.BusinessID,
			&deal.BusinessType,
			&deal.ProductName,
			&deal.DealDescription,
			&deal.DealStart,
			&deal.Price,
			&deal.IsActive.DealID,
			&deal.IsActive.BusinessID,
			&deal.IsActive.DealIsActive,
			&deal.IsActive.Sold,
		)
		if err != nil {
			p.Error.Println(err)
			return nil
		}

		dealsType = append(dealsType, deal)
	}

	return &dealsType
}

// GetDelsByIDs retrive single deal by deal_id and business id  or bus_id
func (p *DbPostgres) GetDealsByIDs(dealId, businessId int) models.Deal {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		select 
			nd.deal_id, nd.bus_id, ba.bus_name, nd.bus_type, 
			nd.pro_name, nd.pro_description, nd.created_at, nd.price,
			a.deal_id, a.bus_id, a.active, a.sold
		from
			new_deal as nd
		join 
			active as a
		on 
			nd.deal_id = a.deal_id
		join 
			business_account as ba
		on
			nd.bus_id = ba.bus_id
		where 
			nd.deal_id = $1 and nd.bus_id = $2;
	`
	row := p.Db.QueryRowContext(ctx, query, dealId, businessId)

	var deal models.Deal

	err := row.Scan(
		&deal.DealID,
		&deal.BusinessID,
		&deal.BusinessName,
		&deal.BusinessType,
		&deal.ProductName,
		&deal.DealDescription,
		&deal.DealStart,
		&deal.Price,
		&deal.IsActive.DealID,
		&deal.IsActive.BusinessID,
		&deal.IsActive.DealIsActive,
		&deal.IsActive.Sold,
	)

	if err != nil {
		p.Error.Println(err)
		return deal
	}

	return deal

}

// DeleteDealByIDs
func (p *DbPostgres) DeleteDealByIDs(dealId, businessId int) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := `
		delete from active 
		where deal_id = $1 and  bus_id = $2	
	`
	_, err := p.Db.ExecContext(ctx, query, dealId, businessId)
	if err != nil {
		p.Error.Println(err)
		return false
	}

	return true
}

// UpdateDealStatus
func (p *DbPostgres) UpdateDealStatus(dealStatus *models.ActiveDeals) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := `
		update active
		set active = $1, sold = $2
		where deal_id = $3 and bus_id = $4
	`

	_, err := p.Db.ExecContext(
		ctx,
		query,
		dealStatus.DealIsActive,
		dealStatus.Sold,
		dealStatus.DealID,
		dealStatus.BusinessID,
	)

	if err != nil {
		p.Error.Println(err)
		return false
	}

	return true
}

// RegisterMyBusiness
func (p *DbPostgres) RegisterMyBusiness(newBusiness *models.BusinessAccount) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	tx, err := p.Db.BeginTx(ctx, nil)
	if err != nil {
		p.Error.Println(err)
		return false
	}
	defer tx.Rollback()

	var businessId int

	query := `
		insert into  business_account (bus_name, bus_type, email, founded, password)
		values ($1, $2, $3, $4, $5)
		returning bus_id
	`

	passwordHash, err := utils.HashPassword(newBusiness.Password)
	if err != nil {
		p.Error.Println(err)
		return false
	}

	row := tx.QueryRowContext(
		ctx,
		query,
		newBusiness.BusinessName,
		newBusiness.BusinessType,
		newBusiness.Email,
		newBusiness.Founded,
		passwordHash,
	)

	err = row.Scan(&businessId)
	if err != nil {
		p.Error.Println(err)
		return false
	}

	_, err = tx.ExecContext(ctx, `
		insert into login (bus_id, email, password)
		values ($1, $2, $3)
		`, businessId, newBusiness.Email, passwordHash,
	)

	if err != nil {
		p.Error.Println(err)
		return false
	}

	err = tx.Commit()
	if err != nil {
		p.Error.Println(err)
		return false
	}
	return true
}
