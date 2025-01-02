# CRM Application for Udacity Bootcamp

This is a simple CRM (Customer Relationship Management) application built using Go. The application provides basic customer management functionalities, including creating, retrieving all customers all single customer using his unique Id, updating, and deleting customer records.


## Features

- Customer Registration
- Retrieving All Customers
- Retrieving a Single Customer by ID
- Updating Customer Information
- Deleting Customer Records

## Getting Started

1. Clone the repository:
   
   git clone <https://github.com/Mohrip/CRM.git>
   cd CRM/cmd   

2. Run the application:

   go run main.go


3. Access the API at `http://localhost:8080`.

Important Note: when using POST method make sure you are sending this body without Id because it's uniquely generated:

{
  "name": " delete3312",
  "role": "XXX",
  "email": "test@example.com",
  "phone": "000001",
  "contacted": false
}



