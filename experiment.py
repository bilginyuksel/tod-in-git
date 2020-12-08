import os
import re
import pickle
from datetime import datetime



def getCommandResult(command):
    stream = os.popen(command)
    output = stream.read()
    stream.close()
    return output

class User:
    def __init__(self, username, email):
        self.username = username
        self.email = email


def getUserInformation():
    username = getCommandResult("git config --local user.name")
    email = getCommandResult("git config --local user.email")
    if len(username) == 0:
        username = getCommandResult("git config --global user.name")
        email = getCommandResult("git config --local user.email")

    return User(username.replace("\n", ""), email.replace("\n", ""))


currentUser = getUserInformation()
todos = {}
class Todo:
    def __init__(self, branch, title, description=""):
        self.branch = branch
        self.title = title
        self.description = description
        self.createdTime = datetime.now()
        self.author = currentUser
        self.importance = 2

    def get_filename(self):
        return str(self.branch + "_" + self.author).lower()

    def __str__(self):
        return "branch: %s, note: %s, author: %s, createdTime: %s," % (self.branch, self.note, self.author.username, self.createdTime)


def save(o):
    outputFile = open(".tod/demo", "wb")
    pickle.dump(o, outputFile)
    outputFile.close()


def read():
    readFile = open(".tod/demo", "rb")
    content = pickle.load(readFile)
    print(content)
    return content

class Branch:
    def __init__(self):
        pass




def createTodBranchesIfNotExists(branches):
    currentBranches = listBranches()
    print("currentBranches:", currentBranches)
    print("branches:", branches)

    for branch in branches:
        if branch in currentBranches:
            continue
        # create a new branch
        with open(".tod/"+branch, "w") as f:
            f.write(currentUser.username + "\n" + currentUser.email)


def listBranches():
    branches = []
    # for root, dir, files in os.walk(".tod"):
    for _, _, files in os.walk(".tod"):
        branches.append(files)
    return branches

# tod setup


def init():
    # first control if this project is a git project or not
    branchResult = getCommandResult("git branch")
    if len(branchResult) == 0:
        raise Exception("This is not a git project")

    # otherwise create the file structure
    os.mkdir(".tod")
    # sync with current git.
    cleanedOutput = branchResult.replace("*", "").replace("\t", "").replace(" ", "")
    branches = [branch for branch in cleanedOutput.split("\n") if len(branch) > 0]
    createTodBranchesIfNotExists(branches)

# Get user informations from git.




def getCurrentBranch():
    return getCommandResult("git branch")


# currentUser = getUserInformation()
# note1 = Todo(1, "My new note!")
# print(note1)

# init()
todos["131413"] = Todo("branch", "My first note title", "My first description")
todos["0000"] =  Todo("boranch", "My second note title", "My second description")
todos["123000"] = Todo("branch", "My third note title", "My third description")

save(todos)
content  = read()
