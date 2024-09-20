
# Mintless Toolbox

Mintless Toolbox is a set of tools to work with the new Mintless Jetton standard in TON.

# Installation

```bash
git clone https://github.com/tonkeeper/mintless-toolbox.git
cd mintless-toolbox
make
```

make will compile the code and put the binary in the `bin` folder.

# Usage

```bash
./bin/mintless-cli --help
```

Currently, there are two commands available:

```bash
./bin/mintless-cli dump <airdrop-filename> 
```

dump reads an airdrop file and dumps it to the console in the format 
`wallet address,amount,start_from,expire_at`

```bash
./bin/mintless-cli hash <airdrop-filename> 
```

hash reads an airdrop file and prints its root's hash.
A mintless Jetton Master has to return the same hash with
`get_mintless_airdrop_hashmap_root`.


# Links

The [reference implementation](https://github.com/ton-community/mintless-jetton/tree/main) of the Mintless Jetton.