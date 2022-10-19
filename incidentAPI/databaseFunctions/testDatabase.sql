# Inserting recieverGroups
INSERT INTO `ReceiverGroups` (`Name`, `Info`) VALUES ("Human Resources","Group for everyone working in the HR departments in the company");
INSERT INTO `ReceiverGroups` (`Name`, `Info`) VALUES ("Information Security","Group for everyone working with cybersecurity in the company"); 
INSERT INTO `ReceiverGroups` (`Name`, `Info`) VALUES ("Marketing","Everyone working in the Marketing departments in the company");
INSERT INTO `ReceiverGroups` (`Name`, `Info`) VALUES ("Development","Group for everyone working with development in the company");

#Inserting emails for 4 warning recievers and one System manager
INSERT INTO `Emails` (`Email`) VALUES ("HalvorHr@gmail.com"); #WR 1
INSERT INTO `Emails` (`Email`) VALUES ("IvarIt@gmail.com"); #WR 2
INSERT INTO `Emails` (`Email`) VALUES ("MonicaMarketing@gmail.com"); #WR 3
INSERT INTO `Emails` (`Email`) VALUES ("UlrikUtvikler@gmail.com"); #WR 4
INSERT INTO `Emails` (`Email`) VALUES ("OdaOverordna@gmail.com"); #System manager

#Inserting password Needs hashing
INSERT INTO  `Passwords`(`Password`) VALUES ("123HR"); #WR 1
INSERT INTO  `Passwords`(`Password`) VALUES ("123IT"); #WR 2
INSERT INTO  `Passwords`(`Password`) VALUES ("123Marketing"); #WR 3
INSERT INTO  `Passwords`(`Password`) VALUES ("123Development"); #wR 4
INSERT INTO  `Passwords`(`Password`) VALUES ("123Sjefen"); #System manager


INSERT INTO `Credentials` (`Email`, `Password`) VALUES ("OdaOverordna@gmail.com", "123Sjefen"); #System manager 
INSERT INTO `Credentials` (`Email`, `Password`) VALUES ("HalvorHr@gmail.com", "123HR"); #WR 1
INSERT INTO `Credentials` (`Email`, `Password`) VALUES ("IvarIt@gmail.com", "123IT"); #WR2
INSERT INTO `Credentials` (`Email`, `Password`) VALUES ("MonicaMarketing@gmail.com", "123Marketing"); #WR 3 
INSERT INTO `Credentials` (`Email`, `Password`) VALUES ("UlrikUtvikler@gmail.com", "123Development"); #WR 4

#Creating warning reciever
INSERT INTO `WarningReceiver` (`WriD`, `Name`, `PhoneNumber`, `CredentialId`, `Company`, `ReceiverGroup`, `ReceiverId`) 
VALUES 
('1', 'Monica', '58328261', '4', 'IncidentCorp', 'Marketing', '3'), 
('2', 'Ulrik', '78590153', '5', 'IncidentCorp', 'Development', '4'), 
('3', 'Halvor', '69874569', '2', 'IncidentCorp', 'Human Resources', '1'), 
('4', 'Ivar', '69874569', '3', 'IncidentCorp', 'Information Security', '2');

#Creating System Manager
INSERT INTO `SystemManager` (`SMiD`, `Username`, `Company`, `Credential`) VALUES ('1', 'OdaManager', 'IncidentCorp', '1');

#Creating test incidents
INSERT INTO `Incident` (`IncidentId`, `Tag`, `Name`, `Description`, `Company`, `Credential`, `Receiving_group`, `Countermeasure`, `Sendbymanager`, `Date`) 
VALUES 
('1', 'Hazard', 'Broken door', 'A door has been broken in office five', 'IncidentCorp', NULL, '1', 'Contact janitor, Fix door', 'OdaManager', CURRENT_TIMESTAMP), 
('2', 'Phishing', 'Hack attack!', 'An email from an unknown party has sent out a malicious email containing malware!', 'IncidentCorp', NULL, '2', 'Do not open email, Block sender ', 'OdaManager', '2022-10-18 11:49:55');