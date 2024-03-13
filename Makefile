tunnel_name := DEMO-$(shell date +%Y-%m-%d)
dns := demo.plasny.one
app_name := zsl-demo

.PHONY: free-tunnel login-tunnel build dev

all: fmt build run

fmt:
	go fmt ./...

build:
	tailwindcss --input components/base.css --output assets/style.css
	sqlc generate
	templ generate
	go build -o $(app_name) main.go

run:
	./$(app_name)

dev:
	air run

free-tunnel: 
	cloudflared tunnel --url http://localhost:8080

login-tunnel: 
	trap 'cloudflared tunnel delete $(tunnel_name)' INT; \
	cloudflared tunnel login && \
	cloudflared tunnel create $(tunnel_name) && \
	cloudflared tunnel route dns --overwrite-dns $(tunnel_name) $(dns) && \
	cloudflared tunnel run --url http://localhost:8080 $(tunnel_name)

