# 📈 Go Rate Limiter

Rate limiter configurável por IP ou Token de acesso, construído com Go 1.24, Redis e middleware HTTP.

## 🚀 Objetivo

Controlar o tráfego de requisições a serviços web, com limites por:

- Endereço IP
- Token de acesso (via header `API_KEY`)

As regras baseadas em token sobrepõem as de IP.

## 📦 Funcionalidades

- ⛔ Bloqueio automático quando o limite é excedido (HTTP 429)
- 🔐 Limitação por IP ou Token
- 🧠 Configuração flexível via `.env`
- 💾 Persistência via Redis (com estratégia desacoplada)
- 🧩 Middleware injetável
- ✅ Testes automatizados
- 🐳 Docker & Docker Compose

## ⚙️ Configuração

Crie um arquivo `.env` na raiz com:

```env
# Limite por IP
RATE_LIMIT_PER_IP=5
RATE_LIMIT_BLOCK_DURATION_IP=5m

# Limite por Token
RATE_LIMIT_PER_TOKEN=10
RATE_LIMIT_BLOCK_DURATION_TOKEN=5m

# Redis
REDIS_HOST=redis:6379
REDIS_PASSWORD=
REDIS_DB=0

## ▶️ Como rodar com Docker

```bash
docker-compose up --build
```

Servidor disponível em: [http://localhost:8080](http://localhost:8080)

## 🛠️ Endpoints

| Método | Rota | Descrição    |
| ------ | ---- | ------------ |
| GET    | `/`  | Health check |

Requisições excedentes retornam:

```json
{
  "error": "you have reached the maximum number of requests or actions allowed within a certain time frame"
}
```

## 💡 Estratégia de Persistência

Interface `LimiterStorage` permite troca de mecanismo (ex: Redis, memória, banco SQL).

```go
type LimiterStorage interface {
    Increment(key string, window time.Duration, maxRequests int) (int, error)
    Block(key string, duration time.Duration) error
    IsBlocked(key string) (bool, error)
}
```

## 🧪 Testes

```bash
go test ./...
```

Inclui testes de:

* Limite por IP
* Limite por Token
* Tempo de expiração
* Comportamento sob carga
* Os testes utilizam um memory_store no lugar do redis
## 📌 Requisitos

* Go 1.24+
* Docker
* Redis

## ✨ Exemplo

```http
GET / HTTP/1.1
Host: localhost:8080
API_KEY: abc123
```

* Se `abc123` tiver limite de 100 req/s e IP padrão for 10 req/s, aplica-se o limite do token.

## 📂 Organização

* `middleware/`: Lógica plugável HTTP
* `limiter/`: Regras de controle e abstração
* `config/`: Leitura do `.env`
* `handler/`: Endpoints e resposta 429
* `test/`: Cobertura automatizada