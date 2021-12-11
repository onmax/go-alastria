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

------ Entity issues credentials to subject ------ 

        Step 1: The subject selects credentials
In this case, the subject DID is did:ala:quor:redT:be186b9595504526858b099111a3072b48137b4b and has selected the credentials [passport email home_address]
The subject will send a request to the backend to start the process

        Step 2: The entity receives the request
The credentials that will be issued will be: [passport email]
The entity generates an AT: eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiIzNTZlM2ZjZTQzNWQ4NzI5MDYyZTUyZDI2M2MwYzcwNWIzYzVlMjAxYTlhOTYwOGNkYjA3MDc2NGU2YjhkZjMwYWU4NDIzYjQzOWE3YWYyYmNjMzUyOTc3ODM0MWFiMDZjMWU0NDQxMTM1MmYyMTdiNjhjZTQ0YTY3M2ExZGY2MyIsImtpZCI6ImRpZDphbGE6cXVvcjpyZWR0OmQyZjg2OGYwNTZlZjNhNDhiYmM4ZDQ0NmRmZWQ0MTFlOWJmOTNhYjAja2V5cy0xIn0.eyJqdGkiOiJ5b3VyLXVuaXF1ZS1hdC1pZC1vbmx5LW9uZS11c2UtYWxhc3RyaWEtdG9rZW4iLCJpYXQiOjE2MzY3OTg1NTIxNTksImV4cCI6MjAzNjc5ODU1MjE1OSwiaXNzIjoiZGlkOmFsYTpxdW9yOnJlZFQ6ZDJmODY4ZjA1NmVmM2E0OGJiYzhkNDQ2ZGZlZDQxMWU5YmY5M2FiMCIsInR5cGUiOlsiQWxhc3RyaWFUb2tlbiIsIlVTMjIxIl0sImFuaSI6InJlZFQiLCJjYnUiOiJodHRwczovL3lvdXItYmFja2VuZC11cmwuY29tL2FsYXN0cmlhL2NyZWRlbnRpYWwiLCJnd3UiOiJodHRwczovL3lvdXItcHJvdmlkZXItdXJsLmNvbSJ9.oXS5yGWNqD-VP03ki7xy9u9KtJ45h1TSVp9TLudk4QoMocj9wgbLSgUcJllNf3_p-gO-QmaW9RHKjthS7Zq6Ug. You can check its contents in https://jwt.io
The entity will send the AT to the subject's wallet using QR, Push notification, Deeplink, tinyURL...

        Step 3: The subject generates an AS
The subject generates an AS: eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiI1ZmQ5N2YzYjlmOWRjNzA2MDJmMjgwOGVmNGVhZTEyYTg3NjQ1ZmIyN2Y2YzE1MDY0NGNlOTljNmM1NGVlNDRlYWViNGZhMDM4NjFiYzZlYTc1ZjFkZTA0MDQxOGY4ZWZjOTIzNTNkYWRmODc4YmFjZmI1ZGZkZTc4ZTA0OWY2ZCJ9.eyJqdGkiOiJ5b3VyLXVuaXF1ZS1hdC1pZC1vbmx5LW9uZS11c2UtYWxhc3RyaWEtc2Vzc2lvbiIsImlhdCI6MTYzNjc5ODU1MjE1OSwiZXhwIjoyMDM2Nzk4NTUyMTU5LCJpc3MiOiJkaWQ6YWxhOnF1b3I6cmVkVDpiZTE4NmI5NTk1NTA0NTI2ODU4YjA5OTExMWEzMDcyYjQ4MTM3YjRiIiwiYWxhc3RyaWFUb2tlbiI6ImV5SmhiR2NpT2lKRlV6STFOa3NpTENKMGVYQWlPaUpLVjFRaUxDSnFkMnNpT2lJek5UWmxNMlpqWlRRek5XUTROekk1TURZeVpUVXlaREkyTTJNd1l6Y3dOV0l6WXpWbE1qQXhZVGxoT1RZd09HTmtZakEzTURjMk5HVTJZamhrWmpNd1lXVTROREl6WWpRek9XRTNZV1l5WW1Oak16VXlPVGMzT0RNME1XRmlNRFpqTVdVME5EUXhNVE0xTW1ZeU1UZGlOamhqWlRRMFlUWTNNMkV4WkdZMk15SXNJbXRwWkNJNkltUnBaRHBoYkdFNmNYVnZjanB5WldSME9tUXlaamcyT0dZd05UWmxaak5oTkRoaVltTTRaRFEwTm1SbVpXUTBNVEZsT1dKbU9UTmhZakFqYTJWNWN5MHhJbjAuZXlKcWRHa2lPaUo1YjNWeUxYVnVhWEYxWlMxaGRDMXBaQzF2Ym14NUxXOXVaUzExYzJVdFlXeGhjM1J5YVdFdGRHOXJaVzRpTENKcFlYUWlPakUyTXpZM09UZzFOVEl4TlRrc0ltVjRjQ0k2TWpBek5qYzVPRFUxTWpFMU9Td2lhWE56SWpvaVpHbGtPbUZzWVRweGRXOXlPbkpsWkZRNlpESm1PRFk0WmpBMU5tVm1NMkUwT0dKaVl6aGtORFEyWkdabFpEUXhNV1U1WW1ZNU0yRmlNQ0lzSW5SNWNHVWlPbHNpUVd4aGMzUnlhV0ZVYjJ0bGJpSXNJbFZUTWpJeElsMHNJbUZ1YVNJNkluSmxaRlFpTENKalluVWlPaUpvZEhSd2N6b3ZMM2x2ZFhJdFltRmphMlZ1WkMxMWNtd3VZMjl0TDJGc1lYTjBjbWxoTDJOeVpXUmxiblJwWVd3aUxDSm5kM1VpT2lKb2RIUndjem92TDNsdmRYSXRjSEp2ZG1sa1pYSXRkWEpzTG1OdmJTSjkub1hTNXlHV05xRC1WUDAza2k3eHk5dTlLdEo0NWgxVFNWcDlUTHVkazRRb01vY2o5d2diTFNnVWNKbGxOZjNfcC1nTy1RbWFXOVJIS2p0aFM3WnE2VWciLCJAY29udGV4dCI6WyJodHRwczovL2FsYXN0cmlhLmdpdGh1Yi5pby9pZGVudGl0eS9hcnRpZmFjdHMvdjEiXSwidHlwZSI6WyJBbGFzdHJpYVNlc3Npb24iLCJVUzIyMSJdfQ.0ydsuWPhD_rOHYQhPrnxEI26VFjqbUot0RMGpT1x7do_nZkiyIVsubaZvoMefnOEv1O0tWt4TmZO643WpSz5hw. You can check its contents in https://jwt.io
The subject will send the AS to the entity to the URL in the at.Payload.cbu. In this case: https://your-backend-url.com/alastria/credential

        Step 4: The entity generates and send signed credentials to subject
The entity needs to verify the AS. The entity has the DID of the subject (from previous steps): did:ala:quor:redT:be186b9595504526858b099111a3072b48137b4b. The entity takes the proxy, which is the address of the user and with that address, the entity calls GetCurrentPublicKey
IMPORTANT! The public key of the subject is 5fd97f3b9f9dc70602f2808ef4eae12a87645fb27f6c150644ce99c6c54ee44eaeb4fa03861bc6ea75f1de040418f8efc92353dadf878bacfb5dfde78e049f6d. The verification has been successful -> true
IMPORTANT! The issuer of the AS should be the subject's DID: did:ala:quor:redT:be186b9595504526858b099111a3072b48137b4b. AS.payload.iss is the same value that the entity has in its database? true
The jti of the AS should be the jti of the AT: your-unique-at-id-only-one-use-alastria-token. AS.payload.jti is the same value that the entity has in its temporal database? true. So this way we know which user sent the request.
The entity has generated 2 new credentials: [eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QifQ.eyJqdGkiOiJ5b3VyLXVuaXF1ZS1hdC1pZC1vbmx5LW9uZS11c2UtY3JlZGVudGlhbCIsImlhdCI6MTYzOTIzMjM0MjkxOSwiZXhwIjoyMDM2Nzk4NTUyMTU5LCJpc3MiOiJkaWQ6YWxhOnF1b3I6cmVkVDoyNjkzMjMzMzNjN2FiZjE5MmVmYmYxNzJkMjE3MTEwNmEyMTZhMjc1Iiwic3ViIjoiZGlkOmFsYTpxdW9yOnJlZFQ6YmUxODZiOTU5NTUwNDUyNjg1OGIwOTkxMTFhMzA3MmI0ODEzN2I0YiIsInZjIjp7IkBjb250ZXh0IjpbImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL3YxIiwiaHR0cHM6Ly9hbGFzdHJpYS5naXRodWIuaW8vaWRlbnRpdHkvY3JlZGVudGlhbHMvdjEiXSwidHlwZSI6WyJWZXJpZmlhYmxlQ3JlZGVudGlhbCIsIkFsYXN0cmlhVmVyaWZpYWJsZUNyZWRlbnRpYWwiXSwiY3JlZGVudGlhbFN1YmplY3QiOnsiZGF0YSI6eyJwYXNzcG9ydCI6IjEyMzQ1Njc4QSJ9LCJsZXZlbE9mQXNzdXJhbmNlIjoxfX19.WPGhDpzWyyADlxa6hs2dZygg-K2X2BSnxhqmqyaiVX5sv0pkv79NX1SZZLMutid9PSp2yToc0bdk1tqkyPl-BA eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QifQ.eyJqdGkiOiJ5b3VyLXVuaXF1ZS1hdC1pZC1vbmx5LW9uZS11c2UtY3JlZGVudGlhbCIsImlhdCI6MTYzOTIzMjM0MjkyMCwiZXhwIjoyMDM2Nzk4NTUyMTU5LCJpc3MiOiJkaWQ6YWxhOnF1b3I6cmVkVDoyNjkzMjMzMzNjN2FiZjE5MmVmYmYxNzJkMjE3MTEwNmEyMTZhMjc1Iiwic3ViIjoiZGlkOmFsYTpxdW9yOnJlZFQ6YmUxODZiOTU5NTUwNDUyNjg1OGIwOTkxMTFhMzA3MmI0ODEzN2I0YiIsInZjIjp7IkBjb250ZXh0IjpbImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL3YxIiwiaHR0cHM6Ly9hbGFzdHJpYS5naXRodWIuaW8vaWRlbnRpdHkvY3JlZGVudGlhbHMvdjEiXSwidHlwZSI6WyJWZXJpZmlhYmxlQ3JlZGVudGlhbCIsIkFsYXN0cmlhVmVyaWZpYWJsZUNyZWRlbnRpYWwiXSwiY3JlZGVudGlhbFN1YmplY3QiOnsiZGF0YSI6eyJlbWFpbCI6ImV4YW1wbGVAbWFpbC5vcmcifSwibGV2ZWxPZkFzc3VyYW5jZSI6MX19fQ.3ZHGHFguM-0DL8QwklhqicDwailPGq2-AI7r5JT5G5Bg3I1xObQL7-06hI9DfWP2cL5TZiMrKe_YzKcorPZ4tQ]
The entity has added 2 credentials to the blockchain
The transaction hashes are: [  0x23a9a8d36583c274fd2eef1895d8ef99936893d97dc6f6c663e5d5874353af5c 0xc1404ffd2b82856161d9e67d224718d6b594b202494e96bedfa1201e1be656ec]
The PSMHashes are: [  0x0f7ce997dd07652063d0b80c44a6bb732fd4150a89aa0d2502018772d20b1b2c 0xc74bb9a19f1e79f1dccbd92bfe82957d77312b784a71683038259e049f11a078]
As a response of the HTTP request made by the subject, the credential will send the signed credentials

        Step 5: The subject receives the signed credentials and store them somewhere secure. Optionally, the user could write in blockchain that he has received the credentials

The entity has added 2 credentials to the blockchain
The transaction hashes are: [  0xd82c8d029e77f690ef6d0a94cf2b5f0c46b99939b6d7650ae7de78c678f32437 0xc2032e8c0e7f215e4e65ba23c4952d716854d7c7cb741831ef25a56e2a4ee26a]
The PSMHashes are: [  0x24b3a1d2b65db3606f04bba85b557e1bd28c25e05b9ddd04d02c8fbf626c41c5 0x09676da221c04dc56a221e159341097dc79cb69668c1a4e77b76cad37f13cbf4]
</details>

<details>
 <summary>Only Blockchain</summary>

------ Entity issues credentials to subject ------ 

        Omiting step 1-3 as they are only for establishing a secure handshake between subject and issuer
        Step 4: The entity generates and send signed credentials to subject
The entity has generated 2 new credentials: [eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QifQ.eyJqdGkiOiJ5b3VyLXVuaXF1ZS1hdC1pZC1vbmx5LW9uZS11c2UtY3JlZGVudGlhbCIsImlhdCI6MTYzOTIzMjM3MDUyOSwiZXhwIjoyMDM2Nzk4NTUyMTU5LCJpc3MiOiJkaWQ6YWxhOnF1b3I6cmVkVDoyNjkzMjMzMzNjN2FiZjE5MmVmYmYxNzJkMjE3MTEwNmEyMTZhMjc1Iiwic3ViIjoiZGlkOmFsYTpxdW9yOnJlZFQ6YmUxODZiOTU5NTUwNDUyNjg1OGIwOTkxMTFhMzA3MmI0ODEzN2I0YiIsInZjIjp7IkBjb250ZXh0IjpbImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL3YxIiwiaHR0cHM6Ly9hbGFzdHJpYS5naXRodWIuaW8vaWRlbnRpdHkvY3JlZGVudGlhbHMvdjEiXSwidHlwZSI6WyJWZXJpZmlhYmxlQ3JlZGVudGlhbCIsIkFsYXN0cmlhVmVyaWZpYWJsZUNyZWRlbnRpYWwiXSwiY3JlZGVudGlhbFN1YmplY3QiOnsiZGF0YSI6eyJwYXNzcG9ydCI6IjEyMzQ1Njc4QSJ9LCJsZXZlbE9mQXNzdXJhbmNlIjoxfX19.ImWwyXGB5N7ATzMOOh4pRhlJL-xmqr4iEHDF-DhKlvU7eU1HRgYkhyKKm5swsIeCVGw26VBj5lNHkYfIEpybKA eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QifQ.eyJqdGkiOiJ5b3VyLXVuaXF1ZS1hdC1pZC1vbmx5LW9uZS11c2UtY3JlZGVudGlhbCIsImlhdCI6MTYzOTIzMjM3MDUyOSwiZXhwIjoyMDM2Nzk4NTUyMTU5LCJpc3MiOiJkaWQ6YWxhOnF1b3I6cmVkVDoyNjkzMjMzMzNjN2FiZjE5MmVmYmYxNzJkMjE3MTEwNmEyMTZhMjc1Iiwic3ViIjoiZGlkOmFsYTpxdW9yOnJlZFQ6YmUxODZiOTU5NTUwNDUyNjg1OGIwOTkxMTFhMzA3MmI0ODEzN2I0YiIsInZjIjp7IkBjb250ZXh0IjpbImh0dHBzOi8vd3d3LnczLm9yZy8yMDE4L2NyZWRlbnRpYWxzL3YxIiwiaHR0cHM6Ly9hbGFzdHJpYS5naXRodWIuaW8vaWRlbnRpdHkvY3JlZGVudGlhbHMvdjEiXSwidHlwZSI6WyJWZXJpZmlhYmxlQ3JlZGVudGlhbCIsIkFsYXN0cmlhVmVyaWZpYWJsZUNyZWRlbnRpYWwiXSwiY3JlZGVudGlhbFN1YmplY3QiOnsiZGF0YSI6eyJlbWFpbCI6ImV4YW1wbGVAbWFpbC5vcmcifSwibGV2ZWxPZkFzc3VyYW5jZSI6MX19fQ.18Gi8BPuMDp54Mz1xk4sI8wtYZXy_u3mK_vKOGU_hjsUd2ckngcruFyeefTys1FgwllBxwv2uY8GCkBXfJxqMQ]
The entity has added 2 credentials to the blockchain
The transaction hashes are: [  0x89f8089334fb83a9c9fc6352e8c7d29b2bc6369635a6aa49c6144e7a8d8a764a 0xd00408a8bfe17a93130a58cd35622f99de6675bffbc6912bd0086e75ba30bce8]
The PSMHashes are: [  0x0791bbf5d023276bbb0ba694e0fd980d421da248f490576ba0a99fdad47c496d 0xf873d3fca4ede8acdd893144a2e7b835e72fb8c71333a0f1e521e4aee2040b89]
As a response of the HTTP request made by the subject, the credential will send the signed credentials

        Step 5: The subject receives the signed credentials and store them somewhere secure. Optionally, the user could write in blockchain that he has received the credentials

The entity has added 2 credentials to the blockchain
The transaction hashes are: [  0x5764c06147405e0e46d979c7743c2085f4f3a21d9cf2313396a21cb8239f6810 0x5217d24b12f2be6da853693ece051d7b769e07d1c6b5aaac59cfc8fd5a8dc7b3]
The PSMHashes are: [  0xaa2293622015df742b964f1f064230e9c6fa6347cebfd5601cb669e9ecf6e53f 0x7a176fb35cb55821d4081ff459b5fea8f5786c3af2633f87c069757be8bef4e4]
</details>


## Other resources

- The code in this folder is heavily based on the [official Alastria examples/exampleCredentials](https://github.com/alastria/alastria-identity-example/tree/master/exampleCredentials)
- [Official documentation](https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#us-221-credential-issuance)
- [JWTs docs](https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions)