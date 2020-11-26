import os
import re
from datetime import datetime


def getCommandResult(command):
    stream = os.popen(command)
    output = stream.read()
    stream.close()
    return output


"""
tod template --create
press y for add command
1. git add .
2. git commit -m "[TicketNo:] AR000000
[Description:] $TEMPLATE$
[Binary Source:] NA"
3. git push -u origin $branch$
attend commit

tod commit
Description $TEMPLATE$: asıjdaıjdsa
$branch$: dev/init

"""


class Todo:
    def __init__(self, branch, note):
        self.branch = branch
        self.note = note
        self.createdTime = datetime.now()
        self.author = currentUser

    def __str__(self):
        return "branch: %s, note: %s, author: %s, createdTime: %s," % (self.branch, self.note, self.author.username, self.createdTime)


class Branch:
    def __init__(self):
        pass


class User:
    def __init__(self, username, email):
        self.username = username
        self.email = email


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
    cleanedOutput = branchResult.replace(
        "*", "").replace("\t", "").replace(" ", "")
    branches = [branch for branch in cleanedOutput.split(
        "\n") if len(branch) > 0]
    createTodBranchesIfNotExists(branches)

# Get user informations from git.


def getUserInformation():
    username = getCommandResult("git config --local user.name")
    email = getCommandResult("git config --local user.email")
    if len(username) == 0:
        username = getCommandResult("git config --global user.name")
        email = getCommandResult("git config --local user.email")

    return User(username.replace("\n", ""), email.replace("\n", ""))


def getCurrentBranch():
    return getCommandResult("git branch")


currentUser = getUserInformation()
note1 = Todo(1, "My new note!")
print(note1)

init()
