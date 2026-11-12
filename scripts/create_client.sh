#!/usr/bin/env bash
set -euo pipefail

CA_DIR="certs/ca"
CA_KEY="${CA_DIR}/ca.key"
CA_CERT="${CA_DIR}/ca.crt"
CLI_DIR="client"
CLI_KEY="${CLI_DIR}/client.key"
CLI_CSR="${CLI_DIR}/client.csr"
CLI_CERT="${CLI_DIR}/client.crt"
DAYS=365
KEY_LEN=4096

read -rp "Введите имя (CN) для клиента: " CLI_CN
read -rsp "Введите пароль для приватного ключа CA: " CA_PASS
echo

if [[ ! -f "$CA_KEY" || ! -f "$CA_CERT" ]]; then
  echo "Ошибка: Файлы CA не найдены в $CA_DIR" >&2
  exit 1
fi

mkdir -p "$CLI_DIR"

openssl genrsa -out "$CLI_KEY" "$KEY_LEN"

openssl req -new -key "$CLI_KEY" -out "$CLI_CSR" -subj "/CN=${CLI_CN}"

EXT_FILE=$(mktemp)
cat > "$EXT_FILE" <<EOF
extendedKeyUsage=clientAuth
keyUsage=digitalSignature
EOF

openssl x509 -req -in "$CLI_CSR" \
  -CA "$CA_CERT" -CAkey "$CA_KEY" -CAcreateserial \
  -out "$CLI_CERT" -days "$DAYS" \
  -extfile "$EXT_FILE" \
  -passin pass:"$CA_PASS"

rm -f "$EXT_FILE"

echo "Клиентский сертификат создан:"
echo "  Ключ: $CLI_KEY"
echo "  Сертификат: $CLI_CERT"
echo "  CN: $CLI_CN"
