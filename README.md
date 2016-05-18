ec2
==========

Go bindings for Amazon EC2 Java SDK

####Example
See example in source code.

####Install
Download/unzip Amazon EC2 Java SDK from https://aws.amazon.com/sdk-for-java/

````
go get github.com/timob/ec2

export LD_LIBRARY_PATH=/usr/lib/jvm/default-java/jre/lib/amd64/server
export CLASSPATH=$(find /aws-java-sdk-dir/ -name \*.jar | head -c -1 | tr \\n :)
cd github.com/timob/ec2/example
go build example1.go
./example1
````
Needs libjvm.so, so that is why LD_LIBRARY_PATH is set (shown here is where this on Ubuntu).

The example prints the list of EC2 instance ids and other info.

