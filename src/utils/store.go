package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"rconn/src/models"
	"sync"

	"github.com/billgraziano/dpapi"
)

type Store struct {
	path string
	list []models.RDPConnectionParams
}

var (
	storeInstance *Store
	once          sync.Once
	storeErr      error
)

func GetStore(customPath string) (*Store, error) {
	once.Do(func() {
		var path string

		if customPath != "" {
			abs, err := filepath.Abs(customPath)
			if err != nil {
				storeErr = fmt.Errorf("invalid path: %w", err)
				return
			}
			path = abs
		} else {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				storeErr = fmt.Errorf("could not get user config dir: %w", err)
				return
			}
			path = filepath.Join(homeDir, ".rconn", "connections.json")
		}

		storeInstance = &Store{path: path}
		storeErr = storeInstance.load()
	})

	return storeInstance, storeErr
}

func (s *Store) load() error {
	f, err := os.Open(s.path)
	if errors.Is(err, os.ErrNotExist) {
		s.list = []models.RDPConnectionParams{}
		return nil
	}
	if err != nil {
		return err
	}
	defer f.Close()

	var rawList []models.RDPConnectionParams
	if err := json.NewDecoder(f).Decode(&rawList); err != nil {
		return err
	}

	for i, conn := range rawList {
		if conn.Password != "" {
			dec, err := dpapi.Decrypt(conn.Password)
			if err != nil {
				return fmt.Errorf("failed to decrypt password for %q: %w", conn.Name, err)
			}
			rawList[i].Password = dec
		}
	}

	s.list = rawList
	return nil
}

func (s *Store) save() error {
	if err := os.MkdirAll(filepath.Dir(s.path), 0755); err != nil {
		return err
	}

	f, err := os.OpenFile(s.path, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	out := make([]models.RDPConnectionParams, len(s.list))
	copy(out, s.list)

	for i, conn := range out {
		if conn.Password != "" {
			enc, err := dpapi.Encrypt(conn.Password)
			if err != nil {
				return fmt.Errorf("failed to encrypt password for %q: %w", conn.Name, err)
			}
			out[i].Password = enc
		}
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	return enc.Encode(out)
}

func (s *Store) List() []models.RDPConnectionParams {
	return s.list
}

func (s *Store) Get(name string) (*models.RDPConnectionParams, error) {
	for i := range s.list {
		if s.list[i].Name == name {
			return &s.list[i], nil
		}
	}
	return nil, fmt.Errorf("no connection found with name %q", name)
}

func (s *Store) Has(name string) bool {
	for _, conn := range s.list {
		if conn.Name == name {
			return true
		}
	}
	return false
}

func (s *Store) Add(conn models.RDPConnectionParams) error {
	for _, existing := range s.list {
		if existing.Name == conn.Name {
			return fmt.Errorf("connection with name %q already exists", conn.Name)
		}
	}
	s.list = append(s.list, conn)
	return s.save()
}

func (s *Store) Update(conn models.RDPConnectionParams) error {
	for i := range s.list {
		if s.list[i].Name == conn.Name {
			s.list[i] = conn
			return s.save()
		}
	}
	return fmt.Errorf("no connection found with name %q", conn.Name)
}

func (s *Store) Remove(name string) error {
	for i, conn := range s.list {
		if conn.Name == name {
			s.list = append(s.list[:i], s.list[i+1:]...)
			return s.save()
		}
	}
	return fmt.Errorf("no connection found with name %q", name)
}
