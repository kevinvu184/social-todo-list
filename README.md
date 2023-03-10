# Social Todo

1. Install [Docker Desktop for Mac](https://docs.docker.com/desktop/install/mac-install/)
2. Pull the Dynamodb-local image `docker pull amazon/dynamodb-local`
3. At working directory, create [`docker-compose.yml`](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html)
   ```
   version: '3.8'
   services:
   dynamodb-local:
       command: "-jar DynamoDBLocal.jar -sharedDb -dbPath ./data"
       image: "amazon/dynamodb-local:latest"
       container_name: dynamodb-local
       ports:
       - "8000:8000"
       volumes:
       - "./docker/dynamodb:/home/dynamodblocal/data"
       working_dir: /home/dynamodblocal
   ```
4. At the working directory, execute `docker compose up`
5. Create dummy AWS credentials at `~/.aws/credential`
   ```
   [default]
   aws_access_key_id=DUMMYIDEXAMPLE
   aws_secret_access_key=DUMMYEXAMPLEKEY
   ```
6. To test the running container, execute `aws dynamodb list-tables --endpoint-url http://localhost:8000`
7. Install a DynamoDB client, we will use [NoSQL Workbench](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.settingup.install.html)
8. Create Todo table and dummy todos via NoSQL Workbench - Operation builder or via CLI

## Important link

- [DynamoDB Local](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.UsageNotes.html)
- [NoSQL Workbench](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/workbench.html)
- [AWS DynamoDB CLI](https://awscli.amazonaws.com/v2/documentation/api/latest/reference/dynamodb/index.html)
- [API design](https://stackoverflow.blog/2020/03/02/best-practices-for-rest-api-design/#h-accept-and-respond-with-json)
- [AWS Go SDK](https://aws.github.io/aws-sdk-go-v2/docs)
  - [AWS Go SDK API](https://pkg.go.dev/github.com/aws/aws-sdk-go-v2)
