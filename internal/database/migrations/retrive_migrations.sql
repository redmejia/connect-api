-- SELECT 
-- 	nd.deal_id, nd.bus_id, nd.bus_type, nd.pro_name, nd.pro_description, nd.created_at, nd.price,
-- 	a.deal_id, a.bus_id, a.active, a.sold
-- FROM 
-- 	new_deal AS nd
-- JOIN 
-- 	active AS a
-- ON 
-- 	nd.deal_id = a.deal_id
-- WHERE 
-- 	nd.bus_type = 'food & drinks';

-- SELECT 
-- 	nd.deal_id, nd.bus_id, ba.bus_name, nd.bus_type, nd.pro_name, nd.pro_description, nd.created_at, nd.price,
-- 	a.deal_id, a.bus_id, a.active, a.sold
-- FROM 
-- 	new_deal AS nd
-- JOIN 
-- 	active AS a
-- ON 
-- 	nd.deal_id = a.deal_id
-- JOIN
-- 	business_account AS ba
-- ON 
-- 	nd.bus_id = ba.bus_id
-- WHERE 
-- 	nd.deal_id = 17 and nd.bus_id = 3;

-- DELETE FROM active WHERE deal_id = 13 AND  bus_id = 1;
-- UPDATE active
-- SET active = true
-- WHERE deal_id = 17 AND bus_id = 3;

-- SELECT 
-- 	l.bus_id,
-- 	ba.bus_name,
-- 	l.email,
-- 	l.password
-- FROM
-- 	login AS l
-- JOIN 
-- 	business_account AS ba
-- ON
-- 	l.bus_id = ba.bus_id
-- WHERE l.email = 'connect@mail.com';


-- SELECT 
-- 	nd.deal_id, nd.bus_id, ba.bus_name, nd.bus_type, nd.pro_name, nd.pro_description, nd.created_at, nd.price,
-- 	a.deal_id, a.bus_id, a.active, a.sold
-- FROM 
-- 	new_deal AS nd
-- JOIN 
-- 	active AS a
-- ON 
-- 	nd.deal_id = a.deal_id
-- JOIN
-- 	business_account AS ba
-- ON 
-- 	nd.bus_id = ba.bus_id
-- WHERE 
-- 	nd.bus_id = 1;

-- DELETE FROM active WHERE deal_id = 13 AND  bus_id = 1;
-- UPDATE active
-- SET active = true
-- WHERE deal_id = 17 AND bus_id = 3;

-- UPDATE new_deal 
-- SET pro_name = 'SSD', pro_description = 'for more info go to https://google.com', price = 12.53
-- WHERE deal_id = 6 AND bus_id = 2;

SELECT
	ba.bus_id, ba.bus_name, ba.bus_type, ba.email, ba.founded,
	nd.deal_id, nd.bus_id, ba.bus_name, nd.bus_type, nd.pro_name, nd.pro_description, nd.created_at, nd.price
	-- a.deal_id, a.bus_id, a.active, a.sold
FROM business_account AS ba
JOIN
	new_deal AS nd
ON	
	ba.bus_id = nd.bus_id
-- JOIN
-- 	active AS a
-- ON 
	-- ba.bus_id = a.deal_id
WHERE ba.bus_id = 1
GROUP BY 
   ba.bus_id, nd.bus_id, nd.deal_id, nd.bus_type, nd.pro_name, nd.pro_description, nd.created_at, nd.price
