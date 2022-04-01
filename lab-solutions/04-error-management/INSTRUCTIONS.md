# Error Groups

The program in this lab is intended to navigate to a web page and discover all of the destinations that it links to via anchor tags. It will then recursively visit those pages until a specified navigation depth is reached. The titles of the resulting tree of pages is then to be exported to a JSON file for future processing.

## Task
Implement the `createGoFunc` and `visitSite` functions to finish this program. Also, a placeholder comment is present in `main()`. Replace this comment with the correct code to kick-off the analysis.

* `createGoFunc` needs to return a function that is compatible with `errgroup.Group#Go` but also has access to the parameters that are passed to it. Take advantage of closures to accomplish this.
* `visitSite` has several responsibilities
    * retrieve the title and children from the `p` 
        * `getTitleAndChildren` can perform this operation
    * the `page` object should be updated with the title and children for the current page
    * while the current navigation depth is less than the maximum depth (controlled by the `maxDepth` constant), `createGoFunc` should be called for each child and registered with the `errgroup.Group`.
        * Remember to increase the current depth!
    * Any errors that are generated should be returned to the caller for handling


* An example of the program's output can be found in `./sample.json` 

