{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
  ];
  shellHook = ''
    echo "Go: $(go version)"
  '';
}
