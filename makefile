default: build

build: fmtcheck
	go build -o examples/terraform/terraform-provider-toy
	go build -o examples/terragrunt/terraform-provider-toy

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

clean:
	find . -type d -name ".terragrunt-cache" -prune -exec rm -rf {} \;
	find . -type d -name ".terraform" -prune -exec rm -rf {} \;
	find . -name "terraform.tfstate" -prune -exec rm -rf {} \;
	find . -name "terraform.tfstate.backup" -prune -exec rm -rf {} \;
	find . -name "terraform-provider-toy" -prune -exec rm -rf {} \;
