CREATE TABLE pessoas (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL
);

CREATE TABLE restaurantes (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(100) NOT NULL
);

CREATE TABLE pedidos (
    id SERIAL PRIMARY KEY,
    pessoa_id INT REFERENCES pessoas(id),
    restaurante_id INT REFERENCES restaurantes(id),
    descricao TEXT
);

