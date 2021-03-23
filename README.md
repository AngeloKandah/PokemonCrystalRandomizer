# PokemonCrystalRandomizer
This Randomizer will read your ROM and savefile and tell you all the stats of your party pokemon, including their name, nickname, IVs, held item, moves, and more.
Currently this is a work in progress, it only outputs party data and does not write anything to memory.

# How it works
This works by initially scanning your save for your player name, once it has that, it will be able to find in memory where this process is being run.
Then we utilizes knowing the trainer name to grab data for party pokemon and their stats, which we can then use to overwrite with other data, such as changing their moves or held item.

# Could this lead to corrupt saves?
Potentially if there is incorrect data being written in certain areas of memory, however this can be prevented with error handling.
