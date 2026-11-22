# Tests HTTP (HURL) - GLAM Estetica Integral
# Requiere HURL 6+ (en versiones anteriores, [Form] se llama [FormParams]).
# Supuestos: éxito = 200, validación fallida = 400, recurso inexistente = 404.
# Las respuestas son filas HTML sin tags de id: los ids de tratamiento y voucher
# se capturan de la primera <td> de la fila (columna Id).
# cliente_documentos de prueba: 11111111 (valentina) y 22222222 (rocio).

# CLIENTE
POST http://localhost:8080/cliente
[Form]
cliente_documento: 11111111
name: valentina
surname: bisogni
debt: 1000
phone: 12345678
HTTP 200
[Asserts]
body contains "valentina"
body contains "bisogni"

POST http://localhost:8080/cliente
[Form]
cliente_documento: 22222222
name: rocio
surname: lopez
debt: 0
phone: 2494
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
body contains "valentina"
body contains "rocio"

# UPDATE
PUT http://localhost:8080/cliente/11111111
[Form]
name: valentina
surname: bisogni
debt: 1000
phone: 12345678
HTTP 200

GET http://localhost:8080/cliente/11111111
HTTP 200
[Asserts]
body contains "1000"

# Casos de Create sin información necesaria

#   Sin número de teléfono
POST http://localhost:8080/cliente
[Form]
cliente_documento: 33333333
name: rocio
surname: lopez
debt: 0
HTTP 200
[Asserts]
body contains "El teléfono debe ser un número entero válido"

#   Sin surname
POST http://localhost:8080/cliente
[Form]
cliente_documento: 33333333
name: rocio
debt: 0
phone: 12345
HTTP 200
[Asserts]
body contains "Faltan campos por completar"

#   Sin name
POST http://localhost:8080/cliente
[Form]
cliente_documento: 33333333
surname: Lopez
debt: 0
phone: 12345
HTTP 200
[Asserts]
body contains "Faltan campos por completar"

#   Sin cliente_documento
POST http://localhost:8080/cliente
[Form]
name: rocio
surname: lopez
debt: 0
phone: 12345
HTTP 200
[Asserts]
body contains "El documento debe ser un número entero válido."

#   Sin deuda
POST http://localhost:8080/cliente
[Form]
cliente_documento: 33333333
name: rocio
surname: lopez
phone: 12345
HTTP 200
[Asserts]
body contains "La deuda debe ser un número entero válido"


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


# create sin name
POST http://localhost:8080/tratamiento
[Form]
descripcion_corta: masaje atleta
costo: 5000
HTTP 200
[Asserts]
body contains "Faltan campos por completar."

# create sin costo
POST http://localhost:8080/tratamiento
[Form]
nombre: masaje
descripcion_corta: masaje atleta
HTTP 200
[Asserts]
body contains "El costo debe ser un número válido"

# create sin descripcion
POST http://localhost:8080/tratamiento
[Form]
nombre: masaje
costo: 5000
HTTP 200
[Asserts]
body contains "Faltan campos por completar."


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
body contains "Masaje"
body contains "valentina"

PUT http://localhost:8080/turno/2026-10-11/17
[Form]
tratamiento_id: {{tratamiento_id}}
cliente_documento: 22222222
HTTP 200
[Asserts]
body contains "2026-10-11"
body contains "Masaje"
body contains "rocio"

DELETE http://localhost:8080/turno/2026-10-11/17
HTTP 200

GET http://localhost:8080/turno/2026-10-11/17
HTTP 200
[Asserts]
body contains "no existe"
body not contains "<tr"

# --- Turnos con información inválida ---

# Sin información
POST http://localhost:8080/turno
HX-Request: true
HTTP 200
[Asserts]
body contains "mensaje-error"
body not contains "<tr"

# Cliente inexistente
POST http://localhost:8080/turno
[Form]
fecha: 2026-10-11
hora: 17
tratamiento_id: {{tratamiento_id}}
cliente_documento: 99999999
HTTP 200
[Asserts]
body contains "mensaje-error"
body not contains "<tr"

# tratamiento inexistente
POST http://localhost:8080/turno
[Form]
fecha: 2026-10-11
hora: 17
cliente_documento: 22222222
tratamiento_id: -1
cliente_documento: 99999999
HTTP 200
[Asserts]
body contains "mensaje-error"
body not contains "<tr"


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
xpath "count(//tr)" == 1.0
xpath "normalize-space(//tr[1]/td[1])" == "{{voucher_id}}"

# Cliente con voucher asociado: no se puede borrar
DELETE http://localhost:8080/cliente/22222222
HTTP 200
[Asserts]
body contains "No se puede eliminar"

PUT http://localhost:8080/voucher/{{voucher_id}}
[Form]
regalador_documento: 11111111
tratamiento_id: {{tratamiento_id}}
cliente_documento: 11111111
HTTP 200

DELETE http://localhost:8080/voucher/{{voucher_id}}
HTTP 200

# ===================== LIMPIEZA =====================
# Orden: voucher (ya borrado), tratamiento, clientes (por las claves foráneas)

DELETE http://localhost:8080/tratamiento/{{tratamiento_id}}
HX-Request: true
HTTP 200

GET http://localhost:8080/tratamiento/{{tratamiento_id}}
HX-Request: true
HTTP 200
[Asserts]
body contains "Tratamiento no encontrado"
body not contains "<tr"

DELETE http://localhost:8080/cliente/11111111
HX-Request: true
HTTP 200

GET http://localhost:8080/cliente/11111111
HX-Request: true
HTTP 200
[Asserts]
body contains "Cliente no encontrado"
body not contains "<tr"

DELETE http://localhost:8080/cliente/22222222
HX-Request: true
HTTP 200

GET http://localhost:8080/cliente/22222222
HX-Request: true
HTTP 200
[Asserts]
body contains "Cliente no encontrado"
body not contains "<tr"