#!/bin/bash

# Base URL
BASE_URL="http://localhost:8080/api/transactions"

# Create a transaction first
echo "Creating a transaction..."
CREATE_RESPONSE=$(curl -s -X POST $BASE_URL \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 100,
    "description": "Test Transaction",
    "category": "Test",
    "date": "2024-01-01"
  }')

# Extract the ID from the response (you might need to adjust this based on your actual response structure)
TRANSACTION_ID=$(echo $CREATE_RESPONSE | jq -r '.id')

echo "Created Transaction ID: $TRANSACTION_ID"

# Test GET transaction
echo "Getting transaction..."
curl -v $BASE_URL/$TRANSACTION_ID

# Test PUT transaction
echo "Updating transaction..."
curl -v -X PUT $BASE_URL/$TRANSACTION_ID \
  -H "Content-Type: application/json" \
  -d '{
    "amount": 200,
    "description": "Updated Test Transaction"
  }'

# Test DELETE transaction
echo "Deleting transaction..."
curl -v -X DELETE $BASE_URL/$TRANSACTION_ID
