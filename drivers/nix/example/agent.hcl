#log_level = "TRACE"

client {
}

plugin "nix-driver" {
  config {
    default_nixpkgs = "github:nixos/nixpkgs/nixos-23.05"
  }
}
