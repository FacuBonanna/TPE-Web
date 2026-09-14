--TABLA CLIENTE

-- name: GetUser :one
SELECT id, nombre, apellido, deuda, nro_telefono
FROM cliente
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO cliente (nombre, apellido, deuda, nro_telefono)
VALUES ($1, $2, $3, $4)
RETURNING id, nombre, apellido, deuda, nro_telefono;   

-- name: UpdateUser :one
UPDATE cliente
SET nombre = $2, apellido = $3, deuda = $4, nro_telefono = $5
WHERE id = $1
RETURNING id, nombre, apellido, deuda, nro_telefono;

-- name: DeleteUser :exec
DELETE FROM cliente
WHERE id = $1;

-- name: ListUsers :many
SELECT id, nombre, apellido, deuda, nro_telefono
FROM cliente
ORDER BY id;  

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
SELECT id_voucher, regalador_id, tratamiento_id, cliente_id
FROM voucher
WHERE id_voucher = $1;

-- name: CreateVoucher :one
INSERT INTO voucher (regalador_id, tratamiento_id, cliente_id)
VALUES ($1, $2, $3)
RETURNING id_voucher, regalador_id, tratamiento_id, cliente_id;

-- name: UpdateVoucher :one
UPDATE voucher
SET regalador_id = $2, tratamiento_id = $3, cliente_id = $4
WHERE id_voucher = $1
RETURNING id_voucher, regalador_id, tratamiento_id, cliente_id;

-- name: DeleteVoucher :exec
DELETE FROM voucher
WHERE id_voucher = $1;

--TABLA turno

-- name: GetTurno :one
SELECT fecha, hora, tratamiento_id, cliente_id
FROM turno 
WHERE fecha = $1 AND hora = $2;

-- name: CreateTurno :one
INSERT INTO turno (fecha, hora, tratamiento_id, cliente_id)
VALUES ($1, $2, $3, $4)
RETURNING fecha, hora, tratamiento_id, cliente_id;

-- name: UpdateTurno :one
UPDATE turno
SET tratamiento_id = $3, cliente_id = $4
WHERE fecha = $1 AND hora = $2
RETURNING fecha, hora, tratamiento_id, cliente_id;

-- name: DeleteTurno :exec
DELETE FROM turno
WHERE fecha = $1 AND hora = $2;

-- name: ListTurno :many
SELECT fecha, hora, tratamiento_id, cliente_id
FROM turno
ORDER BY fecha, hora;   



