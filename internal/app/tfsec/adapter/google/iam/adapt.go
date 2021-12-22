package iam

import (
	"github.com/aquasecurity/defsec/provider/google/iam"
	"github.com/aquasecurity/tfsec/internal/app/tfsec/block"
)

func Adapt(modules []block.Module) iam.IAM {
	return iam.IAM{}
}