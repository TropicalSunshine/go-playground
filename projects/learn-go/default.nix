{ pkgs, ... }:
pkgs.buildGo125Module {
  pname = "experiments";
  version = "0.1";
  src = ./.;
  vendorHash = null;
}
