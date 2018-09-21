
tf_test:
	go fmt
	terraform fmt
	go build -o terraform-provider-onix && terraform init && terraform apply

test:
	go fmt
	go test -v

