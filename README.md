# isCool GPT - API Educacional de Programação

API simples em Go que utiliza o Gemini 2.0 Flash para responder perguntas sobre programação.

## Funcionalidades

- Healthcheck endpoint
- Endpoint para enviar prompts sobre programação
- Rate limiting (10 requisições por minuto por IP)
- Assistente educacional focado em programação
- Configuração via arquivo `.env`

## Requisitos

- Go 1.21 ou superior
- API Key do Google Gemini

## Instalação

1. Clone o repositório:
```bash
git clone <seu-repo>
cd iscool-gpt
```

2. Instale as dependências:
```bash
go mod download
```

3. Configure o arquivo `.env`:
```bash
cp .env.example .env
```

4. Edite o arquivo `.env` e adicione sua API key do Gemini:
```
GEMINI_API_KEY=sua_key_aqui
PORT=8080
RATE_LIMIT_REQUESTS=10
RATE_LIMIT_WINDOW=60s
```

## Como obter a API Key do Gemini

1. Acesse [Google AI Studio](https://makersuite.google.com/app/apikey)
2. Faça login com sua conta Google
3. Clique em "Create API Key"
4. Copie a chave gerada e cole no arquivo `.env`

## Executando

```bash
go run main.go
```

O servidor irá iniciar em `http://localhost:8080`

## Endpoints

### GET /

Healthcheck endpoint que retorna o status da API.

**Resposta:**
```json
{
  "ok": true
}
```

### POST /prompt

Envia uma pergunta sobre programação para o assistente.

**Request:**
```json
{
  "prompt": "Como funciona o garbage collector em Go?"
}
```

**Resposta de sucesso:**
```json
{
  "response": "O garbage collector em Go é..."
}
```

**Resposta de erro:**
```json
{
  "error": "Mensagem de erro"
}
```

## Exemplos de uso

### Usando curl

```bash
# Healthcheck
curl http://localhost:8080/

# Enviar prompt
curl -X POST http://localhost:8080/prompt \
  -H "Content-Type: application/json" \
  -d '{"prompt": "O que são goroutines?"}'
```

### Usando httpie

```bash
# Healthcheck
http GET localhost:8080/

# Enviar prompt
http POST localhost:8080/prompt prompt="Explique o que é um ponteiro em Go"
```

## Rate Limiting

A API possui rate limiting de 10 requisições por minuto por IP no endpoint `/prompt`.

Se o limite for excedido, você receberá:
```json
{
  "error": "Rate limit exceeded. Try again later."
}
```
Status: `429 Too Many Requests`

## Escopo do Assistente

O assistente responde apenas perguntas sobre:
- Linguagens de programação
- Desenvolvimento de software
- Algoritmos e estruturas de dados
- Boas práticas de código
- Debugging e troubleshooting técnico
- Frameworks e bibliotecas
- Ferramentas de desenvolvimento
- Bancos de dados e SQL
- APIs e desenvolvimento web/mobile

Perguntas fora deste escopo receberão uma mensagem educada informando a limitação.

## Estrutura do Projeto

```
iscool-gpt/
├── .env.example           # Template de configuração
├── main.go                # Entry point da aplicação
├── internal/
│   ├── config/
│   │   └── config.go      # Carrega configurações
│   ├── handlers/
│   │   ├── health.go      # Handler GET /
│   │   └── prompt.go      # Handler POST /prompt
│   ├── middleware/
│   │   └── ratelimit.go   # Rate limiting middleware
│   └── gemini/
│       └── client.go      # Cliente Gemini
└── README.md
```

## Tecnologias Utilizadas

- Go (stdlib para HTTP)
- [google/generative-ai-go](https://github.com/google/generative-ai-go) - SDK do Gemini
- [joho/godotenv](https://github.com/joho/godotenv) - Gerenciamento de .env

## Licença

MIT
