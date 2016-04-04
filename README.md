# Actuatorbeat

Welcome to Actuatorbeat.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/kussj`

## Getting Started with Actuatorbeat

### Init Project
To get running with Actuatorbeat, run the following command:

```
make init
```

To commit the first version before you modify it, run:

```
make commit
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Actuatorbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/kussj/actuatorbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Actuatorbeat run the command below. This will generate a binary
in the same directory with the name actuatorbeat.

```
make
```


### Run

To run Actuatorbeat with debugging output enabled, run:

```
./actuatorbeat -c actuatorbeat.yml -e -d "*"
```


### Test

To test Actuatorbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`


### Package

To cross-compile and package Actuatorbeat for all supported platforms, run the following commands:

```
cd dev-tools/packer
make deps
make images
make
```

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/actuatorbeat.template.json and etc/actuatorbeat.asciidoc

```
make update
```


### Cleanup

To clean  Actuatorbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Actuatorbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/kussj
cd ${GOPATH}/github.com/kussj
git clone https://github.com/kussj/actuatorbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).
