# Profiling

The program created in the goroutines lab contains a large amount of concurrent tasks. In this lab, you'll analyze a slightly modified version of that lab's solution to what aspects are using the most resources.

* write a benchmarking test that only tests the time required to run the `execute` function
    * use the `bald-mountain_co.csv` file for raw data, but load it's data only one time (e.g. outside of the benchmark loop)
    * Note: the application closes channels when it's done with them. You will need to recreate the channels for each iteration
        * make sure that this time is excluded from the benchmark time

    
* Run the benchmarking test and determine how much time the analysis takes as well as how much memory is required
* Determine which part of the program consumes the largest amount of memory
* Determine which part of the program requires the greatest amount of CPU time
*  Determine what part of the program is the source of the greatest delays 
    * Hint: use a block profile
