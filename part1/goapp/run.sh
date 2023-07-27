#!/bin/sh
# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0


echo "Starting Vault dev server.."
container_id="ad1796319ccb"

echo "Running quickstart example."
go run main.go

echo "Stopping Vault dev server.."
# docker stop "${container_id}" > /dev/null

# echo "Vault server has stopped."
