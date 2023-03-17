# Client Go

## Generate password database in Linux

```commandline
gpg --generate-key

pass init [GPG_KEY]
```

## Create the Docker Image

```commandline
docker build -t k8s-lister-go:0.1.0 .
```