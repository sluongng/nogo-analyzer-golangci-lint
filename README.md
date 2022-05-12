Note: deprecated over https://github.com/sluongng/nogo-analyzer

# NoGo Analyzer GolangCI-Lint

This is a code generator tool that helps generate golang packages.
Each of these packages export a single analysis.Analyzer variable which
reflect the corresponding linter in `golangci-lint`

Bazel's rules_go projects can then import this project and use the generated packages
with the `nogo` static analysis framework.

## Status

Very much WIP.
I am doing a quick brain dump here on a Friday to make sure that I don't forget about these on Monday.

Also depending on https://github.com/golangci/golangci-lint/pull/2710 to get merged.
Currently I am using a local go.mod replace

## Design

There are a few commands to implement:

1. First we have to parse the golang package to find all linter constructor functions.
   - For early stage, we can skip out constructors with parameters.
   - Linters with custom configs may need manually crafted with the correct set of configuration.
   - Possibly tell golangci-lint to use Analyzer's flagSet instead of config passing through constructor?
   - For quick implementation, I am also using `go mod vendor` to download golangci-lint packages.
   - Currently used go/ast packages but tree-sitter might be a better/faster approach?

2. Implement a quick counting function to check if we can safely assume that all linters only have 1 analyzer.
   It seems like that assumptions is true (for now).
   This helps us avoid having to sort and iterate through the analyzers slice which is nice.

3. Implement a code generator that create packages and the Go file using existing template.

4. Generate starlark and nogo config JSON files programatically for easier consumption.

Current TODOs: 3 and 4

## References

See https://github.com/sluongng/staticcheck-codegen which followed similar approach to expose staticcheck analyzers to nogo.
