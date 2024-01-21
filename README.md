# SecretsManager
This is my base pet project of application for keep secrets.
This small application should have a client and a web application connected to the backend and store the secrets encrypted in the database.
In the process of completing this project, I would like to get acquainted with the Golang programming language and the vue.js web development framework.

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

## 4. Analysis

## 5. DDD

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

## 8. Build

As build system, I decided to use the Gilbert tool, which is an analogue of Ant, Gradle, Maven for those who develop their projects in Go. 
:heavy_minus_sign: [Gilbert](https://github.com/go-gilbert/gilbert)<br />


Gilber provides ample opportunities for creating custom scripts for project assembly, testing, check project coverage, publishing to the server, and so on. <br />

A Yaml file is used for configuration. The examples of the configuration Yaml files can be found there: <br />
:heavy_minus_sign: [Server build](https://github.com/40104/SecretsManager/blob/main/cmd/server/gilbert.yaml)<br />
:heavy_minus_sign: [CLI client build](https://github.com/40104/SecretsManager/blob/main/cmd/cli_client/gilbert.yaml)<br />


## 9. Unit tests

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



