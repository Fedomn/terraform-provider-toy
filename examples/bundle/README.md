## install terraform-bundle

```sh
cd terraform

git checkout v0.12.26

go install ./tools/terraform-bundle 
```

## package commands
```sh
terraform-bundle package terraform-bundle.hcl

terraform-bundle package -os=linux -arch=amd64 terraform-bundle.hcl
```
