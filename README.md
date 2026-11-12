# Terraform AD provider · OU DACL fix

This fork is based on [isometry/terraform-provider-ad v0.6.0](https://github.com/isometry/terraform-provider-ad/tree/v0.6.0) by Robin Breathe and contributors. The upstream documentation is available at the source repository. The original [MIT license](LICENSE) applies.

The LDAP read used to refresh an OU now requests the DACL explicitly via the AD security descriptor flags control. This lets Terraform read the accidental deletion protection it has configured. The regression tests cover the LDAP control and the detection of the protection ACEs.

## Install

Download the Linux AMD64 ZIP and its checksum from [release v0.6.0-ou-dacl.1](https://github.com/AnatolyED/terraform-provider-ad-ou-dacl/releases/tag/v0.6.0-ou-dacl.1).

```sh
sha256sum -c SHA256SUMS
```

Place the ZIP under `<mirror>/registry.terraform.io/isometry/ad/` and configure a Terraform filesystem mirror for `isometry/ad`. Pin `0.6.0-ou-dacl.1` and use the dependency lock file from the consuming project.

## Build

```sh
go test ./internal/ldap ./internal/provider -run 'Test.*(OU|DACL|SecurityDescriptor)' -count=1
version=0.6.0-ou-dacl.1
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath \
  -ldflags "-X main.version=$version" -o "terraform-provider-ad_v$version" .
chmod 0755 "terraform-provider-ad_v$version"
zip -X "terraform-provider-ad_${version}_linux_amd64.zip" "terraform-provider-ad_v$version"
```

The release ZIP is the package pinned by Terraform_AD. A rebuilt ZIP can have a different checksum.
