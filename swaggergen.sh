#!/usr/bin/env bash
swagger generate server --spec=pkg/api/api.yaml --name=FooService --target=pkg/gen
