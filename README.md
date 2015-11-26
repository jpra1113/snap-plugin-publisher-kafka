<!--
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->

## Snap Plugin - Kakfa Publisher


### Description

Allows publishing of data to [Apache Kafka](http://kafka.apache.org)

### Dependencies

Uses [sarama](http://shopify.github.io/sarama/) golang client for Kafka by Shopify

## Configuration details

| key      | value type | required | default  | description  |
|----------|------------|----------|----------|--------------|
| topic | string | yes | | The topic to send messages |         
| brokers  | string | yes | | Semicolon delimited list of "server:port" brokers |
