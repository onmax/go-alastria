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

4. At this point, `Create Alastria ID - US1.2` is complete. But, you need to call `IdentityKeys` to get the `did` given the public key of the `newActor`

> Info: The Alastria ID framework is a complex specification that uses many technologies/layers. At least, it has the following layers: blockchain, JWTs(artifacts) and `HTTP API`. The official docs do not emphasize in its differences and its requirements and there is not really much documentation except for the JWTs (See [other resources](#other-resources)). Therefore, in this examples, there are some parts of the algorithm that are highly opinionated and others developers could implement in other ways.

### Code

You can check two examples on how to create new DID on Alastria:
- [Full code](./full-example/main.go): Shows how to use `go-alastria` with tokens(`AT` and `AIC`) to send the information between entity and the new actor through any mean of communication. And then, it shows how to generate and send transactions to the blockchain.
- [Only blockchain interaction](./only-blockchain/main.go): Shows only the code that is required to generates and send the transactions. This is a simplified version of `Full code`

### Output example

<details>
 <summary>Full code</summary>
 ```bash
------ Creating new actor ------ 

To simplify the code, errors are not being checked

        Step 1: Entity generates an AT
AT: eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiIzNTZlM2ZjZTQzNWQ4NzI5MDYyZTUyZDI2M2MwYzcwNWIzYzVlMjAxYTlhOTYwOGNkYjA3MDc2NGU2YjhkZjMwYWU4NDIzYjQzOWE3YWYyYmNjMzUyOTc3ODM0MWFiMDZjMWU0NDQxMTM1MmYyMTdiNjhjZTQ0YTY3M2ExZGY2MyIsImtpZCI6ImRpZDphbGE6cXVvcjpyZWR0OmQyZjg2OGYwNTZlZjNhNDhiYmM4ZDQ0NmRmZWQ0MTFlOWJmOTNhYjAja2V5cy0xIn0.eyJqdGkiOiJ5b3VyLXVuaXF1ZS1hdC1pZCIsImlhdCI6MTYzNjc5ODU1MjE1OSwiZXhwIjoyMDM2Nzk4NTUyMTU5LCJpc3MiOiJkaWQ6YWxhOnF1b3I6cmVkVDpkMmY4NjhmMDU2ZWYzYTQ4YmJjOGQ0NDZkZmVkNDExZTliZjkzYWIwIiwidHlwZSI6WyJBbGFzdHJpYVRva2VuIiwiVVMxMiJdLCJhbmkiOiJyZWRUIiwiY2J1IjoiaHR0cHM6Ly95b3VyLWJhY2tlbmQtdXJsLmNvbSIsImd3dSI6Imh0dHBzOi8veW91ci1wcm92aWRlci11cmwuY29tIn0.fgKWMlm3CAOQGDXT4kn0JLof9rTd-Q4CwqlRFeQGq9VjYanmLZL3zoeqEkgCktqSQmOF3HdjvLfD_MD0TgVtFA. You can check its contents in https://jwt.io
And somehow sends it to the new agent like a QR code for example.

        Step 2: New actor generates AIC
The new actor once receives the AT, he generates an AIC: eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiI1ZmQ5N2YzYjlmOWRjNzA2MDJmMjgwOGVmNGVhZTEyYTg3NjQ1ZmIyN2Y2YzE1MDY0NGNlOTljNmM1NGVlNDRlYWViNGZhMDM4NjFiYzZlYTc1ZjFkZTA0MDQxOGY4ZWZjOTIzNTNkYWRmODc4YmFjZmI1ZGZkZTc4ZTA0OWY2ZCJ9.eyJpYXQiOjE2MzY3OTg1NTIxNTksInB1YmxpY0tleSI6IjB4NWZkOTdmM2I5ZjlkYzcwNjAyZjI4MDhlZjRlYWUxMmE4NzY0NWZiMjdmNmMxNTA2NDRjZTk5YzZjNTRlZTQ0ZWFlYjRmYTAzODYxYmM2ZWE3NWYxZGUwNDA0MThmOGVmYzkyMzUzZGFkZjg3OGJhY2ZiNWRmZGU3OGUwNDlmNmQiLCJqdGkiOiJ5b3VyLXVuaXF1ZS1hdC1pZCIsImNyZWF0ZUFsYXN0cmlhVFgiOiIweDdiMjI3NDc5NzA2NTIyM2EyMjMwNzgzMDIyMmMyMjZlNmY2ZTYzNjUyMjNhMjIzMDc4NjEzNTIyMmMyMjY3NjE3MzUwNzI2OTYzNjUyMjNhMjIzMDc4MzAyMjJjMjI2ZDYxNzg1MDcyNjk2ZjcyNjk3NDc5NDY2NTY1NTA2NTcyNDc2MTczMjIzYTZlNzU2YzZjMmMyMjZkNjE3ODQ2NjU2NTUwNjU3MjQ3NjE3MzIyM2E2ZTc1NmM2YzJjMjI2NzYxNzMyMjNhMjIzMDc4MzkzMjM3NjMzMDIyMmMyMjc2NjE2Yzc1NjUyMjNhMjIzMDc4MzAyMjJjMjI2OTZlNzA3NTc0MjIzYTIyMzA3ODM2NjQzNjM5NjQzOTM5NjEzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzIzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzA2MzM0MzUzMDMzMzgzMjYzMzE2MTMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMjMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDM4MzAzMzM1MzYzNjM2MzQzMzM5MzMzNzM2MzYzMzMzMzYzMjMzMzkzNjM2MzMzOTM2MzQzNjMzMzMzNzMzMzAzMzM2MzMzMDMzMzIzNjM2MzMzMjMzMzgzMzMwMzMzODM2MzUzNjM2MzMzNDM2MzUzNjMxMzYzNTMzMzEzMzMyMzYzMTMzMzgzMzM3MzMzNjMzMzQzMzM1MzYzNjM2MzIzMzMyMzMzNzM2MzYzMzM2MzYzMzMzMzEzMzM1MzMzMDMzMzYzMzM0MzMzNDM2MzMzNjM1MzMzOTMzMzkzNjMzMzMzNjM2MzMzMzM1MzMzNDM2MzUzNjM1MzMzNDMzMzQzNjM1MzYzMTM2MzUzNjMyMzMzNDM2MzYzNjMxMzMzMDMzMzMzMzM4MzMzNjMzMzEzNjMyMzYzMzMzMzYzNjM1MzYzMTMzMzczMzM1MzYzNjMzMzEzNjM0MzYzNTMzMzAzMzM0MzMzMDMzMzQzMzMxMzMzODM2MzYzMzM4MzYzNTM2MzYzNjMzMzMzOTMzMzIzMzMzMzMzNTMzMzMzNjM0MzYzMTM2MzQzNjM2MzMzODMzMzczMzM4MzYzMjM2MzEzNjMzMzYzNjM2MzIzMzM1MzYzNDM2MzYzNjM0MzYzNTMzMzczMzM4MzYzNTMzMzAzMzM0MzMzOTM2MzYzMzM2MzYzNDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAzMDMwMzAyMjJjMjI3NjIyM2EyMjMwNzgzMjM2NjU2MzMxMzEzMDYxNjQzODIyMmMyMjcyMjIzYTIyMzA3ODY2NjU2NjM1MzAzNDY2MzY2MjY2MzUzNTYyNjEzMzMxNjM2NDY2NjUzOTY1MzAzNjM3MzUzNTMyMzQzMjYzNjI2MzM1MzAzMzYxMzU2NDY2MzgzNjM1MzMzMDY1NjY2MzY2NjU2NTM4MzE2NTY2NjU2NjM3NjMzMTY1NjEzMjY1MjIyYzIyNzMyMjNhMjIzMDc4NjQzODMzNjQzOTYxNjYzMTMyMzAzNjM4NjIzODY2MzA2MzM2NjMzMjY2MzE2MjMyMzk2NDY1MzEzNTM0NjM2MjMyMzk2MTYxMzczOTYxNjQzODMyNjY2MzM4MzAzMDY2MzM2NTYzMzk2NjMwNjYzMjM2MzQzNDMxMzczODY2MjIyYzIyNzQ2ZjIyM2EyMjMwNzg2NTYxNjY2MzYxNjYzODYzMzM2NjM5MzAzMTM2NjIzMTM0NjU2MzM2MzUzMTMwMzUzNjM3MzczNjM1Mzg2NjMzNjQzNjY1NjIzOTMwMzczOTIyMmMyMjY4NjE3MzY4MjIzYTIyMzA3ODMzMzMzMjMwMzAzMjM2MzEzMzY0MzczMzMwNjM2MTM4MzczMDM2MzIzMDYyNjQzNzM0MzAzMzY1NjUzODMxNjYzMjMxMzQzNTY1NjIzNDMwNjE2MTMwMzkzOTYxNjEzNjM0MzQzMTM4MzYzNTMwMzQ2MzM4MzczNjYyMzM2MzY2MjI3ZCIsImFsYXN0cmlhVG9rZW4iOiJleUpoYkdjaU9pSkZVekkxTmtzaUxDSjBlWEFpT2lKS1YxUWlMQ0pxZDJzaU9pSXpOVFpsTTJaalpUUXpOV1E0TnpJNU1EWXlaVFV5WkRJMk0yTXdZemN3TldJell6VmxNakF4WVRsaE9UWXdPR05rWWpBM01EYzJOR1UyWWpoa1pqTXdZV1U0TkRJellqUXpPV0UzWVdZeVltTmpNelV5T1RjM09ETTBNV0ZpTURaak1XVTBORFF4TVRNMU1tWXlNVGRpTmpoalpUUTBZVFkzTTJFeFpHWTJNeUlzSW10cFpDSTZJbVJwWkRwaGJHRTZjWFZ2Y2pweVpXUjBPbVF5WmpnMk9HWXdOVFpsWmpOaE5EaGlZbU00WkRRME5tUm1aV1EwTVRGbE9XSm1PVE5oWWpBamEyVjVjeTB4SW4wLmV5SnFkR2tpT2lKNWIzVnlMWFZ1YVhGMVpTMWhkQzFwWkNJc0ltbGhkQ0k2TVRZek5qYzVPRFUxTWpFMU9Td2laWGh3SWpveU1ETTJOems0TlRVeU1UVTVMQ0pwYzNNaU9pSmthV1E2WVd4aE9uRjFiM0k2Y21Wa1ZEcGtNbVk0TmpobU1EVTJaV1l6WVRRNFltSmpPR1EwTkRaa1ptVmtOREV4WlRsaVpqa3pZV0l3SWl3aWRIbHdaU0k2V3lKQmJHRnpkSEpwWVZSdmEyVnVJaXdpVlZNeE1pSmRMQ0poYm1raU9pSnlaV1JVSWl3aVkySjFJam9pYUhSMGNITTZMeTk1YjNWeUxXSmhZMnRsYm1RdGRYSnNMbU52YlNJc0ltZDNkU0k2SW1oMGRIQnpPaTh2ZVc5MWNpMXdjbTkyYVdSbGNpMTFjbXd1WTI5dEluMC5mZ0tXTWxtM0NBT1FHRFhUNGtuMEpMb2Y5clRkLVE0Q3dxbFJGZVFHcTlWallhbm1MWkwzem9lcUVrZ0NrdHFTUW1PRjNIZGp2TGZEX01EMFRnVnRGQSIsIkBjb250ZXh0IjpbImh0dHBzOi8vYWxhc3RyaWEuZ2l0aHViLmlvL2lkZW50aXR5L2FydGlmYWN0cy92MSJdLCJ0eXBlIjpbIkFsYXN0cmlhSWRlbnRpdHlDcmVhdGlvbiIsIlVTMTIiXX0.yS3U-lWu2gq8W3nNJ3sJmRkM6w5YMBuB615kaz4IpRZQ1GJSvNed2HNWstJWhi102vi29BPZ1d-OqSBHCgt27w. You can check its contents in https://jwt.io
The new actor needs to send the AIC to the entity using the cbu field from at.payload.cbu. In this case, it will be: https://your-backend-url.com

        Step 3: Entity sends the transactions
The entity receives the AIC and destructure the 3 important params mentioned before:
        - PublicKey -> To verify the AIC and to generate the address of the new actor
        - Tx of CreateAlastriaID -> The transaction data that the entity needs to run in the SC
        - The original AT -> In case, the entity needs to keep track of the user from the beginning for his own US

        Step 4: Get the full DID
New actor's Alastria DID: did:ala:quor:redT:af6a7777f5a98c2b47a0bc3bf965b4a4014d302c
```
</details>

<details>
 <summary>Only Blockchain</summary>
 ```bash
------ Creating new actor | Only blockchain part ------ 

To simplify the code, errors are not being checked

        Omitting step 1 as it has nothing to do with blockchain
        Step 2: New agent generates CreateAlastriaID transaction
The signed tx CreateAID generated by new agent: &{0xc01023c240 {13863323974600768752 1097841582 0xb22d60} {<nil>} {<nil>} {<nil>}}
The public key of the new agent: 5fd97f3b9f9dc70602f2808ef4eae12a87645fb27f6c150644ce99c6c54ee44eaeb4fa03861bc6ea75f1de040418f8efc92353dadf878bacfb5dfde78e049f6d
Both of them should go in the AIC
        Step 3: Send both transaction
The entity needs first to parse the signedTxCreateAID transaction, and then, the entity send both transactions in a specific order
        Step 4: Get the full DID
DID of the new actor: did:ala:quor:redT:bd741459000200190ca63338358f1de6c9b37a5d
```
</details>


## Other resources

- The code in this folder is heavily based on the [official Alastria examples/CreateAlastriaID](https://github.com/alastria/alastria-identity-example/tree/master/exampleCreateAlastriaID)
- [Official documentation](https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#us-12-alastria-id-creation)
- [JWTs docs](https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions)