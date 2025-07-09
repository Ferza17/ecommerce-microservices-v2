#!/bin/sh

echo ">>> Merging config/gen/*.yml into buf.gen.yaml"

mkdir -p config/buf/plugins/temp_all
yq eval-all '[.]' config/buf/plugins/api-gateway/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/api_gateway_plugins.yml
yq eval-all '[.]' config/buf/plugins/commerce-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/commerce_service_plugins.yml
yq eval-all '[.]' config/buf/plugins/event-store-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/event_store_service_plugins.yml
yq eval-all '[.]' config/buf/plugins/notification-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/notification_service_plugins.yml
yq eval-all '[.]' config/buf/plugins/payment-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/payment_service_plugins.yml
yq eval-all '[.]' config/buf/plugins/product-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/product_service_plugins.yml
yq eval-all '[.]' config/buf/plugins/user-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/user_service_plugins.yml
yq eval-all '[.]' config/buf/plugins/shipping-service/*.yml | yq eval 'flatten' > config/buf/plugins/temp_all/shipping_service_plugins.yml

yq eval-all '[.]' config/buf/plugins/temp_all/*.yml | yq eval 'flatten' > config/buf/plugins/all_plugins.yml
yq eval-all '.plugins = load("config/buf/plugins/all_plugins.yml") | .' config/buf/_base_config.yml > buf.gen.yaml

rm -rf config/buf/plugins/temp_all config/buf/plugins/all_plugins.yml
echo ">>> Done."
