#!/bin/bash


# Create the 'products' index with the defined mapping
curl -X PUT "http://elasticsearch-local:9200/products" \
     -H "Content-Type: application/json" \
     -d @/products-mapping.json
