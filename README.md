# README

## Package Sorter in Go
kirstenruge@gmail.com

## Project description

This is a simple Go application that fulfills the requirements given to me,
it will take in a set of 4 inputs (numbers only) and return a string indicating
which stack a package with those dimensions belongs in: REJECTED, SPECIAL or
STANDARD.


## Instructions for using Package Sorter

Run the program with `` go run cmd/main.go <width> <height> <length> <mass> `` from the root directory of the project,
where width, height and length are numbers measured in cm and mass is a non-zero number measured in kg.

Test the program with ``go test``. Feel free to add more testcases in the ``sorter_test.go`` file.
