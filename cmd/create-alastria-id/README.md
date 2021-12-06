# Create Alastria ID - US1.2

Allows to an entity to create a new Alastria DID. The new DID can be a new subject or a new identity.

## Algorithm

### Goal

An `entity` wants to create a new DID for a new actor(`newActor`). This new actor can be: a new subject (which is the most common case) or a another new entity.

### Assumptions

1. We assume that the `newActor` has already a new `keystore` created.
2. We assume that the `entity` that wants to create the `newActor` has already a DID and has the necessary permissions.

### Steps

1. The `entity` generates an `AlastriaToken`(`AT`), signs it and sends it to the `newActor`
2. The `newActor`:
    1. Receives the `AT` and verifies the signature
    2. Connects to the blockchain network
    3. Generates a `CreateAlastriaIdentity` transaction and signs it. But it does not send it!
        1. Internally, the library, it will first call `AddKey` to add the keys in the `PublicKeyRegistry`
        2. Then, also internally, it will call `CreateAlastriaIdentity` in `IdentityManager` contract
    4. Generates an AIC and signs it. The AT contains among other data: AT, `newActor`'s public key and the signed `CreateAlastriaIdentity` transactions
    5. Sends a POST request to the endpoint specified in the AT, specifically: `at.payload.cbu` and in the body should have the AIC
3. The `entity`:
    1. Receives the AIC and verifies the signature
    2. Connects to the blockchain network
    3. Generates a `PrepareAlastriaId` transaction and signs it. But it does not send it!
        3.3.1 Internally, the library, it will first call `PrepareAlastriaId` in `IdentityManager` contract
        3.3.2 Then, also internally, the library it will call `DelegateCall`  in `IdentityManager` contract to create the proxy
    4. Entity will send then the `PrepareAlastriaId` transaction and `CreateAlastriaIdentity` transaction, in that order.

--- 

4. At this point, `Create Alastria ID - US1.2` is complete it. But, you need to call `IdentityKeys` to get the `did` given a public key of the `newActor`

> Info: The Alastria ID framework is a complex specification that uses many technologies/layer. At least, it has the following layers: blockchain, JWTs(artifacts) and `HTTP API`. The official docs do not emphasize in its differences and its requirements and there is not really much documentation except for the JWTs (See [other resources](#other-resources)). Therefore, in this examples, there are some parts of the algorithm that are highly opinionated and others developers could implement in other ways.

## Other resources

- The code in this folder is heavily based on the [official Alastria examples/CreateAlastriaID](https://github.com/alastria/alastria-identity-example/tree/master/exampleCreateAlastriaID)
- [Official documentation](https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#us-12-alastria-id-creation)
- [JWTs docs](https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions)