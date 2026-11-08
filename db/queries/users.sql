--TABLA CLIENTE

-- name: GetUser :one
SELECT documento, nombre, apellido, deuda, nro_telefono
FROM cliente
WHERE documento = $1;

-- name: CreateUser :one
INSERT INTO cliente (documento, nombre, apellido, deuda, nro_telefono)
VALUES ($1, $2, $3, $4, $5)
RETURNING documento, nombre, apellido, deuda, nro_telefono;

-- name: UpdateUser :one
UPDATE cliente
SET nombre = $2, apellido = $3, deuda = $4, nro_telefono = $5
WHERE documento = $1
RETURNING documento, nombre, apellido, deuda, nro_telefono;

-- name: DeleteUser :exec
DELETE FROM cliente
WHERE documento = $1;

-- name: ListUsers :many
SELECT documento, nombre, apellido, deuda, nro_telefono
FROM cliente
ORDER BY documento;

--TABLA TRATAMIENTO

-- name: GetTreatment :one
SELECT id, nombre, descripcion_corta, costo
FROM tratamiento
WHERE id = $1;

-- name: CreateTreatment :one
INSERT INTO tratamiento (nombre, descripcion_corta, costo)
VALUES ($1, $2, $3)
RETURNING id, nombre, descripcion_corta, costo;

-- name: UpdateTreatment :one
UPDATE tratamiento
SET nombre = $2, descripcion_corta = $3, costo = $4
WHERE id = $1
RETURNING id, nombre, descripcion_corta, costo;

-- name: DeleteTreatment :exec
DELETE FROM tratamiento
WHERE id = $1;

-- name: ListTreatments :many
SELECT id, nombre, descripcion_corta, costo
FROM tratamiento
ORDER BY id;

--TABLA VOUCHER

-- name: GetVoucher :one
SELECT id_voucher, regalador_documento, tratamiento_id, cliente_documento
FROM voucher
WHERE id_voucher = $1;

-- name: CreateVoucher :one
INSERT INTO voucher (regalador_documento, tratamiento_id, cliente_documento)
VALUES ($1, $2, $3)
RETURNING id_voucher, regalador_documento, tratamiento_id, cliente_documento,
(SELECT CONCAT(nombre,' ',apellido) FROM cliente WHERE voucher.regalador_documento = documento) AS regalador,
(SELECT nombre FROM tratamiento WHERE voucher.tratamiento_id = id) AS tratamiento,
(SELECT CONCAT(nombre,' ',apellido) FROM cliente WHERE voucher.cliente_documento = documento) AS cliente;

-- name: UpdateVoucher :one
UPDATE voucher v
SET regalador_documento = $2, tratamiento_id = $3, cliente_documento = $4
WHERE id_voucher = $1
RETURNING id_voucher, regalador_documento, tratamiento_id, cliente_documento,
(SELECT CONCAT(nombre,' ',apellido) FROM cliente WHERE v.regalador_documento = documento) AS regalador,
(SELECT nombre FROM tratamiento WHERE v.tratamiento_id = id) AS tratamiento,
(SELECT CONCAT(nombre,' ',apellido) FROM cliente WHERE v.cliente_documento = documento) AS cliente;

-- name: DeleteVoucher :exec
DELETE FROM voucher
WHERE id_voucher = $1;

-- name: ListVouchers :many
SELECT id_voucher, regalador_documento, tratamiento_id, cliente_documento
FROM voucher
ORDER BY id_voucher;

-- name: GetVouchers :many
SELECT id_voucher, regalador_documento, cliente_documento, tratamiento_id,
(SELECT CONCAT(nombre,' ',apellido) FROM cliente WHERE v.regalador_documento = documento) AS regalador,
(SELECT nombre FROM tratamiento WHERE v.tratamiento_id = id) AS tratamiento,
(SELECT CONCAT(nombre,' ',apellido) FROM cliente WHERE v.cliente_documento = documento) AS cliente
FROM voucher v
ORDER BY id_voucher;

--TABLA TURNO

-- name: GetTurno :one
SELECT fecha, hora, tratamiento_id, cliente_documento
FROM turno
WHERE fecha = $1 AND hora = $2;

-- name: GetTurnosCompleto :many
SELECT fecha, hora, cliente_documento, tratamiento_id,
(SELECT CONCAT(nombre,' ',apellido) FROM cliente WHERE t.cliente_documento = documento) AS nombre_cliente,
(SELECT nombre FROM tratamiento WHERE t.tratamiento_id = id) AS nombre_tratamiento
FROM turno t;

-- name: CreateTurno :one
INSERT INTO turno (fecha, hora, tratamiento_id, cliente_documento)
VALUES ($1, $2, $3, $4)
RETURNING fecha, hora, cliente_documento,
(SELECT nombre FROM tratamiento WHERE id = turno.tratamiento_id) AS nombre_tratamiento,
(SELECT CONCAT(nombre,' ',apellido) FROM cliente WHERE documento = turno.cliente_documento) AS nombre_cliente;

-- name: UpdateTurno :one
UPDATE turno
SET tratamiento_id = $3, cliente_documento = $4
WHERE fecha = $1 AND hora = $2
RETURNING fecha, hora, tratamiento_id, cliente_documento,
(SELECT nombre FROM tratamiento WHERE id = turno.tratamiento_id) AS nombre_tratamiento,
(SELECT CONCAT(nombre,' ',apellido) FROM cliente WHERE documento = turno.cliente_documento) AS nombre_cliente;

-- name: DeleteTurno :exec
DELETE FROM turno
WHERE fecha = $1 AND hora = $2;

-- name: ListTurno :many
SELECT fecha, hora, tratamiento_id, cliente_documento
FROM turno
ORDER BY fecha, hora;