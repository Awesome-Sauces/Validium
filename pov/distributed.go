package validium

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ethereum/go-ethereum/crypto"
)

type Sig struct {
	Signature []byte
	Address   []byte
	ID        int
}

type Transaction struct {
	Sender     []byte
	Payload    []byte
	Originator int
	Signatures []Sig
}

type Node struct {
	ID           int
	Listener     chan Transaction
	Transactions map[string]Transaction
	Peers        map[int]chan Transaction
}

func NewNode(id int) *Node {
	//fmt.Printf("Creating Node %d\n", id)
	return &Node{
		ID:           id,
		Listener:     make(chan Transaction, 100), // Buffer to prevent blocking
		Transactions: make(map[string]Transaction),
		Peers:        make(map[int]chan Transaction),
	}
}

func (n *Node) PeerConnect(id int, channel chan Transaction) {
	//fmt.Printf("Node %d connecting to Node %d\n", n.ID, id)
	n.Peers[id] = channel
}

func (n *Node) Listen(logChan chan string) {
	//fmt.Printf("Node %d started listening\n", n.ID)
	for tx := range n.Listener {
		hash := crypto.Keccak256Hash(tx.Payload, tx.Sender).Hex()
		if _, exists := n.Transactions[hash]; !exists {
			n.Transactions[hash] = tx
			logMsg := fmt.Sprintf("Node %d received transaction from %d\n", n.ID, tx.Originator)
			//fmt.Print(logMsg)
			logChan <- logMsg
			for peerID, peerChan := range n.Peers {
				logMsg := fmt.Sprintf("Node %d sending to Node %d\n", n.ID, peerID)
				//fmt.Print(logMsg)
				logChan <- logMsg
				peerChan <- tx
			}
		}
	}
}

func setupNodes(nodeCount int) []*Node {
	//fmt.Println("Setting up nodes...")
	nodes := make([]*Node, nodeCount)
	for i := range nodes {
		nodes[i] = NewNode(i)
	}
	return nodes
}

func connectRandomPeers(nodes []*Node, peersCount int) {
	//fmt.Println("Connecting peers randomly...")
	rand.Seed(time.Now().UnixNano())
	for _, node := range nodes {
		perm := rand.Perm(len(nodes))
		for _, idx := range perm {
			if len(node.Peers) >= peersCount {
				break
			}
			if nodes[idx].ID != node.ID && node.Peers[nodes[idx].ID] == nil {
				node.PeerConnect(nodes[idx].ID, nodes[idx].Listener)
			}
		}
	}
}

func TestNetworkPropagation() {
	const nodeCount = 1000
	const peersCount = 3 // Number of peers per node
	nodes := setupNodes(nodeCount)
	connectRandomPeers(nodes, peersCount)

	logChan := make(chan string, 10000) // Large buffer to prevent blocking

	for _, node := range nodes {
		go func(n *Node) {
			n.Listen(logChan)
		}(node)
	}

	// Wait a bit before sending the transaction to ensure all nodes are listening
	time.Sleep(1 * time.Second)

	// Sending a transaction from a random node to test propagation
	randNode := nodes[rand.Intn(nodeCount)]
	//fmt.Printf("Sending transaction from Node %d\n", randNode.ID)
	tx := Transaction{
		Sender:     []byte("sender"),
		Payload:    []byte("payload"),
		Originator: randNode.ID,
		Signatures: []Sig{{Signature: []byte("signature"), Address: []byte("address"), ID: 1}},
	}

	randNode.Listener <- tx

	// Let the network process the transaction
	time.Sleep(100 * time.Millisecond)
	close(logChan) // Close log channel to finish logging

	// Wait for all Listen goroutines to finish
	//wg.Wait()

	fmt.Println("HELLO!")

	for _, node := range nodes {
		for hash := range node.Transactions {
			fmt.Printf("Node -| %d |- with tx -| %s |-", node.ID, hash[:5])
		}
	}
}
