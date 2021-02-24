FROM ethereum/client-go:alltools-latest


WORKDIR workdir
COPY . workdir

RUN apk add --update nodejs nodejs-npm make wget
RUN cd workdir && npm i
RUN npm i -g @waves/surfboard@beta solc@0.6.1

RUN wget https://github.com/ethereum/solidity/releases/download/v0.6.11/solc-static-linux
RUN mv solc-static-linux solc
RUN chmod +x solc
RUN cp solc /bin/solc
RUN echo "export PATH=/bin/solc:${PATH}" >> /root/.bashrc

ENTRYPOINT cd workdir/
