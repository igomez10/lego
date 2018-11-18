## lego

### Account

  - publicKey
  - privateKey
  - amount

  Module for account management. Contains public, private keys, amount. The account creation is tested with basic public/private key encryption

### Transaction
  Used to store a transaction between two accounts. The validity of the transaction is tested via checking the sign of both accounts. If both accounts sign the string representation fo the transaction, then the transaction is valid. Because of this, a ```Sign()``` method is required in the account module

### Block
### Blockchain
