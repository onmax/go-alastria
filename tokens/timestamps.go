package tokens

import (
	"fmt"
	"time"
)

func CurrentTimestamp() uint64 {
	return uint64(time.Now().UnixNano()) / uint64(time.Millisecond)
}

func ValidateTimestamps(iat *uint64, exp *uint64, nbf *uint64) error {
	if err := ValidateIssuedAt(iat); err != nil {
		return err
	}
	if err := ValidateNotBefore(nbf); err != nil {
		return err
	}
	if err := ValidateExpiresAt(exp); err != nil {
		return err
	}

	if *exp != 0 && *iat > *exp {
		return fmt.Errorf("the value IssuedAt must be before ExpiresAt. IssuedAt: %v. ExpiresAt: %v", *iat, *exp)
	}

	if *nbf != 0 && *nbf > *exp {
		return fmt.Errorf("the value NotBefore must be before ExpiresAt. ExpiresAt: %v. NotBefore: %v", *exp, *nbf)
	}

	return nil
}

func ValidateIssuedAt(iat *uint64) error {
	if iat == nil {
		return fmt.Errorf(emptyPayloadField, "IssuedAt")
	}
	timestamp := CurrentTimestamp()

	// Checks if iat is in the future
	if *iat > timestamp {
		return fmt.Errorf("the value IssuedAt is greater than the current timestamp. IssuedAt: %v. Current Timestamp: %v", *iat, timestamp)
	}

	if *iat == 0 {
		*iat = timestamp
	}

	return nil
}

func ValidateExpiresAt(exp *uint64) error {
	if exp == nil {
		return fmt.Errorf(emptyPayloadField, "ExpiresAt")
	}

	timestamp := CurrentTimestamp()

	// Checks if exp is in the past
	if *exp != 0 && *exp < timestamp {
		return fmt.Errorf("the value ExpiresAt is less than the current timestamp. ExpiresAt: %v. Current Timestamp: %v", *exp, timestamp)
	}

	return nil
}

func ValidateNotBefore(nbf *uint64) error {
	if nbf == nil {
		return fmt.Errorf(emptyPayloadField, "NotBefore")
	}

	return nil
}
