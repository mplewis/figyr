This is a small Go program to demonstrate how Figyr supports multiple sources
for configuration values, and how they override each other.

Values are parsed from sources in this order (values higher in the list take
priority):

1. `--key=val` command-line flags
2. Values parsed from `--config=cfg-file.yaml` files
3. `KEY=val` environment variables

From the root of this repo, run `bin/demo` to see the output.
