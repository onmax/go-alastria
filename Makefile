generate_contracts_go:
	# git clone https://github.com/alastria/alastriaID-truffle-contracts
	# npm i @chainsafe/truffle-plugin-abigen # Install this plugin and add it to truffle.config
	# ./node-modules/.bin/truffle compile
	# npm run migrateLocal # make sure you have ganache, geth... or similar up and running
	# npm run initLocal
	mkdir ./contracts/admin-upgradeability-proxy
	mkdir ./contracts/alastria-credential-registry
	mkdir ./contracts/alastria-identity-issuer
	mkdir ./contracts/alastria-identity-manager
	mkdir ./contracts/alastria-identity-service-provider
	mkdir ./contracts/alastria-name-service
	mkdir ./contracts/alastria-presentation-registry
	mkdir ./contracts/base-admin-upgradeability-proxy
	mkdir ./contracts/base-upgradeability-proxy
	mkdir ./contracts/context
	mkdir ./contracts/eidas
	mkdir ./contracts/migrations
	mkdir ./contracts/ownable
	mkdir ./contracts/owned
	mkdir ./contracts/proxy
	mkdir ./contracts/upgradeability-proxy
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AdminUpgradeabilityProxy.abi --pkg=alastria --out=./contracts/admin-upgradeability-proxy/AdminUpgradeabilityProxy.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaCredentialRegistry.abi --pkg=alastria --out=./contracts/alastria-credential-registry/AlastriaCredentialRegistry.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaIdentityIssuer.abi --pkg=alastria --out=./contracts/alastria-identity-issuer/AlastriaIdentityIssuer.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaIdentityManager.abi --pkg=alastria --out=./contracts/alastria-identity-manager/AlastriaIdentityManager.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaIdentityServiceProvider.abi --pkg=alastria --out=./contracts/alastria-identity-service-provider/AlastriaIdentityServiceProvider.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaNameService.abi --pkg=alastria --out=./contracts/alastria-name-service/AlastriaNameService.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaPresentationRegistry.abi --pkg=alastria --out=./contracts/alastria-presentation-registry/AlastriaPresentationRegistry.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaProxy.abi --pkg=alastria --out=./contracts/alastria-proxy/AlastriaProxy.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaPublicKeyRegistry.abi --pkg=alastria --out=./contracts/alastria-public-key-registry/AlastriaPublicKeyRegistry.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/BaseAdminUpgradeabilityProxy.abi --pkg=alastria --out=./contracts/base-admin-upgradeability-proxy/BaseAdminUpgradeabilityProxy.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/BaseUpgradeabilityProxy.abi --pkg=alastria --out=./contracts/base-upgradeability-proxy/BaseUpgradeabilityProxy.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Context.abi --pkg=alastria --out=./contracts/context/Context.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Eidas.abi --pkg=alastria --out=./contracts/eidas/Eidas.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Migrations.abi --pkg=alastria --out=./contracts/migrations/Migrations.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Ownable.abi --pkg=alastria --out=./contracts/ownable/Ownable.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Owned.abi --pkg=alastria --out=./contracts/owned/Owned.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Proxy.abi --pkg=alastria --out=./contracts/proxy/Proxy.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/UpgradeabilityProxy.abi --pkg=alastria --out=./contracts/upgradeability-proxy/UpgradeabilityProxy.go