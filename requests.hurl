# ===================== CLIENTE =====================

# Crear cliente 1
POST http://localhost:8080/cliente
HX-Request: true
[FormParams]
cliente_documento: 1001
name: valentina
surname: bisogni
debt: 1000
phone: 12345678
HTTP 200
[Asserts]
body contains "valentina"
body contains "bisogni"

# Crear cliente 2
POST http://localhost:8080/cliente
HX-Request: true
[FormParams]
cliente_documento: 1002
name: rocio
surname: lopez
debt: 0
phone: 2494
HTTP 200
[Asserts]
body contains "rocio"

# Buscar cliente
GET http://localhost:8080/search-client
HX-Request: true
[QueryStringParams]
cliente_documento: 1001
HTTP 200
[Asserts]
body contains "valentina"
body not contains "rocio"

# Estado vacío: buscar uno que no existe
GET http://localhost:8080/search-client
HX-Request: true
[QueryStringParams]
cliente_documento: 999999
HTTP 200
[Asserts]
body contains "No se encontraron clientes"

# Listar todos (ajustá la ruta a la que uses en el botón "Mostrar clientes")
GET http://localhost:8080/clientes-filas
HX-Request: true
HTTP 200
[Asserts]
body contains "valentina"
body contains "rocio"

# Fila de un cliente
GET http://localhost:8080/cliente/1001
HX-Request: true
HTTP 200
[Asserts]
body contains "bisogni"

# Fila en modo edición
GET http://localhost:8080/cliente/1001/edit
HX-Request: true
HTTP 200
[Asserts]
body contains "editing"

# Actualizar
PUT http://localhost:8080/cliente/1001
HX-Request: true
[FormParams]
name: valentina
surname: bisogni
debt: 2000
phone: 12345678
HTTP 200
[Asserts]
body contains "2000"

# --- Crear cliente con datos inválidos ---

# Sin teléfono
POST http://localhost:8080/cliente
HX-Request: true
[FormParams]
cliente_documento: 1003
name: rocio
surname: lopez
debt: 0
HTTP 200
[Asserts]
body contains "Faltan campos por completar"

# Sin apellido
POST http://localhost:8080/cliente
HX-Request: true
[FormParams]
cliente_documento: 1003
name: rocio
debt: 0
phone: 12345
HTTP 200
[Asserts]
body contains "Faltan campos por completar"

# Sin nombre
POST http://localhost:8080/cliente
HX-Request: true
[FormParams]
cliente_documento: 1003
surname: lopez
debt: 0
phone: 12345
HTTP 200
[Asserts]
body contains "Faltan campos por completar"

# Valor no numérico
POST http://localhost:8080/cliente
HX-Request: true
[FormParams]
cliente_documento: abc
name: rocio
surname: lopez
debt: 0
phone: 12345
HTTP 200
[Asserts]
body contains "no es válido"


# ===================== TRATAMIENTO =====================

# Crear (la tabla devuelta tiene una sola fila; capturamos su id del botón Borrar)
POST http://localhost:8080/tratamiento
HX-Request: true
[FormParams]
nombre: Masaje
descripcion_corta: masaje relajante
costo: 12000
HTTP 200
[Captures]
tratamiento_id: regex "hx-delete=\"/tratamiento/(\\d+)\""
[Asserts]
body contains "Masaje"
body contains "masaje relajante"

# Fila de un tratamiento
GET http://localhost:8080/tratamiento/{{tratamiento_id}}
HX-Request: true
HTTP 200
[Asserts]
body contains "Masaje"
body contains "12000"

# Actualizar
PUT http://localhost:8080/tratamiento/{{tratamiento_id}}
HX-Request: true
[FormParams]
nombre: Masaje
descripcion_corta: masaje relajante
costo: 15000
HTTP 200
[Asserts]
body contains "15000"

# --- Crear tratamiento con datos inválidos ---

# Sin nombre ni costo
POST http://localhost:8080/tratamiento
HX-Request: true
[FormParams]
descripcion_corta: masaje atleta
HTTP 200
[Asserts]
body contains "Faltan campos por completar"

# Sin nombre
POST http://localhost:8080/tratamiento
HX-Request: true
[FormParams]
descripcion_corta: masaje atleta
costo: 5000
HTTP 200
[Asserts]
body contains "Faltan campos por completar"

# Sin costo
POST http://localhost:8080/tratamiento
HX-Request: true
[FormParams]
nombre: masaje
descripcion_corta: masaje atleta
HTTP 200
[Asserts]
body contains "Faltan campos por completar"


# ===================== TURNO =====================

POST http://localhost:8080/turno
HX-Request: true
[FormParams]
fecha: 2026-10-11
hora: 17
tratamiento_id: {{tratamiento_id}}
cliente_documento: 1001
HTTP 200
[Asserts]
body contains "Masaje"
body contains "valentina"

# Fila del turno
GET http://localhost:8080/turno/2026-10-11/17
HX-Request: true
HTTP 200
[Asserts]
body contains "Masaje"

# Buscar turno
GET http://localhost:8080/search-turno
HX-Request: true
[QueryStringParams]
fecha: 2026-10-11
hora: 17
HTTP 200
[Asserts]
body contains "Masaje"

# Cambiar el turno al cliente 2
PUT http://localhost:8080/turno/2026-10-11/17
HX-Request: true
[FormParams]
tratamiento_id: {{tratamiento_id}}
cliente_documento: 1002
HTTP 200
[Asserts]
body contains "rocio"

# --- Turnos con información inválida ---

# Sin información
POST http://localhost:8080/turno
HX-Request: true
HTTP 200
[Asserts]
body contains "no es válido"

# Cliente inexistente
POST http://localhost:8080/turno
HX-Request: true
[FormParams]
fecha: 2026-10-12
hora: 17
tratamiento_id: {{tratamiento_id}}
cliente_documento: 999999
HTTP 200
[Asserts]
body contains "no existe"

# Hora negativa (depende de que la base tenga un CHECK)
POST http://localhost:8080/turno
HX-Request: true
[FormParams]
fecha: 2026-10-11
hora: -1
tratamiento_id: {{tratamiento_id}}
cliente_documento: 1001
HTTP 200
[Asserts]
body contains "No se pudo crear el turno"

# Hora fuera de horario permitido (idem)
POST http://localhost:8080/turno
HX-Request: true
[FormParams]
fecha: 2026-10-11
hora: 21
tratamiento_id: {{tratamiento_id}}
cliente_documento: 1001
HTTP 200
[Asserts]
body contains "No se pudo crear el turno"

# Cliente con turno asociado: no se puede borrar
DELETE http://localhost:8080/cliente/1002
HX-Request: true
HTTP 200
[Asserts]
body contains "No se puede eliminar"

# Borrar turno
DELETE http://localhost:8080/turno/2026-10-11/17
HX-Request: true
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

DELETE http://localhost:8080/voucher/{{voucher_id}}
HX-Request: true
HTTP 200


# ===================== LIMPIEZA =====================

DELETE http://localhost:8080/tratamiento/{{tratamiento_id}}
HX-Request: true
HTTP 200

GET http://localhost:8080/tratamiento/{{tratamiento_id}}
HX-Request: true
HTTP 404

DELETE http://localhost:8080/cliente/1001
HX-Request: true
HTTP 200

GET http://localhost:8080/cliente/1001
HX-Request: true
HTTP 404

DELETE http://localhost:8080/cliente/1002
HX-Request: true
HTTP 200

GET http://localhost:8080/cliente/1002
HX-Request: true
HTTP 404