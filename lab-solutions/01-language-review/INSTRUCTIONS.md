# Language Review Lab

In this lab, you will be finishing the implementation of an asynchronous logging service. This service is intended to receiver HTTP messages and write the request body's content out the a log file. You have been provided with the API and partial implementation of many functions and methods. Your goal is to finish the logging service and discuss design decisions that you made.

## Complete log package

### Implement the log#Run function

The `Run` function is responsible for configuring the package and getting it ready to process log writing requests.

* Use the provided `dest` parameter to set the log file's location, stored in the package-level variable `destination`
* Register the package level `logWriter`, `lw` as the destination for log writing requests.

### Implement the logWriter#Write method

The `logWriter#Write` method is used by the `log` package to determine how messages will actually be written. Implement this method in such a way as to send write requests to the file at the location stored in the `destination` variable.

## Implement service handler

The `service` package is designed to handle all HTTP interactions with this service. 

* Examine the package's structure and be prepared to discuss the role of the `RegisterHandlers` function. ***What are the advantages and disadvantages of the approach used?***

### Prevent unsupported methods

* The service handler is intended to only handle POST requests. Add a guard to the `handleMessage` function that will send an appropriate status message if any other method is used.

### Implement the service behavior

* Read the contents of the request body
    * `io` has functions that can simplify this
* Write the request's body the the log
* Send appropriate responses to the requestor depending on the result of the request using the `ResponseWriter#WriteHeader` method


## Analyze main() function
* Review the main() function and be prepared to discuss its use of the following:
    * the flag package
    * contexts
