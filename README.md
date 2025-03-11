# Expense Tracker CLI

A Task Tracker Command Line Interface created with **GO** language as a beginner project while learning the syntax and the structural typing shape of the go language. 
this project is done in reference to [roadmap.sh](https://roadmap.sh/projects/expense-tracker) project recommendation.


# Running the CLI

after cloning into the project, open up a terminal and CD into the program with 
```
cd expense-tracker
```
and you're good to **GO**

## Structure 
the project is using ``.json`` file as the expenses list and budget list. the CLI will update the ``.json`` file

## Command list

to see the command list available for the program you can use
```
./expense-tracker
```
this will show list of available command which is :

 ```
 1. add
 2. delete
 3. list
 ```

for the CRUD of the tasks. with their specified flags

## Notable commands and information

``list`` will show all the tasks in the expense list, and you can also pass in an argument to filter the list by number of ``month`` such as
```
./expense-tracker list 1
```
this will show all the expenses in the list with month of ``January``


##
**Happy Expense-ing**
##


