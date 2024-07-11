// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package logs

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_logs_delivery_source", deliverySourceDataSource)
}

// deliverySourceDataSource returns the Terraform awscc_logs_delivery_source data source.
// This Terraform data source corresponds to the CloudFormation AWS::Logs::DeliverySource resource.
func deliverySourceDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) that uniquely identifies this delivery source.",
		//	  "maxLength": 2048,
		//	  "minLength": 16,
		//	  "pattern": "[\\w#+=/:,.@-]*\\*?",
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) that uniquely identifies this delivery source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: LogType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The type of logs being delivered. Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "[\\w-]*$",
		//	  "type": "string"
		//	}
		"log_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The type of logs being delivered. Only mandatory when the resourceArn could match more than one. In such a case, the error message will contain all the possible options.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The unique name of the Log source.",
		//	  "maxLength": 60,
		//	  "minLength": 1,
		//	  "pattern": "[\\w-]*$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The unique name of the Log source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The ARN of the resource that will be sending the logs.",
		//	  "maxLength": 2048,
		//	  "minLength": 16,
		//	  "pattern": "[\\w#+=/:,.@-]*\\*?",
		//	  "type": "string"
		//	}
		"resource_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The ARN of the resource that will be sending the logs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: ResourceArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "This array contains the ARN of the AWS resource that sends logs and is represented by this delivery source. Currently, only one ARN can be in the array.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "description": "The Amazon Resource Name (ARN) that uniquely identifies this delivery source.",
		//	    "maxLength": 2048,
		//	    "minLength": 16,
		//	    "pattern": "[\\w#+=/:,.@-]*\\*?",
		//	    "type": "string"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"resource_arns": schema.SetAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "This array contains the ARN of the AWS resource that sends logs and is represented by this delivery source. Currently, only one ARN can be in the array.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Service
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS service that is sending logs.",
		//	  "maxLength": 255,
		//	  "minLength": 1,
		//	  "pattern": "[\\w-]*$",
		//	  "type": "string"
		//	}
		"service": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS service that is sending logs.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The tags that have been assigned to this delivery source.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode",
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "The tags that have been assigned to this delivery source.",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::Logs::DeliverySource",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::Logs::DeliverySource").WithTerraformTypeName("awscc_logs_delivery_source")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":           "Arn",
		"key":           "Key",
		"log_type":      "LogType",
		"name":          "Name",
		"resource_arn":  "ResourceArn",
		"resource_arns": "ResourceArns",
		"service":       "Service",
		"tags":          "Tags",
		"value":         "Value",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
