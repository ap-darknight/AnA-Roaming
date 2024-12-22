package crypto_objects

import "encoding/json"

type CryptoObject struct {
	Key1 string
	Key2 string
	Key3 int64
}

// ToBytes serializes the CryptoObject into a byte slice (JSON encoding)
func (c *CryptoObject) ToBytes() ([]byte, error) {
	return json.Marshal(c)
}

// FromBytes deserializes the byte slice into a CryptoObject (JSON decoding)
func FromBytes(data []byte) (*CryptoObject, error) {
	var obj CryptoObject
	err := json.Unmarshal(data, &obj)
	if err != nil {
		return nil, err
	}
	return &obj, nil
}
