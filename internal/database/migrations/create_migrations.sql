CREATE TABLE IF NOT EXISTS business_account (
	bus_id SMALLSERIAL PRIMARY KEY,
	bus_name TEXT NOT NULL,
	bus_type TEXT NOT NULL,
	email TEXT NOT NULL,
	founded SMALLINT NOT NULL,
	password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS login (
	bus_id SMALLINT NOT NULL,
	email TEXT NOT NULL,
	password TEXT NOT NULL,
	CONSTRAINT fk_business_accounts FOREIGN KEY(bus_id)
		REFERENCES business_account(bus_id) ON UPDATE CASCADE ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS active (
	deal_id SMALLSERIAL PRIMARY KEY,
	bus_id SMALLINT NOT NULL,
	active BOOLEAN NOT NULL, -- change to is_active
	sold BOOLEAN NOT NULL
);

CREATE TABLE IF NOT EXISTS new_deal (
	deal_id SMALLINT NOT NULL,
	bus_id SMALLINT NOT NULL,
	bus_type TEXT NOT NULL,
	pro_name TEXT NOT NULL,
	pro_description TEXT NOT NULL,
	created_at TIMESTAMP, 
	price NUMERIC(5, 2),
	CONSTRAINT fk_actives FOREIGN KEY(deal_id) 
		REFERENCES active(deal_id) ON UPDATE CASCADE ON DELETE CASCADE
);

