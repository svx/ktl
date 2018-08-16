#!/bin/bash
set -euo pipefail

# Vars
ESC_SEQ="\x1b["
COL_RESET=$ESC_SEQ"39;49;00m"
COL_YELLOW=$ESC_SEQ"33;01m"
COL_GREEN=$ESC_SEQ"32;01m"

echo -en "$COL_YELLOW Iterating...$COL_RESET\\n"

nodes=$(kubectl get node --no-headers -o custom-columns=NAME:.metadata.name)

for node in $nodes; do
  echo "Node: $node"
  kubectl describe node "$node" | sed '1,/Non-terminated Pods/d'
  echo
done
