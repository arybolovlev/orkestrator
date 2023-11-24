Orkestrator
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/arybolovlev/orkestrator) [![License: MPL 2.0](https://img.shields.io/badge/License-MPL_2.0-brightgreen.svg)](LICENSE)
===

Orkestrator is an application to deploy and manage containers. The project aims to be educational to explore and demonstrate different aspects of container orchestration.

The specification is not fully complete yet and will be updated in the future. In its current state, it aims to provide the bare minimum functionality.

Please, take into account the educational status of this project and DO NOT use it in your production or critical environment.

## CLI Options

### Global options

`-port` -- Port to connect or listen to.

### Client options

`-client` -- Run Orkestrator in the client mode.

`-file` -- Specification file.

### Server options

`-server` -- Run Orkestrator in the server mode.

### Worker options

`-worker` -- Run Orkestrator in the worker mode.

## Specification

Orkestrator consists of a few components that interact with each other to manage containers:

- Worker
- Manager
- Scheduler

The components build a pipeline for working units that describe the desired state of the managed containers:

- Task
- Job

Below is a more detailed explanation of each component and working unit.

### Task

The task is the smallest working unit of the Orkestrator. Essentially, this is a single container running on a worker machine.

A task should specify the following:

- the name and the tag of the container image;
- the resources that the container needs to run(CPU, Memory, Disk);
- a restart policy, or in other words how to handle failuries.

### Job

The job accumulates tasks into a bigger logical group. It consists of a minimum of one task.

A job should specify the following:

- each task that should be a part of the job.

### Worker

The worker is responsible for running the jobs assigned to it by the manager according to the specification. It collects and publishes metrics about the jobs and its own health status for the manager.

The worker should do the following:

- registering on the manager;
- reporting own health status, including available and allocated resources;
- accepting jobs to run from the manager;
- running assigning jobs according to the specification;
- collecting jobs metrics;
- publishing jobs metrics.

### Manager

The manager is the main entry point for all components of the system and external users. The manager should accept job submissions from users. Once a new job specification is available, it calls the scheduler to find a worker to run the job and forwards its decision along with the job to the worker. The manager provides the health status of its components and jobs.

The worker should do the following:

- accepting a job specification from a user;
- accepting worker registration;
- accepting scheduler registration;
- calling the scheduler to assign a job to a worker;
- publishing jobs for the workers;
- collecting jobs metrics from the workers;
- collecting the health status of the workers;
- collecting the health status of the schedulers;
- publishing the health status of its components and jobs.


### Scheduler

The scheduler is responsible to make a decision about where to place a job according to its specification and available workers' resources. It listens to the manager and waits for new unassigned jobs. When the one shows up, it makes a decision on which worker should run the job and report this decision back to the manager.

The scheduler should do the following:

- registering on the manager;
- accepting an unassigned job specification from the manager;
- collecting available workers and their available resources;
- assigning the job to the worker;
- publishing the assigned job;
- reporting own health status.
