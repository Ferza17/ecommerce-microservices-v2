#!/bin/bash
set -e

echo "Starting all ksqldb-cli initialization streams..."

# Run streams init
if [ -f "/streams/init.sh" ]; then
  echo "Running /streams/init.sh..."
  bash /streams/init.sh
fi

echo "Waiting streams configured!!!"
sleep 5

# Run connectors init
if [ -f "/connectors/init.sh" ]; then
  echo "Running /connectors/init.sh..."
  bash /connectors/init.sh
fi

echo "All init streams completed. Keeping container alive..."
# Keep container running
tail -f /dev/null
