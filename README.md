# SecretsManager
This is my base pet project of application for keep secrets.
This small application should have a client connected to the backend and store the secrets encrypted in the database.
In the process of completing this project, I would like to get acquainted with the Golang programming language.

Dependencies:<br />
_go get -v 'package-name'_  <br />

Run:<br />
_go run .cmd/server/main/_  <br />
_go run .cmd/cli_client/main/_  <br />

Build & Run:<br />
_go build -o server .cmd/server/main/_ <br />
_.cmd/server/server_ <br />
_go build -o clsm .cmd/cli_client/main/_ <br />
_.cmd/cli_client/clsm_ <br />


## Information on meeting all requirements can be found below.
## 1. Git
While working on the project I used GitHub

:heavy_minus_sign: [Commit history](https://github.com/40104/SecretsManager/commits/main) <br />
:heavy_minus_sign: [Branches history](https://github.com/40104/SecretsManager/branches) <br />

## 2. UML 
To describe the project, I prepared UML diagrams of classes, components and users. <br />
:heavy_minus_sign: [UML Diagrams](https://github.com/40104/SecretsManager/tree/main/diagrams)<br />

## 3. Requirements Engineering
To analyse the project, I prepared requirements engineering of my project using 2 different platform: <br />
:heavy_minus_sign: [Confuence](https://40104.atlassian.net/wiki/spaces/~615f3ba3bfa2c1006bd815a0/pages/294913/40104+Secrets+Manager)<br />
:heavy_minus_sign: [Notion](https://amusing-snake-490.notion.site/40104-Secrets-Manager-9aa9eac6b6224574b30825b51bad64c2)<br />

## 4. Analysis
:heavy_minus_sign: [Analysis_A(link)](https://github.com/40104/SecretsManager/blob/main/documets/Analyse_A.md) or [Analysis_A(PDF)](https://github.com/40104/SecretsManager/blob/main/documets/Analyse_A.md)<br />
:heavy_minus_sign: [Analysis_B(link)](https://github.com/40104/SecretsManager/blob/main/documets/Analyse_B.md) or [Analysis_B(PDF)](https://github.com/40104/SecretsManager/blob/main/documets/Analyse_B.md)<br />

## 5. DDD
To describe the project,I prepared Visual Event Storming, Core Domain Chart and Relations / Mappings between the Domains  <br />
:heavy_minus_sign: [DDD](https://github.com/40104/SecretsManager/blob/main/diagrams/DDD.pdf)<br />

## 6. Metrics
I used the SonarQube service to get metrics for my project. <br />
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=40104_SecretsManager&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=40104_SecretsManager)
[![Technical Debt](https://sonarcloud.io/api/project_badges/measure?project=40104_SecretsManager&metric=sqale_index)](https://sonarcloud.io/summary/new_code?id=40104_SecretsManager)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=40104_SecretsManager&metric=bugs)](https://sonarcloud.io/summary/new_code?id=40104_SecretsManager)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=40104_SecretsManager&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=40104_SecretsManager)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=40104_SecretsManager&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=40104_SecretsManager)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=40104_SecretsManager&metric=vulnerabilities)](https://sonarcloud.io/summary/new_code?id=40104_SecretsManager)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=40104_SecretsManager&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=40104_SecretsManager)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=40104_SecretsManager&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=40104_SecretsManager)

## 7. Clean Code Development
In the project I used many methods to make my code clean <br />
A: <br />
:heavy_minus_sign: [Detailed comments](https://github.com/40104/SecretsManager/blob/main/cmd/server/main/config.go#L24-L52)<br />
:heavy_minus_sign: [Clear function names](https://github.com/40104/SecretsManager/blob/main/cmd/server/models/folder.go#L43)<br />
:heavy_minus_sign: [Class Design](https://github.com/40104/SecretsManager/blob/main/cmd/server/main/config.go#L12-L23)<br />
:heavy_minus_sign: [Strong typing](https://github.com/40104/SecretsManager/blob/main/cmd/server/controllers/middleware.go#L30)<br />
:heavy_minus_sign: [Error handler](https://github.com/40104/SecretsManager/blob/main/cmd/server/controllers/middleware.go#L47-L63)<br />
:heavy_minus_sign: [Using standard packages](https://github.com/40104/SecretsManager/blob/main/cmd/server/models/crypto.go#L4-L7)<br />
:heavy_minus_sign: [Clear variables names](https://github.com/40104/SecretsManager/blob/main/cmd/server/models/model.go#L11-L45)<br />
:heavy_minus_sign: [Formating](https://github.com/40104/SecretsManager/blob/main/cmd/server/controllers/secret.go#L47-L49)<br />
:heavy_minus_sign: [Using env](https://github.com/40104/SecretsManager/blob/main/cmd/server/configs/app.env#L1-L4)<br />
:heavy_minus_sign: [Testing]()<br />

B: <br />
:heavy_minus_sign: [PDF]()<br />

## 8. Build

As build system, I decided to use the Gilbert tool, which is an analogue of Ant, Gradle, Maven for those who develop their projects in Go. 
:heavy_minus_sign: [Gilbert](https://github.com/go-gilbert/gilbert)<br />


Gilber provides ample opportunities for creating custom scripts for project assembly, testing, check project coverage, publishing to the server, and so on. <br />

A Yaml file is used for configuration. The examples of the configuration Yaml files can be found there: <br />
:heavy_minus_sign: [Server build](https://github.com/40104/SecretsManager/blob/main/cmd/server/gilbert.yaml)<br />
:heavy_minus_sign: [CLI client build](https://github.com/40104/SecretsManager/blob/main/cmd/cli_client/gilbert.yaml)<br />


## 9. Unit tests

For unit testing, the "testing" library built into Go was used. Unit tests cover almost all functions of accessing a database, as well as functions for generating a JWT token and its verification. The Unit testing file can be found below. <br />
:heavy_minus_sign: [Unit test](https://github.com/40104/SecretsManager/blob/main/cmd/server/main/main_test.go)<br />
To start Unit testing you need to use the command:<br />
 _go test -v -o coverage.html_ <br />
Output: <br />
=== RUN   Test_Create_Token <br />
--- PASS: Test_Create_Token (0.01s) <br />
=== RUN   Test_Verify_Token <br />
--- PASS: Test_Verify_Token (0.01s) <br />
=== RUN   Test_Roles <br />
--- PASS: Test_Roles (0.04s) <br />
=== RUN   Test_User <br />
--- PASS: Test_User (0.14s) <br />
=== RUN   Test_Folder <br />
--- PASS: Test_Folder (0.03s) <br />
=== RUN   Test_Secret <br />
--- PASS: Test_Secret (0.05s) <br />
PASS <br />
:heavy_minus_sign: [Coverage file](https://github.com/40104/SecretsManager/blob/main/documents/coverage.html)<br />


## 10. IDE
While working on the project, I used Visual Studio Code to work on the project.
This is a convenient, multi-platform IDE with many plugins for any task.

List of plugins used:<br />
:heavy_minus_sign: Go - for help with writing Go code <br />
:heavy_minus_sign: Remote - SSH - for easy connection to the server on which the work was carried out <br />
:heavy_minus_sign: GitHub Pull Requests and Issues - for help with GitHub commits and branches <br />

List of favorite shortcuts:<br />
:heavy_minus_sign: 'Ctrl + c' - Copy <br />
:heavy_minus_sign: 'Ctrl + v' - Paste <br />
:heavy_minus_sign: 'Ctrl + f' - Search <br />
:heavy_minus_sign: 'Ctrl + '' - Console <br />
:heavy_minus_sign: 'Ctrl + r' - Replace <br />

## 11. Functional Programming
I have prepared a separate file to demonstrate Functional Programming. Below you can find links to lines of code that show the functional aspects in my code <br />
:heavy_minus_sign: [Only final data structures](https://github.com/40104/SecretsManager/blob/main/cmd/examples/Functional_Programming.go#L10-L21)<br />
:heavy_minus_sign: [Side-effect-free functions (Mostly)](https://github.com/40104/SecretsManager/blob/main/cmd/examples/Functional_Programming.go#L23-L35)<br />
:heavy_minus_sign: [The use of higher-order functions](https://github.com/40104/SecretsManager/blob/main/cmd/examples/Functional_Programming.go#L37-L51)<br />
:heavy_minus_sign: [Functions as parameters and return values](https://github.com/40104/SecretsManager/blob/main/cmd/examples/Functional_Programming.go#L53-L75)<br />
:heavy_minus_sign: [Use closures / anonymous functions](https://github.com/40104/SecretsManager/blob/main/cmd/examples/Functional_Programming.go#L77-L87)<br />




