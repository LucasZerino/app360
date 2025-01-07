CREATE EXTENSION IF NOT EXISTS citext;

-- Tabela de usuários
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
    active_company_id UUID REFERENCES companies(id), -- Empresa ativa (FK opcional)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Data/hora de criação
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL, -- Data/hora da última atualização
    deleted_at TIMESTAMP -- Data/hora da exclusão lógica (soft delete)
);

-- Trigger para atualizar automaticamente o campo updated_at
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS set_updated_at_on_users ON users;
CREATE TRIGGER set_updated_at_on_users
BEFORE UPDATE ON users
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();
