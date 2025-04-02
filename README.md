# CCIP Firedrill Deployment

This repository contains smart contracts source code and deployment scripts for CCIP Firedrill process
CCIP Firedrill documentation and process description: https://smartcontract-it.atlassian.net/wiki/spaces/CCIP/pages/967442501/CCIP+Firedrills

## Structure
This repository contains of two main modules: `contracts` and `deployment`
1) `contracts` has smart contracts for CCIP 1.5 (`v1_5`) and CCIP 1.6 (`v1_6`). Contracts in `common` and `v1_0` used for both
`make build-all` also will build smart contracts using forge and will generate abi, bin and geth wrappers for them. 
These files can be found in `generated` folder
2) `deployment` folder contains changesets and views chainlink-deployments, while pkg contains multipurpose go-code. 
