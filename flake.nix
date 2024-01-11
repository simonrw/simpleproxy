{
  description = "Flake utils demo";

  inputs = {
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        overlays = [
        ];

        pkgs = import nixpkgs {
          inherit overlays system;
        };
      in
      {
        devShells = rec {
          default = empty;

          empty = import ./shell.nix { inherit pkgs; };
        };

        packages.default = import ./default.nix { inherit pkgs; };
      }
    );
}
