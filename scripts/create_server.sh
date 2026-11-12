#!/usr/bin/env bash
set -euo pipefail

CA_DIR="certs/ca"
CA_KEY="${CA_DIR}/ca.key"
CA_CERT="${CA_DIR}/ca.crt"
SRV_DIR="server"
SRV_KEY="${SRV_DIR}/server.key"
SRV_CSR="${SRV_DIR}/server.csr"
SRV_CERT="${SRV_DIR}/server.crt"
DAYS=365
KEY_LEN=4096

read -rp "Введите имя (CN) для сервера: " SRV_CN
read -rp "Введите список alt-names через запятую (например, DNS:example.com,DNS:www.example.com,IP:192.168.1.1): " SRV_SANS
read -rsp "Введите пароль для приватного ключа CA: " CA_PASS
echo

if [[ ! -f "$CA_KEY" || ! -f "$CA_CERT" ]]; then
  echo "Ошибка: Файлы CA не найдены в $CA_DIR" >&2
  exit 1
fi

mkdir -p "$SRV_DIR"

openssl genrsa -out "$SRV_KEY" "$KEY_LEN"

openssl req -new -key "$SRV_KEY" -out "$SRV_CSR" -subj "/CN=${SRV_CN}"

EXT_FILE=$(mktemp)
cat > "$EXT_FILE" <<EOF
subjectAltName=${SRV_SANS}
extendedKeyUsage=serverAuth
keyUsage=digitalSignature,keyEncipherment
EOF

openssl x509 -req -in "$SRV_CSR" \
  -CA "$CA_CERT" -CAkey "$CA_KEY" -CAcreateserial \
  -out "$SRV_CERT" -days "$DAYS" \
  -extfile "$EXT_FILE" \
  -passin pass:"$CA_PASS"

rm -f "$EXT_FILE"

echo "Серверный сертификат создан:"
echo "  Ключ: $SRV_KEY"
echo "  Сертификат: $SRV_CERT"
echo "  SAN: $SRV_SANS"
