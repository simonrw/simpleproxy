{ pkgs ? import <nixpkgs> {} }:
pkgs.buildGoModule {
  pname = "simpleproxy";
  version = "0.1.0";
  src = ./.;
  vendorHash = null;

  CGO_ENABLED = "0";
}
