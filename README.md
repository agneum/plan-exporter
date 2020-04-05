# Query plan exporter

[![Build Status](https://github.com/agneum/plan-exporter/workflows/build/badge.svg)](https://github.com/agneum/plan-exporter/actions)
[![License](https://img.shields.io/github/license/agneum/plan-exporter)](/LICENSE)
[![Go Report](https://goreportcard.com/badge/github.com/agneum/plan-exporter)](https://goreportcard.com/badge/github.com/agneum/plan-exporter)

The utility receives Postgres Query Plans from `psql` and sends them to a visualizer for sharing. 
It will compute and highlight the most important information to make them easier to understand.

## Installation

### Binary
It's highly recommended installing a specific version of the `plan-exporter` available on the [releases page](https://github.com/agneum/plan-exporter/releases).

### Go
Version `1.13+` is required.

``` 
go install github.com/agneum/plan-exporter
```

## Usage

* Run `psql`
* Set up output to the query plan exporter:
    ```bash
    postgres=# \o | plan-exporter
    ```
* Run explain to post the plan to visualizer and get a link.
    ```bash
    postgres=# explain select 1;
    postgres=# Posting to visualizer...
    The plan has been posted successfully.
    URL: https://explain.depesz.com/s/XXX
    postgres=#
    ```
  
## Command-Line Options

- `--target` - (string, optional) - defines a visualizer to export query plans. 
  Available targets:
  - `depesz` - https://explain.depesz.com [default]
  - `dalibo` - https://explain.dalibo.com

## Contact Information

Follow the news and releases on [twitter](https://twitter.com/arkartasov).

