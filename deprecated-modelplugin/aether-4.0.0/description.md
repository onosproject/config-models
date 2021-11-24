## Background

The Aether ROC control API is available via REST or via gNMI. It is expected that most external
consumers of the API will use REST.

The REST API supports the typical GET, POST, PATCH, DELETE operations:

* GET. Retrieve an object.
* POST. Create an object.
* PUT,  PATCH. Modify an existing object.
* DELETE. Delete an object.

Endpoints are named based on the type of object. Some examples:

* `GET http://roc/aether/v4.0.0/connectivity-service-v4/enterprise/`. Get a list of enterprises.
* `GET http://roc/aether/v4.0.0/connectivity-service-v4/enterprise/Starbucks`. Get the Starbucks enterprise.
* `POST http://roc/aether/v4.0.0/connectivity-service-v4/enterprise`. Create a new enterprise.
* `PATCH http://roc/aether/v4.0.0/connectivity-service-v4/site/Starbucks-NewYork`. Update the Starbucks New York site.

This document is a high-level description of the objects that can be interacted with. For a
low-level description, see the Swagger API specification.

## Identifying and Referencing Objects

Every object contains an `id` that is used to identify the object. The `id` is only unique within
the scope of a particular type of object. For example, a site may be named `foo` and a device-group
may also be named `foo`, and the two names do not conflict because they are different object types.

Some objects contain references to other objects. For example, many objects contain references to
the `Enterprise` object, which allows them to be associated with a particular enterprise. References
are constructed using the `id` field of the referenced object. It is an error to attempt to create
a reference to an object that does not exist. Deleting an object while there are open references
to it from other objects is also an error.

## Common Model Fields

Several fields are common to all models in Aether:

* `id`. The identifier for objects of this model.
* `description`. A human-readable description, used to store additional context about the object.
* `display-name`. A human-readable name that is shown in the GUI.

As these fields are common to all models, they will be omitted from the per-model descriptions below.

## Key Aether Objects

The following is a list of Aether models, generally organized in a top-down manner.

### Enterprise

`Enterprise` forms the root of a customer-specific Enterprise hierarchy. The `Enterprise` model is
referenced by many other objects, and allows those objects to be scoped to a particular Enterprise
for ownership and role-based access control purposes. `Enterprise` contains the following fields:

* `connectivity-service`. A list of connectivity services that realize connectivity for this
  enterprise. A connectivity service is a reference to the SD-Core, and reflects either a 4G or a
  5G core.

### Site

`Enterprises` are further divided into `Sites`. A site is a point of presence for an `Enterprise` and
may be either physical or logical (i.e. a single geographic location could in theory contain several
logical sites). Site contains the following fields:

* `enterprise`. A link to the `Enterprise` that owns this site.
* `imsi-definition`. A description of how IMSIs are constructed for this site. Contains the following
  sub-fields:

   * `mcc`. Mobile country code.
   * `mnc`. Mobile network code.
   * `enterprise`. A numeric enterprise id.
   * `format`. A mask that allows the above three fields to be embedded into an IMSI. For example
     `CCCNNNEEESSSSSS` will construct IMSIs using a 3-digit MCC, 3-digit MNC, 3-digit ENT, and a
     6-digit subscriber.
   
* `small-cell` A list of 5G gNodeB or Access Point or Radios. Each access point has the following:

  * `address`. Hostname of the small cell.
  * `tac`. Type Allocation Code.
  * `enable`. If set to `true`, the small cell is enabled. Otherwise, it is disabled.

* `monitoring` Configuration of how the monitoring framework of the site can be connected:

  * `edge-cluster-prometheus-url` the URL of the site's Edge cluster Prometheus service
  * `edge-monitoring-prometheus-url` the URL of the site's Edge monitoring Prometheus service
  * `edge-device` a list of monitoring devices that verify end-to-end connectivity

    * `name` the name of the monitoring device (usually a Raspberry Pi). This will usually
      correspond to its short hostname
    * `display-name` a user friendly name
    * `description` an optional description

### Device-Group

`Device-Group` allows multiple devices to be logically grouped together. `Device-Group` contains
the following fields:

* `imsis`. A list of IMSI ranges. Each range has the following
  fields:

   * `name`. Name of the range. Used as a key.
   * `imsi-range-from`. First subscriber in the range.
   * `imsi-range-to`. Last subscriber in the range. Can be omitted if the range only contains one
     IMSI.
* `ip-domain`. Reference to an `IP-Domain` object that describes the IP and DNS settings for UEs
  within this group.
* `site`. Reference to the site where this `Device-Group` may be used. Indirectly identifies the
  `Enterprise` as `Site` contains a reference to `Enterprise`.

* `device` related configuration:

  * `mbr`. The maximum bitrate in bits per second that the application will be limited to:

    * `uplink` the `mbr` from device to slice
    * `downlink` the `mbr` from slice to device

### Virtual Cellular Service

`Virtual Cellular Service (VCS)` connects a `Device-Group` to an `Application`. `VCS` has the
following fields:

* `device-group`. A list of `Device-Group` objects that can participate in this `VCS`. Each
  entry in the list contains both the reference to the `Device-Group` as well as an `enable`
  field which may be used to temporarily remove access to the group.
* `filter`. A list of `Application` objects that are either allowed or denied for this
  `VCS`. Each entry in the list contains both a reference to the `Application` as well as an
  `allow` field which can be set to `true` to allow the application or `false` to deny it. It
  also has a `priority` field which can be used to order the applications when considering the
  enforcing of their `allow` condition.
* `template`. Reference to the `Template` that was used to initialize this `VCS`.
* `upf`. Reference to the User Plane Function (`UPF`) that should be used to process packets
  for this `VCS`. It's permitted for multiple `VCS` to share a single `UPF`.
* `ap`. Reference to an Access Point List (`AP-List`) that lists the access points for this
  `VCS`.
* `enterprise`. Reference to the `Enterprise` that owns this `VCS`.
* `SST`, `SD`, `uplink`, `downlink`, `traffic-class`. Parameters that were initialized using the
  `template`. They are described in the section for the `Template` model.

### Application

`Application` specifies an application and the endpoints for the application. Applications are
the termination point for traffic from the UPF. Contains the following fields:

* `address`. The DNS name or IP address of the endpoint.
* `endpoint`. A list of endpoints. Each has the following fields:

  * `name`. Name of the endpoint. Used as a key.
  * `port-start`. Starting port number.
  * `port-end`. Ending port number.
  * `protocol`. `TCP|UDP`, specifies the protocol for the endpoint.
  * `mbr`. The maximum bitrate in bits per second that the application will be limited to:

    * `uplink` the `mbr` from device to application
    * `downlink` the `mbr` from application to device

* `enterprise`. Link to an `Enterprise` object that owns this application. May be left empty
  to indicate a global application that may be used by multiple enterprises.

## Supporting Aether Objects

### Connectivity-Service

`Connectivity-Service` specifies the URL of an SD-Core control plane.

* `core-5g-endpoint`. Endpoint of a `config4g` or `config5g` core.

### IP-Domain

`IP-Domain` specifies IP and DNS settings and has the following fields:

* `dnn`. Data network name for 5G, or APN for 4G.
* `dns-primary`, `dns-secondary`. IP addresses for DNS servers.
* `subnet`. Subnet to allocate to UEs.
* `admin-status`. Tells whether these ip-domain settings should be used, or whether they
  should be drained from UEs.
* `mtu`. Ethernet maximum transmission unit.
* `enterprise`. `Enterprise that owns this `IP-Domain`.

### Template

`Template` contains connectivity settings that are pre-configured by Aether Operations.
Templates are used to initialize `VCS` objects. `Template` has the following fields:

* `sst`, `sd`. Slice identifiers.
* `uplink`, `downlink`. Guaranteed uplink and downlink bandwidth.
* `traffic-class`. Link to a `Traffic-Class` object that describes the type of traffic.
* `slice` configuration for slice

  * `mbr` configuration of maximum bit rate in slice

    * `uplink` the `mbr` from device to slice
    * `downlink` the `mbr` from slice to device

### Traffic-Class

Specifies the class of traffic. Contains the following:

* `arp`. Allocation and Retention Priority.
* `qci`. QoS class identifier.
* `pelr`. Packet error loss rate.
* `pdb`. Packet delay budget.

### UPF

Specifies the UPF that should forward packets. A UPF can only be used by one VCS at a time.
Has the following fields:

* `address`. Hostname or IP address of UPF.
* `port`. Port number of UPF.
* `enterprise`. Enterprise that owns this UPF.
* `site`. The Site that this UPF is located at.
* `config-endpoint` URL for configuring the UPF
