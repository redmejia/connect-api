# connect-api
connect-api is REST-API for Connect business to business web application

# Signin and Register a business
To Singin you need email and password.\
POST /api/login

To Register a business you need fill a form with business name, business type, email, when the business were founded and the password.\
POST /api/create/account

# Get business information
Retrive business information and business deals by business id.\
GET /api/my/business?bus-id=1
Response will be json business information.\
`
{
    "myBusiness": {
        "my_business": {
            "business_id": 1,
            "business_name": "Connect",
            "business_type": "Technologie",
            "email": "connect@mail.com",
            "founded": 1991,
            "password": ""
        },
        "my_deals": [
            {
                "deal_id": 6,
                "business_id": 1,
                "business_name": "Connect",
                "business_type": "Technologie",
                "product_name": "NVIDIA GeForce RTX 3060 Ti",
                "deal_desciption": "For more information email us to connect@mail.com",
                "deal_start": "2022-06-25T21:17:55.066853Z",
                "is_active": {
                    "deal_id": 6,
                    "business_id": 1,
                    "deal_is_active": false,
                    "sold": false
                },
                "price": 455
            }
        ]
    }
}
`

# Update deal
Update deal business id is require you can update one or more business information.\
PATCH /api/my/business
Response will be the business information json format.\
`
{
    "myBusiness": {
        "my_business": {
            "business_id": 1,
            "business_name": "Connect",
            "business_type": "Technologie",
            "email": "connect@mail.com",
            "founded": 1989,
            "password": ""
        }
}
`

# Create deal
Creating new deal business id is require. The form also must have information of the deal such as business type, product name, deal descriptio (short description) and the price.\
POST /api/my/business
Reponse json format with the created date(deal start).\
`
{
    "myDeal": {
        "deal_id": 7,
        "business_id": 1,
        "business_name": "",
        "business_type": "Technologie",
        "product_name": "Printer",
        "deal_desciption": "more info call to 444.4444.5553",
        "deal_start": "2022-06-25T22:22:04.823065332-07:00",
        "is_active": {
            "deal_id": 0,
            "business_id": 0,
            "deal_is_active": false,
            "sold": false
        },
        "price": 100.08
    }
}
`
# Get deal by business type
Retrive business by type for now there are three business type such as Technologie, Food and Drink and Agriculture.\
GET /api/my/business/deals?type=Technologie
Reponse with json format for all the business of the same type on the example Technologie that was pass on the url query.\
`
{
    "deals": [
        {
            "deal_id": 7,
            "business_id": 1,
            "business_name": "Connect",
            "business_type": "Technologie",
            "product_name": "Printer",
            "deal_desciption": "more info call to 444.4444.5553",
            "deal_start": "2022-06-25T22:22:04.818353Z",
            "is_active": {
                "deal_id": 7,
                "business_id": 1,
                "deal_is_active": false,
                "sold": false
            },
            "price": 100.08
        },
        {
            "deal_id": 6,
            "business_id": 1,
            "business_name": "Connect",
            "business_type": "Technologie",
            "product_name": "NVIDIA GeForce RTX 3060 Ti",
            "deal_desciption": "For more information email us to connect@mail.com",
            "deal_start": "2022-06-25T21:17:55.066853Z",
            "is_active": {
                "deal_id": 6,
                "business_id": 1,
                "deal_is_active": false,
                "sold": false
            },
            "price": 455
        }
    ]
}
`

# Delete deal
For deletin a deal a deal id is require.\
DELETE /api/my/business/del/deal?deal-id=8
\
Reponse 200 OK.



