# Setup
Para configurar um servidor kafka e iniciar o consumidor

```
docker compose up -d

cd kafka-consumer

go build ./cmd/main.go

./kafka-consumer
```

Para iniciar o produtor

```
cd kafka-producer

go build ./cmd/main.go

./kafka-producer
```

O produtor envia mensagens no formato JSON, enquanto o consumidor as decodifica e printa no terminal
