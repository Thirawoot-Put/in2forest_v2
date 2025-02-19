POST http://{{base_url}}/{{prefix}}/{{version}}/employee-role

{
  "name": "ADMIN"
}

DELETE http://{{base_url}}/{{prefix}}/{{version}}/employee-role/4

PUT http://{{base_url}}/{{prefix}}/{{version}}/employee-role/4

{
  "name": "STAFF"
}

GET http://{{base_url}}/{{prefix}}/{{version}}/employee-role/4

GET http://{{base_url}}/{{prefix}}/{{version}}/employee-role
