#!/bin/sh
echo "Waiting for kafka-connect to be ready..."
while ! curl -s http://kafka-connect-local:8083/connectors > /dev/null; do
  sleep 5
done
echo "kafka-connect server is up! Running initialization scripts..."

# Get the script directory
CONNECTOR_DIR="/connectors"

### NAMESPACE EVENT
echo "Running notification initialization..."
sh ${CONNECTOR_DIR}/event/init.sh

### NAMESPACE NOTIFICATION
echo "Running notification initialization..."
sh ${CONNECTOR_DIR}/notification/init.sh

### NAMESPACE PAYMENT
echo "Running payment initialization..."
sh ${CONNECTOR_DIR}/payment/init.sh

### NAMESPACE PRODUCT
echo "Running product initialization..."
sh ${CONNECTOR_DIR}/product/init.sh


### NAMESPACE SHIPPING
echo "Running shipping initialization..."
sh ${CONNECTOR_DIR}/shipping/init.sh

### NAMESPACE USER
echo "Running user initialization..."
sh ${CONNECTOR_DIR}/user/init.sh

echo "Initialization completed."