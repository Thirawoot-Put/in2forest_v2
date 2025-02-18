POST http://{{base_url}}/{{prefix}}/{{version}}/employee-role

{
  "Role": "test3"
}

DELETE http://{{base_url}}/{{prefix}}/{{version}}/employee-role/12

PUT http://{{base_url}}/{{prefix}}/{{version}}/employee-role/99

{
  "Role": "STAFF"
}

GET http://{{base_url}}/{{prefix}}/{{version}}/employee-role/99
