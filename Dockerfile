FROM ethereum/client-go:alltools-latest


WORKDIR workdir
COPY . workdir

RUN apk add --update nodejs nodejs-npm make wget
RUN cd workdir && rm -rf node_modules && npm i
RUN npm i -g @waves/surfboard@beta

RUN cp -r /usr/lib/node_modules/ ./workdir

RUN wget https://github.com/ethereum/solidity/releases/download/v0.6.10/solc-static-linux
RUN mv solc-static-linux solc
RUN chmod +x solc
RUN cp solc /bin/solc
RUN echo "export PATH=/bin/solc:${PATH}" >> /root/.bashrc

ENTRYPOINT cd workdir/
