#!/bin/bash
swagger-codegen generate -l go --api-package openfigi --model-package openfigi --additional-properties packageName=openfigi -i https://api.openfigi.com/schema -o openfigi