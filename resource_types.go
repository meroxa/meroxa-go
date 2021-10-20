package meroxa

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type ResourceType string

const (
	ResourceTypePostgres      ResourceType = "postgres"
	ResourceTypeMysql                      = "mysql"
	ResourceTypeRedshift                   = "redshift"
	ResourceTypeUrl                        = "url"
	ResourceTypeS3                         = "s3"
	ResourceTypeMongodb                    = "mongodb"
	ResourceTypeElasticsearch              = "elasticsearch"
	ResourceTypeSnowflake                  = "snowflake"
	ResourceTypeBigquery                   = "bigquery"
	ResourceTypeSqlserver                  = "sqlserver"
	ResourceTypeCosmosdb                   = "cosmosdb"
)


// ListResourceTypes returns the list of supported resources
func (c *Client) ListResourceTypes(ctx context.Context) ([]string, error) {
	path := fmt.Sprintf("/v1/resource-types")

	resp, err := c.MakeRequest(ctx, http.MethodGet, path, nil, nil)
	if err != nil {
		return nil, err
	}

	err = handleAPIErrors(resp)
	if err != nil {
		return nil, err
	}

	var supportedTypes []string
	err = json.NewDecoder(resp.Body).Decode(&supportedTypes)
	if err != nil {
		return nil, err
	}

	return supportedTypes, nil
}
