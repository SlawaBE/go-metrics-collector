#!/usr/bin/env bash
set -euo pipefail

CA_DIR="certs/ca"
CA_KEY="${CA_DIR}/ca.key"
CA_CERT="${CA_DIR}/ca.crt"
DAYS=3650
KEY_LEN=4096

read -rp "Введите имя (CN) для корневого CA: " CA_CN
read -rsp "Введите пароль для приватного ключа CA: " CA_PASS
echo

mkdir -p "$CA_DIR"

openssl req -x509 -newkey rsa:${KEY_LEN} \
  -keyout "$CA_KEY" \
  -out "$CA_CERT" \
  -days "$DAYS" \
  -subj "/CN=${CA_CN}" \
  -addext "basicConstraints=critical,CA:TRUE" \
  -addext "keyUsage=critical,keyCertSign,cRLSign" \
  -passout pass:"$CA_PASS"

echo "Корневой CA создан:"
echo "  Ключ: $CA_KEY (защищён паролем)"
echo "  Сертификат: $CA_CERT"
