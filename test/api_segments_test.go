/*
Elastic Email REST API

Testing SegmentsApiService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package ElasticEmail

import (
    "context"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
    "testing"
    openapiclient "./openapi"
)

func Test_ElasticEmail_SegmentsApiService(t *testing.T) {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)

    t.Run("Test SegmentsApiService SegmentsByNameDelete", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var name string

        resp, httpRes, err := apiClient.SegmentsApi.SegmentsByNameDelete(context.Background(), name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SegmentsApiService SegmentsByNameGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var name string

        resp, httpRes, err := apiClient.SegmentsApi.SegmentsByNameGet(context.Background(), name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SegmentsApiService SegmentsByNamePut", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        var name string

        resp, httpRes, err := apiClient.SegmentsApi.SegmentsByNamePut(context.Background(), name).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SegmentsApiService SegmentsGet", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SegmentsApi.SegmentsGet(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

    t.Run("Test SegmentsApiService SegmentsPost", func(t *testing.T) {

        t.Skip("skip test")  // remove to run test

        resp, httpRes, err := apiClient.SegmentsApi.SegmentsPost(context.Background()).Execute()

        require.Nil(t, err)
        require.NotNil(t, resp)
        assert.Equal(t, 200, httpRes.StatusCode)

    })

}