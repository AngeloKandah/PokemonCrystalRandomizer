package pokemon

import (
	"encoding/binary"
	"strings"

	"github.com/AngeloKandah/pokemon/lib"
	"github.com/AngeloKandah/pokemon/utils"
)

type PokemonInfo struct {
	Species   string
	Held_Item string
	Move1     string
	Move2     string
	Move3     string
	Move4     string
	ID_NUM    uint16
	EXP       uint32
	HP_EVs    uint16
	ATK_EVs   uint16
	DEF_EVs   uint16
	SPD_EVs   uint16
	SPC_EVs   uint16
	IV_ATK    uint8
	IV_DEF    uint8
	IV_SPD    uint8
	IV_SPC    uint8
	Move1_PP  byte
	Move2_PP  byte
	Move3_PP  byte
	Move4_PP  byte
	Happiness byte
	Pokerus   byte
	Poke_Seer []byte
	Level     byte
	Status    byte
	Unused    byte
	CUR_HP    uint16
	MAX_HP    uint16
	ATK       uint16
	DEF       uint16
	SPD       uint16
	SPC_ATK   uint16
	SPC_DEF   uint16
	DEX_NUM   uint8 //Technically Party Species but using it for other funcs
	OT_NAME   string
	NICKNAME  string
	Gender    string
	Types     string
}

func CreatePokemonInfo(p []byte) (pokemon PokemonInfo, err error) {
	byteReader := binary.BigEndian
	pokemon = PokemonInfo{
		Species:   strings.ToUpper(lib.PokemonSpecies[(p[0])]),
		Held_Item: lib.Items[(p[1])],
		Move1:     lib.Moves[(p[2])],
		Move2:     lib.Moves[(p[3])],
		Move3:     lib.Moves[(p[4])],
		Move4:     lib.Moves[(p[5])],
		ID_NUM:    byteReader.Uint16(p[6:8]),
		EXP:       (uint32(byteReader.Uint16(p[8:10])) << 8) + uint32(p[10]),
		HP_EVs:    byteReader.Uint16(p[11:13]),
		ATK_EVs:   byteReader.Uint16(p[13:15]),
		DEF_EVs:   byteReader.Uint16(p[15:17]),
		SPD_EVs:   byteReader.Uint16(p[17:19]),
		SPC_EVs:   byteReader.Uint16(p[19:21]),
		IV_ATK:    uint8(p[21] >> 4),
		IV_DEF:    (uint8(p[21] << 4)) >> 4,
		IV_SPD:    uint8(p[22] >> 4),
		IV_SPC:    (uint8(p[22] << 4)) >> 4,
		Move1_PP:  (p[23] & 0x3F),
		Move2_PP:  (p[24] & 0x3F),
		Move3_PP:  (p[25] & 0x3F),
		Move4_PP:  (p[26] & 0x3F),
		Happiness: (p[27]),
		Pokerus:   (p[28]),
		Poke_Seer: (p[29:31]),
		Level:     (p[31]),
		Status:    (p[32]),
		Unused:    (p[33]),
		CUR_HP:    byteReader.Uint16(p[34:36]),
		MAX_HP:    byteReader.Uint16(p[36:38]),
		ATK:       byteReader.Uint16(p[38:40]),
		DEF:       byteReader.Uint16(p[40:42]),
		SPD:       byteReader.Uint16(p[42:44]),
		SPC_ATK:   byteReader.Uint16(p[44:46]),
		SPC_DEF:   byteReader.Uint16(p[46:48]),
		DEX_NUM:   uint8(p[48]), //Technically Party Species but using it for other funcs
		OT_NAME:   utils.ConvertTrainerToString(p[49:57]),
		NICKNAME:  utils.ConvertNicknameToString(p[57:68]),
		Gender:    "",
		Types:     "",
	}
	return
}
