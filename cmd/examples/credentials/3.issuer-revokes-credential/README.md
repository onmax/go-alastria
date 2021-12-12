# Get Credential List of subject

###### It does not have a US-X identifier like other User Stories

Any `actor` of the network retrieves the list of PSMHashes that a `subject` has written in the blockchain

## Algorithm

### Goal

Get list of PSMHashes of `subject`

### Assumptions

There are not assumptions, even if the `subject` does not exist, the return value will be an empty array

### Steps

1. The `actor` (anyone in the network) initializes the client and executes `GetSubjectCredentialsList` given an address of the `subject` which will return the PSMHashes of the credentials.
2. Then, the `actor` will use `GetSubjectCredentialStatus` to retrieve the status of all the credentials.


> Info: The Alastria ID framework is a complex specification that uses many technologies/layers. At least, it has the following layers: blockchain, JWTs(artifacts) and `HTTP API`. The official docs do not emphasize in its differences and its requirements and there is not really much documentation except for the JWTs (See [other resources](#other-resources)). Therefore, in this examples, there are some parts of the algorithm that are highly opinionated and others developers could implement in other ways.

### Code

You can check the code [here](./subject/main.go)


### Output example

<details>
 <summary>Full code for the subject</summary>
------ Any actor gets the PSMhashes of credentials of a subject ------ 

The subject 0xD0a0D5A1310A715157c3f81B789d6d9Dc447AEF5 (with did: did:ala:quor:redT:be186b9595504526858b099111a3072b48137b4b) has 4 credentials registered in the blockchain:
        Credential 0x3634333933303330333833333336363236313333 || Status 0
        Credential 0x3333333336353631333833383631333636333634 || Status 0
        Credential 0x3635333136343339363133393632333933343633 || Status 0
        Credential 0x3331333936313632333936323634333933373633 || Status 0
</details>


## Other resources

- The code in this folder is heavily based on the [official Alastria examples/exampleCredentials](https://github.com/alastria/alastria-identity-example/tree/master/exampleCredentials)