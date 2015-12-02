[![Build Status](https://travis-ci.com/intelsdi-x/snap-plugin-publisher-kafka.svg?token=HoxHq3yqBGpySzRd5XUm)](https://travis-ci.com/intelsdi-x/snap-plugin-publisher-kafka)

# snap publisher plugin - Kafka

Allows publishing of data to [Apache Kafka](http://kafka.apache.org)

It's used in the [snap framework](http://github.com/intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

### System Requirements

- Uses [sarama](http://shopify.github.io/sarama/) golang client for Kafka by Shopify

### Installation

#### Download Kafka plugin binary:
You can get the pre-built binaries for your OS and architecture at snap's [GitHub Releases](https://github.com/intelsdi-x/snap/releases) page.

#### To build the plugin binary:
Fork https://github.com/intelsdi-x/snap-plugin-publisher-kafka  
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-publisher-kafka.git
```

Build the plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `/build/rootfs/`

### Configuration and Usage
* Set up the [snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Ensure `$SNAP_PATH` is exported  
`export SNAP_PATH=$GOPATH/src/github.com/intelsdi-x/snap/build`

| key      | value type | required | default  | description  |
|----------|------------|----------|----------|--------------|
| topic | string | yes | | The topic to send messages |         
| brokers  | string | yes | | Semicolon delimited list of "server:port" brokers |

## Documentation
<< @TODO

### Examples
<< @TODO

### Roadmap

There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release.

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-publisher-kafka/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-publisher-kafka/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions! 

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements

* Author: [Nicholas Weaver](https://github.com/lynxbat)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
