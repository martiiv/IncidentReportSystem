# Incident Report System

This is the readme file for our applied computer science project. It is an Incident report system where System Managers can create and send warnings to people in organizations. It consists of an API written in Golang, an MySQL database as well as a React front end website. The project was developed by Group 10 (Aleksander Aaboen, Tormod Mork Müller, Chrysovalantis Pateiniotis and Martin Iversen) in the course IMT4886 at NTNU Gjøvik.

For a full overview of the project please read the project report found in the documentation folder.

This readme file contains:

- Installation guide
- Missing functionality
- Brief API Documentation

## Installation

In order to run the application you need a .go file containing database login credentials.
This file needs to be placed in a config folder inside the incidentAPI folder
This file needs to be formated like so:

```golang
    //Predefined global variables for database and email configuration
    var DB_NAME = "dbName"
    var DB_HOST = "ipForHostServer"
    var DB_USERNAME = "dbUserName"
    var DB_PASSWORD = "dbPassword"
    var SenderEmailAppPassword = "Password"
```

After this simply run the main.go file to run the API
After running the API to access the Front end application run the app.js file
If everything is configured properly your terminal should print out

```cmd
    Now connecting...
    connected...
    Listening on port::8080
```

If you are struggling to run the application contact one of the teammembers via GitLab

## Missing functionality

- We are missing API authentication
- Endpoint structure needs versioning
- There is no log -report functionality, only individual lessons learned can be updated
- Error handling can be improved
- The applications front-end cannot create tags or update countermeasures

## API Documentation

Below you will find explanations of the endpoints in the API for the incidents, warning receivers, system managers, receiving groups, countermeasures and tags

### Incident

---

### GET Incident

Method gets all incidents in the database, pass in an ID or tag variable to get one specific incident or incidents related to a tag:

```cmd
    /incident
    /incident?id={id}
    /incident?tag={tag}
```

```json
    {
        "id":110,
        "tag":"string",
        "name":"string",
        "description":"string",
        "company":"string",
        "receivingGroup":"string",
        "countermeasure":"string",
        "sendbymanager":"string",
        "lessonlearned":"string",
    }
```

#### POST Incident

Method creates new incidents and sends them to a group of warning receivers

```cmd
    /incident
```

Body:

```json
    {
        "tag": "string",
        "name":  "string",
        "description": "string",
        "company": "string",
        "receivingGroup": "string",
        "sendByManager": "string",
        "lessonlearned": "string"
    }
```

The endpoint will return the newly created incidents ID

### DELETE Incident

Endpoint will delete one or more incidents

```cmd
    /incident
```

Body:

```json
   [{
        "incidentId": "string",
        "incidentName" : "string"
    }]
```

The endpoint will return the newly deleted incidents ID

---

## Receiving group

---

### GET Receiving group

Method gets all groups in the database, pass in an ID to get one specific group:

```cmd
    /groups
    /groups?id={id}
```

```json
    {
        "id": 1, 
        "name": "Human Resources", 
        "info": "Group for everyone working in the HR departments in the company"
    }
```

### POST Receiving group

Method creates new groups

```cmd
    /groups
```

Body:

```json
    {
        {
            "name": "string",
            "info": "string"
        }
    }
```

The endpoint will return the newly created groups ID

### DELETE Receiving group

Endpoint will delete one or more groups

```cmd
    /manager
```

Body:

```json
   [{
        "id": "",
        "name" : "TestGroupAPITEST"
    }]
```

The endpoint will return the newly deleted groups ID or name

---

## System manager

---

### GET System manager

Method gets all system managers in the database, pass in an ID to get one specific manager:

```cmd
    /manager
    /manager?id={id}
```

```json
    {
        "id": 1, 
        "userName": "OdaManager",
        "company": "IncidentCorp",
        "credential": "1"
    }
```

### POST System manager

Method creates new managers

```cmd
    /manager
```

Body:

```json
    {
        "userName": "TestManagerAPITEST",
        "company":"IncidentCorp",
        "email":"testManager@gmail.com",
        "password": "1241erreth23e23r1231"
    }
```

The endpoint will return the newly created managers ID

### DELETE System manager

Endpoint will delete one or more managers

```cmd
    /manager
```

Body:

```json
   [{
        "email" : "testManager@gmail.com"
    }]
```

The endpoint will return the newly deleted managers id

---

## Warning Receiver

---

### GET Warning Receiver

Method gets all system receivers in the database, pass in an ID to get one specific receivers:

```cmd
    /receiver
    /receivers?id={id}
```

```json
    {
        "id":2,
        "name":"Ulrik",
        "phoneNumber":"78590153",
        "company":"IncidentCorp",
        "receiverGroup":"Development",
        "receiverEmail":"UlrikUtvikler@gmail.com"
    }
```

### POST Warning Receiver

Method creates new receivers

```cmd
    /receivers
```

Body:

```json
    {
        "name":"TestReceiverAPITEST",
        "phoneNumber":"12345678",
        "company":"IncidentCorp",
        "receiverGroup":"Marketing",
        "receiverEmail":"APITEST@gmail.com"
   }
```

The endpoint will return the newly created receivers ID

### DELETE Warning Receiver

Endpoint will delete one or more receivers

```cmd
    /receivers
```

Body:

```json
   [{
        "id": "",
        "email":"APITEST@gmail.com"
   }]
```

The endpoint will return the newly deleted receivers id

---
