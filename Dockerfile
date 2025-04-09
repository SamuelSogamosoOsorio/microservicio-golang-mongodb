# ---------- Etapa 1: Build ----------
    FROM golang:1.22 AS builder

    WORKDIR /app
    
    # Copiar solo go.mod y go.sum primero para aprovechar caché
    COPY go.mod go.sum ./
    RUN go mod download
    
    # Copiar el resto del código
    COPY . .
    
    # Compilar la app
    RUN CGO_ENABLED=0 GOOS=linux go build -o main .
    
    # ---------- Etapa 2: Run ----------
    FROM alpine:latest
    
    WORKDIR /root/
    
    # Copiar el binario desde la etapa anterior
    COPY --from=builder /app/main .
    
    # Ejecutar el binario
    CMD ["./main"]