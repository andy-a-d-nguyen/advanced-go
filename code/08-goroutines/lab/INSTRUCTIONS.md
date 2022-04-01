# Goroutines

In this lab, you will complete an application that is performing an ETL (extract, transform, and load) operation. The application's goal is to process a data extract from the [NOAA Local Climatological Database (LCD)](https://www.ncdc.noaa.gov/cdo-web/datatools/lcd) and generate a report indicating the minimum, maximum, and average temperature for each day in the dataset. In order to maximize processing speed, this application will perform its calculations using multiple goroutines.

Hint: To simplify the creation of this program, it is recommended to implement a single step of the program at a time and output a representative sample of the output to stdout. To facilitate this `sample.csv`, containing the first 10 records of the dataset, has been provided.

## Implement createRecords function

The `processCSV` function is responsible for iterating through the CSV records and converting them into `Record` types. The function spawns 10 `createRecords` goroutines to actually perform the processing.

Implement the `createRecords` function. Make sure that all conversions are properly error-checked. The constants defined before the `createRecords` function provide the indexes and date conversion format required for this.

To test this step, uncomment the STEP ONE code in `main()`

## Implement processRecords function 

The `processRecords` function receives the `Record` objects and needs to distribute them to `processDay` goroutines that will each be responsible for processing one day's worth of data.

* recomment the STEP ONE code in `main()`, if required
* update the goroutine within `processRecords` as follows:
    * create a `map[string]chan Record` to keep track of the channels that will distribute `Records` to the proper goroutine for analyzing each `Record`'s day.
    * create a for loop that will receive messages from `recordCh`
    * within the loop, check the map for the existence of a channel that is receiving records for the current `Record`'s day, using the date string as the key (see hint below). If an entry doesn't exist, create the channel, store it in the map, increase the WaitGroup's counter by one, and spawn a `processDay` goroutine passing the newly created channel as the second parameter.
        * Hint: calling the `Format()` method on a time object with the string "2006-01-02" will return a string in yyyy-MM-dd format. This is what the `processDay` function expects as its first parameter and what should be used for the map's key.
    * send each `Record` into the correct channel
        * use the map created in the first step to retrieve the correct channel
    * when the for loop that is ranging over the `recordCh` channel has completed, then the channels held by the map have received all of their messages. Close each channel in the map.
    * when the `processDay` goroutines have completed (hint, that's what the sync.WaitGroup is tracking) all of the `resultCh` messages will have been sent. Close the `resultCh` to signal this.
* review the `processRecords` function and be prepared to discuss why the `resultCh` variable was created here instead of being passed as a parameter
