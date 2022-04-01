# Error Groups

The program in this lab is intended to navigate to a web page and discover all of the destinations that it links to via anchor tags. It will then recursively visit those pages until a specified navigation depth is reached. The titles of the resulting tree of pages is then to be exported to a JSON file for future processing.

## Task
Implement the `createGoFunc` and `visitSite` functions to finish this program. Also, a placeholder comment is present in `main()`. Replace this comment with the correct code to kick-off the analysis.


* The `getTitleAndChildren` function has been provided to aid in the extraction of the title and child pages for a URL. 
* An example of the program's output can be found in `./sample.json` 

