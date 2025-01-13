# Distributed Fault-Tolerant MapReduce Implementation

This project is a distributed, fault-tolerant implementation of the MapReduce programming model, based on Google's MapReduce paper.

## Overview

MapReduce is a programming model and an associated implementation for processing and generating large data sets. Users specify a map function that processes a key/value pair to generate a set of intermediate key/value pairs, and a reduce function that merges all intermediate values associated with the same intermediate key. This repo is the programming model to process mapReduce applications in a distributed and fault-tolerant manner. Users can specify there map and reduce functions and pass them as plugins to the system.

## Getting Started

To get started with this implementation, follow these steps:

1. **Clone the Repository**: 
2. **Install Dependencies**:
3. **Run the Coordinatore**: `go run cmd/coordinator.go input_files/pg-grimm.txt` which runs the coordinator and the second argument is the input file divided across files.
4. **Build the plugin loader**: `go build -buildmode=plugin mrapps/wc.go` which builds the plugin loader for a wod count application.
5. **Run Worker Nodes**: `go run cmd/worker.go wc.so` which runs the worker nodes for this mapReduce application.


## Features
TODO
## Components
TODO
## How It Works
TODO

## References/Motivation

- [Google's MapReduce Paper](https://static.googleusercontent.com/media/research.google.com/en//archive/mapreduce-osdi04.pdf)
- MIT's Distributed Systems Course

## License

This project is licensed under the MIT License.
