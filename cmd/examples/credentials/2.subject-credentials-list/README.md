# Get Credential List of subject

###### It does not have a US-X identifier like other User stories

Any `actor` of the network retrieves the list of PSMHashes that a `subject` has written in the blockchain

## Algorithm

### Goal

Get list of PSMHashes of subject

### Assumptions

There are not assumptions, even if the subject does not exist, the return value will be an empty array

### Steps

1. The `actor` (anyone in the network) initializes the client and executes `GetSubjectCredentialList` given an address of the subject.


> Info: The Alastria ID framework is a complex specification that uses many technologies/layers. At least, it has the following layers: blockchain, JWTs(artifacts) and `HTTP API`. The official docs do not emphasize in its differences and its requirements and there is not really much documentation except for the JWTs (See [other resources](#other-resources)). Therefore, in this examples, there are some parts of the algorithm that are highly opinionated and others developers could implement in other ways.

### Code

You can check the code in [main](./main.go).


### Output example

<details>
 <summary>Full code</summary>
------ Any actor get the psm hashes of credentials of a subject ------ 

The subject 0xD0a0D5A1310A715157c3f81B789d6d9Dc447AEF5 has 3 credentials registered in the blockchain:
The PSMHashes are [0x6335306438643930303833366261336235323962 0x3633363463333365613838613663643361316565 0x3531303131303563663264323930326331346532]
</details>


## Other resources

- The code in this folder is heavily based on the [official Alastria examples/exampleCredentials](https://github.com/alastria/alastria-identity-example/tree/master/exampleCredentials)