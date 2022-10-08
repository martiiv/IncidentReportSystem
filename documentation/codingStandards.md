# Coding standard documents

This document contains coding conventions the team will be following throughout the project process. The document will contain examples for how files will be structured, code commenting conventions and other information regarding code structure.

## File structure

---

Each file will contain a header with information regarding the purpose of the file as well as who modified the file last and when.

    /*
    File EXAMPLE, file connects to the database and does FUNCTION ONE ++. The file contains the following functions:
    - F1
    - F2
    - F3
    TODO: File needs code commenting and F2 needs optimization.
    Author Martin Iversen
    Last mod: 03.10.2022
    */

In addition to this header comment, each function will have a small header explaining the functions purpose as well as the return value and parameters.

    /*
    Function F1, the function will take in parameters and send them to the database using an API call
    param: param1 Description
    param: param2 Description
    param: param3 Description
    return: return value
    */

Lastly inline comments will be provided where necessary.
If a function is complex and needs further explanations a comment can be helpful for clearing up confusion.

    //This is an inline comment

## Functions and parameters

---

Function names will be named according to functionality, this also goes for parameters and other variables in the system. We will be using camelCase for naming.

    func importDataFromDataBase(){

    }

## Teamwork

---

Lastly team members will be working in separate files where this is possible to prevent merge conflicts and similar issues. If this is unavoidable pair programming will be used and communication regarding commits and publishing will be enforced to prevent branch conflicts.
