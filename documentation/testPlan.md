# Test plan

This document will contain information regarding the test process throughout the project. We will be testing each component of the project individually before combining the different components.

The database will be tested using unit tests for all the tables in the database. The API will be using functional tests where we simulate calls. Front end will be tested using user testing as the main methodology.

## Database

---

For database testing we will be doing unit tests only. We will be using the standard golang test framework since it is diverse and solid. The tests will mainly focus on checking priviliges and restraints on the different tables.

Since the database is quite small and simple we will not overcomplicate this part of the testing process. Other functionality will be tested through the API.

### Example Database

`INSERT UNIT TEST EXAMPLE HERE`

## API

---

For API testing we will be simulating a server and calling the different parts of the API, ensuring that calls act as expected. These tests will be closely related to the database since the database will communicate with the back-end using an API. We may also do stress testing if we deem this necessary.

### Example API

`INSERT UNIT TEST EXAMPLE HERE`

## Front End

---

For front end testing we will be doing user tests. We will follow methodology from various UX courses¨

- Distribute user agreement form
- Perform curated user tests
- Question the testers using pre defined questions
  
We will collect and analyze the results before revising functionality and design.
