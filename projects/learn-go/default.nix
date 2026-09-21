{ pkgs, ... }:
pkgs.buildGo125Module {
  pname = "learn-go";
  version = "0.1";
  src = ./.;
  vendorHash = null;
  subPackages = [ "." "./package" ];
}
