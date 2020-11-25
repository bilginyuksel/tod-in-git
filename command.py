import os
import re
from fileOperation import writeBytesToFile

ROOT_DIR = "zippedCommands"

def execCmd(command):
    stream = os.popen(command)
    output = stream.read()
    stream.close()
    return output


class LiveCommandTemplate:
    
    def __init__(self, command, userInteractionNeeded = False):
        self.command = command
        self.isUserInteractionNeeded = userInteractionNeeded

    def _updateLiveCommandTemplateWithCustomData(self):
        fields = re.findall("[$]\w*[$]", self.command)
        for field in fields:
            userInput = input(field)
            self.command.replace(field, userInput)

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
            self.cmdQueue[0].execute()
            self.cmdQueue.pop(0)
    

    def save(self):
        writeBytesToFile(ROOT_DIR, self.cmdQueue)
