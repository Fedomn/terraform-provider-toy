## A toy to dive into terraform

### how to debug terragrunt

```shell script
dlv exec /usr/local/bin/terragrunt plan

terragrunt plan -out=./state.json

terragrunt show ./state.json
```

### how to debug provider

* [terrafrom go-plugin-architecture](https://www.terraform.io/docs/internals/internal-plugins.html#go-plugin-architecture)
* [terraform-plugin-sdk issue](https://github.com/hashicorp/terraform-plugin-sdk/issues/88)
* [go-delve issue](https://github.com/go-delve/delve/issues/496)
* [terraform dlv-debugger](https://runsisi.com/2019-04-10/dlv-debugger)

### delve

* [](https://www.jamessturtevant.com/posts/Using-the-Go-Delve-Debugger-from-the-command-line/)
* [](https://golangbot.com/debugging-go-delve/)


### reference
* [gob](https://blog.golang.org/gob)
* [grpc](https://grpc.io/docs/languages/go/basics/)
* [How Terrafrom Wroks](https://www.terraform.io/docs/extend/how-terraform-works.html)
* [Terrafrom providers](https://www.terraform.io/docs/extend/writing-custom-providers.html)
* [Terrafrom bundle](https://github.com/hashicorp/terraform/tree/master/tools/terraform-bundle)
