build:
	docker build -t api-gtw:latest ./api-gtw/
	docker build -t auth-svc:latest ./auth-svc/

compose:
	docker compose --env-file ./auth-svc/configs/envs/local.env -p kit up