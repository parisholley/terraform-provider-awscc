// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package generic

import (
	"context"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

// copyStateValueAtPath copies the value at a specified path from source State to destination State.
func copyStateValueAtPath(ctx context.Context, dst, src *tfsdk.State, p path.Path) diag.Diagnostics {
	var diags diag.Diagnostics
	var val attr.Value

	diags.Append(src.GetAttribute(ctx, p, &val)...)
	if diags.HasError() {
		return diags
	}

	tflog.Debug(ctx, "copyStateValueAtPath", map[string]interface{}{
		"value": hclog.Fmt("%v = %v", p, val),
	})

	diags.Append(dst.SetAttribute(ctx, p, val)...)
	if diags.HasError() {
		return diags
	}

	return diags
}
