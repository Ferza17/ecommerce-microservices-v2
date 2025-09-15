#!/bin/sh
echo "Waiting for ksqlDB server to be ready..."
while ! curl -s http://ksqldb-server-local:8088/info > /dev/null; do
  sleep 5
done
echo "ksqlDB server is up! Running initialization streams..."

# Get the script directory
SCRIPT_DIR="/streams"

### NAMESPACE NOTIFICATION
echo "Running notification initialization..."
#sh ${SCRIPT_DIR}/notification/init.sh

### NAMESPACE PAYMENT
echo "Running payment initialization..."
sh ${SCRIPT_DIR}/payment/init.sh

### NAMESPACE PRODUCT
echo "Running product initialization..."
sh ${SCRIPT_DIR}/product/init.sh

### NAMESPACE SHIPPING
echo "Running shipping initialization..."
sh ${SCRIPT_DIR}/shipping/init.sh

### NAMESPACE USER
echo "Running user initialization..."
sh ${SCRIPT_DIR}/user/init.sh

echo "Initialization completed."