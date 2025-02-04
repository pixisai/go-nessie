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

        go-nessie-nix = pkgs.buildGoModule {
          pname = "go-nessie-nix";
          version = "unstable";
          src = ./.;
          vendorHash = null;
          subPackages = [ "." ];
          installPhase = ''
            mkdir -p $out/bin
            go build -o $out/bin/go-nessie-nix ./main.go
          '';
        };

        nessie = pkgs.stdenv.mkDerivation {
          pname = "nessie";
          version = "0.102.0";
          src = pkgs.fetchurl {
            url = "https://github.com/projectnessie/nessie/releases/download/nessie-0.102.0/nessie-quarkus-0.102.0-runner.jar";
            sha256 = "sha256-64wO5ENs4aEpTOKK+bk+ywNArNm8urqxdx4ggHA5bsI=";
          };
          nativeBuildInputs = [ pkgs.jdk21 ];
          dontUnpack = true;
          installPhase = ''
            mkdir -p $out/bin
            cp $src $out/bin/nessie.jar
            cat > $out/bin/nessie <<EOF
            #!${pkgs.runtimeShell}
            exec ${pkgs.jdk21}/bin/java -jar $out/bin/nessie.jar
            EOF
            chmod +x $out/bin/nessie
          '';
        };

        nessie-cli = pkgs.stdenv.mkDerivation {
          pname = "nessie-cli";
          version = "0.102.0";
          src = pkgs.fetchurl {
            url = "https://github.com/projectnessie/nessie/releases/download/nessie-0.102.0/nessie-cli-0.102.0.jar";
            sha256 = "sha256-LcPeHttzDWdJAPRU0XLDdvbwnpwBDXdxNGYEKfmPYB4";
          };
          nativeBuildInputs = [ pkgs.jdk21 ];
          dontUnpack = true;
          installPhase = ''
            mkdir -p $out/bin
            cp $src $out/bin/nessie-cli.jar
            cat > $out/bin/nessie-cli <<EOF
            #!${pkgs.runtimeShell}
            exec ${pkgs.jdk21}/bin/java -jar $out/bin/nessie-cli.jar
            EOF
            chmod +x $out/bin/nessie-cli
          '';
        };

        minio = pkgs.stdenv.mkDerivation {
          pname = "minio";
          version = "latest";
          src = pkgs.fetchurl {
            url = "https://dl.min.io/server/minio/release/darwin-arm64/minio";
            sha256 = "sha256-xvyx7dht1/fX2qL8x5pNAIXMmfV1VjzwtECfrVTQur0=";
          };
          phases = [ "installPhase" ];
          installPhase = ''
            mkdir -p $out/bin
            cp $src $out/bin/minio
            chmod +x $out/bin/minio
          '';
        };

      in {
        devShell = pkgs.mkShell {
          buildInputs = [
            pkgs.go_1_23
            pkgs.gopls
            pkgs.gotools
          ];

          shellHook = ''
            go mod tidy
            echo "Go development environment is ready!"
          '';
        };

        # Expose Nessie as a package
        packages = {
          go-nessie-nix = go-nessie-nix;
          default = go-nessie-nix;
        };

        # Define apps for nix run
        apps = {
          nessie = flake-utils.lib.mkApp {
            drv = nessie;
            name = "nessie";
          };
          
          nessie-cli = flake-utils.lib.mkApp {
            drv = nessie-cli;
            name = "nessie-cli";
          };

          minio = flake-utils.lib.mkApp {
            drv = minio;
            name = "minio";
          };
        };
      });
}
