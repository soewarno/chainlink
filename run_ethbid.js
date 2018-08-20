#!/usr/bin/env node

const {deployer, abort} = require('./integration/common.js')

const main = async (contractAddress) => {
  console.log('contractAddress', contractAddress)

  const Contract = await deployer.load('internal/fixtures/contracts/ethbid.sol', contractAddress).catch(abort('Error deploying ethbid contract'))
  console.log(`Loaded Eth Bidding contract at: ${Contract.address}`)

  let result = await Contract.requestEthereumPrice()
    .catch((err) => abort(`Failed to send requestEthereumPrice: ${err}`))

  console.log('result', result)
}

main(process.argv[2])
