terraform {
  source = "${get_terragrunt_dir()}/.//server"
}

inputs = {
  address = "1.2.3.1"
}