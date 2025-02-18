POST http://{{base_url}}/{{prefix}}/{{version}}/employee-role

{
  "Role": "STAFF"
}

DELETE http://{{base_url}}/{{prefix}}/{{version}}/employee-role/15

PUT http://{{base_url}}/{{prefix}}/{{version}}/employee-role/14

{
  "Role": "STAFF2"
}

GET http://{{base_url}}/{{prefix}}/{{version}}/employee-role/15

GET http://{{base_url}}/{{prefix}}/{{version}}/employee-role
