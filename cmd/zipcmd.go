package cmd

import (
	"encoding/json"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/google/uuid"
)

const liveDataRegexp = "[$]\\w*[$]"

var dtoArr dtoZipArr

func init() {
	fmt.Println("Package initialized")
}

// ZipCmd ...
type ZipCmd struct {
	uuid        string
	commands    []Cmd
	abbrevation string
	author      string
	createdTime time.Time
	updatedTime time.Time
	isShared    bool
	isEditable  bool
	os          string
}

// ZipCmdList ...
type ZipCmdList struct {
	ZippedCommands []ZipCmd
}

// Cmd ...
type Cmd struct {
	uuid     string
	command  string
	liveData bool
	datas    []string
}

// NewZipCmd ...
func NewZipCmd(abbrevation string, author string, editable bool, shared bool) *ZipCmd {
	uid, _ := uuid.NewUUID()
	return &ZipCmd{
		uuid:        uid.String(),
		abbrevation: abbrevation,
		author:      author,
		createdTime: time.Now(),
		updatedTime: time.Now(),
		os:          runtime.GOOS,
		isEditable:  editable,
		isShared:    shared}
}

// HasLiveData ...
func (cmd *Cmd) HasLiveData() bool {
	return cmd.liveData
}

// AddCommand ...
func (zpCm *ZipCmd) AddCommand(command string) {
	cmd := newCmd(command)
	zpCm.commands = append(zpCm.commands, cmd)
}

// ListOrdered ...
func (zpCm *ZipCmd) ListOrdered() []Cmd {
	return zpCm.commands
}

func fillLiveDataOfCommand(cmd *Cmd) {
	for _, vl := range cmd.datas {
		// get input
		var userInp string
		fmt.Scanf(vl, userInp)
		cmd.command = strings.ReplaceAll(cmd.command, vl, userInp)
	}
}

// Run ...
func (zpCm *ZipCmd) Run() []string {
	commands := zpCm.ListOrdered()
	for i := 0; i < len(commands); i++ {
		if commands[i].HasLiveData() {
			fillLiveDataOfCommand(&commands[i])
		}
	}
	output := []string{}
	for _, value := range commands {
		if zpCm.os == "windows" {
			out, _ := RunPsl(value.command)
			output = append(output, string(out))
		} else {
			out, _ := RunLx(value.command)
			output = append(output, string(out))
		}
	}
	return output
}

type dtoZipCmd struct {
	UUID        string
	Commands    []Cmd
	Abbrevation string
	Author      string
	CreatedTime time.Time
	UpdatedTime time.Time
	IsShared    bool
	IsEditable  bool
	Os          string
}

type dtoZipArr struct {
	DtoZips []dtoZipCmd
}

func fillDTOArray() {
	// read files then fill dto array list.

}

// Save ...
func (zpCm *ZipCmd) Save() {
	cons := &dtoZipCmd{
		UUID:        zpCm.uuid,
		Commands:    zpCm.commands,
		Abbrevation: zpCm.abbrevation,
		Author:      zpCm.author,
		CreatedTime: zpCm.createdTime,
		UpdatedTime: zpCm.updatedTime,
		IsShared:    zpCm.isShared,
		IsEditable:  zpCm.isEditable,
		Os:          zpCm.os,
	}
	// No better way to do it currently because of I don't want to mess with file operations
	// I am using json encoding/decoding
	// byteData, _ := json.Marshal(&cons)
	// fillDTOArrayIfNot()
	dtoArr.DtoZips = append(dtoArr.DtoZips, *cons)
	if len(dtoArr.DtoZips) == 0 {
		fillDTOArray()
	}
	byteData, _ := json.Marshal(&dtoArr)
	fmt.Println(byteData)
}

func newCmd(command string) Cmd {
	uid, _ := uuid.NewUUID()
	compiledRegexp, _ := regexp.Compile(liveDataRegexp)
	hasLiveData := compiledRegexp.MatchString(command)
	foundLiveData := compiledRegexp.FindAllString(command, -1)
	return Cmd{liveData: hasLiveData, command: command, datas: foundLiveData, uuid: uid.String()}
}
