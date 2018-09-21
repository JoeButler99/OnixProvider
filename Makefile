
tf_test:
	go fmt
	go build -o terraform-provider-onix && terraform init && terraform apply

test:
	go test
