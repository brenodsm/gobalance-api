````markdown
# GoBalance API - Gerencie suas finanças pessoais com Golang

**GoBalance** é uma API REST escrita em **Golang** para controle de finanças pessoais.  
Oferece autenticação segura, gerenciamento de contas e transações, e organização por categorias, com foco em escalabilidade e boas práticas.

---

## ⚙️ Instalação

```bash
# Clone o repositório
git clone https://github.com/seu-usuario/gobalance-api.git
cd gobalance-api
````

```bash
# Suba o banco de dados e a API com Docker Compose
docker-compose up --build
```

> ⚠️ Certifique-se de ter o Docker e o Docker Compose instalados.

---

## 🚀 Objetivo

Desenvolver um **MVP funcional** de uma API de finanças pessoais que permita aos usuários:

* Realizar cadastro e autenticação.
* Criar e gerenciar contas financeiras (corrente, poupança, cartão de crédito, etc.).
* Registrar e consultar transações (receitas, despesas e transferências).
* Personalizar e organizar categorias financeiras.
* Visualizar saldos e obter listagens de transações com filtros flexíveis.

---

## 📌 Funcionalidades Implementadas (MVP)

### 🔐 Autenticação

* Cadastro de usuários (`nome`, `email`, `senha`).
* Login com emissão de token JWT.
* Middleware de proteção para rotas autenticadas.

### 💼 Contas Financeiras

* Criação, edição e remoção de contas:

  * Campos: `nome`, `tipo`, `saldo_inicial`.
* Restrição para exclusão de contas com transações vinculadas.

### 🗂️ Categorias

* Categorias padrão pré-definidas (Alimentação, Transporte, Salário, etc.).
* CRUD completo de categorias personalizadas:

  * Classificação por tipo (`receita`, `despesa`).

### 💸 Transações

* Registro de transações:

  * Campos: `data`, `valor`, `descricao` (opcional), `tipo` (`receita`, `despesa`, `transferencia`), `conta`, `categoria`.
* Edição e exclusão de transações.
* Transferência entre contas com débito/crédito automático.

### 📊 Consultas e Relatórios

* Consulta de saldo atual por conta.
* Listagem de transações com filtros por:

  * Conta
  * Categoria
  * Intervalo de datas

---

## 📚 Documentação da API

A documentação completa estará disponível em:
[http://localhost:8080/docs](http://localhost:8080/docs) *(em desenvolvimento)*

---

## 🛠️ Tecnologias Utilizadas

* [Go (Golang)](https://golang.org/)
* [Chi](https://github.com/go-chi/chi) — roteador HTTP leve e idiomático
* [JWT (JSON Web Token)](https://jwt.io/) — autenticação baseada em token
* [SQLX](https://github.com/jmoiron/sqlx) — acesso a banco com SQL puro
* [PostgreSQL](https://www.postgresql.org/) — banco de dados relacional
* [Docker](https://www.docker.com/) — ambiente containerizado para desenvolvimento

---

## 🔒 Considerações de Segurança e Qualidade

* **Autenticação e autorização:** uso de JWT e escopo por usuário.
* **Validações:** integridade de dados, valores positivos, datas válidas e tipos obrigatórios.
* **Isolamento de dados:** cada usuário possui acesso apenas às suas informações.
* **Consistência:** atualização automática e segura dos saldos em todas as transações.

---

## 📈 Funcionalidades Planejadas (Pós-MVP)

* Exportação de dados em formato `.csv`
* Relatórios mensais por categoria
* Recuperação de senha via e-mail
* Notificações por e-mail ou push (ex: alerta de saldo baixo)
* Suporte a múltiplas moedas
* Integração com instituições financeiras
* Dashboard com visualizações gráficas

---

## 🧠 Conceitos Aplicados

* Organização modular com foco em boas práticas
* Padrões de projeto (MVC simplificado ou Clean Architecture)
* Middlewares reutilizáveis para autenticação e controle de acesso
* Rotas RESTful com tratamento padronizado de erros
* Manutenção de consistência de dados entre entidades relacionadas

---

## 📄 Licença

Distribuído sob a licença [MIT](LICENSE)

---

## 🤝 Contribuindo

Contribuições são bem-vindas!

1. Fork este repositório
2. Criar uma branch com sua feature (`git checkout -b feature/nome`)
3. Fazer o commit das suas alterações (`git commit -m 'feat: nova funcionalidade'`)
4. Enviar para o GitHub (`git push origin feature/nome`)
5. Abrir um Pull Request

---

## 🐳 docker-compose.yml

```yaml
version: '3.8'

services:
  db:
    image: postgres:14
    container_name: gobalance-db
    environment:
      - POSTGRES_USER=admin
      - POSTGRES_PASSWORD=secret
      - POSTGRES_DB=gobalance
    volumes:
      - db_data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  api:
    build:
      context: .
      dockerfile: Dockerfile
    container_name: gobalance-api
    depends_on:
      - db
    environment:
      - DB_HOST=db
      - DB_PORT=5432
      - DB_USER=admin
      - DB_PASSWORD=secret
      - DB_NAME=gobalance
      - JWT_SECRET=your_jwt_secret
    ports:
      - "8080:8080"
    volumes:
      - .:/app

volumes:
  db_data:
```

```
```
