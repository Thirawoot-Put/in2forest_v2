POST http://{{base_url}}/{{prefix}}/{{version}}/auth/employee/register

{
  "email": "admin2@mail.com",
  "password": "1234",
  "name": "admin1",
  "mobile": "0990909090"
}

POST http://{{base_url}}/{{prefix}}/{{version}}/auth/employee/login

{
  "email": "admin2@mail.com",
  "password": "1234"
}
