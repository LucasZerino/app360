# Bolt 360 App

## Descrição
O Bolt 360 é uma aplicação web completa para gerenciamento de negócios, desenvolvida com Next.js no frontend e Go no backend. A aplicação oferece uma interface moderna e responsiva para gerenciar vendas, estoque, clientes e muito mais.

## Tecnologias Principais

### Frontend
- Next.js 15.1
- Material UI 6.3
- TypeScript
- Redux Toolkit
- React Hook Form
- i18n para internacionalização

### Backend
- Go
- PostgreSQL
- JWT para autenticação
- Migrations para controle do banco de dados

## Funcionalidades Principais

- Sistema completo de autenticação (Login, Registro, Recuperação de senha)
- Internacionalização (Português, Inglês, Árabe, Chinês)
- Dashboard administrativo
- Gerenciamento de usuários
- Interface responsiva e moderna
- Temas claro/escuro
- Proteção de rotas

## Estrutura do Projeto

├── web/ # Frontend (Next.js)
│ ├── src/
│ │ ├── app/ # Páginas e rotas
│ │ ├── components/ # Componentes reutilizáveis
│ │ ├── sections/ # Seções específicas
│ │ ├── layouts/ # Layouts da aplicação
│ │ └── locales/ # Arquivos de tradução
│ └── public/ # Arquivos estáticos
│
└── cmd/ # Backend (Go)
├── api/ # API principal
└── migrate/ # Migrações do banco de dados


## Configuração do Ambiente

### Pré-requisitos
- Node.js 18+
- Go 1.20+
- PostgreSQL 14+

### Instalação

1. Clone o repositório
2. Instale as dependências do frontend com `npm install`
3. Instale as dependências do backend com `go mod download`
4. Configure o arquivo `.env` com as informações do seu ambiente
5. Execute o comando `make migrate-up` para criar as tabelas no banco de dados
6. Execute o comando `make dev` para iniciar o servidor de desenvolvimento
7. Acesse a aplicação no navegador com `http://localhost:3000`

### Comandos
- `make migrate-up`: Cria as tabelas no banco de dados
- `make migrate-down`: Remove as tabelas do banco de dados
- `make migrate-create`: Cria uma nova migração
- `make dev`: Inicia o servidor de desenvolvimento
- `make migrate-create`: Cria uma nova migração
- `make migrate-up`: Aplica as migrações pendentes
- `make migrate-down`: Remove as migrações aplicadas
- `make migrate-force`: Força a versão de uma migração específica
