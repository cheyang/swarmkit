# abstract resource support

Currently swarm mode only supports CPU and memory as resource type. This idea is borrowed from [http://mesos.apache.org/documentation/latest/attributes-resources/](http://mesos.apache.org/documentation/latest/attributes-resources/)

It'd be great to allow user specify any resource, bandwidth, GPU count, disk space, or artificial ones. It could be scalar like 0.25 CPU, or discrete like GPU-1 (don't divide GPU). Swarm Mode can handle these resource without knowing the physical meaning. CPU and memory are 2 resource instances.

## Abstract Resource message

Abstract Resource are defined by the `Abstract Resource` protobuf message. A simplified version of this
message, showing only the fields described in this document, is presented below:

```
/**
 * Describes an Attribute or Resource "value". A value is described
 * using the standard protocol buffer "union" trick.
 */
message Value {
  enum Type {
    SCALAR = 0;
    RANGES = 1;
    SET = 2;
    TEXT = 3;
  }

  message Scalar {
    required double value = 1;
  }

  message Range {
    required uint64 begin = 1;
    required uint64 end = 2;
  }

  message Ranges {
    repeated Range range = 1;
  }

  message Set {
    repeated string item = 1;
  }

  message Text {
    required string value = 1;
  }

  required Type type = 1;
  optional Scalar scalar = 2;
  optional Ranges ranges = 3;
  optional Set set = 4;
  optional Text text = 5;
}

/**
 * Describes a resource on a machine. A resource can take on one of
 * three types: scalar (double), a list of finite and discrete ranges
 * (e.g., [1-10, 20-30]), or a set of items. A resource is described
 * using the standard protocol buffer "union" trick.
 */
message AbstractResource {
  required string name = 1;
  required Value.Type type = 2;
  optional Value.Scalar scalar = 3;
  optional Value.Ranges ranges = 4;
  optional Value.Set set = 5;
}
```