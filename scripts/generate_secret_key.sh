#!/bin/bash

# string base
string="dwjadwadwadawkdjakljdlwakjdkwadajdlawkjdwklajdkljwakdjwalkdjwakdawkaw"

base64_key=$(echo -n "$string" | base64)

secret_key=$(echo "$base64_key" | head -c 64 | tr -d '\n')

echo "Secret Key: $secret_key"

