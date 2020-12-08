package executor

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/uuid"
)

const liveDataRegexp = "[$]\\w*[$]"

func init() {
	fmt.Println("Package initialized")
	uid, _ := uuid.NewUUID()
	fmt.Println(uid)
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

func newCmd(command string) Cmd {
	uid, _ := uuid.NewUUID()
	compiledRegexp, _ := regexp.Compile(liveDataRegexp)
	hasLiveData := compiledRegexp.MatchString(command)
	foundLiveData := compiledRegexp.FindAllString(command, -1)
	return Cmd{liveData: hasLiveData, command: command, datas: foundLiveData, uuid: uid.String()}
}
