#!/bin/bash

mkdir -p ../services/backend/internal/frontend/assets
cp -r ../frontend/dist/frontend/browser/* ../services/backend/internal/frontend/assets

mkdir -p ../services/backend/internal/docs/assets
cp -r ../docs/.retype/* ../services/backend/internal/docs/assets