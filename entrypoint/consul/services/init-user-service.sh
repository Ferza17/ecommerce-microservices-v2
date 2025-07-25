#!/bin/sh


echo "INIT USER SERVICE"
## Local
consul kv put local/services/user/ENV 'local'
consul kv put local/services/user/SERVICE_NAME 'user-service'
consul kv put local/services/user/RPC_HOST 'localhost'
consul kv put local/services/user/RPC_PORT '50056'
consul kv put local/services/user/HTTP_HOST 'localhost'
consul kv put local/services/user/HTTP_PORT '40056'
consul kv put local/services/user/METRIC_HTTP_PORT '30056'
consul kv put local/services/user/JWT_ACCESS_TOKEN_SECRET 'ecommerce-service-v2'
consul kv put local/services/user/JWT_ACCESS_TOKEN_EXPIRATION_TIME '1h'
consul kv put local/services/user/JWT_REFRESH_TOKEN_SECRET 'v2-service-ecommerce'
consul kv put local/services/user/JWT_REFRESH_TOKEN_EXPIRATION_TIME '2d'
consul kv put local/services/user/OTP_EXPIRATION_TIME '10m'
consul kv put local/services/user/VERIFICATION_USER_LOGIN_URL 'http://localhost:4000?access_token=%s?refresh_token=?'

## Production
consul kv put production/services/user/ENV 'production'
consul kv put production/services/user/SERVICE_NAME 'user-service'
consul kv put production/services/user/RPC_HOST 'user-service'
consul kv put production/services/user/RPC_PORT '50056'
consul kv put production/services/user/HTTP_HOST 'user-service'
consul kv put production/services/user/HTTP_PORT '40056'
consul kv put production/services/user/METRIC_HTTP_PORT '30056'
consul kv put production/services/user/JWT_ACCESS_TOKEN_SECRET 'ecommerce-service-v2'
consul kv put production/services/user/JWT_ACCESS_TOKEN_EXPIRATION_TIME '1h'
consul kv put production/services/user/JWT_REFRESH_TOKEN_SECRET 'v2-service-ecommerce'
consul kv put production/services/user/JWT_REFRESH_TOKEN_EXPIRATION_TIME '1h'
consul kv put production/services/user/OTP_EXPIRATION_TIME '10m'
consul kv put production/services/user/VERIFICATION_USER_LOGIN_URL 'http://user-service:4000?access_token=%s?refresh_token=?'
echo "DONE INIT USER SERVICE"
