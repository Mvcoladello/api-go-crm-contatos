#!/bin/bash

# Script de teste da API CRM Contatos
# Certifique-se de que a API está rodando em http://localhost:3000

echo "=== Testando API CRM Contatos ==="
echo ""

echo "1. Verificando saúde da API..."
curl -s http://localhost:3000/health | jq
echo ""

echo "2. Listando contatos (deve estar vazio inicialmente)..."
curl -s http://localhost:3000/api/v1/contatos | jq
echo ""

echo "3. Criando contato com CPF..."
RESPONSE=$(curl -s -X POST http://localhost:3000/api/v1/contatos \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "João Silva",
    "email": "joao@example.com",
    "cpf_cnpj": "11144477735",
    "telefone": "11999999999"
  }')
echo $RESPONSE | jq
CONTACT_ID=$(echo $RESPONSE | jq -r '.data.id')
echo "ID do contato criado: $CONTACT_ID"
echo ""

echo "4. Criando contato com CNPJ..."
curl -s -X POST http://localhost:3000/api/v1/contatos \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Empresa ABC Ltda",
    "email": "contato@empresaabc.com",
    "cpf_cnpj": "11222333000181",
    "telefone": "1133334444"
  }' | jq
echo ""

echo "5. Listando todos os contatos..."
curl -s http://localhost:3000/api/v1/contatos | jq
echo ""

echo "6. Buscando contato por ID..."
curl -s http://localhost:3000/api/v1/contatos/$CONTACT_ID | jq
echo ""

echo "7. Tentando criar contato com dados inválidos..."
curl -s -X POST http://localhost:3000/api/v1/contatos \
  -H "Content-Type: application/json" \
  -d '{
    "nome": "Test",
    "email": "email-inválido",
    "cpf_cnpj": "12345678901",
    "telefone": "123"
  }' | jq
echo ""

echo "8. Tentando buscar contato com ID inválido..."
curl -s http://localhost:3000/api/v1/contatos/id-invalido | jq
echo ""

echo "9. Deletando contato..."
curl -s -X DELETE http://localhost:3000/api/v1/contatos/$CONTACT_ID | jq
echo ""

echo "10. Verificando se contato foi deletado..."
curl -s http://localhost:3000/api/v1/contatos/$CONTACT_ID | jq
echo ""

echo "=== Teste finalizado ==="
