// pkg/utils/helpers.go
package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// HasWhatsAppPrefix checks if a phone number has the WhatsApp prefix
func HasWhatsAppPrefix(phoneNumber string) bool {
	return strings.HasPrefix(phoneNumber, "whatsapp:")
}

// GetPlaceholderIndex converts an index to a string for SQL placeholders
func GetPlaceholderIndex(index int) string {
	return strconv.Itoa(index)
}

// IsValidPhoneNumber checks if a phone number is valid
// This is a simplified implementation and should be replaced with
// proper phone number validation in a production environment
func IsValidPhoneNumber(phoneNumber string) bool {
	// Remove WhatsApp prefix if present
	normalizedNumber := phoneNumber
	if HasWhatsAppPrefix(phoneNumber) {
		normalizedNumber = strings.TrimPrefix(phoneNumber, "whatsapp:")
	}

	// Remove any non-digit characters
	digitsOnly := strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, normalizedNumber)

	// Check if the number has at least 10 digits (simplified validation)
	return len(digitsOnly) >= 10
}

// FormatPhoneNumber formats a phone number for WhatsApp
func FormatPhoneNumber(phoneNumber string) string {
	// Already has WhatsApp prefix
	if HasWhatsAppPrefix(phoneNumber) {
		return phoneNumber
	}

	// Add WhatsApp prefix
	return "whatsapp:" + phoneNumber
}

// AnyToString converts any value to a string representation
func AnyToString(value interface{}) string {
	if value == nil {
		return ""
	}

	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case float64:
		return strconv.FormatFloat(v, 'f', -1, 64)
	case bool:
		return strconv.FormatBool(v)
	default:
		return fmt.Sprintf("%v", v)
	}
}