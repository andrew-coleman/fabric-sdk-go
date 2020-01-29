/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package gateway

import "errors"

// FileSystemWallet stores identity information used to connect to a Hyperledger Fabric network.
// Instances are created using NewFileSystemWallet()
type FileSystemWallet struct {
	idhandler IDHandler
	path      string
	storage   map[string]map[string]string
}

// NewFileSystemWallet creates an instance of a wallet, held in memory.
// This implementation is not backed by a persistent store.
func NewFileSystemWallet(path string) *FileSystemWallet {
	return &FileSystemWallet{newX509IdentityHandler(), path, make(map[string]map[string]string, 10)}
}

// Put an identity into the wallet.
func (f *FileSystemWallet) Put(label string, id IdentityType) error {
	elements := f.idhandler.GetElements(id)
	f.storage[label] = elements
	return nil
}

// Get an identity from the wallet.
func (f *FileSystemWallet) Get(label string) (IdentityType, error) {
	if elements, ok := f.storage[label]; ok {
		return f.idhandler.FromElements(elements), nil
	}
	return nil, errors.New("label doesn't exist: " + label)
}

// Remove an identity from the wallet. If the identity does not exist, this method does nothing.
func (f *FileSystemWallet) Remove(label string) error {
	if _, ok := f.storage[label]; ok {
		delete(f.storage, label)
		return nil
	}
	return nil // what should we do here ?
}

// Exists returns true if the identity is in the wallet.
func (f *FileSystemWallet) Exists(label string) bool {
	_, ok := f.storage[label]
	return ok
}

// List all of the labels in the wallet.
func (f *FileSystemWallet) List() []string {
	labels := make([]string, 0, len(f.storage))
	for label := range f.storage {
		labels = append(labels, label)
	}
	return labels
}
