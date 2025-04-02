CONTRACTS_V1_0=FiredrillToken FiredrillRevertMessageReceiver
CONTRACTS_ALL=FiredrillEntrypoint FiredrillCompound FiredrillOnRamp FiredrillOffRamp

.PHONY: abigen
abigen: ## Install abigen.
	go install github.com/ethereum/go-ethereum/cmd/abigen@latest

.PHONY: foundry
foundry: ## Install foundry.
	foundryup --version v0.3.0

.PHONY: foundry-refresh
foundry-refresh: foundry
	git submodule deinit -f .
	git submodule update --init --recursive

.PHONY: build
build: ## Build the contracts and generate abis and bins
	forge build
	mkdir -p generated/$(VERSION)/abi
	mkdir -p generated/$(VERSION)/bin
	mkdir -p generated/$(VERSION)/gethwrappers
	for contract in ${CONTRACTS}; do \
	  forge inspect contracts/$(VERSION)/$$contract.sol:$$contract abi > generated/$(VERSION)/abi/$$contract.abi; \
	  forge inspect contracts/$(VERSION)/$$contract.sol:$$contract bytecode > generated/$(VERSION)/bin/$$contract.bin; \
	  pkgName=$$(echo $$contract | sed -E 's/([A-Z]+)([A-Z][a-z])/\1_\2/g; s/([a-z0-9])([A-Z])/\1_\2/g' | tr '[:upper:]' '[:lower:]'); \
	  mkdir -p generated/$(VERSION)/gethwrappers/$$pkgName; \
	  abigen --abi=generated/$(VERSION)/abi/$$contract.abi --bin=generated/$(VERSION)/bin/$$contract.bin --pkg=$$pkgName --out=generated/$(VERSION)/gethwrappers/$$pkgName/$$contract.go; \
	done

.PHONY: build-all
build-all:
	VERSION=v1_0 CONTRACTS="${CONTRACTS_V1_0}" make build
	VERSION=v1_5 CONTRACTS="${CONTRACTS_ALL}" make build
	VERSION=v1_6 CONTRACTS="${CONTRACTS_ALL}" make build

help:
	@echo ""
	@echo "         .__           .__       .__  .__        __"
	@echo "    ____ |  |__ _____  |__| ____ |  | |__| ____ |  | __"
	@echo "  _/ ___\|  |  \\\\\\__  \ |  |/    \|  | |  |/    \|  |/ /"
	@echo "  \  \___|   Y  \/ __ \|  |   |  \  |_|  |   |  \    <"
	@echo "   \___  >___|  (____  /__|___|  /____/__|___|  /__|_ \\"
	@echo "       \/     \/     \/        \/             \/     \/"
	@echo ""
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
	awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
