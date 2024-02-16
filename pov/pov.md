# Proof of Validation

A gossip network,

n1 -> n2 -> n3

n3 -> n4 -> n5

n5 -> n6 -> n7

TX1 - 100001
TX2 - 100006
TX3 - 100008

BLOCK :

1 - TX1
2 - TX3
3 - TX2

n2, n3, n4, n5, n6, n7

node <- TX3 100008 
    -> RECEIVE TIME: 100018

node <- BLOCK
    -> TX1, TX3, TX2

transactions : 
    -> TX1 100001 : 100004
    -> TX2 100006 : 100018
    -> TX3 100008 : 100018

TRANSACTION!!! -> SIGNATURE OF THE SENDER

n1 -> n2
        -> n3
            -> n5
            -> n6
        -> n4
            -> n7
            -> n8


BLOCK : HASH
LIST OF TRANSACTIONS -> 100 TRANSACTIONS

1000 BYTES EACH TX

POW (SHA256 hash with a subsequent number of zeros at the begin)

NODE

tx1 = 10003 -> 10005
tx2 = 10004 -> 10015
tx5 = 10008 -> 10030

BLOCK = More of an update to the network, an execution of all the gathered valid transactions in the valid order

proof-of-validation

tx1->tx2->tx3 = 256HASH 
    -> LIST OF TRANSACTIONS IN SUBSEQUENT ORDER

STATE = 256HASH

PREVIOUS_STATE_HASH, UPDATED_STATE_HASH, ORDER_HASH, (LIST OF TRANSACTIONS: LIST OF TRANSACTION HASHES)

1mb -> 100 kb

SHA256, SHA256, SHA256, LIST[SHA256]

VALIDATOR_LIST = [
    node1,
    node2,
    node7,
    node8,
    node9,
    node10,
    node11,
    node13
]

VALIDATOR_LIST.SEND(SHA256, SHA256, SHA256, LIST[SHA256])
    -> VALIDATION(SIGNATURES AND PROOFS), ARGUMENTS




MONETARY REWARDS FOR PARTICIPATING IN THE NETWORK!!!