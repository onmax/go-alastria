# Credential Issuance - US-2.2.1

`Entity` issues a credential(s) to the specific subject that requested the action.

## Algorithm

### Goal

Save credentials into the subject's wallet.

### Assumptions

1. The `subject` has already a `keystore` and `did` created.
2. The `issuer` (aka `entity`) that will issue the credentials has already a DID and has the necessary permissions.
3. The `subject` was already registered in the `entity`'s CRM and has some data on it. *This means that user has to register using the centralized system of the entity and have some data in their systems*.
4. The `subject` has a legacy ID and the `subject` is logged in.
<!-- 5. The `entity` knows the `did` of the subject given his `legacyID`. -->


### Steps

1. The `subject` selects the credentials that want to store in his/her wallet. Then it makes the request to the `entity` to start the process with the selected credentials and his/her did.
2. The `entity`
	1. Checks that it has the credentials selected by the subject given the `DID`.
        2. Generates an `AT`. The `jti` should be mapped to the `DID` and store that relation temporally. The `AT` should not contain any critical data. The only important data is: `cbu`, `type` and `jti`.
        3. Send it to the wallet using any method: QR, Push notification, Deeplink, TinyURL...
3. The `subject` receives the `AT` in the wallet and generates an `AS` and send it to the `entity` using the field the `cbu` in `AT.payload.cbu`
4. The `entity` 
        1. Needs to verify the signature of the `AS` and therefore the public key is required
                1. Given the DID, the `PublicKeyRegistry` can return the current public key.
        2. Needs to verify that the subject is who he/she claims to be, we compare `AS.payload.iss` with the DID that is saved in the database
        3. Generates and signs credentials
        4. [OPTIONAL] Registers the credentials in the blockchain with a `PSMHash` and the `AddIssuerCredential`
        5. Return the credentials to the wallet as the response of the HTTP request     
5. The `subject`
        1. Receives the credentials and save them in his/her wallet.
        2. [OPTIONAL] Registers the credentials in the blockchain with a `PSMHash` and the `AddSubjectCredential`


> Info: The Alastria ID framework is a complex specification that uses many technologies/layers. At least, it has the following layers: blockchain, JWTs(artifacts) and `HTTP API`. The official docs do not emphasize in its differences and its requirements and there is not really much documentation except for the JWTs (See [other resources](#other-resources)). Therefore, in this examples, there are some parts of the algorithm that are highly opinionated and others developers could implement in other ways.

### Code

You can check two examples on how an issuer can issue a credential:
- [Full code](./full-example/main.go): Shows how to use `go-alastria` with the tokens(`AT`, `AS` and `Credential`) to send the information between entity and the subject through any mean of communication. This also, shows a way to authenticate the subject. And then, it shows how to generate PSMHashes and send transactions to the blockchain.
- [Only blockchain interaction](./only-blockchain/main.go): Shows only the code that is required to generate the `Credential`, generate `PSMHash` and send the transaction. This is a simplified version of `Full code`.


### Output example

<details>
 <summary>Full code</summary>
TODO
</details>

<details>
 <summary>Only Blockchain</summary>
TODO
</details>


## Other resources

- The code in this folder is heavily based on the [official Alastria examples/exampleCredentials](https://github.com/alastria/alastria-identity-example/tree/master/exampleCredentials)
- [Official documentation](https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#us-221-credential-issuance)
- [JWTs docs](https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions)