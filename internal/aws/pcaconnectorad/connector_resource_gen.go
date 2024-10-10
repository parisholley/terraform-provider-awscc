// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package pcaconnectorad

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_pcaconnectorad_connector", connectorResource)
}

// connectorResource returns the Terraform awscc_pcaconnectorad_connector resource.
// This Terraform resource corresponds to the CloudFormation AWS::PCAConnectorAD::Connector resource.
func connectorResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CertificateAuthorityArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 200,
		//	  "minLength": 5,
		//	  "pattern": "^arn:[\\w-]+:acm-pca:[\\w-]+:[0-9]+:certificate-authority(\\/[\\w-]+)$",
		//	  "type": "string"
		//	}
		"certificate_authority_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(5, 200),
				stringvalidator.RegexMatches(regexp.MustCompile("^arn:[\\w-]+:acm-pca:[\\w-]+:[0-9]+:certificate-authority(\\/[\\w-]+)$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// CertificateAuthorityArn is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: ConnectorArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 200,
		//	  "minLength": 5,
		//	  "pattern": "^arn:[\\w-]+:pca-connector-ad:[\\w-]+:[0-9]+:connector(\\/[\\w-]+)$",
		//	  "type": "string"
		//	}
		"connector_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: DirectoryId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^d-[0-9a-f]{10}$",
		//	  "type": "string"
		//	}
		"directory_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.RegexMatches(regexp.MustCompile("^d-[0-9a-f]{10}$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// DirectoryId is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "patternProperties": {
		//	    "": {
		//	      "type": "string"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"tags":              // Pattern: ""
		schema.MapAttribute{ /*START ATTRIBUTE*/
			ElementType: types.StringType,
			Optional:    true,
			// Tags is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: VpcInformation
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "SecurityGroupIds": {
		//	      "items": {
		//	        "maxLength": 20,
		//	        "minLength": 11,
		//	        "pattern": "^(?:sg-[0-9a-f]{8}|sg-[0-9a-f]{17})$",
		//	        "type": "string"
		//	      },
		//	      "maxItems": 5,
		//	      "minItems": 1,
		//	      "type": "array",
		//	      "uniqueItems": true
		//	    }
		//	  },
		//	  "required": [
		//	    "SecurityGroupIds"
		//	  ],
		//	  "type": "object"
		//	}
		"vpc_information": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: SecurityGroupIds
				"security_group_ids": schema.ListAttribute{ /*START ATTRIBUTE*/
					ElementType: types.StringType,
					Required:    true,
					Validators: []validator.List{ /*START VALIDATORS*/
						listvalidator.SizeBetween(1, 5),
						listvalidator.UniqueValues(),
						listvalidator.ValueStringsAre(
							stringvalidator.LengthBetween(11, 20),
							stringvalidator.RegexMatches(regexp.MustCompile("^(?:sg-[0-9a-f]{8}|sg-[0-9a-f]{17})$"), ""),
						),
					}, /*END VALIDATORS*/
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Required: true,
			PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
				objectplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
			// VpcInformation is a write-only property.
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
		Description: "Definition of AWS::PCAConnectorAD::Connector Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::PCAConnectorAD::Connector").WithTerraformTypeName("awscc_pcaconnectorad_connector")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"certificate_authority_arn": "CertificateAuthorityArn",
		"connector_arn":             "ConnectorArn",
		"directory_id":              "DirectoryId",
		"security_group_ids":        "SecurityGroupIds",
		"tags":                      "Tags",
		"vpc_information":           "VpcInformation",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/CertificateAuthorityArn",
		"/properties/DirectoryId",
		"/properties/Tags",
		"/properties/VpcInformation",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
