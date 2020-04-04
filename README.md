## Query plan exporter for `psql`

The utility receives Postgres Query Plans and sends them to a visualizer for sharing. 
It will compute and highlight the most important information to make them easier to understand.

### Pre-requirements

- Go 1.14+

### Installation

``` 
go install github.com/agneum/plan-exporter
```

### Usage  

* Run psql
* Set up output to the query plan exporter:
    ```
    postgres=# \o | plan-exporter
    ```
* Run explain to post the plan to visualizer and get a link.
    ```
    postgres=# explain select 1;
    postgres=# Posting to Dalibo...
    The plan has been posted successfully.
    URL: https://explain.dalibo.com/plan/XXX
    postgres=#
    ```
  
### Options

- `--target` - (string, required) - defines a visualizer to export query plans. 
  Available targets:
  - `dalibo` - https://explain.dalibo.com [default]

