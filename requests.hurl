# Tests HTTP (HURL) - GLAM Estetica Integral
# Requiere HURL 6+ (en versiones anteriores, [Form] se llama [FormParams]).
# Supuestos: éxito = 200, validación fallida = 400, recurso inexistente = 404.
# Las respuestas son filas HTML sin tags de id: los ids de tratamiento y voucher
# se capturan de la primera <td> de la fila (columna Id).
# Documentos de prueba: 11111111 (valentina) y 22222222 (rocio).

# CLIENTE
POST http://localhost:8080/cliente
[Form]
documento: 11111111
nombre: valentina
apellido: bisogni
deuda: 1000
nro_telefono: 12345678
HTTP 200
[Asserts]
body contains "valentina"
body contains "bisogni"

POST http://localhost:8080/cliente
[Form]
documento: 22222222
nombre: rocio
apellido: lopez
nro_telefono: 2494
HTTP 200
[Asserts]
body contains "rocio"
body contains "lopez"

GET http://localhost:8080/cliente/11111111
HTTP 200
[Asserts]
body contains "valentina"
body contains "bisogni"

GET http://localhost:8080/clientes-filas
HTTP 200
[Asserts]
xpath "count(//tr)" == 2

# UPDATE
PUT http://localhost:8080/cliente/11111111
[Form]
nombre: valentina
apellido: bisogni
deuda: 1000
nro_telefono: 12345678
HTTP 200

GET http://localhost:8080/cliente/11111111
HTTP 200
[Asserts]
body contains "1000"

# Casos de Create sin información necesaria

#   Sin número de teléfono
POST http://localhost:8080/cliente
[Form]
documento: 33333333
nombre: rocio
apellido: lopez
HTTP 400

#   Sin apellido
POST http://localhost:8080/cliente
[Form]
documento: 33333333
nombre: rocio
nro_telefono: 12345
HTTP 400

#   Sin nombre
POST http://localhost:8080/cliente
[Form]
documento: 33333333
apellido: Lopez
nro_telefono: 12345
HTTP 400

#   Sin documento
POST http://localhost:8080/cliente
[Form]
nombre: rocio
apellido: lopez
nro_telefono: 12345
HTTP 400


# TRATAMIENTOS
POST http://localhost:8080/tratamiento
[Form]
nombre: Masaje
descripcion_corta: masaje relajante
costo: 12000
HTTP 200
[Captures]
tratamiento_id: xpath "normalize-space(//td[1])"
[Asserts]
body contains "Masaje"

GET http://localhost:8080/tratamiento-filas
HTTP 200
[Asserts]
xpath "count(//tr)" == 1
xpath "normalize-space(//tr[1]/td[1])" == "{{tratamiento_id}}"

# Fila de un tratamiento
GET http://localhost:8080/tratamiento/{{tratamiento_id}}
HX-Request: true
HTTP 200
[Asserts]
body contains "Masaje"
body contains "masaje relajante"
body contains "12000"

# UPDATE
PUT http://localhost:8080/tratamiento/{{tratamiento_id}}
[Form]
nombre: Masaje
descripcion_corta: masaje relajante
costo: 15000
HTTP 200

GET http://localhost:8080/tratamiento/{{tratamiento_id}}
HTTP 200
[Asserts]
body contains "15000"

# create sin nombre ni costo
POST http://localhost:8080/tratamiento
[Form]
descripcion_corta: masaje atleta
HTTP 400

# create sin nombre
POST http://localhost:8080/tratamiento
[Form]
descripcion_corta: masaje atleta
costo: 5000
HTTP 400

# create sin costo
POST http://localhost:8080/tratamiento
[Form]
nombre: masaje
descripcion_corta: masaje atleta
HTTP 400


# TURNO
POST http://localhost:8080/turno
[Form]
fecha: 2026-10-11
hora: 17
tratamiento_id: {{tratamiento_id}}
cliente_documento: 11111111
HTTP 200
[Asserts]
body contains "2026-10-11"
body contains "17"

GET http://localhost:8080/turno-filas
HTTP 200
[Asserts]
xpath "count(//tr)" == 1
body contains "2026-10-11"
body contains "17"

GET http://localhost:8080/turno/2026-10-11/17
HTTP 200
[Asserts]
body contains "2026-10-11"
body contains "17"
body contains "11111111"

PUT http://localhost:8080/turno/2026-10-11/17
[Form]
tratamiento_id: {{tratamiento_id}}
cliente_documento: 22222222
HTTP 200
[Asserts]
body contains "2026-10-11"
body contains "17"
body contains "22222222"

DELETE http://localhost:8080/turno/2026-10-11/17
HTTP 200

GET http://localhost:8080/turno/2026-10-11/17
HTTP 404

# --- Turnos con información inválida ---

# Sin información
POST http://localhost:8080/turno
HX-Request: true
HTTP 200
[Asserts]
body contains "no es válido"

# Cliente inexistente
POST http://localhost:8080/turno
[Form]
fecha: 2026-10-11
hora: -1
tratamiento_id: {{tratamiento_id}}
cliente_documento: 11111111
HTTP 400

# Hora negativa (depende de que la base tenga un CHECK)
POST http://localhost:8080/turno
[Form]
fecha: 2026-10-11
hora: 21
tratamiento_id: {{tratamiento_id}}
cliente_documento: 11111111
HTTP 400


# VOUCHER
POST http://localhost:8080/voucher
[Form]
regalador_documento: 11111111
tratamiento_id: {{tratamiento_id}}
cliente_documento: 22222222
HTTP 200
[Captures]
voucher_id: xpath "normalize-space(//td[1])"

GET http://localhost:8080/vouchers-filas
HTTP 200
[Asserts]
xpath "count(//tr)" == 1
xpath "normalize-space(//tr[1]/td[1])" == "{{voucher_id}}"

# Cliente con turno asociado: no se puede borrar
DELETE http://localhost:8080/cliente/1002
HX-Request: true
HTTP 200
[Asserts]
xpath "normalize-space(//td[1])" == "{{voucher_id}}"
body contains "11111111"
body contains "22222222"

PUT http://localhost:8080/voucher/{{voucher_id}}
[Form]
regalador_documento: 11111111
tratamiento_id: {{tratamiento_id}}
cliente_documento: 11111111
HTTP 200


# ===================== VOUCHER =====================
# AJUSTAR: no vi voucher.go, así que los nombres de campos y rutas son supuestos.

POST http://localhost:8080/voucher
HX-Request: true
[FormParams]
regalador_id: 1001
tratamiento_id: {{tratamiento_id}}
cliente_id: 1002
HTTP 200
[Captures]
voucher_id: regex "hx-delete=\"/voucher/(\\d+)\""

GET http://localhost:8080/voucher/{{voucher_id}}
HX-Request: true
HTTP 200
[Asserts]
body contains "11111111"
body not contains "22222222"

DELETE http://localhost:8080/voucher/{{voucher_id}}
HTTP 200

# ===================== LIMPIEZA =====================


# Limpieza

# elimino el tratamiento
DELETE http://localhost:8080/tratamiento/{{tratamiento_id}}
HTTP 200

GET http://localhost:8080/tratamiento/{{tratamiento_id}}
HX-Request: true
HTTP 404

# elimino cliente 1
DELETE http://localhost:8080/cliente/11111111
HTTP 200

GET http://localhost:8080/cliente/11111111
HTTP 404

# elimino cliente 2
DELETE http://localhost:8080/cliente/22222222
HTTP 200

GET http://localhost:8080/cliente/22222222
HTTP 404
