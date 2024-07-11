// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package ses

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceFactory("awscc_ses_mail_manager_traffic_policy", mailManagerTrafficPolicyDataSource)
}

// mailManagerTrafficPolicyDataSource returns the Terraform awscc_ses_mail_manager_traffic_policy data source.
// This Terraform data source corresponds to the CloudFormation AWS::SES::MailManagerTrafficPolicy resource.
func mailManagerTrafficPolicyDataSource(ctx context.Context) (datasource.DataSource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: DefaultAction
		// CloudFormation resource type schema:
		//
		//	{
		//	  "enum": [
		//	    "ALLOW",
		//	    "DENY"
		//	  ],
		//	  "type": "string"
		//	}
		"default_action": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: MaxMessageSizeBytes
		// CloudFormation resource type schema:
		//
		//	{
		//	  "minimum": 1,
		//	  "type": "number"
		//	}
		"max_message_size_bytes": schema.Float64Attribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: PolicyStatements
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Action": {
		//	        "enum": [
		//	          "ALLOW",
		//	          "DENY"
		//	        ],
		//	        "type": "string"
		//	      },
		//	      "Conditions": {
		//	        "items": {
		//	          "properties": {
		//	            "BooleanExpression": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Evaluate": {
		//	                  "properties": {
		//	                    "Analysis": {
		//	                      "additionalProperties": false,
		//	                      "properties": {
		//	                        "Analyzer": {
		//	                          "pattern": "^[a-zA-Z0-9:_/+=,@.#-]+$",
		//	                          "type": "string"
		//	                        },
		//	                        "ResultField": {
		//	                          "maxLength": 256,
		//	                          "minLength": 1,
		//	                          "pattern": "^[\\sa-zA-Z0-9_]+$",
		//	                          "type": "string"
		//	                        }
		//	                      },
		//	                      "required": [
		//	                        "Analyzer",
		//	                        "ResultField"
		//	                      ],
		//	                      "type": "object"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "Operator": {
		//	                  "enum": [
		//	                    "IS_TRUE",
		//	                    "IS_FALSE"
		//	                  ],
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Evaluate",
		//	                "Operator"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "IpExpression": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Evaluate": {
		//	                  "properties": {
		//	                    "Attribute": {
		//	                      "enum": [
		//	                        "SENDER_IP"
		//	                      ],
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "Operator": {
		//	                  "enum": [
		//	                    "CIDR_MATCHES",
		//	                    "NOT_CIDR_MATCHES"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Values": {
		//	                  "items": {
		//	                    "pattern": "^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)/([0-9]|[12][0-9]|3[0-2])$",
		//	                    "type": "string"
		//	                  },
		//	                  "type": "array"
		//	                }
		//	              },
		//	              "required": [
		//	                "Evaluate",
		//	                "Operator",
		//	                "Values"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "StringExpression": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Evaluate": {
		//	                  "properties": {
		//	                    "Attribute": {
		//	                      "enum": [
		//	                        "RECIPIENT"
		//	                      ],
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "Operator": {
		//	                  "enum": [
		//	                    "EQUALS",
		//	                    "NOT_EQUALS",
		//	                    "STARTS_WITH",
		//	                    "ENDS_WITH",
		//	                    "CONTAINS"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Values": {
		//	                  "items": {
		//	                    "type": "string"
		//	                  },
		//	                  "type": "array"
		//	                }
		//	              },
		//	              "required": [
		//	                "Evaluate",
		//	                "Operator",
		//	                "Values"
		//	              ],
		//	              "type": "object"
		//	            },
		//	            "TlsExpression": {
		//	              "additionalProperties": false,
		//	              "properties": {
		//	                "Evaluate": {
		//	                  "properties": {
		//	                    "Attribute": {
		//	                      "enum": [
		//	                        "TLS_PROTOCOL"
		//	                      ],
		//	                      "type": "string"
		//	                    }
		//	                  },
		//	                  "type": "object"
		//	                },
		//	                "Operator": {
		//	                  "enum": [
		//	                    "MINIMUM_TLS_VERSION",
		//	                    "IS"
		//	                  ],
		//	                  "type": "string"
		//	                },
		//	                "Value": {
		//	                  "enum": [
		//	                    "TLS1_2",
		//	                    "TLS1_3"
		//	                  ],
		//	                  "type": "string"
		//	                }
		//	              },
		//	              "required": [
		//	                "Evaluate",
		//	                "Operator",
		//	                "Value"
		//	              ],
		//	              "type": "object"
		//	            }
		//	          },
		//	          "type": "object"
		//	        },
		//	        "minItems": 1,
		//	        "type": "array"
		//	      }
		//	    },
		//	    "required": [
		//	      "Action",
		//	      "Conditions"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"policy_statements": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Action
					"action": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Conditions
					"conditions": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
						NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
							Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
								// Property: BooleanExpression
								"boolean_expression": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Evaluate
										"evaluate": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Analysis
												"analysis": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
													Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
														// Property: Analyzer
														"analyzer": schema.StringAttribute{ /*START ATTRIBUTE*/
															Computed: true,
														}, /*END ATTRIBUTE*/
														// Property: ResultField
														"result_field": schema.StringAttribute{ /*START ATTRIBUTE*/
															Computed: true,
														}, /*END ATTRIBUTE*/
													}, /*END SCHEMA*/
													Computed: true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Operator
										"operator": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: IpExpression
								"ip_expression": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Evaluate
										"evaluate": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Attribute
												"attribute": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Operator
										"operator": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Values
										"values": schema.ListAttribute{ /*START ATTRIBUTE*/
											ElementType: types.StringType,
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: StringExpression
								"string_expression": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Evaluate
										"evaluate": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Attribute
												"attribute": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Operator
										"operator": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Values
										"values": schema.ListAttribute{ /*START ATTRIBUTE*/
											ElementType: types.StringType,
											Computed:    true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
								// Property: TlsExpression
								"tls_expression": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
									Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
										// Property: Evaluate
										"evaluate": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
											Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
												// Property: Attribute
												"attribute": schema.StringAttribute{ /*START ATTRIBUTE*/
													Computed: true,
												}, /*END ATTRIBUTE*/
											}, /*END SCHEMA*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Operator
										"operator": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
										// Property: Value
										"value": schema.StringAttribute{ /*START ATTRIBUTE*/
											Computed: true,
										}, /*END ATTRIBUTE*/
									}, /*END SCHEMA*/
									Computed: true,
								}, /*END ATTRIBUTE*/
							}, /*END SCHEMA*/
						}, /*END NESTED OBJECT*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "pattern": "^[a-zA-Z0-9/_\\+=\\.:@\\-]+$",
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 0,
		//	        "pattern": "^[a-zA-Z0-9/_\\+=\\.:@\\-]*$",
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
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Computed: true,
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TrafficPolicyArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "type": "string"
		//	}
		"traffic_policy_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TrafficPolicyId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "minLength": 1,
		//	  "type": "string"
		//	}
		"traffic_policy_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
		// Property: TrafficPolicyName
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 63,
		//	  "minLength": 3,
		//	  "pattern": "^[A-Za-z0-9_\\-]+$",
		//	  "type": "string"
		//	}
		"traffic_policy_name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Required:    true,
	}

	schema := schema.Schema{
		Description: "Data Source schema for AWS::SES::MailManagerTrafficPolicy",
		Attributes:  attributes,
	}

	var opts generic.DataSourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SES::MailManagerTrafficPolicy").WithTerraformTypeName("awscc_ses_mail_manager_traffic_policy")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"action":                 "Action",
		"analysis":               "Analysis",
		"analyzer":               "Analyzer",
		"attribute":              "Attribute",
		"boolean_expression":     "BooleanExpression",
		"conditions":             "Conditions",
		"default_action":         "DefaultAction",
		"evaluate":               "Evaluate",
		"ip_expression":          "IpExpression",
		"key":                    "Key",
		"max_message_size_bytes": "MaxMessageSizeBytes",
		"operator":               "Operator",
		"policy_statements":      "PolicyStatements",
		"result_field":           "ResultField",
		"string_expression":      "StringExpression",
		"tags":                   "Tags",
		"tls_expression":         "TlsExpression",
		"traffic_policy_arn":     "TrafficPolicyArn",
		"traffic_policy_id":      "TrafficPolicyId",
		"traffic_policy_name":    "TrafficPolicyName",
		"value":                  "Value",
		"values":                 "Values",
	})

	v, err := generic.NewSingularDataSource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
