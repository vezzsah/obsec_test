Readme

Docker Container Location:
https://hub.docker.com/r/gomaeba/obsec-server


## Running this Project in LocalHost
To Run this project locally, create an .env file with the following variables:
```sh
PORT = "8080"
ENVIRONMENT = "dev"
TOKEN_TYPE = "obsec-api-access"
SECRET_KEY = "Add a random secret key here"
DATABASE_URL = Turso connection string like: "libsql://[Your Project DB].turso.io?authToken=[auth Token]"
```
Once you have your SQLite DB set up, use goose to generate the tables using the following command:
```sh
goose turso ${DATABASE_URL} up
```
You can execute the code with 
```sh
go build -o obsec-server && ./obsec-server
```

For HELM execution, have a minikube service running, and create a configmap.yaml under ./obsec-chart/templates
The file should contain the following:
```sh
apiVersion: "v1"
kind: ConfigMap
metadata:
  name: {{ .Values.configmapName }}
data:
  PORT: "8080"
  ENVIRONMENT: "dev"
  TOKEN_TYPE: {{ .Values.tokenType }}
  SECRET_KEY: "Add a random secret key here"
  DATABASE_URL: ${DATABASE_URL} 
```
Secret management in this project is still pending.
Once that is done, you can use the following command to have HELM create the release.
```sh
./scripts/helminstall.sh
```
Integration with GCP is still pending for this project.

## API Usage
# Clear all users and their projects:
* DELETE http://localhost:8080/v1/users
This will delete all the users. Given that the other tables have the constraint "ON DELETE CASCADE", this will result in clearing the Database. Only allowed on a Dev environment.

# Create a user:
* POST http://localhost:8080/v1/users
example Body:
```sh
{
  "email": "Test01@gmail.com",
  "password": "test01"
}
```
Creates a new user in the DB.

# Get a Log In Token: 
* POST http://localhost:8080/v1/login
example Body:
```sh
{
  "email": "Test01@gmail.com",
  "password": "test01"
}
```
Log In as the given user. If the password is correct, a JWT is generated.

# Create a new Project:
* POST http://localhost:8080/v1/projects
example Body:
```sh
{
  "name": "test_project"
}
```
Creates a new project under the current user.

# Get Project Details: 
* GET http://localhost:8080/v1/projects?project_name=test_project
Gets the project details.

# Assing a CPE to a Project: 
* POST http://localhost:8080/v1/projects/cpes
example Body:
```sh
{
  "project_name": "test_project",
  "cpe_data": [
    {"part": "a",
    "vendor": "mercadopago",
    "product": "mercado_pago_payments_for_woocommerce",
    "version": "2.0.7"}
  ]
}
```
Assing a new CPE to the project. If the CPE is valid (currently checked against mocked data) it will populate the corresponding CVEs also.

# Get all CVEs associated to Project: 
* GET http://localhost:8080/v1/projects/cves?project_name=test_project
Get all the CVE found for the CPEs assigned to the project.

# Resolve a given CVE: 
* POST http://localhost:8080/v1/projects/cves?project_name=test_project
Example Body:
```sh
{
  "project_name": "test_project",
  "cve": "CVE-2022-45068",
  "cpe": "cpe:2.3:a:mercadopago:mercado_pago_payments_for_woocommerce:2.0.7:*:*:*:*:wordpress:*:*"
}
```
Resolve the given CVE. A CPE and a project are required to ensure the correct CVE is resolved.