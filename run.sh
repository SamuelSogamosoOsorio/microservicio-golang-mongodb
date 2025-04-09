#!/bin/bash

echo "🛑 Deteniendo contenedores previos..."
docker compose down -v --remove-orphans

echo "🐳 Construyendo imagen..."
docker compose build

echo "🚀 Levantando contenedores..."
docker compose up -d

echo "✅ Contenedores activos:"
docker ps

echo "🌐 Microservicio disponible en: http://localhost:8080"
