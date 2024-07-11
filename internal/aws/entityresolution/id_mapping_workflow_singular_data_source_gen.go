// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package entityresolution

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_entityresolution_id_mapping_workflow", idMappingWorkflowDataSource)
}

// idMappingWorkflowDataSource returns the Terraform awscc_entityresolution_id_mapping_workflow data source.
// This Terraform data source corresponds to the CloudFormation AWS::EntityResolution::IdMappingWorkflow resource.
func idMappingWorkflowDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: CreatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time of this IdMappingWorkflow got created",
		//	  "type": "string"
		//	}
		"created_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time of this IdMappingWorkflow got created",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The description of the IdMappingWorkflow",
		//	  "maxLength": 255,
		//	  "minLength": 0,
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The description of the IdMappingWorkflow",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: IdMappingTechniques
		// CloudFormation resource type schema:
		//
		//	{
		//	  "additionalProperties": false,
		//	  "properties": {
		//	    "IdMappingType": {
		//	      "enum": [
		//	        "PROVIDER",
		//	        "RULE_BASED"
		//	      ],
		//	      "type": "string"
		//	    },
		//	    "NormalizationVersion": {
		//	      "type": "string"
		//	    },
		//	    "ProviderProperties": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "IntermediateSourceConfiguration": {
		//	          "additionalProperties": false,
		//	          "properties": {
		//	            "IntermediateS3Path": {
		//	              "description": "The s3 path that would be used to stage the intermediate data being generated during workflow execution.",
		//	              "type": "string"
		//	            }
		//	          },
		//	          "required": [
		//	            "IntermediateS3Path"
		//	          ],
		//	          "type": "object"
		//	        },
		//	        "ProviderConfiguration": {
		//	          "additionalProperties": false,
		//	          "description": "Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format",
		//	          "patternProperties": {
		//	            "": {
		//	              "type": "string"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "ProviderServiceArn": {
		//	          "description": "Arn of the Provider Service being used.",
		//	          "pattern": "^arn:(aws|aws-us-gov|aws-cn):entityresolution:([A-Za-z0-9]+(-[A-Za-z0-9]+)+)::providerservice/[A-Za-z0-9]+/[A-Za-z0-9]+$",
		//	          "type": "string"
		//	        }
		//	      },
		//	      "required": [
		//	        "ProviderServiceArn"
		//	      ],
		//	      "type": "object"
		//	    },
		//	    "RuleBasedProperties": {
		//	      "additionalProperties": false,
		//	      "properties": {
		//	        "AttributeMatchingModel": {
		//	          "enum": [
		//	            "ONE_TO_ONE",
		//	            "MANY_TO_MANY"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "RecordMatchingModel": {
		//	          "enum": [
		//	            "ONE_SOURCE_TO_ONE_TARGET",
		//	            "MANY_SOURCE_TO_ONE_TARGET"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "RuleDefinitionType": {
		//	          "enum": [
		//	            "SOURCE",
		//	            "TARGET"
		//	          ],
		//	          "type": "string"
		//	        },
		//	        "Rules": {
		//	          "insertionOrder": false,
		//	          "items": {
		//	            "additionalProperties": false,
		//	            "properties": {
		//	              "MatchingKeys": {
		//	                "insertionOrder": false,
		//	                "items": {
		//	                  "maxLength": 255,
		//	                  "minLength": 0,
		//	                  "pattern": "^[a-zA-Z_0-9- \\t]*$",
		//	                  "type": "string"
		//	                },
		//	                "maxItems": 15,
		//	                "minItems": 1,
		//	                "type": "array"
		//	              },
		//	              "RuleName": {
		//	                "maxLength": 255,
		//	                "minLength": 0,
		//	                "pattern": "^[a-zA-Z_0-9- \\t]*$",
		//	                "type": "string"
		//	              }
		//	            },
		//	            "required": [
		//	              "RuleName",
		//	              "MatchingKeys"
		//	            ],
		//	            "type": "object"
		//	          },
		//	          "maxItems": 25,
		//	          "minItems": 1,
		//	          "type": "array"
		//	        }
		//	      },
		//	      "required": [
		//	        "AttributeMatchingModel",
		//	        "RecordMatchingModel"
		//	      ],
		//	      "type": "object"
		//	    }
		//	  },
		//	  "type": "object"
		//	}
		"id_mapping_techniques": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
			Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
				// Property: IdMappingType
				"id_mapping_type": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: NormalizationVersion
				"normalization_version": schema.StringAttribute{ /*START ATTRIBUTE*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: ProviderProperties
				"provider_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: IntermediateSourceConfiguration
						"intermediate_source_configuration": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: IntermediateS3Path
								"intermediate_s3_path": schema.StringAttribute{ /*START ATTRIBUTE*/
									Description: "The s3 path that would be used to stage the intermediate data being generated during workflow execution.",
									Computed:    true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: ProviderConfiguration
						"provider_configuration": // Pattern: ""
						schema.MapAttribute{      /*START ATTRIBUTE*/
							ElementType: types.StringType,
							Description: "Additional Provider configuration that would be required for the provider service. The Configuration must be in JSON string format",
							Computed:    true,
						}, /*END ATTRIBUTE*/
						// Property: ProviderServiceArn
						"provider_service_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
							Description: "Arn of the Provider Service being used.",
							Computed:    true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
				// Property: RuleBasedProperties
				"rule_based_properties": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
					Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
						// Property: AttributeMatchingModel
						"attribute_matching_model": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: RecordMatchingModel
						"record_matching_model": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: RuleDefinitionType
						"rule_definition_type": schema.StringAttribute{ /*START ATTRIBUTE*/
							Computed: true,
						}, /*END ATTRIBUTE*/
						// Property: Rules
						"rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
							NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: MatchingKeys
									"matching_keys": schema.ListAttribute{ /*START ATTRIBUTE*/
										ElementType: types.StringType,
										Computed:    true,
									}, /*END ATTRIBUTE*/
									// Property: RuleName
									"rule_name": schema.StringAttribute{ /*START ATTRIBUTE*/
										Computed: true,
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
							}, /*END NESTED OBJECT*/
							Computed: true,
						}, /*END ATTRIBUTE*/
					}, /*END SCHEMA*/
					Computed: true,
				}, /*END ATTRIBUTE*/
			}, /*END SCHEMA*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: InputSourceConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "InputSourceARN": {
		//	        "description": "An Glue table ARN for the input source table, MatchingWorkflow arn or IdNamespace ARN",
		//	        "pattern": "^arn:(aws|aws-us-gov|aws-cn):entityresolution:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:(idnamespace/[a-zA-Z_0-9-]{1,255})$|^arn:(aws|aws-us-gov|aws-cn):entityresolution:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:(matchingworkflow/[a-zA-Z_0-9-]{1,255})$|^arn:(aws|aws-us-gov|aws-cn):glue:[a-z]{2}-[a-z]{1,10}-[0-9]:[0-9]{12}:(table/[a-zA-Z_0-9-]{1,255}/[a-zA-Z_0-9-]{1,255})$",
		//	        "type": "string"
		//	      },
		//	      "SchemaArn": {
		//	        "description": "The SchemaMapping arn associated with the Schema",
		//	        "pattern": "^arn:(aws|aws-us-gov|aws-cn):entityresolution:.*:[0-9]+:(schemamapping/.*)$",
		//	        "type": "string"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "SOURCE",
		//	          "TARGET"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "InputSourceARN"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 20,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"input_source_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: InputSourceARN
					"input_source_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "An Glue table ARN for the input source table, MatchingWorkflow arn or IdNamespace ARN",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: SchemaArn
					"schema_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The SchemaMapping arn associated with the Schema",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: OutputSourceConfig
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "KMSArn": {
		//	        "pattern": "^arn:(aws|aws-us-gov|aws-cn):kms:.*:[0-9]+:.*$",
		//	        "type": "string"
		//	      },
		//	      "OutputS3Path": {
		//	        "description": "The S3 path to which Entity Resolution will write the output table",
		//	        "pattern": "^s3://([^/]+)/?(.*?([^/]+)/?)$",
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "OutputS3Path"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"output_source_config": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: KMSArn
					"kms_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: OutputS3Path
					"output_s3_path": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The S3 path to which Entity Resolution will write the output table",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "pattern": "^arn:(aws|aws-us-gov|aws-cn):iam::\\d{12}:role/?[a-zA-Z_0-9+=,.@\\-_/]+$",
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "A key-value pair to associate with a resource",
		//	    "properties": {
		//	      "Key": {
		//	        "description": "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "description": "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
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
		//	  "maxItems": 200,
		//	  "minItems": 0,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"tags": schema.SetNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The key name of the tag. You can specify a value that is 1 to 128 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Description: "The value for the tag. You can specify a value that is 0 to 256 Unicode characters in length and cannot be prefixed with aws:. You can use any of the following characters: the set of Unicode letters, digits, whitespace, _, ., /, =, +, and -.",
						Computed:    true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: UpdatedAt
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The time of this IdMappingWorkflow got last updated at",
		//	  "type": "string"
		//	}
		"updated_at": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The time of this IdMappingWorkflow got last updated at",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WorkflowArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The default IdMappingWorkflow arn",
		//	  "pattern": "^arn:(aws|aws-us-gov|aws-cn):entityresolution:.*:[0-9]+:(idmappingworkflow/.*)$",
		//	  "type": "string"
		//	}
		"workflow_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The default IdMappingWorkflow arn",
			Computed:    true,
		}, /*END ATTRIBUTE*/
		// Property: WorkflowName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The name of the IdMappingWorkflow",
		//	  "maxLength": 255,
		//	  "minLength": 0,
		//	  "pattern": "^[a-zA-Z_0-9-]*$",
		//	  "type": "string"
		//	}
		"workflow_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The name of the IdMappingWorkflow",
			Computed:    true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::EntityResolution::IdMappingWorkflow",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::EntityResolution::IdMappingWorkflow").WithTerraformTypeName("awscc_entityresolution_id_mapping_workflow")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"attribute_matching_model":          "AttributeMatchingModel",
		"created_at":                        "CreatedAt",
		"description":                       "Description",
		"id_mapping_techniques":             "IdMappingTechniques",
		"id_mapping_type":                   "IdMappingType",
		"input_source_arn":                  "InputSourceARN",
		"input_source_config":               "InputSourceConfig",
		"intermediate_s3_path":              "IntermediateS3Path",
		"intermediate_source_configuration": "IntermediateSourceConfiguration",
		"key":                               "Key",
		"kms_arn":                           "KMSArn",
		"matching_keys":                     "MatchingKeys",
		"normalization_version":             "NormalizationVersion",
		"output_s3_path":                    "OutputS3Path",
		"output_source_config":              "OutputSourceConfig",
		"provider_configuration":            "ProviderConfiguration",
		"provider_properties":               "ProviderProperties",
		"provider_service_arn":              "ProviderServiceArn",
		"record_matching_model":             "RecordMatchingModel",
		"role_arn":                          "RoleArn",
		"rule_based_properties":             "RuleBasedProperties",
		"rule_definition_type":              "RuleDefinitionType",
		"rule_name":                         "RuleName",
		"rules":                             "Rules",
		"schema_arn":                        "SchemaArn",
		"tags":                              "Tags",
		"type":                              "Type",
		"updated_at":                        "UpdatedAt",
		"value":                             "Value",
		"workflow_arn":                      "WorkflowArn",
		"workflow_name":                     "WorkflowName",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
