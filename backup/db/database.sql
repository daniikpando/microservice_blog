

CREATE USER daniel WITH PASSWORD '123456';
ALTER ROLE daniel WITH SUPERUSER;

CREATE DATABASE basic_microservice WITH OWNER daniel;

\c basic_microservice daniel

CREATE TABLE articles(
  id serial primary key not null,
  title varchar(50) not null,
  content text not null,
  description varchar(500) not null,
  image varchar(1000),
  created_at timestamp default now(),
  updated_at timestamp default now()
);

INSERT INTO articles(title, content, description)
VALUES('titulo 1', 'contenido 1', 'descripcion 1');