/*
Copyright 2025 Guided Traffic.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package generator

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// Generator defines the interface for secret generation
type Generator interface {
	// GenerateString generates a random string of the specified length
	GenerateString(length int) (string, error)
	// GenerateBase64 generates a random base64 encoded string
	GenerateBase64(length int) (string, error)
	// Generate generates a value based on the specified type
	Generate(genType string, length int) (string, error)
}

// SecretGenerator implements the Generator interface using crypto/rand
type SecretGenerator struct {
	// charset is the character set used for string generation
	charset string
}

// DefaultCharset is the default character set for generating random strings
const DefaultCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()_+-=[]{}|;:,.<>?"

// AlphanumericCharset contains only alphanumeric characters
const AlphanumericCharset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// NewSecretGenerator creates a new SecretGenerator with default settings
func NewSecretGenerator() *SecretGenerator {
	return &SecretGenerator{
		charset: AlphanumericCharset,
	}
}

// NewSecretGeneratorWithCharset creates a new SecretGenerator with a custom charset
func NewSecretGeneratorWithCharset(charset string) *SecretGenerator {
	return &SecretGenerator{
		charset: charset,
	}
}

// GenerateString generates a random string of the specified length
func (g *SecretGenerator) GenerateString(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be positive, got %d", length)
	}

	result := make([]byte, length)
	charsetLen := len(g.charset)

	// Generate random bytes
	randomBytes := make([]byte, length)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Map random bytes to charset characters
	for i := 0; i < length; i++ {
		result[i] = g.charset[int(randomBytes[i])%charsetLen]
	}

	return string(result), nil
}

// GenerateBase64 generates a random base64 encoded string
// The length parameter specifies the number of random bytes before encoding
func (g *SecretGenerator) GenerateBase64(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be positive, got %d", length)
	}

	randomBytes := make([]byte, length)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	return base64.StdEncoding.EncodeToString(randomBytes), nil
}

// Generate generates a value based on the specified type
func (g *SecretGenerator) Generate(genType string, length int) (string, error) {
	switch genType {
	case "string", "":
		return g.GenerateString(length)
	case "base64":
		return g.GenerateBase64(length)
	default:
		return "", fmt.Errorf("unknown generation type: %s", genType)
	}
}
