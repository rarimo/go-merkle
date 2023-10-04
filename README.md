# Go Merkle tree light implementation

This implementation will construct the Merkle tree where parent hash are hash of concatenation of its leafs,
but firstly will go lexicographically smaller hash. Also, to check the path ypu should concatenate hashes in
lexicographically order. 

More precisely:

According to Openzeppelin contracts "<https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/utils/cryptography/MerkleProof.sol>",
Hash of `a` and `b` will be:
- `hash(ab)`, if a < b  
- `hash(ba)`, otherwise 


