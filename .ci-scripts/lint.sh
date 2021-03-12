#!/bin/sh

kubeval $(find .k8s -not -name 'kustomization.*' -not -name 'metrics.yaml' -name '*.yaml')
