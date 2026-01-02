#!/usr/bin/env bash
set -euo pipefail
source .env

# Usage:
#   INTERNAL_API_SECRET=... ./scripts/signed_request.sh GET /projects/1/users_list '?limit=10&page=1'
#
# Notes:
# - Provide the query string as the 3rd arg (include the leading '?') or leave it empty.
# - For methods with a body, pass JSON as the 4th arg (e.g. '{"name":"test"}').

METHOD="${1:-GET}"
PATH_ONLY="${2:-/health}"
QUERY_STRING="${3:-}"
BODY="${4:-}"

if [[ -z "${INTERNAL_API_SECRET:-}" ]]; then
  echo "INTERNAL_API_SECRET is required" >&2
  exit 1
fi

TS="$(date +%s)"
PATH_WITH_QUERY="${PATH_ONLY}${QUERY_STRING}"
BASE="${TS}\n${METHOD}\n${PATH_WITH_QUERY}\n${BODY}"
SIG="$(printf "%b" "$BASE" | openssl dgst -sha256 -hmac "$INTERNAL_API_SECRET" -hex | awk '{print $2}')"

URL="${GO_API_URL:-http://localhost:8080}${PATH_WITH_QUERY}"

if [[ -n "$BODY" ]]; then
  curl -s "$URL" \
    -X "$METHOD" \
    -H "Content-Type: application/json" \
    -H "X-Timestamp: ${TS}" \
    -H "X-Signature: ${SIG}" \
    -d "$BODY"
else
  curl -s "$URL" \
    -X "$METHOD" \
    -H "X-Timestamp: ${TS}" \
    -H "X-Signature: ${SIG}"
fi
