/*
Nudm_SDM

Testing SubscriptionCreationAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_SubscriptionCreationAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubscriptionCreationAPIService Subscribe", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var ueId string

		resp, httpRes, err := apiClient.SubscriptionCreationAPI.Subscribe(context.Background(), ueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
