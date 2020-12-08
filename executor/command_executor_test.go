package executor

import "testing"

func TestNewCommand_WithoutLiveData(t *testing.T) {
	cmdString := "git branch -a"
	expected := Cmd{command: cmdString}
	given := newCmd(cmdString)

	if !sameCmd(expected, given) {
		t.Error()
	}
}

func TestNewCommand_WithOneLiveData(t *testing.T) {
	cmdString := "git commit -m \"[TicketNo:] AR0202020\n[Description:] $DESCRIPTION$\n[Binary Source:] NA"
	expected := Cmd{command: cmdString, liveData: true, datas: []string{"$DESCRIPTION$"}}
	given := newCmd(cmdString)

	if !sameCmd(expected, given) {
		t.Error()
	}
}

func TestNewCommand_WithTwoLiveData(t *testing.T) {
	cmdString := "git commit -m \"[TicketNo:] AR0202020\n[Description:] $DESC$\n[Binary Source:] $BIN$"
	expected := Cmd{command: cmdString, liveData: true, datas: []string{"$DESC$", "$BIN$"}}
	given := newCmd(cmdString)

	if !sameCmd(expected, given) {
		t.Error()
	}
}

func TestNewCommand_WithThreeLiveData(t *testing.T) {
	cmdString := "git commit -m \"[TicketNo:] $AR$\n[Description:] $AR$\n[Binary Source:] $BIN$"
	expected := Cmd{command: cmdString, liveData: true, datas: []string{"$AR$", "$AR$", "$BIN$"}}
	given := newCmd(cmdString)

	if !sameCmd(expected, given) {
		t.Error()
	}
}

func TestNewZipCommand_ExpectZipCommand(t *testing.T) {
	expected := &ZipCmd{abbrevation: "gcom", author: "bilginyuksel", isShared: false, isEditable: true}
	given := NewZipCmd("gcom", "bilginyuksel", true, false)
	if !sameZipCmd(expected, given) {
		t.Error()
	}
}

func TestAddCommand_AddTwoCommandWithoutLiveData(t *testing.T) {
	expected := &ZipCmd{abbrevation: "gcom", author: "bilginyuksel", isShared: false, isEditable: true,
		commands: []Cmd{newCmd("git branch -a"), newCmd("git init")}}
	given := NewZipCmd("gcom", "bilginyuksel", true, false)
	given.AddCommand("git branch -a")
	given.AddCommand("git init")
	if !sameZipCmd(expected, given) {
		t.Error()
	}
}

func TestAddCommand_AddThreeCommandWithLiveData(t *testing.T) {
	cmds1 := "git commit -m \"[TicketNo:] AR0202020\n[Description:] $DESC$\n[Binary Source:] $BIN$"
	cmds2 := "git push -u origin $BRANCH$"
	cmds3 := "git config --global user.name"
	expected := &ZipCmd{abbrevation: "gcom", author: "bilginyuksel", isShared: false, isEditable: true,
		commands: []Cmd{newCmd(cmds1), newCmd(cmds2), newCmd(cmds3)}}
	given := NewZipCmd("gcom", "bilginyuksel", true, false)
	given.AddCommand(cmds1)
	given.AddCommand(cmds2)
	given.AddCommand(cmds3)
	if !sameZipCmd(expected, given) {
		t.Error()
	}
}

func sameCmd(expected Cmd, given Cmd) bool {
	isLiveDatasEqual := expected.HasLiveData() == given.HasLiveData()
	isCommandsEqual := expected.command == given.command
	isDatasEqual := len(expected.datas) == len(given.datas)
	for idx, val := range expected.datas {
		isDatasEqual = isDatasEqual && (val == given.datas[idx])
	}
	return isLiveDatasEqual && isCommandsEqual && isDatasEqual
}

func sameZipCmd(expected *ZipCmd, given *ZipCmd) bool {
	isAuthorsSame := expected.author == given.author
	isAbbrevationsSame := expected.abbrevation == given.abbrevation
	isEditablesSame := expected.isEditable == given.isEditable
	isSharedSame := expected.isShared == given.isShared
	isCommandsSame := len(expected.commands) == len(given.commands)
	givenCommands := given.ListOrdered()
	for idx, val := range expected.ListOrdered() {
		isCommandsSame = isCommandsSame && sameCmd(val, givenCommands[idx])
	}
	sameZip := isAuthorsSame && isAbbrevationsSame && isEditablesSame && isSharedSame && isCommandsSame
	return sameZip
}
