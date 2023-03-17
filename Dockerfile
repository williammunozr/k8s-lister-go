FROM ubuntu

COPY ./k8s-lister-go ./k8s-lister-go

ENTRYPOINT ["./k8s-lister-go"]
