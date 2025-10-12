#!/bin/sh

echo ">>> Merging config/gen/*.yml into buf.gen.yaml"

mkdir -p config/buf/plugins/temp_all
yq eval-all '[.]' config/buf/plugins/notification-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/notification_service_plugins.yml
yq eval-all '[.]' config/buf/plugins/payment-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/payment_service_plugins.yml
yq eval-all '[.]' config/buf/plugins/product-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/product_service_plugins.yml
yq eval-all '[.]' config/buf/plugins/user-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/user_service_plugins.yml

yq eval-all '[.]' config/buf/plugins/temp_all/*.yml | yq eval 'flatten' > config/buf/plugins/all_plugins.yml
yq eval-all '.plugins = load("config/buf/plugins/all_plugins.yml") | .' config/buf/_base_config.yml > buf.gen.yaml

rm -rf config/buf/plugins/temp_all config/buf/plugins/all_plugins.yml
echo ">>> Done."
