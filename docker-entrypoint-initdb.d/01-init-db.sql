
-- CREATE TABLE
DROP TABLE IF EXISTS client;
CREATE TABLE client (
    cpf INT PRIMARY KEY,
    nome VARCHAR NOT NULL,
    sobrenome VARCHAR NOT NULL,
    contato VARCHAR NOT NULL,
    endereco VARCHAR NOT NULL,
    dataNascimento DATETIME NOT NULL,
);