CREATE TABLE public.usuario (
	id serial NOT NULL,
	nombre varchar NOT NULL,
	apellido varchar,
	CONSTRAINT pk_usuario PRIMARY KEY (id)

);
ALTER TABLE public.usuario OWNER TO postgres;


CREATE TABLE public.rol (
	id serial NOT NULL,
	aplicacion varchar NOT NULL,
	usuario_id integer NOT NULL,
	CONSTRAINT pk_rol PRIMARY KEY (id)

);
ALTER TABLE public.rol OWNER TO postgres;

ALTER TABLE public.rol ADD CONSTRAINT fk_rol_usuario FOREIGN KEY (usuario_id)
REFERENCES public.usuario (id) MATCH FULL
ON DELETE NO ACTION ON UPDATE NO ACTION;