import os
import re
import subprocess
import sys
from fileOperation import writeBytesToFile

ROOT_DIR = "zippedCommands"
GIT_BASH = 'start "" "%ProgramFiles%\Git\git-bash.exe" -c "echo 1 && echo 2 && /usr/bin/bash --login -i"'
POWERSHELL = 'powershell.exe'
CMD = 'cmd.exe'

def execCmd(command):
    p = subprocess.Popen([POWERSHELL, command], stdout=sys.stdout)
    p.communicate()

class LiveCommandTemplate:
    
    def __init__(self, command, userInteractionNeeded = False):
        self.command = command
        self.isUserInteractionNeeded = userInteractionNeeded

    def _updateLiveCommandTemplateWithCustomData(self):
        fields = re.findall("[$]\w*[$]", self.command)
        for field in fields:
            userInput = input(field+": ")
            print(userInput, field, self.command)
            self.command = self.command.replace(field, userInput)

    def execute(self):
        if self.isUserInteractionNeeded:
            self._updateLiveCommandTemplateWithCustomData()
        return execCmd(self.command)


class ZippedCommand:

    def __init__(self, zipName):
        self.name = zipName
        self.cmdQueue = []
    
    def add(self, command):
        self.cmdQueue.append(command)

    def run(self):
        while len(self.cmdQueue) > 0:
            print(self.cmdQueue[0].execute())
            self.cmdQueue.pop(0)
    
    def save(self):
        writeBytesToFile(ROOT_DIR, self.cmdQueue)


# example code it is not for production only for experiment purposes
option = 1
zippedCommand = ZippedCommand("test")
while option != 0:
    option = int(input("inp"))
    command = input("cmd")
    isUserInteractionNeeded = False
    if '$' in command: isUserInteractionNeeded = True 
    zippedCommand.add(LiveCommandTemplate(command, isUserInteractionNeeded))

else:
    zippedCommand.run()
