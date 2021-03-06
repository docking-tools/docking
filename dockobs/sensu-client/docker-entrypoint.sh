#!/bin/sh

if [ ! -z "$ADDITIONAL_INFO" ]; then
  ADDITIONAL_INFO=",$*"
fi

if [ -z "$CLIENT_ADDRESS" ]; then
  echo "\$CLIENT_ADDRESS must be provided" 
  exit 1
fi

if [ -z "$CLIENT_NAME" ]; then
  echo "\$CLIENT_NAME must be provided" 
  exit 1
fi

if [ -z "$SUB" ]; then
  echo "\$SUB must be provided" 
  exit 1
fi

if [ -z "$RABBITMQ_HOST" ]; then
  echo "\$RABBITMQ_HOST must be provided" 
  exit 1
fi

SUBSCRIPTIONS="`echo $SUB|sed s/,/\\",\\"/g`"

cat << EOF > /etc/sensu/config.json
{
  "client": {
    "name": "$CLIENT_NAME",
    "address": "$CLIENT_ADDRESS",
    "subscriptions": ["$SUBSCRIPTIONS"],
    "socket": {
      "bind":"0.0.0.0",
      "port":3030
    }
    $ADDITIONAL_INFO
  },
  "rabbitmq": {
    "host": "$RABBITMQ_HOST",
    "port": "${RABBITMQ_PORT:=5672}",
    "vhost": "${RABBITMQ_VHOST:=/sensu}",
    "user": "${RABBITMQ_USER:=sensu}",
    "password": "${RABBITMQ_PASS:=secret}"
  }
}
EOF

echo "Running sensu config:"
cat /etc/sensu/config.json

exec "$@" 
