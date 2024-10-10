// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package datasync

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	fwvalidators "github.com/hashicorp/terraform-provider-awscc/internal/validators"
)

func init() {
	registry.AddResourceFactory("awscc_datasync_location_azure_blob", locationAzureBlobResource)
}

// locationAzureBlobResource returns the Terraform awscc_datasync_location_azure_blob resource.
// This Terraform resource corresponds to the CloudFormation AWS::DataSync::LocationAzureBlob resource.
func locationAzureBlobResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AgentArns
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Names (ARNs) of agents to use for an Azure Blob Location.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "maxLength": 128,
		//	    "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$",
		//	    "type": "string"
		//	  },
		//	  "maxItems": 4,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"agent_arns": schema.ListAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Description: "The Amazon Resource Names (ARNs) of agents to use for an Azure Blob Location.",
			Required:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 4),
				listvalidator.ValueStringsAre(
					stringvalidator.LengthAtMost(128),
					stringvalidator.RegexMatches(regexp.MustCompile("^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:agent/agent-[0-9a-z]{17}$"), ""),
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AzureAccessTier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "HOT",
		//	  "description": "Specifies an access tier for the objects you're transferring into your Azure Blob Storage container.",
		//	  "enum": [
		//	    "HOT",
		//	    "COOL",
		//	    "ARCHIVE"
		//	  ],
		//	  "type": "string"
		//	}
		"azure_access_tier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies an access tier for the objects you're transferring into your Azure Blob Storage container.",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString("HOT"),
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"HOT",
					"COOL",
					"ARCHIVE",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AzureBlobAuthenticationType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "SAS",
		//	  "description": "The specific authentication type that you want DataSync to use to access your Azure Blob Container.",
		//	  "enum": [
		//	    "SAS"
		//	  ],
		//	  "type": "string"
		//	}
		"azure_blob_authentication_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The specific authentication type that you want DataSync to use to access your Azure Blob Container.",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString("SAS"),
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"SAS",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: AzureBlobContainerUrl
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the Azure Blob container that was described.",
		//	  "maxLength": 325,
		//	  "pattern": "^https://[A-Za-z0-9]((.|-+)?[A-Za-z0-9]){0,252}/[a-z0-9](-?[a-z0-9]){2,62}$",
		//	  "type": "string"
		//	}
		"azure_blob_container_url": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the Azure Blob container that was described.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(325),
				stringvalidator.RegexMatches(regexp.MustCompile("^https://[A-Za-z0-9]((.|-+)?[A-Za-z0-9]){0,252}/[a-z0-9](-?[a-z0-9]){2,62}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
				stringplanmodifier.RequiresReplaceIfConfigured(),
			}, /*END PLAN MODIFIERS*/
			// AzureBlobContainerUrl is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: AzureBlobSasConfiguration
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "description": "Specifies the shared access signature (SAS) that DataSync uses to access your Azure Blob Storage container.",
		//	  "properties": {
		//	    "AzureBlobSasToken": {
		//	      "description": "Specifies the shared access signature (SAS) token, which indicates the permissions DataSync needs to access your Azure Blob Storage container.",
		//	      "maxLength": 255,
		//	      "minLength": 1,
		//	      "pattern": "(^.+$)",
		//	      "type": "string"
		//	    }
		//	  },
		//	  "required": [
		//	    "AzureBlobSasToken"
		//	  ],
		//	  "type": "object"
		//	}
		"azure_blob_sas_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: AzureBlobSasToken
				"azure_blob_sas_token": schema.StringAttribute{ /*START ATTRIBUTE*/
					Description: "Specifies the shared access signature (SAS) token, which indicates the permissions DataSync needs to access your Azure Blob Storage container.",
					Required:    true,
					Validators: []validator.String{ /*START VALIDATORS*/
						stringvalidator.LengthBetween(1, 255),
						stringvalidator.RegexMatches(regexp.MustCompile("(^.+$)"), ""),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Description: "Specifies the shared access signature (SAS) that DataSync uses to access your Azure Blob Storage container.",
			Optional:    true,
			// AzureBlobSasConfiguration is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: AzureBlobType
		// CloudFormation resource type schema:
		//
		//	{
		//	  "default": "BLOCK",
		//	  "description": "Specifies a blob type for the objects you're transferring into your Azure Blob Storage container.",
		//	  "enum": [
		//	    "BLOCK"
		//	  ],
		//	  "type": "string"
		//	}
		"azure_blob_type": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "Specifies a blob type for the objects you're transferring into your Azure Blob Storage container.",
			Optional:    true,
			Computed:    true,
			Default:     stringdefault.StaticString("BLOCK"),
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.OneOf(
					"BLOCK",
				),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocationArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The Amazon Resource Name (ARN) of the Azure Blob Location that is created.",
		//	  "maxLength": 128,
		//	  "pattern": "^arn:(aws|aws-cn|aws-us-gov|aws-iso|aws-iso-b):datasync:[a-z\\-0-9]+:[0-9]{12}:location/loc-[0-9a-z]{17}$",
		//	  "type": "string"
		//	}
		"location_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The Amazon Resource Name (ARN) of the Azure Blob Location that is created.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: LocationUri
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The URL of the Azure Blob Location that was described.",
		//	  "maxLength": 4356,
		//	  "pattern": "^(azure-blob)://[a-zA-Z0-9./\\-]+$",
		//	  "type": "string"
		//	}
		"location_uri": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The URL of the Azure Blob Location that was described.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Subdirectory
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The subdirectory in the Azure Blob Container that is used to read data from the Azure Blob Source Location.",
		//	  "maxLength": 1024,
		//	  "pattern": "^[\\p{L}\\p{M}\\p{Z}\\p{S}\\p{N}\\p{P}\\p{C}]*$",
		//	  "type": "string"
		//	}
		"subdirectory": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The subdirectory in the Azure Blob Container that is used to read data from the Azure Blob Source Location.",
			Optional:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(1024),
				stringvalidator.RegexMatches(regexp.MustCompile("^[\\p{L}\\p{M}\\p{Z}\\p{S}\\p{N}\\p{P}\\p{C}]*$"), ""),
			}, /*END VALIDATORS*/
			// Subdirectory is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An array of key-value pairs to apply to this resource.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource.",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:/-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for an AWS resource tag.",
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9\\s+=._:@/-]+$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Key",
		//	      "Value"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 50,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key for an AWS resource tag.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:/-]+$"), ""),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for an AWS resource tag.",
						Optional:    true,
						Computed:    true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
							stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9\\s+=._:@/-]+$"), ""),
							fwvalidators.NotNullString(),
						}, /*END VALIDATORS*/
						PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
							stringplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An array of key-value pairs to apply to this resource.",
			Optional:    true,
			Computed:    true,
			Validators: []validator.Set{ /*START VALIDATORS*/
				setvalidator.SizeAtMost(50),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.Set{ /*START PLAN MODIFIERS*/
				setplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Resource schema for AWS::DataSync::LocationAzureBlob.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::DataSync::LocationAzureBlob").WithTerraformTypeName("awscc_datasync_location_azure_blob")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"agent_arns":                     "AgentArns",
		"azure_access_tier":              "AzureAccessTier",
		"azure_blob_authentication_type": "AzureBlobAuthenticationType",
		"azure_blob_container_url":       "AzureBlobContainerUrl",
		"azure_blob_sas_configuration":   "AzureBlobSasConfiguration",
		"azure_blob_sas_token":           "AzureBlobSasToken",
		"azure_blob_type":                "AzureBlobType",
		"key":                            "Key",
		"location_arn":                   "LocationArn",
		"location_uri":                   "LocationUri",
		"subdirectory":                   "Subdirectory",
		"tags":                           "Tags",
		"value":                          "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Subdirectory",
		"/properties/AzureBlobSasConfiguration",
		"/properties/AzureBlobContainerUrl",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
