build:
	docker build -t api-gtw ./api-gtw/
	docker build -t auth-svc ./auth-svc/

compose:
	docker compose --env-file ./auth-svc/configs/envs/dev.env -p kit up