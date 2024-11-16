aconf() {
  local output
  output=$(/usr/local/bin/aconf "$@")
  if [ $? -ne 0 ]; then
    echo "$output"
    return 1
  fi
  eval "$output"
}
