package pov

import (
	"crypto/sha256"
	"errors"
	"log"

	"github.com/cbergoon/merkletree"
)

// TestContent implements the Content interface provided by merkletree and represents the content stored in the tree.
type MerkleTransaction struct {
	x string
}

// CalculateHash hashes the values of a TestContent
func (t MerkleTransaction) CalculateHash() ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write([]byte(t.x)); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals tests for equality of two Contents
func (t MerkleTransaction) Equals(other merkletree.Content) (bool, error) {
	otherTC, ok := other.(MerkleTransaction)
	if !ok {
		return false, errors.New("value is not of type TestContent")
	}
	return t.x == otherTC.x, nil
}

func TestMerkle() {
	//Build list of Content to build tree
	var list []merkletree.Content
	list = append(list, MerkleTransaction{x: "0x6265617665726275696c642e6f72670x6265617665726275696c642e6f7267"})
	list = append(list, MerkleTransaction{x: "0x6265617665726275696c642e6f7267"})
	list = append(list, MerkleTransaction{x: "0x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f7267"})
	list = append(list, MerkleTransaction{x: "0x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f72670x6265617665726275696c642e6f7267"})

	//Create a new Merkle Tree from the list of Content
	t, err := merkletree.NewTree(list)
	if err != nil {
		log.Fatal(err)
	}

	//Get the Merkle Root of the tree
	mr := t.MerkleRoot()
	log.Println(mr)

	//Verify the entire tree (hashes for each node) is valid
	vt, err := t.VerifyTree()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Verify Tree: ", vt)

	//Verify a specific content in in the tree
	vc, err := t.VerifyContent(list[0])
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Verify Content: ", vc)

	//String representation
	//log.Println(t)
}
