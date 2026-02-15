{
  description = "Control Themes Monitors and Wallpaper on X using xrandr and xwallpaper";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
  let
    system = "x86_64-linux";
    pkgs = import nixpkgs { inherit system; };
  in {
    packages.${system}.default = pkgs.buildGoModule {
      pname = "themr";
      version = "0.1.1";
      src = self;
      subPackages = [ "cmd/themr" ];
      vendorHash = "sha256-g+yaVIx4jxpAQ/+WrGKxhVeliYx7nLQe/zsGpxV4Fn4=";
    };
  };
}
