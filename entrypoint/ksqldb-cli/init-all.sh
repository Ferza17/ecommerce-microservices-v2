#!/bin/bash
set -e

echo "Starting all ksqldb-cli initialization scripts..."

# Run connectors init
if [ -f "/connectors/init.sh" ]; then
  echo "Running /connectors/init.sh..."
  bash /connectors/init.sh
fi

# Run streams init
if [ -f "/scripts/init.sh" ]; then
  echo "Running /scripts/init.sh..."
  bash /scripts/init.sh
fi

echo "All init scripts completed. Keeping container alive..."
# Keep container running
tail -f /dev/null
