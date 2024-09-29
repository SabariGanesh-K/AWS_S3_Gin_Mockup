<h1>Complete Test Driven Development of Backend using Go Lang and AWS  S3 </h1>

<h2>Features and Stacks </h2>
- Used Gin framework for Backend 

- Complete test driven development using testify
  
- Migrations based DB setup using golang migrate

- Redis with background processors for mail handling
  
- PostgreSQL and Redis  Dockerized at Alpine lightweight image
  
- DB Mock being used to isolate dependencies  [Mockgen]
  
- SQLC generation

- AWS
  
- Cron Job for asset scheduled deletion [Hourly]

- Graceful shutdown to prevent crashes

  ![image](https://github.com/user-attachments/assets/863abb2f-b24f-4fdd-949d-fdb817ebfd82)

  
 and many more...

 Feel free to checkout the code.

<h2>Technical Implementation</h2>

Pointer wise explanation :-
Requirements:
1. User Authentication & Authorization:
- Implemented an advanced safe and secured Paseto token based auth with tests in token package . Helps offer clear distinction between private and public tokens

  ([https://permify.co/post/jwt-paseto/](https://permify.co/post/jwt-paseto/))

2. File Upload & Management:
- Used AWS SDK for Go to upload objects and retrieve location and metadatas
- Localised aws config for safe access key auth
  

3. File Retrieval & Sharing:
- Location is stored in postgresql for further retrieval

4. File Search:
- Has search function based on Coalace and sqlc args for a flexible search of file based on need


5. Database Interaction:
- Used Postgresql and sqlc with mocks for file details handling and mapping
- Tests to keep DB operations in sync
![image](https://github.com/user-attachments/assets/cb518360-bda3-4172-906a-19a7522ebf66)


6. Background Job for File Deletion:
- Utilized cron job to run deleter function @hourley

7. Testing:
Complete app is test driven which is pretty evident
- Makefile has every commands u need including test to facilitate all unit testing on multiple layers
![image](https://github.com/user-attachments/assets/9461ed0e-361e-4b58-8f2f-769571bb20fc)

![image](https://github.com/user-attachments/assets/84753341-6363-4781-9c7e-365bac68faa9)

8. Redis:
- Redis setup with Task processor and Distributor using Asynq for background workers
  
<h1> Running and Installation Instructions in Local </h2>

Clone the repo to your local

Start your Docker Engine 

Install Docker dependency Images for Postgres and Redis Alpine Images [ If not installed already ] 
```code
make docker
make redis
```

Configure AWS IAM config details to access AWS Services .
- Install aws cli
- Create a S3 Bucket in AWS Console
- Create a IAM User role and attach policy to access S3 Buckets with the IAM role user
- Download and record the Access key ID and Secret Access key
- Now enter in the prompt when asked after 
```code
aws config
```

Launch the images
```code
make launch
```
Start the server
```code
make server
```










