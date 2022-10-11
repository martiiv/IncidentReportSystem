/*
SQL File for the incident report system 
Currently contains 7 tables for basic functionality

TODO: Passwords need hashing and script needs testing

Last update 2.10.2022
*/

CREATE TABLE SystemManager (
    SMiD            int             NOT NULL AUTO_INCREMENT,
    Username        varchar(50)     NOT NULL,
    Company         varchar(100),
    Credential      int,

    PRIMARY KEY (SMiD)
);

CREATE TABLE Incident (
    IncidentId      int             NOT NULL AUTO_INCREMENT,
    Name            varchar(50)     NOT NULL,
    Context         varchar(255),
    Company         varchar(100),
    Credential      varchar(255),
    Receiving_group int,
    Countermeasure  varchar(255),
    Sendbymanager   varchar(255),

    PRIMARY KEY (IncidentId)
);

CREATE TABLE WarningReciever(
    WriD            int             NOT NULL AUTO_INCREMENT,
    Name            varchar(255)    NOT NULL,
    PhoneNumber     varchar(255)    NOT NULL,
    Credential      varchar(255),
    Company         varchar(255),
    RecieverGroup   varchar(255),

    PRIMARY KEY (WriD)

);

CREATE TABLE Credentials(
    CiD             int             NOT NULL AUTO_INCREMENT, 
    Email           varchar(255)    NOT NULL,
    Password        varchar(255)    NOT NULL,
    
    PRIMARY KEY(CiD)

);

CREATE TABLE Passwords(
    Password varchar(255)           NOT NULL,
    
    PRIMARY KEY (Password)
);

CREATE TABLE Emails(
    Email	varchar(255)	NOT NULL,
    PRIMARY KEY(Email)
);

CREATE TABLE RecieverGroups (
    Groupid         int             NOT NULL AUTO_INCREMENT,
    Name            varchar(255)     NOT NULL,
    Info            varchar(255),

    PRIMARY KEY (Groupid)
);


    ALTER TABLE SystemManager ADD FOREIGN KEY (Credential)        REFERENCES Credentials(CiD);
        
    ALTER TABLE Incident ADD FOREIGN KEY (Sendbymanager)     REFERENCES SystemManager(Username);
    
    ALTER TABLE Incident ADD FOREIGN KEY (Receiving_group)   REFERENCES RecieverGroups(Groupid);
    
    ALTER TABLE WarningReciever ADD FOREIGN KEY(Credential)         REFERENCES Credentials(CiD);
    
    ALTER TABLE WarningReciever ADD FOREIGN KEY(RecieverGroup)      REFERENCES RecieverGroups(Groupid);
    
    ALter TABLE Credentials ADD FOREIGN KEY(Email)              REFERENCES Emails(Email);
    
    ALTER TABLE Credentials ADD FOREIGN KEY(Password)           REFERENCES Passwords(Password);