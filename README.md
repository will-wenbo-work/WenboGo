# golang project
happy holiday
# How to run 
1, make sure you have go envrionment in your local machine

2, download this package and cd to src

3, run 'go build'

4, run './src'

or you can open this project in vs code and hit f5 and start local debug
# Open APIs
1  POST: http://localhost:8080/event
this api take yaml payload, here's an example payload:
```yaml
title: Valid App 1
version: 0.0.1
maintainers:
- name: firstmaintainer app1
  email: firstmaintainer@hotmail.com
- name: secondmaintainer app1
  email: secondmaintainer@gmail.com
company: Random Inc.
website: https://website.com
source: https://github.com/random/repo
license: Apache-2.0
description: |
 ### Interesting Title
 Some application content, and description
``` 

2  POST: http://localhost:8080/events/?{parameter}={value}&{parameter}={value}..
it returns all saved payloads which meet the search requirement 


# What's cool about this project
1, golang service with YAML payload

2, data deduplicate, we deduplicate data if it's already saved once.

3, data flattener, we flatten the payload to make seach faster.

4, UUID generator

5, concurrent save in data saving, (but not read committed)

6, localing search indexing

7, request/email validation

8, avoid over-engineering

# What's not so cool about this project
1, test coverage (WIP)

2, service not able to scale up. 
    We dont have data presistence, only cache data in local. Therefore if more than one service running, those services wont sharing data between each other

3, go code style. 
    I dont have much industrial experience in go, so code may not look very pretty to you.

4, we pass search in url parameter, however http url has a limitation of length(by differernt browser). so if you search parameter is too many, or email addresss is tooooo long, then request may not make it to server side.

# What we can do in future
1, use Elastic Search and Kibana in stead of local indexing. ES is prefect solution for this project.

2, If we dont presis any data, only do things in cache layer, redis is a good choice which can share data between machines.

3，We can configure the local service in k8s/vm/ec2 host

4, implement Create/Update/Read/Delete/Batch_Save APIs

5, change UUID generator if we have mutiple machines, or we have delete event API
