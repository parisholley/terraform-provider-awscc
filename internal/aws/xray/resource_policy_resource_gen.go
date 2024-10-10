// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package xray

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_xray_resource_policy", resourcePolicyResource)
}

// resourcePolicyResource returns the Terraform awscc_xray_resource_policy resource.
// This Terraform resource corresponds to the CloudFormation AWS::XRay::ResourcePolicy resource.
func resourcePolicyResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: BypassPolicyLockoutCheck
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A flag to indicate whether to bypass the resource policy lockout safety check",
		//	  "type": "boolean"
		//	}
		"bypass_policy_lockout_check": schema.BoolAttribute{ /*START ATTRIBUTE*/
			Description: "A flag to indicate whether to bypass the resource policy lockout safety check",
			Optional:    true,
			// BypassPolicyLockoutCheck is a write-only property.
		}, /*END ATTRIBUTE*/
		// Property: PolicyDocument
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The resource policy document, which can be up to 5kb in size.",
		//	  "maxLength": 5120,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"policy_document": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The resource policy document, which can be up to 5kb in size.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 5120),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: PolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the resource policy. Must be unique within a specific AWS account.",
		//	  "maxLength": 128,
		//	  "minLength": 1,
		//	  "pattern": "[\\w+=,.@-]+",
		//	  "type": "string"
		//	}
		"policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the resource policy. Must be unique within a specific AWS account.",
			Required:    true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(1, 128),
				stringvalidator.RegexMatches(regexp.MustCompile("[\\w+=,.@-]+"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
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
		Description: "This schema provides construct and validation rules for AWS-XRay Resource Policy resource parameters.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::XRay::ResourcePolicy").WithTerraformTypeName("awscc_xray_resource_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"bypass_policy_lockout_check": "BypassPolicyLockoutCheck",
		"policy_document":             "PolicyDocument",
		"policy_name":                 "PolicyName",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/BypassPolicyLockoutCheck",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
