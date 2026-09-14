#request para tests

# CLIENTE
POST http://localhost:8080/cliente 
{
    "nombre" :"valentina",
    apellido": "bisogni"
}
HTTP 201
[Asserts]
jsonpath "$.id" exists
[captures]
cliente_id : jsonpath "$.id"

GET http://localhost:8080/cliente/{{cliente_id}}
HTTP 200
[Asserts]
jsonpath "$.nombre" == "valentina"
jsonpath "$.apellido" == "bisogni"

#UPDATE 
PUT http://localhost:8080/cliente/{{cliente_id}}
{
    "nombre": "valentina",
    "apellido": "bisogni",
    "deuda" : 1000
}
HTTP 200

GET http://localhost:8080/cliente/{{cliente_id}}
HTTP 200
[Asserts]
jsonpath "$.deuda" == 1000

DELETE http://localhost:8080/cliente/{{cliente_id}}
HTTP 204

GET http://localhost:8080/cliente/{{cliente_id}}
HTTP 404

#TRATAMIENTOS 


