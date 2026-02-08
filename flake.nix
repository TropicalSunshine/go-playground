{
  description = "Go project with Nix flakes";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let pkgs = nixpkgs.legacyPackages.${system};
      in {
        packages = {
          httpService = pkgs.callPackage ./projects/simple-http-service { };
          experiments = pkgs.callPackage ./projects/experiments { };
        };
        app.default = self.packages.${system}.default;
        devShells.default = self.packages.${system}.default;
      });
}
