# Mindcli-for-controlling-multi-robots

## Introduction
This code is for controlling multi-robots by a mindcli in a PC or sever. It has changed  a small of the open source mindcli by Vincross.

## Usage
The Usage of controlling multi robots is:

`$ sudo mind group <OPTION> <ROBOT/GROUP NAME> <ROBOT/GROUP NAME>`

>(PS: `sudo` is essential, otherwise the opreations will not be saved. When use `$ mind run`, you should make sure `cd` to the dir which has a `skill.mpk` file.)

### Options:
The options and their means in command usage is:

  1. **addr**:    Add robot into group (eg. `$ sudo mind group addr HEXA ALL` means add robot HEXA into group ALL.

  2. **addg**:    Creat a new group (eg. `$ sudo mind group addg all` means creat a new group ALL.

  3. **delr**:	Delet robot in a specific group (eg. `$ sudo mind group delr HEXA ALL` means delet robot HEXA from group ALL.

  4. **delg**:	Delet a existing group (eg. `$ sudo mind group delg ALL` means delet group ALL.

  5. **listr**:	List robot in a specific group (eg. `$ sudo mind group listr ALL` means list all robots in group ALL.

  6. **listg**:	List all existing groups (eg. `$ sudo mind group listg` means list all groups' information.

  7. **runn**:	Run skill in a group without install (eg. `$ sudo mind group runn ALL` means run skill in group ALL without install.

  8. **run**:	Run skill in a group with install (eg. `$ sudo mind group runn ALL` means run skill in group ALL with install.

## Changes I made
The changes I made in the mindcli are as follows:

1. Add file `group.go` in package `mindcli` to support the command .

2. Add file `portconfig.go` in package `mindcli` to manage the allocated ports, the infomation of ports is saved in `homeDir()/.mindport.json`.

3. Change the strut `MindCli` in `mindcli.go` and strut `UserConfig` in `userconfig.go` basis on new strut `Group` and `PortConfig`.

4. Change the functions in `mindcli.go`: add `mindcli.AllocatePort(robotIp)` in function `RunSkill()`, and change the function `NewMindCli()`

## Test GUI
I also write a GUI program `GUI.py` to control multi-robots by python. It is used to click button instead of openning browser, connecting localhost port and clicking buttons.

## Others
There is still some bugs in my program, all suggestions are welcome. Lets make it better!
