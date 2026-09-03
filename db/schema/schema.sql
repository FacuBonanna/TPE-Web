-- Created by Redgate Data Modeler (https://datamodeler.redgate-platform.com)
-- Last modification date: 2026-08-31 19:18:58.011

-- tables
-- Table: calendario
CREATE TABLE calendario (
    fecha date  NOT NULL,
    hora timestamp  NOT NULL,
    tratamiento_id bigserial  NOT NULL,
    cliente_id bigserial  NOT NULL,
    CONSTRAINT PK_CALENDARIO PRIMARY KEY (fecha,hora)
);

-- Table: cliente
CREATE TABLE cliente (
    id bigserial  NOT NULL,
    nombre varchar(20)  NOT NULL,
    apellido varchar(20)  NOT NULL,
    deuda decimal(10,10)  NOT NULL,
    nro_telefono int  NOT NULL,
    CONSTRAINT PK_CLIENTE PRIMARY KEY (id)
);

-- Table: tratamiento
CREATE TABLE tratamiento (
    id bigserial  NOT NULL,
    nombre varchar(20)  NOT NULL,
    descripcion_corta varchar(20)  NOT NULL,
    costo int  NOT NULL,
    CONSTRAINT PK_TRATAMIENTO PRIMARY KEY (id)
);

-- Table: voucher
CREATE TABLE voucher (
    id_voucher bigserial  NOT NULL,
    regalador_id bigserial  NOT NULL,
    tratamiento_id bigserial  NOT NULL,
    cliente_id bigserial  NOT NULL,
    CONSTRAINT PK_VOUCHER PRIMARY KEY (id_voucher)
);

-- foreign keys
-- Reference: calendario_cliente (table: calendario)
ALTER TABLE calendario ADD CONSTRAINT calendario_cliente
    FOREIGN KEY (cliente_id)
    REFERENCES cliente (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: calendario_tratamiento (table: calendario)
ALTER TABLE calendario ADD CONSTRAINT calendario_tratamiento
    FOREIGN KEY (tratamiento_id)
    REFERENCES tratamiento (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: vocuher_cliente (table: voucher)
ALTER TABLE voucher ADD CONSTRAINT vocuher_cliente
    FOREIGN KEY (regalador_id)
    REFERENCES cliente (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: vocuher_cliente (table: voucher)
ALTER TABLE voucher ADD CONSTRAINT vocuher_cliente
    FOREIGN KEY (cliente_id)
    REFERENCES cliente (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;

-- Reference: vocuher_tratamiento (table: voucher)
ALTER TABLE voucher ADD CONSTRAINT vocuher_tratamiento
    FOREIGN KEY (tratamiento_id)
    REFERENCES tratamiento (id)  
    NOT DEFERRABLE 
    INITIALLY IMMEDIATE
;



