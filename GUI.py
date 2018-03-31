# -*- coding: utf-8 -*-
import easygui as g
import sys
import json
import os.path
from splinter import Browser
from selenium import webdriver

MINDPORT_FILE_NAME = '/home/zxt/.mindport.json'
MIND_FILE_NAME = '/home/zxt/.mind.json'

def GetJson(filename):
    if not os.path.isfile(filename):
        print('{} does not exit'.format(filename))
        sys.exit()
    load_dict = json.load(open(filename,'r'))
    return load_dict

# return dictionary array of robots information
def GetRobots():
    return GetJson(MIND_FILE_NAME).get('Robots')

#return string array of robots name
def GetRobotsName():
    robotname = []
    for i in GetRobots():
        robotname.append(i.get('Name'))
    return robotname
def GetRobotIp(name):
    for i in GetRobots():
        if i.get('Name') == name :
            robotip = i.get('IP')
    return robotip

# return dictionary array of ports information
def GetPortsInfo():
    return GetJson(MINDPORT_FILE_NAME).get('Robotsport')
def GetServeRemotePort(name):
    for i in GetPortsInfo():
        if i.get('RobotIp') == GetRobotIp(name) :
            ServeRemotePort = i.get('ServeRemotePort')
    return ServeRemotePort


# use shell to reach these function
def init():
    browser = webdriver.Chrome()
    browser.set_window_size(0,0)
    return browser

def ClickAll(button_id,browser):
    for port_dict in GetPortsInfo():
        url = "localhost:"+str(port_dict.get("ServeRemotePort"))
        browser.get(url)
        browser.find_element_by_id(button_id).click()

def ClickSignle(Robotname,button_id,browser):
    url = "localhost:"+str(GetServeRemotePort(Robotname))
    browser.get(url)
    browser.find_element_by_id(button_id).click()

def exit(browser):
    browser.quit()

# design the gui for testing
def gui():
    browser = init()
    #welcom page
    g.msgbox("welcom to the group HEXAController!")
    #information in Index page
    IndexMsg = "Chose your control command:"
    IndexTitle = "HEXAController"
    IndexItem = ["start all","stop all","single control","exit"]

    # information in chosing robot page
    SingleControlChoiceMsg = "Chose the robot name you want control:"
    SingleControlChoiceTitle = "RobotChose"
    SingleControlChoicesItem = GetRobotsName()

    #information in controlling single robot
    SingleControlMsg = "Chose your command to the controlled robot:"
    SingleControlTitle = "SingleController_"
    SingleControlItem = ["start","stop","exit"]

    # gui
    while 1:
        # index page
        IndexChoice = g.buttonbox(IndexMsg, IndexTitle, IndexItem )
        if IndexChoice == "start all":
            ClickAll("start",browser)
        elif IndexChoice == "stop all":
            ClickAll("stop",browser)
        elif IndexChoice == "single control":
            # new control page
            Robotname = g.choicebox(SingleControlMsg, SingleControlTitle, SingleControlChoicesItem )
            while 1:
                SignleChioce = g.buttonbox(SingleControlMsg, SingleControlTitle + Robotname, SingleControlItem)
                if SignleChioce == "start":
                    ClickSignle(Robotname,"start",browser)
                elif SignleChioce == "stop":
                    ClickSignle(Robotname,"stop",browser)
                elif SignleChioce == "exit":
                    break
        elif IndexChoice == "exit":
            exit(browser)
            break

gui()
