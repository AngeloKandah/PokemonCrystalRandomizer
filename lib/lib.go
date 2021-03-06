package lib

var CharacterTable = [16][16]string{
	//  0   1   2   3   4   5   6   7   8   9   A   B   C   D   E   F
	/*0*/ {" ", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
	/*1*/ {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
	/*2*/ {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
	/*3*/ {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
	/*4*/ {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
	/*5*/ {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
	/*6*/ {"", "", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
	/*7*/ {"PO", "Ké", "", "", "", "", "", "", "", "", "", "", "", "", "", ""},
	/*8*/ {"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"},
	/*9*/ {"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "(", ")", ":", ";", "[", "]"},
	/*A*/ {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p"},
	/*B*/ {"q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "", "", "", "", "", ""},
	/*C*/ {"Ä", "Ö", "Ü", "ä", "ö", "ü", "", "", "", "", "", "", "", "", "", ""},
	/*D*/ {"'d", "'l", "'m", "'r", "'s", "'t", "'v", "", "", "", "", "", "", "", "", ""},
	/*E*/ {"'", "PK", "MN", "-", "", "", "?", "!", ".", "&", "é", "→", "▷", "▶", "▼", "♂"},
	/*F*/ {"", "×", ".", "/", ",", "♀", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"},
}

var PokemonSpecies = map[uint8]string{
	0x01: "Bulbasaur",
	0x02: "Ivysaur",
	0x03: "Venusaur",
	0x04: "Charmander",
	0x05: "Charmeleon",
	0x06: "Charizard",
	0x07: "Squirtle",
	0x08: "Wartortle",
	0x09: "Blastoise",
	0x0A: "Caterpie",
	0x0B: "Metapod",
	0x0C: "Butterfree",
	0x0D: "Weedle",
	0x0E: "Kakuna",
	0x0F: "Beedrill",
	0x10: "Pidgey",
	0x11: "Pidgeotto",
	0x12: "Pidgeot",
	0x13: "Rattata",
	0x14: "Raticate",
	0x15: "Spearow",
	0x16: "Fearow",
	0x17: "Ekans",
	0x18: "Arbok",
	0x19: "Pikachu",
	0x1A: "Raichu",
	0x1B: "Sandshrew",
	0x1C: "Sandslash",
	0x1D: "NidoranF",
	0x1E: "Nidorina",
	0x1F: "Nidoqueen",
	0x20: "NidoranM",
	0x21: "Nidorino",
	0x22: "Nidoking",
	0x23: "Clefairy",
	0x24: "Clefable",
	0x25: "Vulpix",
	0x26: "Ninetales",
	0x27: "Jigglypuff",
	0x28: "Wigglytuff",
	0x29: "Zubat",
	0x2A: "Golbat",
	0x2B: "Oddish",
	0x2C: "Gloom",
	0x2D: "Vileplume",
	0x2E: "Paras",
	0x2F: "Parasect",
	0x30: "Venonat",
	0x31: "Venomoth",
	0x32: "Diglett",
	0x33: "Dugtrio",
	0x34: "Meowth",
	0x35: "Persian",
	0x36: "Psyduck",
	0x37: "Golduck",
	0x38: "Mankey",
	0x39: "Primeape",
	0x3A: "Growlithe",
	0x3B: "Arcanine",
	0x3C: "Poliwag",
	0x3D: "Poliwhirl",
	0x3E: "Poliwrath",
	0x3F: "Abra",
	0x40: "Kadabra",
	0x41: "Alakazam",
	0x42: "Machop",
	0x43: "Machoke",
	0x44: "Machamp",
	0x45: "Bellsprout",
	0x46: "Weepinbell",
	0x47: "Victreebel",
	0x48: "Tentacool",
	0x49: "Tentacruel",
	0x4A: "Geodude",
	0x4B: "Graveler",
	0x4C: "Golem",
	0x4D: "Ponyta",
	0x4E: "Rapidash",
	0x4F: "Slowpoke",
	0x50: "Slowbro",
	0x51: "Magnemite",
	0x52: "Magneton",
	0x53: "Farfetch'd",
	0x54: "Doduo",
	0x55: "Dodrio",
	0x56: "Seel",
	0x57: "Dewgong",
	0x58: "Grimer",
	0x59: "Muk",
	0x5A: "Shellder",
	0x5B: "Cloyster",
	0x5C: "Gastly",
	0x5D: "Haunter",
	0x5E: "Gengar",
	0x5F: "Onix",
	0x60: "Drowzee",
	0x61: "Hypno",
	0x62: "Krabby",
	0x63: "Kingler",
	0x64: "Voltorb",
	0x65: "Electrode",
	0x66: "Exeggcute",
	0x67: "Exeggutor",
	0x68: "Cubone",
	0x69: "Marowak",
	0x6A: "Hitmonlee",
	0x6B: "Hitmonchan",
	0x6C: "Lickitung",
	0x6D: "Koffing",
	0x6E: "Weezing",
	0x6F: "Rhyhorn",
	0x70: "Rhydon",
	0x71: "Chansey",
	0x72: "Tangela",
	0x73: "Kangaskhan",
	0x74: "Horsea",
	0x75: "Seadra",
	0x76: "Goldeen",
	0x77: "Seaking",
	0x78: "Staryu",
	0x79: "Starmie",
	0x7A: "Mr.Mime",
	0x7B: "Scyther",
	0x7C: "Jynx",
	0x7D: "Electabuzz",
	0x7E: "Magmar",
	0x7F: "Pinsir",
	0x80: "Tauros",
	0x81: "Magikarp",
	0x82: "Gyarados",
	0x83: "Lapras",
	0x84: "Ditto",
	0x85: "Eevee",
	0x86: "Vaporeon",
	0x87: "Jolteon",
	0x88: "Flareon",
	0x89: "Porygon",
	0x8A: "Omanyte",
	0x8B: "Omastar",
	0x8C: "Kabuto",
	0x8D: "Kabutops",
	0x8E: "Aerodactyl",
	0x8F: "Snorlax",
	0x90: "Articuno",
	0x91: "Zapdos",
	0x92: "Moltres",
	0x93: "Dratini",
	0x94: "Dragonair",
	0x95: "Dragonite",
	0x96: "Mewtwo",
	0x97: "Mew",
	0x98: "Chikorita",
	0x99: "Bayleef",
	0x9A: "Meganium",
	0x9B: "Cyndaquil",
	0x9C: "Quilava",
	0x9D: "Typhlosion",
	0x9E: "Totodile",
	0x9F: "Croconaw",
	0xA0: "Feraligatr",
	0xA1: "Sentret",
	0xA2: "Furret",
	0xA3: "Hoothoot",
	0xA4: "Noctowl",
	0xA5: "Ledyba",
	0xA6: "Ledian",
	0xA7: "Spinarak",
	0xA8: "Ariados",
	0xA9: "Crobat",
	0xAA: "Chinchou",
	0xAB: "Lanturn",
	0xAC: "Pichu",
	0xAD: "Cleffa",
	0xAE: "Igglybuff",
	0xAF: "Togepi",
	0xB0: "Togetic",
	0xB1: "Natu",
	0xB2: "Xatu",
	0xB3: "Mareep",
	0xB4: "Flaaffy",
	0xB5: "Ampharos",
	0xB6: "Bellossom",
	0xB7: "Marill",
	0xB8: "Azumarill",
	0xB9: "Sudowoodo",
	0xBA: "Politoed",
	0xBB: "Hoppip",
	0xBC: "Skiploom",
	0xBD: "Jumpluff",
	0xBE: "Aipom",
	0xBF: "Sunkern",
	0xC0: "Sunflora",
	0xC1: "Yanma",
	0xC2: "Wooper",
	0xC3: "Quagsire",
	0xC4: "Espeon",
	0xC5: "Umbreon",
	0xC6: "Murkrow",
	0xC7: "Slowking",
	0xC8: "Misdreavus",
	0xC9: "Unown",
	0xCA: "Wobbuffet",
	0xCB: "Girafarig",
	0xCC: "Pineco",
	0xCD: "Forretress",
	0xCE: "Dunsparce",
	0xCF: "Gligar",
	0xD0: "Steelix",
	0xD1: "Snubbull",
	0xD2: "Granbull",
	0xD3: "Qwilfish",
	0xD4: "Scizor",
	0xD5: "Shuckle",
	0xD6: "Heracross",
	0xD7: "Sneasel",
	0xD8: "Teddiursa",
	0xD9: "Ursaring",
	0xDA: "Slugma",
	0xDB: "Magcargo",
	0xDC: "Swinub",
	0xDD: "Piloswine",
	0xDE: "Corsola",
	0xDF: "Remoraid",
	0xE0: "Octillery",
	0xE1: "Delibird",
	0xE2: "Mantine",
	0xE3: "Skarmory",
	0xE4: "Houndour",
	0xE5: "Houndoom",
	0xE6: "Kingdra",
	0xE7: "Phanpy",
	0xE8: "Donphan",
	0xE9: "Porygon2",
	0xEA: "Stantler",
	0xEB: "Smeargle",
	0xEC: "Tyrogue",
	0xED: "Hitmontop",
	0xEE: "Smoochum",
	0xEF: "Elekid",
	0xF0: "Magby",
	0xF1: "Miltank",
	0xF2: "Blissey",
	0xF3: "Raikou",
	0xF4: "Entei",
	0xF5: "Suicune",
	0xF6: "Larvitar",
	0xF7: "Pupitar",
	0xF8: "Tyranitar",
	0xF9: "Lugia",
	0xFA: "Ho-oh",
	0xFB: "Celebi",
}

var Moves = map[uint8]string{
	0x00: "x",
	0x01: "Pound",
	0x02: "Karate Chop",
	0x03: "Doubleslap",
	0x04: "Comet Punch",
	0x05: "Mega Punch",
	0x06: "Pay Day",
	0x07: "Fire Punch",
	0x08: "Ice Punch",
	0x09: "Thunderpunch",
	0x0A: "Scratch",
	0x0B: "Vicegrip",
	0x0C: "Guillotine",
	0x0D: "Razor Wind",
	0x0E: "Swords Dance",
	0x0F: "Cut",
	0x10: "Gust",
	0x11: "Wing Attack",
	0x12: "Whirlwind",
	0x13: "Fly",
	0x14: "Bind",
	0x15: "Slam",
	0x16: "Vine Whip",
	0x17: "Stomp",
	0x18: "Double Kick",
	0x19: "Mega Kick",
	0x1A: "Jump Kick",
	0x1B: "Rolling Kick",
	0x1C: "Sand-Attack",
	0x1D: "Headbutt",
	0x1E: "Horn Attack",
	0x1F: "Fury Attack",
	0x20: "Horn Drill",
	0x21: "Tackle",
	0x22: "Body Slam",
	0x23: "Wrap",
	0x24: "Take Down",
	0x25: "Thrash",
	0x26: "Double-Edge",
	0x27: "Tail Whip",
	0x28: "Poison Sting",
	0x29: "Twineedle",
	0x2A: "Pin Missile",
	0x2B: "Leer",
	0x2C: "Bite",
	0x2D: "Growl",
	0x2E: "Roar",
	0x2F: "Sing",
	0x30: "Supersonic",
	0x31: "Sonicboom",
	0x32: "Disable",
	0x33: "Acid",
	0x34: "Ember",
	0x35: "Flamethrower",
	0x36: "Mist",
	0x37: "Water Gun",
	0x38: "Hydro Pump",
	0x39: "Surf",
	0x3A: "Ice Beam",
	0x3B: "Blizzard",
	0x3C: "Psybeam",
	0x3D: "Bubblebeam",
	0x3E: "Aurora Beam",
	0x3F: "Hyper Beam",
	0x40: "Peck",
	0x41: "Drill Peck",
	0x42: "Submission",
	0x43: "Low Kick",
	0x44: "Counter",
	0x45: "Seismic Toss",
	0x46: "Strength",
	0x47: "Absorb",
	0x48: "Mega Drain",
	0x49: "Leech Seed",
	0x4A: "Growth",
	0x4B: "Razor Leaf",
	0x4C: "Solarbeam",
	0x4D: "Poisonpowder",
	0x4E: "Stun Spore",
	0x4F: "Sleep Powder",
	0x50: "Petal Dance",
	0x51: "String Shot",
	0x52: "Dragon Rage",
	0x53: "Fire Spin",
	0x54: "Thundershock",
	0x55: "Thunderbolt",
	0x56: "Thunder Wave",
	0x57: "Thunder",
	0x58: "Rock Throw",
	0x59: "Earthquake",
	0x5A: "Fissure",
	0x5B: "Dig",
	0x5C: "Toxic",
	0x5D: "Confusion",
	0x5E: "Psychic",
	0x5F: "Hypnosis",
	0x60: "Meditate",
	0x61: "Agility",
	0x62: "Quick Attack",
	0x63: "Rage",
	0x64: "Teleport",
	0x65: "Night Shade",
	0x66: "Mimic",
	0x67: "Screech",
	0x68: "Double Team",
	0x69: "Recover",
	0x6A: "Harden",
	0x6B: "Minimize",
	0x6C: "Smokescreen",
	0x6D: "Confuse Ray",
	0x6E: "Withdraw",
	0x6F: "Defense Curl",
	0x70: "Barrier",
	0x71: "Light Screen",
	0x72: "Haze",
	0x73: "Reflect",
	0x74: "Focus Energy",
	0x75: "Bide",
	0x76: "Metronome",
	0x77: "Mirror Move",
	0x78: "Selfdestruct",
	0x79: "Egg Bomb",
	0x7A: "Lick",
	0x7B: "Smog",
	0x7C: "Sludge",
	0x7D: "Bone Club",
	0x7E: "Fire Blast",
	0x7F: "Waterfall",
	0x80: "Clamp",
	0x81: "Swift",
	0x82: "Skull Bash",
	0x83: "Spike Cannon",
	0x84: "Constrict",
	0x85: "Amnesia",
	0x86: "Kinesis",
	0x87: "Softboiled",
	0x88: "Hi-Jump Kick",
	0x89: "Glare",
	0x8A: "Dream Eater",
	0x8B: "Poison Gas",
	0x8C: "Barrage",
	0x8D: "Leech Life",
	0x8E: "Lovely Kiss",
	0x8F: "Sky Attack",
	0x90: "Transform",
	0x91: "Bubble",
	0x92: "Dizzy Punch",
	0x93: "Spore",
	0x94: "Flash",
	0x95: "Psywave",
	0x96: "Splash",
	0x97: "Acid Armor",
	0x98: "Crabhammer",
	0x99: "Explosion",
	0x9A: "Fury Swipes",
	0x9B: "Bonemerang",
	0x9C: "Rest",
	0x9D: "Rock Slide",
	0x9E: "Hyper Fang",
	0x9F: "Sharpen",
	0xA0: "Conversion",
	0xA1: "Tri Attack",
	0xA2: "Super Fang",
	0xA3: "Slash",
	0xA4: "Substitute",
	0xA5: "Struggle",
	0xA6: "Sketch",
	0xA7: "Triple Kick",
	0xA8: "Thief",
	0xA9: "Spider Web",
	0xAA: "Mind Reader",
	0xAB: "Nightmare",
	0xAC: "Flame Wheel",
	0xAD: "Snore",
	0xAE: "Curse",
	0xAF: "Flail",
	0xB0: "Conversion2",
	0xB1: "Aeroblast",
	0xB2: "Cotton Spore",
	0xB3: "Reversal",
	0xB4: "Spite",
	0xB5: "Powder Snow",
	0xB6: "Protect",
	0xB7: "Mach Punch",
	0xB8: "Scary Face",
	0xB9: "Faint Attack",
	0xBA: "Sweet Kiss",
	0xBB: "Belly Drum",
	0xBC: "Sludge Bomb",
	0xBD: "Mud-Slap",
	0xBE: "Octazooka",
	0xBF: "Spikes",
	0xC0: "Zap Cannon",
	0xC1: "Foresight",
	0xC2: "Destiny Bond",
	0xC3: "Perish Song",
	0xC4: "Icy Wind",
	0xC5: "Detect",
	0xC6: "Bone Rush",
	0xC7: "Lock-On",
	0xC8: "Outrage",
	0xC9: "Sandstorm",
	0xCA: "Giga Drain",
	0xCB: "Endure",
	0xCC: "Charm",
	0xCD: "Rollout",
	0xCE: "False Swipe",
	0xCF: "Swagger",
	0xD0: "Milk Drink",
	0xD1: "Spark",
	0xD2: "Fury Cutter",
	0xD3: "Steel Wing",
	0xD4: "Mean Look",
	0xD5: "Attract",
	0xD6: "Sleep Talk",
	0xD7: "Heal Bell",
	0xD8: "Return",
	0xD9: "Present",
	0xDA: "Frustration",
	0xDB: "Safeguard",
	0xDC: "Pain Split",
	0xDD: "Sacred Fire",
	0xDE: "Magnitude",
	0xDF: "Dynamicpunch",
	0xE0: "Megahorn",
	0xE1: "Dragonbreath",
	0xE2: "Baton Pass",
	0xE3: "Encore",
	0xE4: "Pursuit",
	0xE5: "Rapid Spin",
	0xE6: "Sweet Scent",
	0xE7: "Iron Tail",
	0xE8: "Metal Claw",
	0xE9: "Vital Throw",
	0xEA: "Morning Sun",
	0xEB: "Synthesis",
	0xEC: "Moonlight",
	0xED: "Hidden Power",
	0xEE: "Cross Chop",
	0xEF: "Twister",
	0xF0: "Rain Dance",
	0xF1: "Sunny Day",
	0xF2: "Crunch",
	0xF3: "Mirror Coat",
	0xF4: "Psych Up",
	0xF5: "Extremespeed",
	0xF6: "Ancientpower",
	0xF7: "Shadow Ball",
	0xF8: "Future Sight",
	0xF9: "Rock Smash",
	0xFA: "Whirlpool",
	0xFB: "Beat Up",
}

var Items = map[uint8]string{
	0x01: "Master Ball",
	0x02: "Ultra Ball",
	0x03: "Brightpowder",
	0x04: "Great Ball",
	0x05: "Pok? Ball",
	0x06: "Teru-Sama",
	0x07: "Bicycle",
	0x08: "Moon Stone",
	0x09: "Antidote",
	0x0A: "Burn Heal",
	0x0B: "Ice Heal",
	0x0C: "Awakening",
	0x0D: "Parlyz Heal",
	0x0E: "Full Restore",
	0x0F: "Max Potion",
	0x10: "Hyper Potion",
	0x11: "Super Potion",
	0x12: "Potion",
	0x13: "Escape Rope",
	0x14: "Repel",
	0x15: "Max Elixer",
	0x16: "Fire Stone",
	0x17: "Thunderstone",
	0x18: "Water Stone",
	0x19: "Teru-Sama",
	0x1A: "HP Up",
	0x1B: "Protein",
	0x1C: "Iron",
	0x1D: "Carbos",
	0x1E: "Lucky Punch",
	0x1F: "Calcium",
	0x20: "Rare Candy",
	0x21: "X Accuracy",
	0x22: "Leaf Stone",
	0x23: "Metal Powder",
	0x24: "Nugget",
	0x25: "Pok? Doll",
	0x26: "Full Heal",
	0x27: "Revive",
	0x28: "Max Revive",
	0x29: "Guard Spec.",
	0x2A: "Super Repel",
	0x2B: "Max Repel",
	0x2C: "Dire Hit",
	0x2D: "Teru-Sama",
	0x2E: "Fresh Water",
	0x2F: "Soda Pop",
	0x30: "Lemonade",
	0x31: "X Attack",
	0x32: "Teru-Sama",
	0x33: "X Defend",
	0x34: "X Speed",
	0x35: "X Special",
	0x36: "Coin Case",
	0x37: "Itemfinder",
	0x38: "Teru-Sama",
	0x39: "Exp.Share",
	0x3A: "Old Rod",
	0x3B: "Good Rod",
	0x3C: "Silver Leaf",
	0x3D: "Super Rod",
	0x3E: "PP Up",
	0x3F: "Ether",
	0x40: "Max Ether",
	0x41: "Elixer",
	0x42: "Red Scale",
	0x43: "Secretpotion",
	0x44: "S.S.Ticket",
	0x45: "Mystery Egg",
	0x46: "Clear Bell",
	0x47: "Silver Wing",
	0x48: "Moomoo Milk",
	0x49: "Quick Claw",
	0x4A: "Psncureberry",
	0x4B: "Gold Leaf",
	0x4C: "Soft Sand",
	0x4D: "Sharp Beak",
	0x4E: "Przcureberry",
	0x4F: "Burnt Berry",
	0x50: "Ice Berry",
	0x51: "Poison Barb",
	0x52: "King's Rock",
	0x53: "Bitter Berry",
	0x54: "Mint Berry",
	0x55: "Red Apricorn",
	0x56: "Tinymushroom",
	0x57: "Big Mushroom",
	0x58: "Silverpowder",
	0x59: "Blu Apricorn",
	0x5A: "Teru-Sama",
	0x5B: "Amulet Coin",
	0x5C: "Ylw Apricorn",
	0x5D: "Grn Apricorn",
	0x5E: "Cleanse Tag",
	0x5F: "Mystic Water",
	0x60: "Twistedspoon",
	0x61: "Wht Apricorn",
	0x62: "Blackbelt",
	0x63: "Blk Apricorn",
	0x64: "Teru-Sama",
	0x65: "Pnk Apricorn",
	0x66: "Blackglasses",
	0x67: "Slowpoketail",
	0x68: "Pink Bow",
	0x69: "Stick",
	0x6A: "Smoke Ball",
	0x6B: "Nevermeltice",
	0x6C: "Magnet",
	0x6D: "Miracleberry",
	0x6E: "Pearl",
	0x6F: "Big Pearl",
	0x70: "Everstone",
	0x71: "Spell Tag",
	0x72: "Ragecandybar",
	0x73: "GS Ball",
	0x74: "Blue Card",
	0x75: "Miracle Seed",
	0x76: "Thick Club",
	0x77: "Focus Band",
	0x78: "Teru-Sama",
	0x79: "Energypowder",
	0x7A: "Energy Root",
	0x7B: "Heal Powder",
	0x7C: "Revival Herb",
	0x7D: "Hard Stone",
	0x7E: "Lucky Egg",
	0x7F: "Card Key",
	0x80: "Machine Part",
	0x81: "Egg Ticket",
	0x82: "Lost Item",
	0x83: "Stardust",
	0x84: "Star Piece",
	0x85: "Basement Key",
	0x86: "Pass",
	0x87: "Teru-Sama",
	0x88: "Teru-Sama",
	0x89: "Teru-Sama",
	0x8A: "Charcoal",
	0x8B: "Berry Juice",
	0x8C: "Scope Lens",
	0x8D: "Teru-Sama",
	0x8E: "Teru-Sama",
	0x8F: "Metal Coat",
	0x90: "Dragon Fang",
	0x91: "Teru-Sama",
	0x92: "Leftovers",
	0x93: "Teru-Sama",
	0x94: "Teru-Sama",
	0x95: "Teru-Sama",
	0x96: "Mysteryberry",
	0x97: "Dragon Scale",
	0x98: "Berserk Gene",
	0x99: "Teru-Sama",
	0x9A: "Teru-Sama",
	0x9B: "Teru-Sama",
	0x9C: "Sacred Ash",
	0x9D: "Heavy Ball",
	0x9E: "Flower Mail",
	0x9F: "Level Ball",
	0xA0: "Lure Ball",
	0xA1: "Fast Ball",
	0xA2: "Teru-Sama",
	0xA3: "Light Ball",
	0xA4: "Friend Ball",
	0xA5: "Moon Ball",
	0xA6: "Love Ball",
	0xA7: "Normal Box",
	0xA8: "Gorgeous Box",
	0xA9: "Sun Stone",
	0xAA: "Polkadot Bow",
	0xAB: "Teru-Sama",
	0xAC: "Up-Grade",
	0xAD: "Berry",
	0xAE: "Gold Berry",
	0xAF: "Squirtbottle",
	0xB0: "Teru-Sama",
	0xB1: "Park Ball",
	0xB2: "Rainbow Wing",
	0xB3: "Teru-Sama",
	0xB4: "Brick Piece",
	0xB5: "Surf Mail",
	0xB6: "Litebluemail",
	0xB7: "Portraitmail",
	0xB8: "Lovely Mail",
	0xB9: "Eon Mail",
	0xBA: "Morph Mail",
	0xBB: "Bluesky Mail",
	0xBC: "Music Mail",
	0xBD: "Mirage Mail",
	0xBE: "Teru-Sama",
	0xBF: "TM01",
	0xC0: "TM02",
	0xC1: "TM03",
	0xC2: "TM04",
	0xC3: "Teru-Sama",
	0xC4: "TM05",
	0xC5: "TM06",
	0xC6: "TM07",
	0xC7: "TM08",
	0xC8: "TM09",
	0xC9: "TM10",
	0xCA: "TM11",
	0xCB: "TM12",
	0xCC: "TM13",
	0xCD: "TM14",
	0xCE: "TM15",
	0xCF: "TM16",
	0xD0: "TM17",
	0xD1: "TM18",
	0xD2: "TM19",
	0xD3: "TM20",
	0xD4: "TM21",
	0xD5: "TM22",
	0xD6: "TM23",
	0xD7: "TM24",
	0xD8: "TM25",
	0xD9: "TM26",
	0xDA: "TM27",
	0xDB: "TM28",
	0xDC: "Teru-Sama",
	0xDD: "TM29",
	0xDE: "TM30",
	0xDF: "TM31",
	0xE0: "TM32",
	0xE1: "TM33",
	0xE2: "TM34",
	0xE3: "TM35",
	0xE4: "TM36",
	0xE5: "TM37",
	0xE6: "TM38",
	0xE7: "TM39",
	0xE8: "TM40",
	0xE9: "TM41",
	0xEA: "TM42",
	0xEB: "TM43",
	0xEC: "TM44",
	0xED: "TM45",
	0xEE: "TM46",
	0xEF: "TM47",
	0xF0: "TM48",
	0xF1: "TM49",
	0xF2: "TM50",
	0xF3: "HM01",
	0xF4: "HM02",
	0xF5: "HM03",
	0xF6: "HM04",
	0xF7: "HM05",
	0xF8: "HM06",
	0xF9: "HM07",
	0xFA: "Teru-Sama",
	0xFB: "Teru-Sama",
	0xFC: "Teru-Sama",
	0xFD: "Teru-Sama",
	0xFE: "Teru-Sama",
	0xFF: "Teru-Sama",
}

var Types = map[uint8]string{
	0x00: "Normal",
	0x01: "Fighting",
	0x02: "Flying",
	0x03: "Poison",
	0x04: "Ground",
	0x05: "Rock",
	0x06: "Bird",
	0x07: "Bug",
	0x08: "Ghost",
	0x09: "Steel",
	//0x0a:-0x12: blank
	0x13: "???",
	0x14: "Fire",
	0x15: "Water",
	0x16: "Grass",
	0x17: "Electric",
	0x18: "Psychic",
	0x19: "Ice",
	0x1a: "Dragon",
	0x1b: "Dark",
}
