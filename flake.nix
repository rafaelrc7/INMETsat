{
  description = "Fetch animated satellite GIF from INMET";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-unstable";
    flake-utils = {
      url = "github:numtide/flake-utils";
      inputs.nixpkgs.follows = "nixpkgs";
    };
  };

  outputs = { self, nixpkgs, flake-utils }:
    let lastModifiedDate = self.lastModifiedDate or self.lastModified or "19700101";
        version = builtins.substring 0 8 self.lastModifiedDate;
    in flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs {
            inherit system;
            overlays = [ (final: prev: { go = final.go_1_18; }) ];
          };
          pkg = pkgs.buildGo118Module {
            pname = "inmetsat";
            inherit version;
            src = ./.;
            vendorSha256 = "sha256-pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";
          };
      in {
        defaultPackage = pkg;
        devShell = pkgs.mkShell {
          buildInputs = [
            pkgs.go
            pkgs.gotools
            pkgs.gopls
            pkgs.gopkgs
          ];
        };
      }
    );

}

