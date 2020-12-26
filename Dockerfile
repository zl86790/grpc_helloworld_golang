FROM golang as golang   

RUN mkdir -p /root/grpc
COPY ./ /root/grpc

WORKDIR /root/grpc/output
RUN go build /root/grpc/server.go
RUN go build /root/grpc/client.go

FROM ubuntu as ubuntu

RUN mkdir -p /root/grpc
COPY --from=golang --chown=root:root /root/grpc /root/grpc
COPY --from=golang --chown=root:root /root/grpc/output /root/grpc

CMD /bin/bash
