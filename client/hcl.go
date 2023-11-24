package client

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclsyntax"
)

func Decode(filename string, src []byte, target interface{}) error {
	var file *hcl.File
	var diags hcl.Diagnostics

	switch suffix := strings.ToLower(filepath.Ext(filename)); suffix {
	case ".or":
		file, diags = hclsyntax.ParseConfig(src, filename, hcl.Pos{Line: 1, Column: 1})
	default:
		diags = diags.Append(&hcl.Diagnostic{
			Severity: hcl.DiagError,
			Summary:  "Unsupported file format",
			Detail:   fmt.Sprintf("Cannot read from %s: unrecognized file format suffix %q. It must be %q.", filename, suffix, ".or"),
		})
		return diags
	}
	if diags.HasErrors() {
		return diags
	}

	diags = gohcl.DecodeBody(file.Body, nil, target)
	if diags.HasErrors() {
		return diags
	}
	return nil
}

func DecodeFile(filename string, target interface{}) error {
	src, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return hcl.Diagnostics{
				{
					Severity: hcl.DiagError,
					Summary:  "Configuration file not found",
					Detail:   fmt.Sprintf("The configuration file %s does not exist.", filename),
				},
			}
		}
		return hcl.Diagnostics{
			{
				Severity: hcl.DiagError,
				Summary:  "Failed to read configuration",
				Detail:   fmt.Sprintf("Can't read %s: %s.", filename, err),
			},
		}
	}

	return Decode(filename, src, target)
}
