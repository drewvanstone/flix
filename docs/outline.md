* Why project structure matters
    * Not about directory structure
    * About package design
* Project structure patterns
    * Flat
    * Domain package
* Flat Project Structure
    * Everything in package main
    * Pros:
        * Don't have to worry about packages
        * No circular dependencies
        * Maintainable up for 1k-10k codebases
    * Cons:
        * Doesn't enforce good code structure
        * Domain logic usually not isolated
        * Code is generally tightly coupled
        * Can become unweildy with dozens of files
* Domain Package Project Structure
    * All domain logic and types in single package
    * Domain package has no external dependencies
    * Domain package defines interfaces that dependencys implement
    * Main should primarily call the domain package
    * Pros:
        * Dependencies are abstracted (easy to swap out)
        * Loosely coupled and easy to mock
        * Scales hundreds of thousands of lines of code
        * Easy to see what the core system is doing
    * Cons:
        * More boiler plate for smal projects
        * Need to know abstraction patterns to work effectively
* Exampe project

