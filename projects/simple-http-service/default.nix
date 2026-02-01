{
  pkgs,
  ...
}: pkgs.buildGo125Module {
  pname = "myapp";
  version = "0.1.0";
  src = ./.;
  vendorHash = null;
}