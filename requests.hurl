#request para tests

# CLIENTE
POST http://localhost:8080/cliente 
{
    "nombre" :"valentina",
    "apellido": "bisogni",
    "deuda": 1000,
    "nro_telefono" : 12345678
}
HTTP 201
[Asserts]
jsonpath "$.id" exists
[Captures]
cliente_id : jsonpath "$.id"


POST http://localhost:8080/cliente 
{
    "nombre" :"rocio",
    "apellido": "lopez",
    "nro_telefono": 2494
}
HTTP 201
[Asserts]
jsonpath "$.id" exists
[Captures]
cliente2_id : jsonpath "$.id"



GET http://localhost:8080/cliente/{{cliente_id}}
HTTP 200
[Asserts]
jsonpath "$.nombre" == "valentina"
jsonpath "$.apellido" == "bisogni"

GET http://localhost:8080/cliente
HTTP 200
[Asserts]
jsonpath "$" count == 2

#UPDATE 
PUT http://localhost:8080/cliente/{{cliente_id}}
{
    "nombre": "valentina",
    "apellido": "bisogni",
    "deuda" : 1000,
    "nro_telefono" : 12345678
}
HTTP 200

GET http://localhost:8080/cliente/{{cliente_id}}
HTTP 200
[Asserts]
jsonpath "$.deuda" == 1000

#Casos de Create sin información necesaria

#   Sin número de telefono
POST http://localhost:8080/cliente 
{
    "nombre" :"rocio",
    "apellido": "lopez"
}
HTTP 400

#   Sin apellido
POST http://localhost:8080/cliente 
{
    "nombre" :"rocio",
    "nro_telefono": 12345

}
HTTP 400

#   Sin nombre
POST http://localhost:8080/cliente 
{
    "apellido" :"Lopez",
    "nro_telefono": 12345

}
HTTP 400


#TRATAMIENTOS 
POST http://localhost:8080/tratamiento
{
    "nombre" : "Masaje",
    "descripcion_corta" : "masaje relajante",
    "costo" : 12000
}
HTTP 201
[Asserts]
jsonpath "$.id" exists
[Captures]
tratamiento_id : jsonpath "$.id"

GET http://localhost:8080/tratamiento
HTTP 200
[Asserts]
jsonpath "$" count == 1
jsonpath "$[0].id" == {{tratamiento_id}}

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

#create sin nombre ni costo
POST http://localhost:8080/tratamiento
{
    "descripcion_corta" : "masaje atleta"
}
HTTP 400

#create sin nombre
POST http://localhost:8080/tratamiento
{
    "descripcion_corta" : "masaje atleta",
    "costo" : 5000
}
HTTP 400

#create sin costo
POST http://localhost:8080/tratamiento
{
    "nombre" : "masaje",
    "descripcion_corta" : "masaje atleta"
}
HTTP 400


#TURNO 

POST http://localhost:8080/turno
{
    "fecha" : "2026-10-11",
    "hora" : 17,
    "tratamiento_id" : {{tratamiento_id}},
    "cliente_id" : {{cliente_id}}

}
HTTP 201
[Asserts]
jsonpath "$.fecha" == "2026-10-11"
jsonpath "$.hora" == 17
jsonpath "$.tratamiento_id" == {{tratamiento_id}}
jsonpath "$.cliente_id" == {{cliente_id}}
[Captures]
turno_fecha : jsonpath "$.fecha"
turno_hora : jsonpath "$.hora"



GET http://localhost:8080/turno
HTTP 200
[Asserts]
jsonpath "$" count == 1
jsonpath "$[0].fecha" == {{turno_fecha}}
jsonpath "$[0].hora" == {{turno_hora}}


GET http://localhost:8080/turno/2026-10-11&17
HTTP 200
[Asserts]
jsonpath "$.fecha" == "2026-10-11"
jsonpath "$.hora" == 17
jsonpath "$.tratamiento_id" == {{tratamiento_id}}
jsonpath "$.cliente_id" == {{cliente_id}}


PUT http://localhost:8080/turno/2026-10-11&17
{
    "tratamiento_id" : {{tratamiento_id}},
    "cliente_id" : {{cliente2_id}}
}
HTTP 200
[Asserts]
jsonpath "$.fecha" == "2026-10-11" 
jsonpath "$.hora" == 17
jsonpath "$.tratamiento_id" == {{tratamiento_id}} 
jsonpath "$.cliente_id" == {{cliente2_id}}

DELETE http://localhost:8080/turno/2026-10-11&17
HTTP 204

GET http://localhost:8080/turno/2026-10-11&17
HTTP 404

# Turnos creados con información inválida

#   Sin información
POST http://localhost:8080/turno
HTTP 400

#   Con hora negativa
POST http://localhost:8080/turno
{
    "fecha" : "2026-10-11",
    "hora" : -1,
    "tratamiento_id" : {{tratamiento_id}},
    "cliente_id" : {{cliente_id}}
}
HTTP 400

#   Con hora pasando los horarios permitidos
POST http://localhost:8080/turno
{
    "fecha" : "2026-10-11",
    "hora" : 21,
    "tratamiento_id" : {{tratamiento_id}},
    "cliente_id" : {{cliente_id}}
}
HTTP 400

# VOUCHER
POST http://localhost:8080/voucher
{
    "regalador_id" : {{cliente_id}},
    "tratamiento_id": {{tratamiento_id}},
    "cliente_id" : {{cliente2_id}}
}
HTTP 201
[Asserts]
jsonpath "$.id_voucher" exists
[Captures]
voucher_id : jsonpath "$.id_voucher"

GET http://localhost:8080/voucher
HTTP 200
[Asserts]
jsonpath "$" count == 1
jsonpath "$[0].id_voucher" == {{voucher_id}}

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

# elimino el tartamiento
DELETE http://localhost:8080/tratamiento/{{tratamiento_id}}
HTTP 204

GET http://localhost:8080/tratamiento/{{tratamiento_id}}
HTTP 404

#elimino cliente 1
DELETE http://localhost:8080/cliente/{{cliente_id}}
HTTP 204

GET http://localhost:8080/cliente/{{cliente_id}}
HTTP 404

#elimino cliente 2 
#elimino cliente 1
DELETE http://localhost:8080/cliente/{{cliente2_id}}
HTTP 204

GET http://localhost:8080/cliente/{{cliente2_id}}
HTTP 404