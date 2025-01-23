{
  description = "Go development environment with Nix Flakes";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs {
          inherit system;
        };
      in {
        devShell = pkgs.mkShell {
          buildInputs = [
            pkgs.go_1_23  # Go programming language
            pkgs.gopls    # Go language server
            pkgs.gotools  # Additional Go tools
          ];

          shellHook = ''
            cd ./nessie-client
            go mod tidy
            echo "Go development environment is ready!"
          '';
        };
      });
}
