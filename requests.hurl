#request para tests

# CLIENTE
POST http://localhost:8080/cliente 
{
    "nombre" :"valentina",
    "apellido": "bisogni",
    "deuda": 1000,
    "nro_telefono" : 12345678
}
HTTP 200
[Asserts]
jsonpath "$.id" exists
[Captures]
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
POST http://localhost:8080/tratamiento
{
    "nombre" : "Masaje",
    "descripcion_corta" : "masaje relajante",
    "costo" : 12000
}
HTTP 200
[Asserts]
jsonpath "$.id" exists
[Captures]
tratamiento_id : jsonpath "$.id"

GET http://localhost:8080/tratamiento/{{tratamiento_id}}
HTTP 200
[Asserts]
jsonpath "$.nombre" == "Masaje"
jsonpath "$.descripcion_corta" == "masaje relajante"
jsonpath "$.costo" == 12000

#UPDATE
PUT http://localhost:8080/tratamiento/{{tratamiento_id}}
{
    "nombre" : "Masaje",
    "descripcion_corta" : "masaje relajante",
    "costo" : 15000
}
HTTP 200

GET http://localhost:8080/tratamiento/{{tratamiento_id}}
HTTP 200
[Asserts]
jsonpath "$.costo" == 15000

DELETE http://localhost:8080/tratamiento/{{tratamiento_id}}
HTTP 204

GET http://localhost:8080/tratamiento/{{tratamiento_id}}
HTTP 404

#VOUCHER-- hay que hacer un cliente y un tratamiento para que funcione bien 
# CLIENTE
POST http://localhost:8080/cliente 
{
    "nombre" :"valentina",
    "apellido": "bisogni"
}
HTTP 200
[Asserts]
jsonpath "$.id" exists
[Captures]
cliente_id : jsonpath "$.id"

#TURNO 
POST http://localhost:8080/tratamiento
{
    "nombre" : "Masaje",
    "descripcion_corta" : "masaje relajante",
    "costo" : 12000
}
HTTP 200
[Asserts]
jsonpath "$.id" exists
[Captures]
tratamiento_id : jsonpath "$.id"

POST http://localhost:8080/turno
{
    "fecha" : "2026-10-11",
    "hora" : 1,
    "tratamiento_id" : {{tratamiento_id}},
    "cliente_id" : {{cliente_id}}

}
HTTP 200
[Asserts]
jsonpath "$.fecha" == "2026-10-11"
jsonpath "$.hora" == 1
jsonpath "$.tratamiento_id" == {{tratamiento_id}}
jsonpath "$.cliente_id" == {{cliente_id}}
[Captures]
turno_fecha : jsonpath "$.fecha"
turno_hora : jsonpath "$.hora"

GET http://localhost:8080/turno
{
    "fecha" : "2026-10-11",
    "hora" : 1
}
HTTP 200
[Asserts]
jsonpath "$.fecha" == "2026-10-11"
jsonpath "$.hora" == 1
jsonpath "$.tratamiento_id" == {{tratamiento_id}}
jsonpath "$.cliente_id" == {{cliente_id}}

#INICIO OTRO CLIENTE PARA ACTUALIZAR, LO DEJO PARA LAS PRUEBAS DE VOUCHER
POST http://localhost:8080/cliente 
{
    "nombre" :"rocio",
    "apellido": "lopez"
}
HTTP 200
[Asserts]
jsonpath "$.id" exists
[Captures]
cliente2_id : jsonpath "$.id"

PUT http://localhost:8080/turno
{
    "fecha" : "2026-10-11",
    "hora" : 1,
    "tratamiento_id" : {{tratamiento_id}},
    "cliente_id" : {{cliente2_id}}
}
HTTP 200
[Asserts]
jsonpath "$.fecha" == "2026-10-11" 
jsonpath "$.hora" == 1 
jsonpath "$.tratamiento_id" == {{tratamiento_id}} 
jsonpath "$.cliente_id" == {{cliente2_id}}

DELETE http://localhost:8080/turno
{
    "fecha" : "2026-10-11",
    "hora" : 1
}
HTTP 204

GET http://localhost:8080/turno
{
    "fecha" : "2026-10-11",
    "hora" : 1
}
HTTP 404

# VOUCHER
POST http://localhost:8080/voucher
{
    "regalador_id" : {{cliente_id}},
    "tratamiento_id": {{tratamiento_id}},
    "cliente_id" : {{cliente2_id}}
}
HTTP 200
[Asserts]
jsonpath "$.id_voucher" exists
[Captures]
voucher_id : jsonpath "$.id_voucher"

GET http://localhost:8080/voucher/{{voucher_id}}
HTTP 200
[Asserts]
jsonpath "$.id_voucher" == {{voucher_id}}
jsonpath "$.regalador_id" == {{cliente_id}}
jsonpath "$.tratamiento_id" == {{tratamiento_id}}
jsonpath "$.cliente_id" == {{cliente2_id}}

PUT http://localhost:8080/voucher/{{voucher_id}}
{
    "regalador_id" : {{cliente_id}},
    "tratamiento_id" : {{tratamiento_id}},
    "cliente_id" : {{cliente_id}}
}
HTTP 200

GET http://localhost:8080/voucher/{{voucher_id}}
HTTP 200
[Asserts]
jsonpath "$.cliente_id" == {{cliente_id}}

DELETE http://localhost:8080/voucher/{{voucher_id}}
HTTP 204

GET http://localhost:8080/voucher/{{voucher_id}}
HTTP 404