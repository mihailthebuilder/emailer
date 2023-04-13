# Emailer

Run the server locally:

```
go run .
```

Run Dockerised application locally, in production mode:

```
docker build -t company-data-backend-container .
docker run -p 8080:8080 --env-file .env --env GIN_MODE=release company-data-backend-container
```

Deploy to Hetzner server with Caprover:

```
caprover deploy --default
```

Test email endpoint

```
curl -X POST http://localhost:8080/email -d '{"EmailAddress":"hello@test.com", "Title":"CompanyHound"}'
```
