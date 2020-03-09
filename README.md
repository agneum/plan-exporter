## Query plan sender

The utility receives PostgreSQL Query Plans and sends them to a visualizer for sharing. 
It will compute and highlight the most important information to make them easier to understand.

### Pre-requirements

- Go 1.13+

### Installation

``` 
go install github.com/agneum/plan-sender
```

### Usage  

* Run psql
* Set up output to the plan sender:
    ```
    postgres=# \o | plan-sender
    ```
* Run explain to post the plan to visualizer and get a link.
    ```
    postgres=# explain select 1;
    postgres=# Posting to Dalibo...
    The plan has been posted successfully.
    URL: https://explain.dalibo.com/plan/XXX
    postgres=#
    ```

