# CrowdCompute

<!--
repo name: CrowdCompute
description: An distrubuted computing platform. People can share their CPU and GPU power to help with tasks requiring heavy load.
github name:  Basileus1990
link: https://github.com/Basileus1990/CrowdCompute
email: pawelb021@gmail.com
-->

<!-- TABLE OF CONTENTS -->
## Table of Contents

* [About the Project](#about-the-project)
    * [Built With](#built-with)
    * [Architecture graph](#architecture-graph)
    * [TODO](#TODO)
    * [Ideas](#ideas)
* [Getting Started](#getting-started)
    * [Prerequisites](#prerequisites)
    * [Installation](#installation)
* [Usage](#usage)
* [Roadmap](#roadmap)
* [Contributing](#contributing)
* [License](#license)
* [Contact](#contact)
* [Acknowledgements](#acknowledgements)



<!-- ABOUT THE PROJECT -->
## About The Project

An distrubuted computing platform. People can share their CPU and GPU power to help with tasks requiring heavy load.
    * Task creators can upload their tasks to the platform
    * Task executors can contribute to the tasks through primarly the web browser, but also CLI tool or windowed application

### Built With
* [GO](https://golang.org/)
* [Typescript](https://www.typescriptlang.org/)
* [Vue.js](https://vuejs.org/)

### Architecture graph
![Architecture graph](./docs/architecture_graph.png)
### TODO:
0. Planning phase
    * [X] Check how to do computations in the browser
        * create sandboxed environment and run given code there
        * In future check how to do it with GPU
    * [X] Check how to break big tasks into smaller ones, so they can be shared
        * Basically the task creator will write JS code which will be executed in the user's browser
    * [X] Create graph of the project
    
1. Frontend
    * [ ] Create basic website UI (purely functional)
        * [X] Part for adding new tasks
        * [X] Part for viewing tasks and selecting them
        * [X] executing of tasks
        * [ ] viewing results of tasks
        * [ ] make them call the server
2. Backend 
    * [ ] Create basic website backend
        * [X] Make API for adding tasks and add them (temporarily) to map in memory
        * [ ] Make API for getting all tasks
        * [ ] Make API for receiving results of tasks
        * [ ] Make API for seeing results of tasks

3. Database
    * [ ] Create database
        * [X] Make it possible to add tasks to the database
        * [X] Make it possible to get tasks from the database
            * [X] Consider tasks that have the same name. Should title be primary key or something else?
        * [ ] Make it possible to add results of tasks to the database
        * [ ] Make it possible to get results of tasks from the database

#### First step
    * Make it possible to add tasks, see the list of them as a user and as an author
    * Make it possible to execute tasks as a user
    * Make it possible to see the results of the tasks as an author
    * Start learning about postgreSQL and how to use it with GO

### Ideas
1. Possibly task creators could pay for the task execution. The payment would be split between the task executors.
2. Create an API for using the platform outside the browser
3. Add possiblity to use other languages than JS
4. Add possiblity to use GPU for computations


<!-- GETTING STARTED -->
## Getting Started
% To be added

### Prerequisites
% To be added

### Installation
% To be added

<!-- USAGE EXAMPLES -->
## Usage
% To be added


<!-- ROADMAP -->
## Roadmap
% To be added



<!-- CONTRIBUTING -->
## Contributing
% To be added



<!-- LICENSE -->
## License
% To be added



<!-- CONTACT -->
## Contact
% To be added



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
% To be added
