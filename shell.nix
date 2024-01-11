{ pkgs ? import <nixpkgs> { } }:
with pkgs;
mkShell {
  packages = [
    go
    gopls
    delve
  ];
}

