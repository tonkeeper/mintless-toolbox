

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

Currently, there is a single command available:

```bash
./bin/mintless-cli dump <airdrop-filename> 
```

dump reads an airdrop file and dump it to the console in the format 
`address,amount,start_from,expire_at`
