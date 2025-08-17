#!/bin/bash

echo "🌱 Executando seed de dados do CRM Contatos..."

if [ ! -f "go.mod" ]; then
    echo "❌ Erro: Execute este script a partir da raiz do projeto"
    exit 1
fi

echo "📦 Compilando script de seed..."
go build -o seed scripts/seed.go

if [ $? -eq 0 ]; then
    echo "🚀 Executando seed..."
    ./seed
    
    echo "🧹 Limpando arquivos temporários..."
    rm -f seed
    
    echo "✅ Seed executado com sucesso!"
else
    echo "❌ Erro na compilação do script de seed"
    exit 1
fi
