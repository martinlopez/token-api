run-token-api:
	docker build -f Dockerfile.api -t token-api:1.0.0 .
	docker run -i -e ENV=local -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} -p 8080:8080 -d token-api:1.0.0

run-token-job:
	docker build -f Dockerfile.job -t token-job:1.0.0 .
	docker run -i -e ENV=local -e AWS_ACCESS_KEY_ID=${AWS_ACCESS_KEY_ID} -e AWS_SECRET_ACCESS_KEY=${AWS_SECRET_ACCESS_KEY} -d token-job:1.0.0

execute-migration:
	docker run -v /Users/martindev/interviews/blockparty/ifps-data/db/migrations:/flyway/sql flyway/flyway:9.5.1 -url=jdbc:mysql://dev-ifps.cfwo0avpvdh0.us-east-1.rds.amazonaws.com:3306/core -user=admin -password=${DBPASSWORD}  migrate

push-image-latest:
	aws ecr get-login-password --region us-east-1 | docker login --username AWS --password-stdin 556891675300.dkr.ecr.us-east-1.amazonaws.com
	docker build -t token-api:latest .
	docker tag ifpsdata:latest 556891675300.dkr.ecr.us-east-1.amazonaws.com/ifpsdata:latest
 	docker push 556891675300.dkr.ecr.us-east-1.amazonaws.com/ifpsdata:latest
