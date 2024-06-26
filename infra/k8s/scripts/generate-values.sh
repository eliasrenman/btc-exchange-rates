# generate-values.sh
#!/bin/bash

# Load environment variables from .env file
export $(grep -v '^#' ../../.env | xargs)

# Substitute environment variables in values-template.yaml
envsubst < Values-template.yaml > values.yaml
