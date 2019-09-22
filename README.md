# envoy-grpc-json
Running the example in this repository is, but make sure you have the following prerequisites.

1. Get hold of the Envoy 1.12 binary. One easy way is to build [this sample](https://github.com/envoyproxy/envoy/tree/master/examples/front-proxy), then start the envoy container and just copy the `envoy` binary out using `docker cp`. If you're familiar with envoy, you can use other means of running it like out of a container, etc. Make a note of the directory in which you copied the Envoy binary. Let's save it in a variable called ENVOYDIR.

2. Build the examples in this directory. It's sample:

        make         # to build everything
        make clean   # to clean

3. Run the examples using:

        cd envoy-grpc-json
        ./conf/test.sh $ENVOYDIR

    Look at the conf/test.sh file to understand better how we make the `curl` calls.
