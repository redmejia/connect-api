
INSERT INTO  business_account (bus_name, bus_type, email, founded, password)
VALUES
 ('Connect Inc', 'Software service', 'connect@mail.com', 1989, 'fdsafjkjweqwkelkjewqlkje');
 

 INSERT INTO active (bus_id, active, sold)
 VALUES 
 (2, TRUE, FALSE),
 (2, TRUE, FALSE);


 INSERT INTO new_deal (deal_id, bus_id, bus_type, pro_name, pro_description, created_at, price)
 VALUES 
 	(4, 1, 'agro', 'compost bag', 'I am selling thash bags', NOW(), 50.53),
 	(5, 1, 'agro', 'compost plants', ' I am selling a compost for coffee plants 2 bags', NOW(), 153.8);


