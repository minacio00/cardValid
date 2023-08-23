# cardValid
This is a microservice to check is a credit card number is valid. It expects a POST request to /validateCard with a body containing a json object with only one field called "number" wich value is the card
number to be tested.

The server will respond with an json object containing a field called "valid" which will have the values ​​true or false depending on the validity of the given number
