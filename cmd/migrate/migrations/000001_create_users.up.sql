CREATE EXTENSION IF NOT EXISTS citext;

CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY, -- UUID único para identificação do usuário
    name VARCHAR(255) NOT NULL, -- Nome do usuário
    email CITEXT UNIQUE NOT NULL, -- E-mail tratado como case-insensitive
    password BYTEA NOT NULL, -- Senha do usuário (hasheada)
    image_path TEXT, -- Caminho para a foto de perfil (opcional)
    last_login_at TIMESTAMP, -- Última data/hora de login (opcional)
    is_verified BOOLEAN DEFAULT FALSE NOT NULL, -- Indica se o e-mail foi confirmado
    auth_token TEXT, -- Token de autenticação (para Magic Link)
    auth_token_expiry TIMESTAMP, -- Validade do token de autenticação
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Data/hora de criação
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Data/hora da última atualização
    deleted_at TIMESTAMP -- Data/hora da exclusão lógica (soft delete)
);