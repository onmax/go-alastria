generate_contracts_go:
	# git clone https://github.com/alastria/alastriaID-truffle-contracts
	# npm i @chainsafe/truffle-plugin-abigen # Install this plugin and add it to truffle.config
	# ./node-modules/.bin/truffle compile
	# npm run migrateLocal # make sure you have ganache, geth... or similar up and running
	# npm run initLocal
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AdminUpgradeabilityProxy.abi --pkg=alastria --out=./contracts/AdminUpgradeabilityProxy.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaCredentialRegistry.abi --pkg=alastria --out=./contracts/AlastriaCredentialRegistry.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaIdentityIssuer.abi --pkg=alastria --out=./contracts/AlastriaIdentityIssuer.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaIdentityManager.abi --pkg=alastria --out=./contracts/AlastriaIdentityManager.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaIdentityServiceProvider.abi --pkg=alastria --out=./contracts/AlastriaIdentityServiceProvider.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaNameService.abi --pkg=alastria --out=./contracts/AlastriaNameService.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaPresentationRegistry.abi --pkg=alastria --out=./contracts/AlastriaPresentationRegistry.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaProxy.abi --pkg=alastria --out=./contracts/AlastriaProxy.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/AlastriaPublicKeyRegistry.abi --pkg=alastria --out=./contracts/AlastriaPublicKeyRegistry.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/BaseAdminUpgradeabilityProxy.abi --pkg=alastria --out=./contracts/BaseAdminUpgradeabilityProxy.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/BaseUpgradeabilityProxy.abi --pkg=alastria --out=./contracts/BaseUpgradeabilityProxy.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Context.abi --pkg=alastria --out=./contracts/Context.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Eidas.abi --pkg=alastria --out=./contracts/Eidas.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Migrations.abi --pkg=alastria --out=./contracts/Migrations.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Ownable.abi --pkg=alastria --out=./contracts/Ownable.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Owned.abi --pkg=alastria --out=./contracts/Owned.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/Proxy.abi --pkg=alastria --out=./contracts/Proxy.go
	abigen --abi=./alastriaID-truffle-contracts/abigenBindings/abi/UpgradeabilityProxy.abi --pkg=alastria --out=./contracts/UpgradeabilityProxy.go