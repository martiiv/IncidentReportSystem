/*
SQL File for the incident report system 
Currently contains 7 tables for basic functionality

TODO: Passwords need hashing and script needs testing

Last update 2.10.2022
*/

--Table System Manager, main user of the syste
CREATE TABLE SystemManager (
    SMiD            int             NOT NULL AUTOINCREMENT,
    Username        varchar(50)     NOT NULL,
    Company         varchar(100),
    Credential      varchar(255)
    
    PRIMARY KEY (SMiD),
    FOREIGN KEY (Credential)        REFERENCES Credentials(CiD)
);

--Table Incident, the system manager will distribute these to WarningRecievers
CREATE TABLE Incident (
    IncidentId      int             NOT NULL AUTOINCREMENT,
    Name            varchar(50)     NOT NULL,
    Context         varchar(max)
    Company         varchar(100),
    Credential      varchar(255),
    Receiving_group int,
    Countermeasure  varchar(max),
    Sendbymanager   varchar(255),
    
    PRIMARY KEY (IncidentId),
    FOREIGN KEY (Sendbymanager)     REFERENCES SystemManager(Username),
    FOREIGN KEY (Receiving_group)   REFERENCES Groups(Groupid)
);

--Table Groups, SystemManager will group WarningRecievers and send incidents
CREATE TABLE Groups (
    Groupid         int             NOT NULL AUTOINCREMENT,
    Name            varchar(50)     NOT NULL,
    Info            varchar(max),
    
    PRIMARY KEY (id)

);


--Table WarningReciever, will recieve Incidents from SystemManager
CREATE TABLE WarningReciever(
    WriD            int             NOT NULL AUTOINCREMENT,
    Name            varchar(255)    NOT NULL,
    PhoneNumber     varchar(255)    NOT NULL,
    Credential      varchar(255),
    Company         varchar(255),
    Group           varchar(255),

    PRIMARY KEY(WriD),
    FOREIGN KEY(Credential)         REFERENCES Credentials(CiD),
    FOREIGN KEY(Group)              REFERENCES Groups(Groupid)
)

--Table Credentials, contains the email and password for SystemManagers
CREATE TABLE Credentials (
    CiD             int             NOT NULL AUTOINCREMENT, 
    Email           varchar(255)    NOT NULL,
    Password        varchar(255)    NOT NULL, 

    PRIMARY KEY(CiD),
    FOREIGN KEY(Email)              REFERENCES Emails(Email),
    FOREIGN KEY(Password)           REFERENCES Passwords(Password)
)

--Table Password, NEEDS HASHING
CREATE TABLE Passwords(
    Password varchar(255)           NOT NULL

    PRIMARY KEY(Password)
)

--Table Email
CREATE TABLE Emails(
    Email   varchar(255)            NOT NULL

    PRIMARY KEY(Email)
)