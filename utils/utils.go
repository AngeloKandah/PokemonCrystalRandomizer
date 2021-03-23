package utils

import (
	"encoding/binary"
	"fmt"
	"os"
	"strings"

	"github.com/AngeloKandah/pokemon/lib"
)

func ConvertNicknameToString(bit []byte) (nickname string) {
	var nick_bits [11]string
	for i := 0; i < len(bit); i++ {
		nick_bits[i] = lib.CharacterTable[(bit[i] / 16)][(bit[i] % 16)]
	}
	nickname = strings.Join(nick_bits[:], "")
	return
}

func ConvertTrainerToString(bit []byte) (trainer string) {
	var trainer_bits [8]string
	for i := 0; i < len(bit); i++ {
		trainer_bits[i] = lib.CharacterTable[(bit[i] / 16)][(bit[i] % 16)]
	}
	trainer = strings.Join(trainer_bits[:], "")
	return
}

func GetSearchAddress() (search_addr string) {
	f, err := os.OpenFile("C:/Users/Angelo/Desktop/mGBA-0.8.4-win64/Pokemon - Crystal Version (USA, Europe) (Rev 1).sav", os.O_RDONLY, 0777)
	//f, err := os.OpenFile("C:/Users/Angelo/Desktop/mGBA-0.8.4-win64/Pokemon Perfect Crystal (05.04.2019).sav", os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 7)       //Reads 7 bytes
	_, err = f.ReadAt(buf, 4619) //Reads the bytes in buf at the offset 4619 for the name to scan for
	if err != nil {
		panic(err)
	}
	var save_name [7]string
	for i := 0; i < 7; i++ {
		if err != nil {
			fmt.Printf("Failed reading memory. %s", err)
			break
		}
		save_name[i] = lib.CharacterTable[(buf[i] / 16)][(buf[i] % 16)] //Convert hex values into chars from char table
	}
	fmt.Println(strings.Join(save_name[:], "")) //This is how you can convert array into a single string for printing

	byteReader := binary.LittleEndian
	search_addr = fmt.Sprintf("%X", byteReader.Uint32(buf[0:4]))
	return
}

func GetGender(dex uint8, ATK uint8)(gender string){
	f, err := os.OpenFile("C:/Users/Angelo/Desktop/mGBA-0.8.4-win64/Pokemon - Crystal Version (USA, Europe) (Rev 1).gbc", os.O_RDONLY, 0777)
	//f, err := os.OpenFile("C:/Users/Angelo/Desktop/mGBA-0.8.4-win64/Pokemon Perfect Crystal (05.04.2019).gbc", os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 1)       //Reads 7 bytes
	_, err = f.ReadAt(buf, 0x51431+((int64(dex)-1)*32)) //Reads the bytes in buf at the offset 51431 for the gender ratio
	if err != nil {
		panic(err)
	}
	if buf[0] == 255{
		gender = "Genderless"
		return
	}else if (buf[0] >> 4) < ATK || buf[0] == 0{
		gender = "Male"
		return
	} else if (buf[0] >> 4 >= ATK) || buf[0] == 254{
		gender = "Female"
		return
	}
	return
}

func GetTypes(dex uint8)(types string){
	f, err := os.OpenFile("C:/Users/Angelo/Desktop/mGBA-0.8.4-win64/Pokemon - Crystal Version (USA, Europe) (Rev 1).gbc", os.O_RDONLY, 0777)
	//f, err := os.OpenFile("C:/Users/Angelo/Desktop/mGBA-0.8.4-win64/Pokemon Perfect Crystal (05.04.2019).gbc", os.O_RDONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf := make([]byte, 2)       //Reads 7 bytes
	_, err = f.ReadAt(buf, 0x5142B+((int64(dex)-1)*32)) //Reads the bytes in buf at the offset 51431 for the gender ratio
	if err != nil {
		panic(err)
	}
	type1 := lib.Types[buf[0]]
	type2 := lib.Types[buf[1]]
	if type1 == type2{
		types = type1
		return
	}else{
		types = type1 + " " + type2 
		return
	}
}
