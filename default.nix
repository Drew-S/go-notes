{ pkgs ? import <nixpkgs> {} }:

pkgs.buildGoModule {
  pname = "notes";
  version = "0.0.0";
  src = ./.;
  vendorHash = "sha256-WbtXpASD+69vOjhC0nMEfYpd5cbP0aPt1nLSL4ZxFio=";
}
