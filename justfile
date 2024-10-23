entrypoint := "./cmd/bank-system/main.go"

alias r := run

run:
  go run {{entrypoint}}
