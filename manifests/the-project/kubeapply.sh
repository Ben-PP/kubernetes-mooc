#!/bin/bash

IGNORE_PATTERNS=(
    "secret.enc.yaml"
	"namespace.yaml"
	"secret.yaml"
)

set -euo pipefail

DELETE=false
while getopts ":d" opt; do
    case "$opt" in
        d) DELETE=true ;;
        \?)
            echo "Error: Invalid option '-$OPTARG'" >&2
            echo "Usage: $0 [-d] /path/to/age-key-file"
            exit 1
            ;;
    esac
done
shift $((OPTIND-1))

if ! $DELETE; then
	if [[ $# -ne 1 ]]; then
	    echo "Error: Exactly one argument (the path to the AGE key file) is required."
	    echo "Usage: $0 /path/to/age-key-file"
	    exit 1
	fi
	
	AGE_KEY_FILE="$1"
	
	if [[ ! -r "$AGE_KEY_FILE" ]]; then
	    echo "Error: Cannot read file '$AGE_KEY_FILE'. Check the path and permissions."
	    exit 1
	fi
	
	export SOPS_AGE_KEY_FILE="$AGE_KEY_FILE"
fi

should_ignore() {
    local file="$1"
    for pat in "${IGNORE_PATTERNS[@]}"; do
        if [[ "$file" == $pat ]]; then
            return 0
        fi
    done
    return 1
}
shopt -s globstar nullglob

if ! $DELETE; then
	kubectl apply -f namespace.yaml
	sops --decrypt secret.enc.yaml | kubectl apply -f -
else
	kubectl delete -f secret.yaml
fi

found_any=false
for yaml_file in **/*.yaml **/*.yml; do
    # Skip if the glob didn't match anything
    [[ -e "$yaml_file" ]] || continue

    # Apply ignore filter
    if should_ignore "$yaml_file"; then
        echo "Skipping ignored file: $yaml_file"
        continue
    fi

    found_any=true
    echo "Processing: $yaml_file"

	if $DELETE; then
		kubectl delete -f "$yaml_file"
	else
    	kubectl apply -f "$yaml_file"
	fi
done

if ! $found_any; then
    echo "No matching YAML files were found (or all were ignored)."
fi
