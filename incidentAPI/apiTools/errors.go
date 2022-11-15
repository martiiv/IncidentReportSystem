package apitools

/*
File erros.go
*File contains predefined error messages used in all errorhandling
? Last revision Martin Iversen 15.11.2022
*/

//Generic error message
var UnexpectedError = "Sorry, unexpected error"

//Error message for json encoding and similar code
var EncodeError = "Oops, something went wrong"

//Error message for json decoding
var DecodeError = "Error in passed in JSON body Please see endpoint documentation"

//Error message for query execution
var QueryError = "Invalid Query!"
