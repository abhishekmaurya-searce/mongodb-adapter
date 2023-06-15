# mongo-adapter
    Currently I have tested it on one mongo database but I have designed it to handle any mongo database without the nesting, nesting I will implement after I test it on other databases.
## Prereq
    Golang v1.19,podman,make
## Steps for runnning
    Step 0: `go mod tidy`, in cmd/library/part2.go file delete everything but the function signature and make "n"==1 in times.json file.
    Step 1: `make gorun`
    step 2: `make emulator_run`
    step 3: In another terminal at same directry `make gorun`
    step 4: To check if inserted `make spanner_cli` it will take to to spanner command line interface
    step 5: Write SQL quries to show items `SELECT * FROM author;` `SELECT * FROM books`
