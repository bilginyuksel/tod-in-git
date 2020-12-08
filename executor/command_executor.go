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

// UTIL METHODS ARE BELOW

// TrieNode ...
type TrieNode struct {
	isEnd    bool
	children map[rune]TrieNode
	fuuid    string
}

/*
Use this trie to understand given uuid is unique or not
if the uuid is unique then get the actual uuid and use
the asked todo or command.
*/
func buildTrie(list []string) *TrieNode {
	root := TrieNode{}
	for _, str := range list {
		node := root
		for _, char := range str {
			node.children[char] = TrieNode{}
			node = node.children[char]
		}
		node.isEnd = true
		node.fuuid = str
	}
	return &root
}

// FindUUIDFromPrefix ...
func FindUUIDFromPrefix(uid string) (bool, string) {
	root := *buildTrie([]string{})
	for _, char := range uid {
		if _, ok := root.children[char]; ok {
			root = root.children[char]
		} else {
			return false, ""
		}
	}

	// Go if there is only one children exists.
	for len(root.children) == 1 {
		for key := range root.children {
			root = root.children[key]
		}
	}
	// Return the uuid if it is end and there is no children.
	if root.isEnd && len(root.children) == 0 {
		return true, root.fuuid
	}
	return false, ""
}
