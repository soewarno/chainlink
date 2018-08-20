#!/usr/bin/env node

const {deployer, abort} = require('./integration/common.js')
const web3 = require('web3')

const main = async (linkTokenAddress, oracleContractAddress) => {
  console.log('linkTokenAddress', linkTokenAddress, 'oracleContractAddress', oracleContractAddress)

  const LinkToken = await deployer.load('node_modules/linkToken/contracts/LinkToken.sol', linkTokenAddress)
    .catch(abort(`Error loading LinkToken at address ${linkTokenAddress}`))

  const Contract = await deployer.perform(
    'internal/fixtures/contracts/ethbid.sol',
    LinkToken.address,
    oracleContractAddress
  ).catch(abort('Error deploying ethbid contract'))
  console.log(`Deployed Eth Bidding contract at: ${Contract.address}`)

  const amount = web3.utils.toWei('10000')
  await LinkToken.transfer(Contract.address, amount, {gas: 100000})
    .catch(abort('Error transferring link to RunLog'))
  console.log(`Transferred ${amount} to Contract at: ${Contract.address}`)

  let result = await Contract.requestEthereumPrice()
    .catch((err) => abort(`Failed to send requestEthereumPrice: ${err}`))

  console.log('result', result)
}

main(process.argv[2], process.argv[3])
