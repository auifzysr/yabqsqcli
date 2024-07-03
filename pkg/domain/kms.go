package domain

import "fmt"

type KMS struct {
	ProjectID string
	Location  string
	KeyRing   string
	Key       string
}

// https://cloud.google.com/bigquery/docs/customer-managed-encryption#key_resource_id
// projects/KMS_PROJECT_ID/locations/LOCATION/keyRings/KEY_RING/cryptoKeys/KEY

func (kms *KMS) ResourceID() (string, error) {
	if kms.ProjectID == "" {
		return "", fmt.Errorf("projectID is required")
	}
	if kms.Location == "" {
		return "", fmt.Errorf("location is required")
	}
	if kms.KeyRing == "" {
		return "", fmt.Errorf("keyRing is required")
	}
	if kms.Key == "" {
		return "", fmt.Errorf("key is required")
	}

	return fmt.Sprintf(
		"projects/%s/locations/%s/keyRings/%s/cryptoKeys/%s",
		kms.ProjectID,
		kms.Location,
		kms.KeyRing,
		kms.Key,
	), nil
}
