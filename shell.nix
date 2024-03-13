with (import <nixpkgs> {});
mkShell {
  buildInputs = [
    cloudflared
    tailwindcss
    go
    templ
    sqlc
  ];
}
