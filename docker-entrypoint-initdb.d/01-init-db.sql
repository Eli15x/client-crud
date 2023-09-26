
-- CREATE TABLE
CREATE TABLE client (
    cpf VARCHAR PRIMARY KEY,
    nome VARCHAR NOT NULL,
    sobrenome VARCHAR NOT NULL,
    contato VARCHAR NOT NULL,
    endereco VARCHAR NOT NULL,
    dataNascimento DATETIME NOT NULL,
)