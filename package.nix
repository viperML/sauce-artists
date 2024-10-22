{
  buildGoModule,
  lib,
}:
buildGoModule {
  name = "sauce-artists";
  src = lib.cleanSource ./.;
  vendorHash = null;
}
