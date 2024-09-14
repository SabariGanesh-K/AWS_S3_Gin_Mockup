<h1>Complete Test Driven Development of Backend using Go Lang " File Sharing & Management System " </h1>

<h2>Features and Stacks </h2>
- Used Gin framework for Backend 
- Complete test driven development using testify
- Migrations based DB setup using golang migrate
- PostgreSQL from Docker
- DB Mock being used to isolate dependencies  [Mockgen]
- SQLC generation
- AWS 
- Cron Job for asset scheduled deletion [Hourly]
 and many more...

 Feel free to checkout the code.


<h2>Before we go further , I would like to introduce myself </h2>

This is Sabari , A tech enthusiast known for my tech skills and versatility in learning new technologies. My past experience revolves around multiple domains with major ones:
- Fullstack Intern at Nethermind , United Kingdom [Remote] . I was involved in multiple projects including go lang major project called sedge , which is a commandline based project for blockchain POS networks. The day where I was
- introduced to the world of Go lang
- SDE Intern at Hashstack Finance , Bengaluru . A finance based protocol trading application that revoved around Next , typescript , Next API , AWS DynamoDBs etc. Was involved fullstack in contributing to project
- Backend and Infrastructure Intern at EDMYN - A Edtech startup By leading National level Computer science Academic author Ms. Sumita Arora . Involved in designing backend , architecture in firebase and developing it with CTO [I was the first hire :) ]
- Hvae worked with other startups like Golddust Finance as well with freelancing experince in freshman + sophomore year
- Experienced and skilled in Go lang. Have also taken courses revolving around Go lang. 
- Also a fun , empathetic , motivated coder with a will to do anything that is required . Have always been a all rounder as well being passionate on extracurrcilurs in college. Will be a fun to have extrovert around the office

- You may know more about me here :- [https://sabari.codes](url)

<h2>Technical Implementation</h2>

Pointer wise explanation :-
Requirements:
1. User Authentication & Authorization:
- Implemented an advanced safe and secured Paseto token based auth with tests in token package . Helps offer clear distinction between private and public tokens
  [https://permify.co/post/jwt-paseto/](url)

2. File Upload & Management:
- Used AWS SDK for Go to upload objects and retrieve location and metadatas
- Localised aws config for safe access key auth
  

3. File Retrieval & Sharing:
- Location is stored in postgresql for further retrieval

4. File Search:
- Has search function based on Coalace and sqlc args for a flexible search of file based on need


6. Database Interaction:
- Used Postgresql and sqlc with mocks for file details handling and mapping
- Tests to keep DB operations in sync


7. Background Job for File Deletion:
- Utilized cron job to run deleter function @hourley

8. Testing:
Complete app is test driven which is pretty evident
- Makefile has every commands u need including test to facilitate all unit testing on multiple layers



