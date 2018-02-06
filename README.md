# Wiki
Create an http server that allows a user to create, edit, and save wiki pages.

## I. Create functionality 
Use [this](https://golang.org/doc/articles/wiki/) tutorial to implement the web application functionality.

## II. Implement an external database for this Web Application
Web applications commonly use an external database server instead of managing their own data storage directly for a number of reasons:

1. Concurrency: databases servers are designed to deal with many concurrent clients in a correct and predictable manner.
2. Safety: most modern databases are crash safe. An unexpected system failure won't corrupt all your data.
3. Operations: it's often desirable to scale application servers separately from the underlying data storage & querying layer.

While Go makes it _easier_ to deal with concurrent clients than many languages, sharing data safely and correctly is still a complicated problem. Unless you have exceptional circumstances & requirements, the pragmatic choice is to use an existing open source database server.

### The task

In part 1 you built a simple wiki server that allows users to create & edit text files in a directory. In part 2 you will replace the filesystem with a PostgreSQL database.

Unlike the first part, the second will contain a lot less step-by-step instruction, so _please_ ask questions as they come up! An overview of the basic steps:

1. Install Postgres, on a Mac we recommend Postgres.app (there are similar distributions for Windows), on Linux use your system package manager.
2. Create an empty database.
3. Create a database table to hold wiki pages.
4. Import "database/sql" in your Go code, and ensure you have working database connection.
4. Replace the parts of your Go code that load/save wiki pages to use the database table instead of the file system.
5. ???
6. Profit!
