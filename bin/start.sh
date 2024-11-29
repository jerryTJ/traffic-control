#!/bin/sh
/app/proxy-traffic  --url="${DB_URL}  --user=${DB_USER} --name="${DB_NAME}" --ports="${LISTENER_PORTS}" 