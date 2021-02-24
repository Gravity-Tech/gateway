const { seedUtils } = require("@waves/waves-transactions");
const { Seed } = seedUtils;

type DeploymentSeeds<V> = {
  Distributor: V;
  Gravity: V;
  Nebula: V;
  Port: V;
};

type AccountData = {
  address: string;
  seed: string;
  publicKey: string;
  privateKey: string;
}

function generateSeeds(chain_id): DeploymentSeeds<AccountData> {
  const keys = "Distributor,Gravity,Nebula,Port".split(",").reduce((acc, key) => {
    const chainId = chain_id.charCodeAt(0);
    const instance = new Seed(Seed.create().phrase, chainId);
    let { address, phrase: seed, keyPair } = instance;
    const { publicKey, privateKey } = keyPair;

    return {
      ...acc,
      [key]: { address, seed, publicKey, privateKey },
    };
  }, {} as DeploymentSeeds<AccountData>);

  return keys
  // return {
  //   ...keys,
  //   Seeds: {
  //     DistributionSeed: keys.Distributor.seed,
  //     GravityContractSeed: keys.Gravity.seed,
  //     NebulaContractSeed: keys.Nebula.seed,
  //     SubscriberContractSeed: keys.Port.seed,
  //   },
  // };
}
