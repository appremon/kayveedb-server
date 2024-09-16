
package core

import (
	"github.com/rickcollette/kayveedb"
	"log"
)

var (
	HmacKey       = []byte("32-byte-hmac-key")
	EncryptionKey = []byte("32-byte-encryption-key")
	Nonce         = []byte("24-byte-nonce")
)

func LoadBtree(dbPath, dbName, logName string) (*kayveedb.BTree, error) {
	btree, err := kayveedb.NewBTree(3, dbPath, dbName, logName, HmacKey, EncryptionKey, Nonce, 100)
	if err != nil {
		log.Fatalf("Failed to load B-tree: %v", err)
	}
	return btree, nil
}
